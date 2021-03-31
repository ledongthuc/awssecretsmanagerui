# AWS Secrets Manager UI
[![build](https://github.com/ledongthuc/awssecretsmanagerui/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/ledongthuc/awssecretsmanagerui/actions/workflows/build.yml)
A Web application to manage AWS Secrets manager

## Features
 - Listing secrets grid with many columns
 - Support region switching
 - Secrets attributes sorting
 - All grid column searching
 - Editing text-based secrets directly
 - Download and upload binary secrets

## Quick starts

Run:

```
docker run -ti -p 30301:30301 -v $HOME/.aws:/root/.aws ledongthuc/awssecretsmanagerui:unstable
```

Make sure you configure [AWS credential](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/setup-credentials.html). Then you can access http://localhost:30301

![Quickstart](https://user-images.githubusercontent.com/1828895/113198932-cd15a780-9266-11eb-98e3-2a13487c3bd0.png)

## Build from source

**Requirement:**
 - Go 1.16
 - Makefile
 - npm 6.14

```
make build;
```

Binary will available in folder "./build/". Run it and you can access through http://localhost:30301

## Screenshoots


## TODO
 - [ ] Update docker to build with HEADLESS
 - [ ] Local storage sort, filter, search
 - [ ] Release macapp
 - [ ] Support IBM Carbon
 - [ ] Support filter by tags
 - [ ] Create new secret
 - [ ] Delete secret
