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

### Build from source

**Requirement:**
 - Go 1.16
 - Makefile
 - npm 6.14

```
make build;
```

Binary will available in folder "./build/"

### Install from binary

Download from [Download Page](https://github.com/ledongthuc/awssecretsmanagerui/wiki/Download)

## Features

### Support rich table



## TODO
 - [ ] Update docker to build with HEADLESS
 - [ ] Local storage sort, filter, search
 - [ ] Release macapp
 - [ ] Support IBM Carbon
 - [ ] Support filter by tags
 - [ ] Create new secret
 - [ ] Delete secret
