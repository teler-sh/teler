FROM golang:1.18-buster as build

ARG VERSION

LABEL description="Real-time HTTP Intrusion Detection"
LABEL repository="https://github.com/kitabisa/teler"
LABEL maintainer="dwisiswant0"

WORKDIR /app
COPY ./go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags \
	"-s -w -X teler.app/common.Version=${VERSION}" \
	-o ./bin/teler .

FROM alpine:latest

COPY --from=build /app/bin/teler /bin/teler
ENV HOME /

ENTRYPOINT ["/bin/teler"]
