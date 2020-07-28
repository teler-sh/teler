teler
========

<img src="https://user-images.githubusercontent.com/25837540/88009716-594ec700-cb3d-11ea-9863-0e44c2e9ea7c.jpg" height="150">

## Resources
<!-- TODO -->
- [Resources](#resources)
- [What is it?](#what-is-it)
- [Features](#features)
  - [Nearly all web log formats...](#nearly-all-web-log-formats)
- [Why teler?](#why-teler)
- [Usage](#usage)
  - [Flags](#flags)
	  - [Config](#config)
	  - [Input](#input)
	  - [Concurrency](#concurrency)
- [Pronunciation](#pronunciation)

## What is it?

teler is an **real-time threat alert** and huting based on web log that runs in a **terminal** on &ast;nix systems. teler is working like an [Intrusion Detection System](https://en.wikipedia.org/wiki/Intrusion_detection_system), with resources that are compact enough to "_dictate_" a threat.

<img src="https://user-images.githubusercontent.com/25837540/88009657-3e7c5280-cb3d-11ea-9b49-e301187ef21f.jpg" height="200">

## Features

* **Completely Real Time**<br>
  Hunts are timed to be processed every 20 line-buffer _(can be configured)_ on the terminal
  input and every second on the output. Wagelaseh!

* **Minimal Configuration needed**<br>
  You can just run it against your log file, write the log format and let
  teler analyze the log and show you the threats!

* **Nearly All Web Log Formats**<br>
  teler allows any custom log format string. Predefined options include,
  Apache, Nginx, Amazon S3, Elastic Load Balancing, CloudFront, etc.

* **Incremental Log Processing**<br>
  Need data persistence rather than [buffer stream](https://linux.die.net/man/1/stdbuf)?
  teler has the ability to process logs incrementally through the on-disk persistence options.

### Nearly all web log formats...
teler allows any custom log format string. Predefined options include, but
not limited to:

* Amazon CloudFront (Download Distribution).
* Amazon Simple Storage Service (S3)
* AWS Elastic Load Balancing
* Combined Log Format (XLF/ELF) Apache | Nginx
* Common Log Format (CLF) Apache
* Google Cloud Storage.
* Apache virtual hosts
* Squid Native Format.
* W3C format (IIS).

It all depends on how you fill the log format in configuration file.

## Why teler?

<img src="https://user-images.githubusercontent.com/25837540/88010437-1130a400-cb3f-11ea-9089-b6a1e2fb1ae5.jpg" height="400">

teler was designed to be a fast, terminal-based threat analyzer. Its core idea
is to quickly analyze and hunt threats in real time without
needing to use your browser (_great if you want to do a quick analysis of your
access log via SSH, or if you simply love working in the terminal_).

## Usage

Simply, teler can be run with:

```bash
▶ [buffer] | teler -c /path/to/config/teler.yaml
# or
▶ teler -i /path/to/access.log -c /path/to/config/teler.yaml
```

### Flags

```bash
▶ teler -h
```

This will display help for the tool.

<img src="https://user-images.githubusercontent.com/25837540/88668885-66efe800-d10d-11ea-95b6-038cc3a82406.png">

Here are all the switches it supports.

| Flag             	| Description                                                 	| Examples                                                	|
|------------------	|-------------------------------------------------------------	|---------------------------------------------------------	|
| -c, --config     	| teler configuration file                                    	| kubectl logs nginx \| teler -c /path/to/config/teler.yaml |
| -i, --input      	| Analyze logs from data persistence rather than buffer stream 	| teler -i /var/log/nginx/access.log                      	|
| -x, --concurrent 	| Set the concurrency level to analyze logs (default: 20)      	| tail -f /var/log/nginx/access.log \| teler -x 50        	|
| -v, --version    	| Show current teler version                                  	| teler -v                                                	|

#### Config

The `-c` flag is to specify teler configuration file. See [teler.yaml.sample](https://github.com/kitabisa/teler/blob/development/teler.yaml.sample) for an example.

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

## Pronunciation

/télér/ bagaimana bisa seorang pemuda itu teler hanya dengan meminum 1 sloki ciu _(?)_

## License

teler is released under MIT. See [LICENSE.md](https://github.com/kitabisa/teler/blob/development/LICENSE).