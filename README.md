# CI Scripts

A collection of modular scripts that are commonly run in CI. The goal of this project is to reduce the number of CI configuration files that have duplicate code. Environment variables are used to configure the scripts. To include a script add the following to the CI config:


```
gem install ci-scripts
ci-scripts SCRIPT_NAME
```

<!--Since this allows remove code execuation in the CI environment, it is suggested that this repo is forked so -->
