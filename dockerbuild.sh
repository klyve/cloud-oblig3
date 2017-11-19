docker build -t cloud3 .
docker build -t mongodb -f DockerfileDB .

docker container rm cloud3
docker container rm mongodb
docker container create -t --name cloud3 -i cloud3
docker container create -t --name mongodb -i mongodb