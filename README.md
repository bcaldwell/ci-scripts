# CI Scripts

A collection of modular scripts that are commonly run in CI. The goal of this project is to reduce the number of CI configuration files that have duplicate code. Environment variables are used to configure the scripts. To include a script add the following to the CI config:

```
wget -qO- https://raw.githubusercontent.com/benjamincaldwell/ci-scripts/master/script_path | ruby 
```

```
wget -qO- https://raw.githubusercontent.com/benjamincaldwell/ci-scripts/master/run | bash -s \
  docker/build \
  docker/lastest_master
```

<!--Since this allows remove code execuation in the CI environment, it is suggested that this repo is forked so -->
