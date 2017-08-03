# CI Scripts

A collection of modular scripts that are commonly run in CI. The goal of this project is to reduce the number of CI configuration files that have duplicate code. Environment variables are used to configure the scripts. To include a script add the following to the CI config:

## Installation
```
gem install ci-scripts
ci-scripts SCRIPT_NAME
```

## Scripts

### demo/test
TODO: Some description of what I do

### docker/build
TODO: Some description of what I do

#### Environment Variables
| Variable | Default | Required | Description |
|:--|:--|:--|
| DOCKER_IMAGE |  | ✔ | |
| IMAGE_TAG | git tag |  | My favoruite thing|
| BUILD_DOCKERFILE | Dockerfile |  | |

### docker/herokuish
TODO: Some description of what I do

### docker/login
TODO: Some description of what I do

#### Environment Variables
| Variable | Default | Required | Description |
|:--|:--|:--|
| DOCKER_USERNAME |  | ✔ | |
| DOCKER_PASSWORD |  | ✔ | |
| DOCKER_EMAIL | ci@ci-runner.com |  | |
| DOCKER_REGISTRY | i |  | |

### docker/push_branch
TODO: Some description of what I do

#### Environment Variables
| Variable | Default | Required | Description |
|:--|:--|:--|
| DOCKER_IMAGE |  | ✔ | |
| IMAGE_TAG |  |  | |

### docker/push_latest
TODO: Some description of what I do

#### Environment Variables
| Variable | Default | Required | Description |
|:--|:--|:--|
| DOCKER_LATEST_BRANCH | master |  | |
| DOCKER_IMAGE |  | ✔ | |
| IMAGE_TAG |  |  | |

### git/ssh_keys
TODO: Some description of what I do

### ruby/bundler
TODO: Some description of what I do

#### Environment Variables
| Variable | Default | Required | Description |
|:--|:--|:--|
| BUNDLER_INSTALL_PATH | vendor |  | |

### ruby/publish_gem
TODO: Some description of what I do

### ruby/rake_test
TODO: Some description of what I do

### ruby/rubocop
TODO: Some description of what I do

