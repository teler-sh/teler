<p align="center">
  <a href="#"><img alt="teler" src="https://user-images.githubusercontent.com/25837540/97091757-7200d880-1668-11eb-82c4-e5c4971d2bc8.png" height="400" /></a>
  <h3 align="center"><b>teler</b></h3>
</p>

[![Kitabisa SecLab](https://img.shields.io/badge/kitabisa-security%20project-blue)](#)
[![License](https://img.shields.io/badge/License-Apache%202.0-yellowgreen)](https://github.com/kitabisa/teler/blob/development/LICENSE)
[![made with Go](https://img.shields.io/badge/made%20with-Go-brightgreen)](http://golang.org)
[![Version](https://img.shields.io/badge/version-0.0.1--beta3.2-blueviolet)](https://github.com/kitabisa/teler/releases)
[![Platform](https://img.shields.io/badge/platform-osx%2Flinux%2Fwindows-green)](#)
[![GitHub issues](https://img.shields.io/github/issues/kitabisa/teler)](https://github.com/kitabisa/teler/issues)

<!-- [![Gitter](https://badges.gitter.im/kitabisa-teler/community.svg)](https://gitter.im/kitabisa-teler/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) -->
<!-- [![Go report](https://goreportcard.com/badge/ktbs.dev/teler)](https://goreportcard.com/report/ktbs.dev/teler) -->

<p align="center">
  Real-time HTTP Intrusion Detection
  <br />
  <a href="https://github.com/kitabisa/teler/blob/development/.github/CONTRIBUTING.md">Contribute</a>
  ·
  <a href="https://github.com/kitabisa/teler/blob/development/CHANGELOG.md">What's new</a>
  ·
  <a href="https://github.com/kitabisa/teler/issues/new/choose">Report Bug</a>
  ·
  <a href="https://github.com/kitabisa/teler/issues/new/choose">Request Feature</a>
</p>

---

`teler` is an **real-time intrusion detection** and threat alert based on web log that runs in a **terminal** with resources that we collect and provide by the community. :heart:

[![teler](https://user-images.githubusercontent.com/25837540/97096468-f8ccaa00-1696-11eb-8830-0d3a7be45a2d.gif)](#)

## Table of Contents
- [Table of Contents](#table-of-contents)
- [Features](#features)
- [Why teler?](#why-teler)
- [Demo](#demo)
- [Installation](#installation)
  - [from Binary](#from-binary)
  - [using Docker](#using-docker)
  - [from Source](#from-source)
  - [from GitHub](#from-github)
- [Usage](#usage)
  - [Flags](#flags)
    - [Config](#config)
    - [Input](#input)
    - [Concurrency](#concurrency)
    - [Output](#output)
- [Configuration](#configuration)
  - [Log formats](#log-formats)
    - [Apache](#apache)
    - [Nginx](#nginx)
    - [Nginx Ingress](#nginx-ingress)
    - [Amazon S3](#amazon-s3)
    - [Elastic LB](#elastic-lb)
    - [CloudFront](#cloudfront)
  - [Threat rules](#threat-rules)
    - [Excludes](#excludes)
    - [Whitelists](#whitelists)
  - [Notification](#notification)
- [Contributors](#contributors)
  - [Resources](#resources)
- [Pronunciation](#pronunciation)
- [Changes](#changes)
- [License](#license)

## Features

* **Real-time**: Analyze logs and identify suspicious activity in real-time.

* **Alerting** _(roadmap)_: teler provides alerting when a threat is detected, push notifications include Slack, Telegram and Discord.

* **Latest resources**: Collections is continuously up-to-date, see [resources](#resources).

* **Minimal configuration**: You can just run it against your log file, write the log format and let
  teler analyze the log and show you alerts!

* **Flexible log formats**: teler allows any custom log format string! It all depends on how you write the log format in configuration file.

* **Incremental log processing**: Need data persistence rather than [buffer stream](https://linux.die.net/man/1/stdbuf)?
  teler has the ability to process logs incrementally through the on-disk persistence options.

## Why teler?

teler was designed to be a fast, terminal-based threat analyzer. Its core idea is to quickly analyze and hunt threats in real time!

## Demo

Here is a preview of `teler` with conditions of use as:

| **Buffer-streams** 	| **Incremental** 	|
|--------------------	|-----------------	|
| <a href="https://asciinema.org/a/367616" alt="teler"><img src="https://asciinema.org/a/367616.svg"></a>	| <a href="https://asciinema.org/a/367610" alt="teler"><img src="https://asciinema.org/a/367610.svg"></a>	|

## Installation

### from Binary

The installation is easy. You can download a prebuilt binary from [releases page](https://github.com/kitabisa/teler/releases), unpack and run! or run with:

```bash
▶ curl -sSfL 'https://ktbs.dev/get-teler.sh' | sh -s -- -b /usr/local/bin
```

### using Docker

Build the Docker image with:

```bash
▶ docker build -t teler https://github.com/kitabisa/teler.git
```

### from Source

If you have go1.13+ compiler installed and configured:

```bash
▶ GO111MODULE=on go get -v -u ktbs.dev/teler/cmd/teler
```

In order to update the tool, you can use `-u` flag with `go get` command.

### from GitHub

```bash
▶ git clone https://github.com/kitabisa/teler
▶ cd teler
▶ make build
▶ mv ./bin/teler /usr/local/bin
```

## Usage

Simply, teler can be run with:

```bash
▶ [buffers] | teler -c /path/to/config/teler.yaml
# or
▶ teler -i /path/to/access.log -c /path/to/config/teler.yaml
```

If you've built teler with a Docker image:

```bash
▶ [buffers] | docker run -i --rm -e TELER_CONFIG=/path/to/config/teler.yaml teler
# or
▶ docker run -i --rm -e TELER_CONFIG=/path/to/config/teler.yaml teler --input /path/to/access.log
```

### Flags

```bash
▶ teler -h
```

This will display help for the tool.

<img src="https://user-images.githubusercontent.com/25837540/95018744-d218e600-068b-11eb-8fa1-27ce2458a696.png">

Here are all the switches it supports.

| Flag                  | Description                                                   | Examples                                                  |
|---------------------- |-------------------------------------------------------------  |---------------------------------------------------------  |
| -c,<br> --config      | teler configuration file                                      | kubectl logs nginx \| teler -c /path/to/config/teler.yaml |
| -i,<br> --input       | Analyze logs from data persistence rather than buffer stream  | teler -i /var/log/nginx/access.log |
| -x,<br> --concurrent  | Set the concurrency level to analyze logs<br>(default: 20)    | tail -f /var/log/nginx/access.log \| teler -x 50 |
| -o,<br> --output      | Save detected threats to file                                 | teler -i /var/log/nginx/access.log -o /tmp/threats.log |
| -v,<br> --version     | Show current teler version                                    | teler -v |

#### Config

The `-c` flag is to specify teler configuration file.

```bash
▶ tail -f /var/log/nginx/access.log | teler -c /path/to/config/teler.yaml
```

This is **required**, but if you have defined `TELER_CONFIG` environment you don't need to use this flag, e.g.:

```bash
▶ export TELER_CONFIG="/path/to/config/teler.yaml"
▶ tail -f /var/log/nginx/access.log | teler
# or
▶ tail -f /var/log/nginx/access.log | TELER_CONFIG="/path/to/config/teler.yaml" teler
```

#### Input

Need log analysis incrementally? This `-i` flag is useful for that.

```bash
▶ teler -i /var/log/nginx/access.log
```

#### Concurrency

Concurrency is the number of logs analyzed at the same time. Default value teler provide is 20, you can change it by using `-x` flag.

```bash
▶ teler -i /var/log/nginx/access.log -x 50
```

#### Output

You can also save the detected threats into a file with `-o` flag.

```bash
▶ teler -i /var/log/nginx/access.log -o nginx-threat.log
```

## Configuration

`teler` requires a minimum of configuration to process and/or log analysis, and execute threats and/or alerts. See [teler.example.yaml](https://github.com/kitabisa/teler/blob/development/teler.example.yaml) for an example.

### Log Formats

Because we use `gonx` package to parse the log, you can write any log format. As an example:

#### Apache
```yaml
log_format: |
  $remote_addr - $remote_user [$time_local] "$request_method $request_uri $request_protocol" $status $body_bytes_sent
```

#### Nginx
```yaml
log_format: |
  $remote_addr $remote_user - [$time_local] "$request_method $request_uri $request_protocol" 
  $status $body_bytes_sent "$http_referer" "$http_user_agent"
```

#### Nginx Ingress
```yaml
log_format: |
  $remote_addr - [$remote_addr] $remote_user - [$time_local] 
  "$request_method $request_uri $request_protocol" $status $body_bytes_sent 
  "$http_referer" "$http_user_agent" $request_length $request_time 
  [$proxy_upstream_name] $upstream_addr $upstream_response_length $upstream_response_time $upstream_status $req_id
```

#### Amazon S3
```yaml
log_format: |
  $bucket_owner $bucket [$time_local] $remote_addr $requester $req_id $operationration $key 
  "$request_method $request_uri $request_protocol" $status $error_code $body_bytes_sent - 
  $total_time - "$http_referer" "$http_user_agent" $version_id $host_id 
  $signature_version $cipher_suite $http_auth_type $http_host_header $tls_version
```

#### Elastic LB
```yaml
log_format: |
  $time_local $elb_name $remote_addr $upstream_addr $request_processing_time 
  $upstream_processing_time $response_processing_time $status $upstream_status $body_received_bytes $body_bytes_sent 
  "$request_method $request_uri $request_protocol" "$http_user_agent" $cipher_suite $tls_version
```

#### CloudFront
```yaml
log_format: |
  $date $time $edge_location  $body_bytes_sent  $remote_addr  
  $request_method $http_host_header $requst_uri $status 
  $http_referer $http_user_agent  $request_query  $http_cookie  $edge_type  $req_id 
  $http_host_header $ssl_protocol $body_bytes_sent  $response_processing_time $http_host_forwarded  
  $tls_version  $cipher_suite $edge_result_type $request_protocol $fle_status $fle_encrypted_fields 
  $http_port  $time_first_byte  $edge_detail_result_type  
  $http_content_type  $request_length $request_length_start $request_length_end
```

### Threat rules

#### Excludes

We include resources for predetermined threats, including:
- Common Web Attack
- Bad IP Address
- Bad Referrer
- Bad Crawler
- Directory Bruteforce

You can disable any type of threat in the `excludes` configuration _(case-sensitive)_.

```yaml
rules:
  threat:
    excludes:
      - "Bad IP Address"
```

The above format detects threats that are not included as bad IP address, and will not analyze logs/ send alerts for that type.

#### Whitelists

You can also add whitelists to teler configuration.

```yaml
rules:
  threat:
    whitelists:
      - "(curl|Go-http-client|okhttp)/*"
      - "^/wp-login\\.php"
```

It covers the entire HTTP request and processed as _regExp_, please write it with caution!

### Notification

We provide alert notification options:
- Slack,
- Telegram _(roadmap)_, or
- Discord _(roadmap)_

Configure the notification alerts needed on:

```yaml
notifications:
  slack:
    token: "xoxb-..."
    color: "#ffd21a"
    channel: "G30SPKI"
```

You can also choose to disable alerts or want to be sent where the alerts are.

```yaml
alert:
  active: true
  provider: "slack"
```

## Contributors

[![contributions](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/kitabisa/teler/issues)

This project exists thanks to all the people who contribute. To learn how to setup a development environment and for contribution guidelines, see [CONTRIBUTING.md](https://github.com/kitabisa/teler/blob/development/.github/CONTRIBUTING.md).

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://dw1.io"><img src="https://avatars0.githubusercontent.com/u/25837540?v=4" width="100px;" alt=""/><br /><sub><b>Dwi Siswanto</b></sub></a><br /><a href="https://github.com/kitabisa/teler/commits?author=dwisiswant0" title="Code">💻</a> <a href="https://github.com/kitabisa/teler/commits?author=dwisiswant0" title="Documentation">📖</a> <a href="https://github.com/kitabisa/teler/commits?author=dwisiswant0" title="Tests">⚠️</a> <a href="#ideas-dwisiswant0" title="Ideas, Planning, & Feedback">🤔</a></td>
    <td align="center"><a href="https://projectdiscovery.io/open-source"><img src="https://avatars1.githubusercontent.com/u/50994705?v=4" width="100px;" alt=""/><br /><sub><b>ProjectDiscovery</b></sub></a><br /><a href="#tool-projectdiscovery" title="Tools">🔧</a></td>
    <td align="center"><a href="https://twitter.com/satyrius"><img src="https://avatars2.githubusercontent.com/u/278630?v=4" width="100px;" alt=""/><br /><sub><b>Anton Egorov</b></sub></a><br /><a href="#tool-satyrius" title="Tools">🔧</a></td>
    <td align="center"><a href="https://github.com/0ktavandi"><img src="https://avatars0.githubusercontent.com/u/26356781?v=4" width="100px;" alt=""/><br /><sub><b>0ktavandi</b></sub></a><br /><a href="#ideas-0ktavandi" title="Ideas, Planning, & Feedback">🤔</a></td>
    <td align="center"><a href="http:///instagram.com/fikcompany"><img src="https://avatars3.githubusercontent.com/u/73404079?v=4" width="100px;" alt=""/><br /><sub><b>Fik</b></sub></a><br /><a href="#design-fikridhiyau" title="Design">🎨</a></td>
    <td align="center"><a href="https://github.com/fairyhunter13"><img src="https://avatars3.githubusercontent.com/u/12372147?v=4" width="100px;" alt=""/><br /><sub><b>fairyhunter13</b></sub></a><br /><a href="https://github.com/kitabisa/teler/commits?author=fairyhunter13" title="Tests">⚠️</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

### Resources

All external resources used in this teler are **NOT** provided by us. See all peoples who involved in this resources at [teler Resource Collections](https://github.com/kitabisa/teler-resources).

## Pronunciation

/télér/ bagaimana bisa seorang pemuda itu teler hanya dengan meminum 1 sloki ciu _(?)_

## Changes

For changes, see the [CHANGELOG.md](https://github.com/kitabisa/teler/blob/development/CHANGELOG.md).

## License

This program is free software: you can redistribute it and/or modify it under the terms of the [Apache license](https://github.com/kitabisa/teler/blob/development/LICENSE). Kitabisa teler and any contributions are Copyright © by Dwi Siswanto 2020.

[![Stargazers over time](https://starchart.cc/kitabisa/teler.svg)](https://starchart.cc/kitabisa/teler)