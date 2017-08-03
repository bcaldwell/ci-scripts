# CI Scripts

A collection of modular scripts that are commonly run in CI. The goal of this project is to reduce the number of CI configuration files that have duplicate code. Environment variables are used to configure the scripts. To include a script add the following to the CI config:

## Installation
```
gem install ci-scripts
ci-scripts SCRIPT_NAME
```

## Scripts

### demo/test



### docker/build
Uses docker to build the docker image for

This script assumes the following binaries are installed:
- docker

#### Environment Variables

| Variable | Default | Required | Description |
|:---|:---|:---:|:---|
| DOCKER_IMAGE |  | ✔ | |
| IMAGE_TAG | git tag |  | My favoruite thing|
| BUILD_DOCKERFILE | Dockerfile |  | |

### docker/herokuish


This script depends on and will run the following other scripts:
- [docker/build](#docker/build)

### docker/login



#### Environment Variables

| Variable | Default | Required | Description |
|:---|:---|:---:|:---|
| DOCKER_USERNAME |  | ✔ | |
| DOCKER_PASSWORD |  | ✔ | |
| DOCKER_EMAIL | ci@ci-runner.com |  | |
| DOCKER_REGISTRY | hub.docker.com |  | |

### docker/push_branch



#### Environment Variables

| Variable | Default | Required | Description |
|:---|:---|:---:|:---|
| DOCKER_IMAGE |  | ✔ | |
| IMAGE_TAG | current git hash |  | |

### docker/push_latest



#### Environment Variables

| Variable | Default | Required | Description |
|:---|:---|:---:|:---|
| DOCKER_LATEST_BRANCH | master |  | |
| DOCKER_IMAGE |  | ✔ | |
| IMAGE_TAG | current git hash |  | |

### git/ssh_keys



### ruby/bundler



#### Environment Variables

| Variable | Default | Required | Description |
|:---|:---|:---:|:---|
| BUNDLER_INSTALL_PATH | vendor |  | |

### ruby/publish_gem



### ruby/rake_test



### ruby/rubocop



