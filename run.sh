# build the docker image
docker build -t client_ip .

# run the docker image
docker run -d -p 8080:8080 client_ip
