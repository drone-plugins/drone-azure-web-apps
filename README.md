# drone-azure-web-apps

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-azure-web-apps/status.svg)](http://beta.drone.io/drone-plugins/drone-azure-web-apps)
[![](https://badge.imagelayers.io/plugins/drone-azure-web-apps:latest.svg)](https://imagelayers.io/?images=plugins/drone-azure-web-apps:latest 'Get your own badge on imagelayers.io')

Drone plugin for deploying to Azure Web Apps

## Usage

```sh
./drone-azure-web-apps <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
      "username": "octocat",
      "password": "my_password",
      "site": "awesome",
      "force": true
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```sh
make deps build
docker build --rm=true -t plugins/drone-azure-web-apps .
```

### Example

```sh
docker run -i plugins/drone-azure-web-apps <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
      "username": "octocat",
      "password": "my_password",
      "site": "awesome",
      "force": true
    }
}
EOF
```
