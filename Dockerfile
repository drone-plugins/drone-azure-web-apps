# Docker image for the Drone Azure Web Apps plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-azure-web-apps
#     make deps build
#     docker build --rm=true -t plugins/drone-azure-web-apps .

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates \
    git \
    openssh \
    curl \
    perl && \
  rm -rf /var/cache/apk/*

ADD drone-azure-web-apps /bin/
ENTRYPOINT ["/bin/drone-azure-web-apps"]
