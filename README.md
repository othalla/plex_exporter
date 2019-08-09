# plex_exporter

[![Go Report Card](https://goreportcard.com/badge/github.com/othalla/plex_exporter)](https://goreportcard.com/report/github.com/othalla/plex_exporter)
[![Build Status](https://travis-ci.org/othalla/plex_exporter.svg?branch=master)](https://travis-ci.org/othalla/plex_exporter)

## Description

Prometheus exporter for Plex Media Server.
This exporter query your Plex Media Server installation directly without passing by Plex.tv.

## Metrics

- Sessions : total active
- Transcoding Sessions : total active
- Libraries : number of medias

## Configuration

```json
{
  "server": {
    "address": "plex",
    "port": 32400,
    "token": "token"
  }
}
```

## Usage

```console
# ./plex_exporter --config config.json
2019/08/09 19:58:34 Starting exporter...
```

## TODO

Secure transport (with TLS) while requesting metrics. Server token is actually passed through
headers and can be easilly haced if used with http.
First thing is to move it to https without verification and check is in the futur.
