## Handy function
define docker_cmd
	@docker run \
		-ti --rm \
		--name bentobox-crawler \
		--net=host \
		-v ${PWD}:/go/src/github.com/consensys/event-crawler \
		-v ${PWD}/build/pkg:/go/pkg \
		-w /go/src/github.com/consensys/event-crawler \
		infura/golang-dev \
		$(1);
endef

## Need to get into the container and run bash?
bash:
	$(call docker_cmd, /bin/bash)

## Install all your dependencies
bootstrap:
	$(call docker_cmd, glide install)

run-psql:
	docker run \
		-ti --rm \
		--net=host \
		--name psql \
		-e POSTGRES_PASSWORD=mysecretpassword \
		-v ${PWD}:/workdir \
		-v ${HOME}/.psql:/var/lib/postgresql/data \
		postgres

# Will compile for debian. That's the container
crawler:
	$(call docker_cmd, go build -o build/bin/crawler -i main.go)

# Will run inside the dev container
run-crawler:
	$(call docker_cmd, ./build/bin/crawler)
