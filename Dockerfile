LABEL dev.ktbs.project.name="teler"
LABEL description="Real-time HTTP Intrusion Detection"
LABEL version="0.0.1-dev3"
LABEL repository="https://github.com/kitabisa/teler"
LABEL maintainer="dwisiswant0"

FROM golang:1.14.2-alpine3.11 as build
RUN apk --no-cache add git
ENV GO111MODULE on
RUN go get ktbs.dev/teler/cmd/teler; exit 0

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /go/bin/teler /bin/teler
ENV HOME /
ENTRYPOINT ["/bin/teler"]
