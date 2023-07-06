## Create a docker file

A Dockerfile is a text file that contains a series of instructions used to create a docker images

## Docker file vs Docker compose file

### A docker file

A Dockerfile is a text document that contains all the commands or instructions a user could call on the command line to `assemble an image`. Using the docker build command, users can create an automated build that executes several command-line instructions in succession. The end result of a Dockerfile is a `Docker image` that can be run in a Docker container.

For example, a Dockerfile might include instructions to:

- Set up the base OS (e.g., Ubuntu, Alpine).
- Install the necessary system libraries.
- Install the necessary programming languages (e.g., Python, Java).
- Set environment variables.
- Copy the application code into the container.
- Specify the default command to run when the container starts.

### Docker compose file

On the other hand, Docker Compose is a tool for defining and managing `multi-container Docker applications`. It uses a YAML file (docker-compose.yml) to configure the application's services, networks, and volumes. Then, with a single command (docker-compose up), you can create and start all the services specified in the configuration.

For example, a docker-compose.yml file might include instructions to:

- Create three services: a web server, a database, and a Redis instance.
- Set up network connections between them.
- Set up shared volumes.

## Some commands

```
docker images
```

- Remove an docker images

```
docker rmi <image-id>
```

- Run a docker image

```
docker run --name simplebank -p 8080:8080 -e GIN_MODE=release simplebank:latest
```

- Stop && remove a docker container

```
docker remove image-name
```

- Inspect a docker image container

```
docker container inspect simplebank
```

### Connect two stand-alone containers

- Create a network

```
docker network create bank-network
```

- Connect the `postgres15` containers to the network

```
docker network connect bank-network postgres15
```

- Run the `simplebank` container and connect it to the `bank-network`

```
docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:kien6034@postgres15:5432/simple_bank?sslmode=disable" simplebank:latest
```
