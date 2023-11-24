package status

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	MINER_IMAGE = "cesslab/cess-bucket"
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
	cli DockerCli
}

func NewClient() (*Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, errors.Wrap(err, "create new docker client error")
	}

	return &Client{cli}, nil
}

func (cli *Client) ListContainers() ([]Container, error) {
	list, err := cli.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "list containers error")
	}
	containers := make([]Container, len(list))
	for i, c := range list {
		name := "no name"
		if len(c.Names) > 0 {
			name = strings.TrimPrefix(c.Names[0], "/")
		}
		containers[i] = Container{
			ID:      c.ID,
			Names:   c.Names,
			Name:    name,
			Image:   c.Image,
			ImageID: c.ImageID,
			Command: c.Command,
			Created: c.Created,
			State:   c.State,
			Status:  c.Status,
		}
	}
	return containers, nil
}

func (cli *Client) ContainerStats(ctx context.Context, cid string, stats chan<- ContainerStat) error {
	response, err := cli.cli.ContainerStats(ctx, cid, true)

	if err != nil {
		return err
	}

	go func() {
		log.Debugf("starting to stream stats for: %s", cid)
		defer response.Body.Close()
		decoder := json.NewDecoder(response.Body)
		var v *types.StatsJSON
		for {
			if err := decoder.Decode(&v); err != nil {
				if err == context.Canceled || err == io.EOF {
					log.Debugf("stopping stats streaming for container %s", cid)
					close(stats)
					return
				}
				log.Errorf("decoder for stats api returned an unknown error %v", err)
			}

			ncpus := uint8(v.CPUStats.OnlineCPUs)
			if ncpus == 0 {
				ncpus = uint8(len(v.CPUStats.CPUUsage.PercpuUsage))
			}

			var (
				cpuDelta    = float64(v.CPUStats.CPUUsage.TotalUsage) - float64(v.PreCPUStats.CPUUsage.TotalUsage)
				systemDelta = float64(v.CPUStats.SystemUsage) - float64(v.PreCPUStats.SystemUsage)
				cpuPercent  = int64((cpuDelta / systemDelta) * float64(ncpus) * 100)
				memUsage    = int64(calculateMemUsageUnixNoCache(v.MemoryStats))
				memPercent  = int64(float64(memUsage) / float64(v.MemoryStats.Limit) * 100)
			)

			if cpuPercent > 0 || memUsage > 0 {
				select {
				case <-ctx.Done():
					close(stats)
					return
				case stats <- ContainerStat{
					ID:            cid,
					CPUPercent:    cpuPercent,
					MemoryPercent: memPercent,
					MemoryUsage:   memUsage,
				}:
				}
			}
		}
	}()

	return nil
}

func (cli *Client) ExeCommand(cid string, config types.ExecConfig) ([]byte, error) {
	execId, err := cli.cli.ContainerExecCreate(context.Background(), cid, config)
	if err != nil {
		return nil, errors.Wrap(err, "exe cmd in container error")
	}
	resp, err := cli.cli.ContainerExecAttach(context.Background(), execId.ID, types.ExecStartCheck{})
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
	return cli.cli.Ping(ctx)
}
