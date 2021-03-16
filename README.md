# AWS Secrets Manager UI
A Web application to manage AWS Secrets manager

## Features
 - Listing secrets grid with rich attributes
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
 - [x] Region selector
 - [ ] Support IBM Carbon
 - [ ] Support simple authentication
 - [ ] Support OAuth 2.0
 - [ ] Support filter by tags
 - [ ] Local storage sort, filter, search
 - [ ] Loading state
 - [ ] Create new secret
 - [ ] Delete secret
