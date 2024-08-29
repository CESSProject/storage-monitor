# sh build.sh v0.0.1
docker build -t cesslab/watchdog:"$1" .
docker tag cesslab/watchdog:"$1" cesslab/watchdog:latest