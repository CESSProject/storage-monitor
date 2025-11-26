# storage-miner-monitor

This service is designed to make it easier for users to view the working status of the storage node cluster. 

It is divided into two modules: watchdog runs on the server together with the storage node to collect storage-node statistics, and finnaly display statistics by watchdog-web service.

## How to deploy

```sh
# watchdog-web

 docker run -d --name watchdog-web -p 13080:13080 cesslab/watchdog-web:latest
```

```sh
# watchdog-web
# config.yaml demo: https://doc.cess.network/cess-miners/storage-miner/multi-miner

docker run -d --name watchdog -p 13081:13081 --volume=/opt/cess/mineradm/build/config.yaml:/opt/miner/config.yaml --volume=/var/run/docker.sock:/var/run/docker.sock cesslab/watchdog:latest
```

## Deploy by mineradm

Doc: https://doc.cess.network/cess-miners/storage-miner/multi-miner
