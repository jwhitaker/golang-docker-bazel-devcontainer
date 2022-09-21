# golang-docker-bazel-devcontainer

Playing around with Golang, Docker and Bazel in a devcontainer.

## Building

To Generate the docker image:

```bash
bazelisk run //cmd/server:docker
```

To run the built image:

```bash
docker run -p 3000:3000 golang-docker-bazel-devcontainer/cmd/server
```