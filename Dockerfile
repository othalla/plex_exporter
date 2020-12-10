FROM golang:alpine as build
ADD . /source
RUN cd /source && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o plex_exporter -v
FROM alpine
COPY --from=build /source/plex_exporter /plex_exporter
ENTRYPOINT ["/plex_exporter"]
