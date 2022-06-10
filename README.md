# plex_exporter

[![Go Report Card](https://goreportcard.com/badge/github.com/othalla/plex_exporter)](https://goreportcard.com/report/github.com/othalla/plex_exporter)
[![Build Status](https://travis-ci.org/othalla/plex_exporter.svg?branch=master)](https://travis-ci.org/othalla/plex_exporter)

## Description

Prometheus exporter for Plex Media Server.
This exporter query your Plex Media Server installation directly without passing by Plex.tv.

## Metrics

- Plex Media Server information: version
- Sessions: total active
- Transcoding Sessions: total active
- Libraries: number of medias

output

```
# HELP plex_info Plex media server information
# TYPE plex_info Gauge
plex_info{version="1.26.2.5797-5bd057d2b"} 1
# HELP plex_media_server_library_media_count Number of medias in a plex library
# TYPE plex_media_server_library_media_count gauge
plex_media_server_library_media_count{name="Animes",type="show"} 33
plex_media_server_library_media_count{name="Cartoons",type="show"} 6
plex_media_server_library_media_count{name="Movies",type="movie"} 283
plex_media_server_library_media_count{name="TV Shows",type="show"} 36
# HELP plex_sessions_active_count Number of active Plex sessions
# TYPE plex_sessions_active_count gauge
plex_sessions_active_count 2
# HELP plex_transcode_sessions_active_count Number of active Plex transcoding sessions
# TYPE plex_transcode_sessions_active_count gauge
plex_transcode_sessions_active_count 2
```

## Configuration

```json
{
 "exporter": {
    "port": 9594
  },
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
2021/10/22 08:18:11 Starting exporter on port 9594 ...
```

## TODO

Secure transport (with TLS) while requesting metrics. Server token is actually passed through
headers and can easilly be hacked if used with http.
First thing is to move it to https without verification and check is in the future.
