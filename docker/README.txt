To build the Docker image:
  % docker build -t johansteyn/helloworld .

To help debug a build:
  % docker build -t johansteyn/helloworld --progress=plain --no-cache .

To list images:
  % docker images

To remove the image:
  % docker rmi johansteyn/helloworld
In case of error, try to force the removal:
  % docker rmi --force johansteyn/helloworld

To run the container:
  % docker run -p 8080:8080 johansteyn/helloworld

To run the container in the background (detached mode), mapping host port 80 to container port 8080:
  % docker run -dp 80:8080 johansteyn/helloworld

To see what containers are running:
  % docker ps

To stop a container, specify the container ID or name:
  % docker stop <container-id>|<name>

Similarly, specify the container ID or name to start a stopped container restart a running container, or remove a stopped container:
  % docker start <container-id>|<name>
  % docker restart <container-id>|<name>
  % docker rm <container-id>|<name>

Use the "force" flag as a shortcut to stop and remove a running contaiuner: 
  % docker rm -f <container-id>|<name>

