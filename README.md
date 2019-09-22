#Prerequisites
```
install golang
https://golang.org/doc/install
```
#How to run
```
go run src/app/client_ip.go
```
# run locally
```
go build client_ip/src/app/client_ip.go
```

# build and run docker
```
bash run.sh
```


# debugging
If you are already running a container
you can run
```
docker container ls
docker container stop {whatever name of the container you need to stop}
```

# deployment Strategy
For deployment strategies, I think that it's probably wise to just use docker hub. It's also the first thing that came up on google lol.

# Considerations for high availibility 
On first breadth it seem like docker swarm is the industry standard for high availability 

theres a tutorial here
```
https://medium.com/brian-anstett-things-i-learned/high-availability-and-horizontal-scaling-with-docker-swarm-76e69845825e
```

I'm sure you could probably deploy it to any of the cloud platforms using GCP, AWS or Azure. 
If you wanted to do it locally, I would just use 3 intel nucs or even some raspberry pis. 

You could probably scale with multiple containers per vm with a number of ports that would connect to container via some kind of load balancer.
