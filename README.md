# AWS Secrets Manager UI
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
./build/aws-secrets-manager-ui
```

System serve through `localhost:3000`

### Docker

```
docker run -p 3000:3000 ledongthuc/awssecretsmanagerui:latest
```
System serve through `localhost:3000`

## Configuration

### AWS Authentication

### Application setting

**Change port**

**Change host binding**


## TODO
 - [ ] Update docker to build with HEADLESS
 - [ ] Local storage sort, filter, search
 - [ ] Release macapp
 - [ ] Support IBM Carbon
 - [ ] Support filter by tags
 - [ ] Create new secret
 - [ ] Delete secret
