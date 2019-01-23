# drone-azure-web-apps

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-azure-web-apps/status.svg)](http://cloud.drone.io/drone-plugins/drone-azure-web-apps)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![](https://images.microbadger.com/badges/image/plugins/azure-web-apps.svg)](https://microbadger.com/images/plugins/azure-web-apps "Get your own image badge on microbadger.com")
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-azure-web-apps?status.svg)](http://godoc.org/github.com/drone-plugins/drone-azure-web-apps)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-azure-web-apps)](https://goreportcard.com/report/github.com/drone-plugins/drone-azure-web-apps)

Drone plugin to deploy or update a project on Azure Web Apps. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go test -cover ./...
go build -v -a -tags netgo -o release/linux/amd64/drone-azure-web-apps
```

## Docker

Build the Docker image with the following commands:

```
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.linux.amd64 --tag plugins/azure-web-apps .
```

### Usage

```
docker run --rm \
  -e PLUGIN_USERNAME=octocat \
  -e PLUGIN_PASSWORD=p455w0rd \
  -e PLUGIN_SITE=awesome \
  -e PLUGIN_FORCE=true \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/azure-web-apps
```
