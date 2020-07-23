teler
========

<img src="https://user-images.githubusercontent.com/25837540/88009716-594ec700-cb3d-11ea-9863-0e44c2e9ea7c.jpg" height="150" />

## What is it?

Teler is an **real-time threat alert** and huting based on web log that runs in a **terminal** on &ast;nix systems. Teler is working like an Intrusion Detection System, with resources that are compact enough to "dictate" a threat.

<img src="https://user-images.githubusercontent.com/25837540/88009657-3e7c5280-cb3d-11ea-9b49-e301187ef21f.jpg" height="200" />

## Features _(read: TODO, LMAO)_

<img src="https://user-images.githubusercontent.com/25837540/88281608-33c6e680-cd12-11ea-91d0-ad4a67ee1a3f.gif" height="350" />

* **Completely Real Time**<br>
  Hunts are timed to be processed every 20 line-buffer _(can be configured)_ on the terminal
  input and every second on the output. Wagelaseh!

* **Minimal Configuration needed**<br>
  You can just run it against your access log file, pick the log format and let
  Teler parse the access log and show you the threats.

* **Nearly All Web Log Formats**<br>
  Teler allows any custom log format string. Predefined options include,
  Apache, Nginx, Amazon S3, Elastic Load Balancing, CloudFront, etc.

* **Incremental Log Processing**<br>
  Need data persistence rather than [buffer stream](https://linux.die.net/man/1/stdbuf)?
  Teler has the ability to process logs incrementally through the on-disk persistence options.

### Nearly all web log formats...
Teler allows any custom log format string. Predefined options include, but
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

_Bismillah aja dulu._

## Why Teler?

<img src="https://user-images.githubusercontent.com/25837540/88010437-1130a400-cb3f-11ea-9089-b6a1e2fb1ae5.jpg" height="400" />

Teler was designed to be a fast, terminal-based log analyzer. Its core idea
is to quickly analyze and hunt threats in real time without
needing to use your browser (_great if you want to do a quick analysis of your
access log via SSH, or if you simply love working in the terminal_).

## So on...
Sabar dong, baru juga init.

## References
- https://en.wikipedia.org/wiki/Intrusion_detection_system
