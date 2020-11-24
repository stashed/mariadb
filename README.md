[![Go Report Card](https://goreportcard.com/badge/stash.appscode.dev/mariadb)](https://goreportcard.com/report/stash.appscode.dev/mariadb)
[![Build Status](https://travis-ci.org/stashed/mariadb.svg?branch=master)](https://travis-ci.org/stashed/mariadb)
[![Docker Pulls](https://img.shields.io/docker/pulls/stashed/stash-mariadb.svg)](https://hub.docker.com/r/stashed/stash-mariadb/)
[![Slack](https://slack.appscode.com/badge.svg)](https://slack.appscode.com)
[![Twitter](https://img.shields.io/twitter/follow/appscodehq.svg?style=social&logo=twitter&label=Follow)](https://twitter.com/intent/follow?screen_name=AppsCodeHQ)

# MariaDB

MariaDB backup and restore plugin for [Stash by AppsCode](https://appscode.com/products/stash).

## Install

Install MariaDB 10.5.8 backup or restore plugin for Stash as below.

```console
helm repo add appscode https://charts.appscode.com/stable/
helm repo update
helm install appscode/stash-mariadb --name=stash-mariadb-10.5.8 --version=10.5.8
```

To install catalog for all supported MariaDB versions, please visit [here](https://github.com/stashed/catalog).

## Uninstall

Uninstall MariaDB 10.5.5 backup or restore plugin for Stash as below.

```console
helm delete stash-mariadb-10.5.8
```

## Support

To speak with us, please leave a message on [our website](https://appscode.com/contact/).

To join public discussions with the Stash community, join us in the [AppsCode Slack team](https://appscode.slack.com/messages/C8NCX6N23/details/) channel `#stash`. To sign up, use our [Slack inviter](https://slack.appscode.com/).

To receive product annoucements, follow us on [Twitter](https://twitter.com/KubeStash).

If you have found a bug with Stash or want to request new features, please [file an issue](https://github.com/stashed/project/issues/new).
