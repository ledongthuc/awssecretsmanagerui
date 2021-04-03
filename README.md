# AWS Secrets Manager UI
[![build](https://github.com/ledongthuc/awssecretsmanagerui/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/ledongthuc/awssecretsmanagerui/actions/workflows/build.yml)

A Web application to manage AWS Secrets manager

![AWS Secrets Manager UI](https://user-images.githubusercontent.com/1828895/113342648-07e31280-932f-11eb-9949-9a2ec0c08e40.png)

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

## AWS Credential

AWS Secrets Manager UI tool uses AWS configuration credential to authenticate requests. 

### Credential file

More detail: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html

### Credential environment variables

More detail: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html

## Screenshoots

![Table](https://user-images.githubusercontent.com/1828895/113339958-8ccc2d00-932b-11eb-9738-4c7fea424985.png)

![Detail](https://user-images.githubusercontent.com/1828895/113339995-9c4b7600-932b-11eb-80a3-b550ad6893b0.png)

![Text data](https://user-images.githubusercontent.com/1828895/113340032-abcabf00-932b-11eb-9b2b-568f7d443873.png)

![Binary data](https://user-images.githubusercontent.com/1828895/113340095-bc7b3500-932b-11eb-8976-b0c6efdcd02a.png)

## TODO
 - [ ] Local storage sort, filter, search
 - [ ] Release macapp
 - [ ] Support IBM Carbon
 - [ ] Support filter by tags
 - [ ] Create new secret
 - [ ] Delete secret

# Contributor

 - Thuc Le
 - sinthithuthao@gmail.com (many thanks for logo)
