package core

import (
	"context"
	"encoding/json"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/CESSProject/watchdog/internal/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"strings"
)

type DockerCli interface {
	ContainerList(context.Context, types.ContainerListOptions) ([]types.Container, error)
	ContainerLogs(context.Context, string, types.ContainerLogsOptions) (io.ReadCloser, error)
	Events(context.Context, types.EventsOptions) (<-chan events.Message, <-chan error)
	ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error)
	ContainerStats(ctx context.Context, containerID string, stream bool) (types.ContainerStats, error)
	Ping(ctx context.Context) (types.Ping, error)

	//exec
	ContainerExecAttach(ctx context.Context, execID string, config types.ExecStartCheck) (types.HijackedResponse, error)
	ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
}

type Client struct {
	dockerCli DockerCli
}

func NewClient(host model.HostItem) (*Client, error) {
	dockerHost := "tcp://" + host.IP + ":" + host.Port
	ip := net.ParseIP(host.IP)
	if util.IsPrivateIP(ip) {
		cli, err := client.NewClientWithOpts(
			client.WithHost(dockerHost),
			client.WithAPIVersionNegotiation(),
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return &Client{cli}, nil
	} else {
		cli, err := client.NewClientWithOpts(
			client.WithHost(dockerHost),
			client.WithTLSClientConfig(host.CAPath, host.CertPath, host.KeyPath),
			client.WithAPIVersionNegotiation(),
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return &Client{cli}, nil
	}
}

func (cli *Client) ListContainers() ([]model.Container, error) {
	list, err := cli.dockerCli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "list containers error")
	}
	containers := make([]model.Container, len(list))
	for i, c := range list {
		name := "no name"
		if len(c.Names) > 0 {
			name = strings.TrimPrefix(c.Names[0], "/")
		}
		containers[i] = model.Container{
			ID:            c.ID,
			Names:         c.Names,
			Name:          name,
			Image:         c.Image,
			ImageID:       c.ImageID,
			Command:       c.Command,
			Created:       c.Created,
			State:         c.State,
			Status:        c.Status,
			CPUPercent:    float64(0),
			MemoryPercent: float64(0),
			MemoryUsage:   uint64(0),
		}
	}
	return containers, nil
}

func (cli *Client) ContainerStats(ctx context.Context, cid string) (model.ContainerStat, error) {
	response, err := cli.dockerCli.ContainerStats(ctx, cid, false)
	if err != nil {
		return model.ContainerStat{}, nil
	}
	var v *types.StatsJSON
	decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&v); err == io.EOF {
		log.Fatalf("Docker Stats API Response io.EOF")
		return model.ContainerStat{}, nil
	} else if err != nil {
		panic(err)
	}
	cpuPercent := calculateCPUPercentUnix(v)
	memUsage := v.MemoryStats.Usage
	memLimit := v.MemoryStats.Limit
	memPercent := float64(memUsage) / float64(memLimit) * 100.0
	res := model.ContainerStat{
		CPUPercent:    cpuPercent,
		MemoryPercent: memPercent,
		MemoryUsage:   memUsage / 1048576,
	}
	return res, nil
}

func calculateCPUPercentUnix(v *types.StatsJSON) float64 {
	var (
		cpuPercent  = 0.0
		cpuDelta    = float64(v.CPUStats.CPUUsage.TotalUsage) - float64(v.PreCPUStats.CPUUsage.TotalUsage)
		systemDelta = float64(v.CPUStats.SystemUsage) - float64(v.PreCPUStats.SystemUsage)
		onlineCPUs  = float64(v.CPUStats.OnlineCPUs)
	)

	if onlineCPUs == 0.0 {
		onlineCPUs = float64(len(v.CPUStats.CPUUsage.PercpuUsage))
	}
	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * onlineCPUs * 100.0
	}
	return cpuPercent
}

func (cli *Client) ExeCommand(cid string, config types.ExecConfig) ([]byte, error) {
	execId, err := cli.dockerCli.ContainerExecCreate(context.Background(), cid, config)
	if err != nil {
		return nil, errors.Wrap(err, "exe cmd in container error")
	}
	resp, err := cli.dockerCli.ContainerExecAttach(context.Background(), execId.ID, types.ExecStartCheck{})
	if err != nil {
		return nil, errors.Wrap(err, "exe cmd in container error")
	}
	defer resp.Close()

	buf, err := io.ReadAll(resp.Reader)
	if err != nil {
		return nil, errors.Wrap(err, "exe cmd in container error")
	}

	return buf, nil
}

func (cli *Client) Ping(ctx context.Context) (types.Ping, error) {
	return cli.dockerCli.Ping(ctx)
}
