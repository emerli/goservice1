#build image
sudo docker container stop service1 && sudo docker container rm service1 && sudo docker build -t service -f Dockerfile  .

#run of services master and slave
sudo docker run -d -p 8080:8080 --name service1 --env SERVICE_ADDRESS=:8080 --env SERVICE_ROLE=MASTER service:latest
sudo docker run -d -p 8081:8080 --name service2 --env SERVICE_ADDRESS=:8080 --env SERVICE_ROLE=SLAVE service:latest
