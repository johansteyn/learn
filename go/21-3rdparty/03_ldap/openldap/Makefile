IMAGE = osixia/extend-osixia-openldap
VERSION = 1.4.0
CONTAINER = openldap-test-server

usage:
	@ echo "MongoDB OpenLDAP test server, preconfigured with users/groups"
	@ echo
	@ echo "Usage:"
	@ echo "  delete_container:	stops and deletes the container"
	@ echo "  clean: 	        deletes any existing containers/images"
	@ echo "  build: 	        builds/re-builds the Docker image"
	@ echo "  clean-build:      runs the clean target and rebuilds the image from scratch"
	@ echo "  run: 		        runs a previously built Docker image (foreground)"
	@ echo "  fresh:	        performs a clean rebuild and runs the OpenLDAP container (foreground)"
	@ echo "  start: 		    launches the Docker image in the background"
	@ echo "  shell: 		    connects to a running container"
	@ echo

delete_container:
	- docker rm -f $(CONTAINER)

clean: delete_container
	- docker rmi $(IMAGE):$(VERSION)

build:
	docker build $(EXTRA_PARAM) -t $(IMAGE):$(VERSION) .

clean-build:
	@ $(MAKE) EXTRA_PARAM="--force-rm --rm" clean build

run:
	@ docker run --rm -i --hostname $(CONTAINER) --name $(CONTAINER) -p 1389:389 -p 636:636 -t $(IMAGE):$(VERSION)

start:
	@ docker run --rm -d --hostname $(CONTAINER) --name $(CONTAINER) -p 1389:389 -p 636:636 -t $(IMAGE):$(VERSION)

fresh:
	$(MAKE) clean build run

shell:
	docker exec -it $(CONTAINER) /bin/bash
