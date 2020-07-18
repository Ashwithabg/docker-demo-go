# go-app

run following command to start golang server

```
go run main.go
```
open browser and hit `localhost:8080` ,you should see `hello` message

##Docker commands

```text
docker ps
docker ps -a
docker container ls 
docker images
```

Docker run image
```text
docker run test-nginx-docker-ninja
docker run -d --name custom-server test-nginx-docker-ninja
```

get inside container
```text
docker exec -it fbb62e6e37d9 bash
cd /usr/share/nginx/html/
echo "Hi docker ninja" > docker.html
exit
```

commit changes
```text
docker commit testNginx7 test-nginx-docker-ninja
docker images
```

or can use dockerfile.
create dockerfile  
```dockerfile
FROM nginx

WORKDIR /usr/share/nginx/html
RUN echo "hi docker" > docker1.html
```

then build image
```text
docker build -t nginx:auto .
```

##DockerFile
base image

```dockerfile
FROM nginx
```

mkdir /usr/app/src and cd /usr/app/src

```dockerfile
WORKDIR /usr/app/src
```

copy files from host machine to container
using add we can also give url. it will download the file and put it here

```dockerfile
ADD /src/path/from/host /dest/path/to/image
COPY /src/path/from/host /dest/path/to/image
```

pass env variables
```dockerfile
ENV key value
```

execute 
```dockerfile
RUN apt-get install python3 git
```

Expose ports from container from inter container communication
```dockerfile
EXPOSE 8080 9090
```

```dockerfile
CMD ["executable", "param1", "param2", ...]
ENTRYPOINT ["executable", "param1", "param2", ...]
```
