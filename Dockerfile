FROM golang:1.14.2-alpine3.11 as build

ARG VERSION

LABEL description="Real-time HTTP Intrusion Detection"
LABEL repository="https://github.com/kitabisa/teler"
LABEL maintainer="dwisiswant0"

WORKDIR /app
COPY ./go.mod .
RUN go mod download

COPY . .
RUN go build -ldflags "-s -w -X ktbs.dev/teler/common.Version=${VERSION}" \
	-o ./bin/teler ./cmd/teler 

FROM alpine:latest

COPY --from=build /app/bin/teler /bin/teler
ENV HOME /
ENTRYPOINT ["/bin/teler"]
