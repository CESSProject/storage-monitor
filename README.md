# storage-monitor

The storage-monitor is designed to make it easier for users to view the working status of the storage node cluster. It is divided into two modules: watchdog runs on the server together with the storage node to collect storage-node information, while node-monitor as a data server used to summarize the storage-node information collected by each watchdog and display it to the user in the form of web page.

You can deploy the node-monitor program through the following command:

```sh
 docker run -d --name watchdog-web -p 13080:13080 cesslab/watchdog-web:latest
```
The default port of the program in docker is 13080. You can map it to any port on the host machine and configure access policies such as firewalls to ensure that the port can be accessed from the outside.

Both watchdog and node-monitor can be deployed through docker containers, and only one command is needed to start each container. First, run the watchdog program on the server where you are running the storage node container with the following command:

```sh
 docker run -d --name watchdog -p 13081:13081 --volume=/opt/cess/mineradm/build/config.yaml:/opt/miner/config.yaml --volume=/var/run/docker.sock:/var/run/docker.sock cesslab/watchdog:latest
```

In subsequent version iterations, we will continue to optimize the program and bring richer functions.

Doc: https://doc.cess.network/cess-miners/storage-miner/multi-miner
