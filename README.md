# AWS Secrets Manager UI

[![Built with WeBuild](https://raw.githubusercontent.com/webuild-community/badge/master/svg/WeBuild.svg)](https://webuild.community) [![build](https://github.com/ledongthuc/awssecretsmanagerui/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/ledongthuc/awssecretsmanagerui/actions/workflows/build.yml)

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
docker run -ti \
 -p 30301:30301 \
 -e AWS_ACCESS_KEY_ID=123456789012 \
 -e AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY \
 ledongthuc/awssecretsmanagerui:unstable
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

### Credential environment variables (recommend)

More detail: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html

### Minimum Permission

 - Easy policy name: SecretsManagerReadWrite : https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-policies.html

 - Or custom action permission: `secretsmanager:*`


## Authentication

Default, AWS Secrets manager UI disable authentication.

AWS Secrets manager supports basic auth through two variable environments, in order enable it, try with 2 variable environments:

 - `AUTH_ENABLED=true`
 - `AUTH_ACCOUNTS="{\"admin\":\"5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5\"}"`

`AUTH_ACCOUNTS` define accounts' username and password. Password follows SHA-256, https://emn178.github.io/online-tools/sha256.html

## Filter secret by names

Default, AWS Secrets manager UI get all secrets.

AWS Secrets manager support filter secrets by names through variable environment, in order to use it, you can try with variable environment:
- `FILTER_NAMES=production,development`

`FILTER_NAMES` defines which secrets you want to provide with AWS Secrets manager, each secret name is separated by comma. In the above example, you want to show only secrets with name `production` and `development`.

## Screenshoots

![Table](https://user-images.githubusercontent.com/1828895/113339958-8ccc2d00-932b-11eb-9738-4c7fea424985.png)

![Detail](https://user-images.githubusercontent.com/1828895/113339995-9c4b7600-932b-11eb-80a3-b550ad6893b0.png)

![Text data](https://user-images.githubusercontent.com/1828895/113340032-abcabf00-932b-11eb-9b2b-568f7d443873.png)

![Binary data](https://user-images.githubusercontent.com/1828895/113340095-bc7b3500-932b-11eb-8976-b0c6efdcd02a.png)

## TODO
 - [ ] Local storage sort, filter, search
 - [ ] Support MoaiJS
 - [ ] Support filter by tags
 - [ ] All POST API
 - [ ] Login page
 - [ ] Create new secret
 - [ ] Delete secret

# Contributor

 - Thuc Le
 - sinthithuthao@gmail.com (many thanks for logo)
 - [duysmile](github.com/duysmile)
