## Create a docker file

A Dockerfile is a text file that contains a series of instructions used to create a docker images

## Multi stage docker file

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
