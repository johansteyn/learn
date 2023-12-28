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

