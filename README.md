# pmsx_exporter

## Description

Prometheus exporter for Plex Media Server.
This exporter query your Plex Media Serverinstallation directly without passing by Plex.tv.

## Metrics

- Sessions : total active
- Libraries : number of medias

## TODO

Secure transport (with TLS) while requesting metrics. Server token is actually passed through
headers and can be easilly haced if used with http.
First thing is to move it to https without verification and check is in the futur.
