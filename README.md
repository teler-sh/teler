<p align="center">
  <a href="#"><img alt="teler" src="https://user-images.githubusercontent.com/25837540/97091757-7200d880-1668-11eb-82c4-e5c4971d2bc8.png" height="300" /></a>
  <h3 align="center"><b>teler</b></h3>
</p>

<p align="center">
  <a href="#"><img alt="Kitabisa Security" src="https://img.shields.io/badge/kitabisa-security%20project-blue" /></a>
  <a href="/LICENSE"><img alt="License" src="https://img.shields.io/badge/License-Apache%202.0-yellowgreen" /></a>
  <a href="http://golang.org"><img alt="made with Go" src="https://img.shields.io/badge/made%20with-Go-brightgreen" /></a>
  <a href="https://github.com/kitabisa/teler/releases"><img alt="Release" src="https://img.shields.io/github/v/release/kitabisa/teler?color=blueviolet" /></a>
  <a href="#"><img alt="Platform" src="https://img.shields.io/badge/platform-osx%2Flinux%2Fwindows-green" /></a>
  <a href="https://github.com/kitabisa/teler/issues"><img alt="GitHub issues" src="https://img.shields.io/github/issues/kitabisa/teler" /></a>
  <h3 align="center"><b>teler</b></h3>
</p>

<!-- [![Gitter](https://badges.gitter.im/kitabisa-teler/community.svg)](https://gitter.im/kitabisa-teler/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) -->
<!-- [![Go report](https://goreportcard.com/badge/teler.app)](https://goreportcard.com/report/teler.app) -->

<p align="center">
  Real-time HTTP Intrusion Detection
  <br />
  <a href="/.github/CONTRIBUTING.md">Contribute</a>
  ·
  <a href="/CHANGELOG.md">What's new</a>
  ·
  <a href="https://github.com/kitabisa/teler/issues/new/choose">Report Bug</a>
  ·
  <a href="https://github.com/kitabisa/teler/issues/new/choose">Request Feature</a>
</p>

---

`teler` is an **real-time intrusion detection** and threat alert based on web log that runs in a **terminal** with resources that we collect and provide by the community. :heart:

| **CLI**  | **Dashboard**  |
|--------- |--------------- |
| [![teler](https://user-images.githubusercontent.com/25837540/97096468-f8ccaa00-1696-11eb-8830-0d3a7be45a2d.gif)](#) | [![dashboard](https://user-images.githubusercontent.com/25837540/175797412-1921c0e8-c4dc-4e2f-a29d-1c0208a86d22.gif)](#) |

> **Note**:
> If you upgrade from prior to v2 frontwards there will be some **break changes** that affect configuration files. 
> Appropriate adaptations can refer to [teler.example.yaml](https://github.com/kitabisa/teler/blob/v2/teler.example.yaml) file.

## Table of Contents
- [Features](#features)
- [Why teler?](#why-teler)
- [Demo](#demo)
- [Documentation](#documentation)
- [Supporting Materials](#supporting-materials)
- [Contributors](#contributors)
  - [Resources](#resources)
- [Pronunciation](#pronunciation)
- [Changes](#changes)
- [License](#license)

## Features

* **Real-time**: Analyze logs and identify suspicious activity in real-time.

* **Alerting**: teler provides alerting when a threat is detected, push notifications include Slack, Mattermost, Telegram and Discord.

* **Monitoring**: We've our own metrics if you want to monitor threats easily, and we use Prometheus for that.

* **Logging**: is also provided in file form or sends detected threats to the Zinc logs search engine.

* **Latest resources**: Collections is continuously up-to-date.

* **Minimal configuration**: You can just run it against your log file, write the log format and let
  teler analyze the log and show you alerts!

* **Flexible log formats**: teler allows any custom log format string! It all depends on how you write the log format in configuration file.

* **Custom threat rules**: Want to reach a wider range of threats instead of engine-based _(default)_ rules? You can customize threat rules!

* **Incremental log processing**: Need data persistence rather than [buffer stream](https://linux.die.net/man/1/stdbuf)?
  teler has the ability to process logs incrementally through the on-disk persistence options.

## Why teler?

teler was designed to be a fast, terminal-based threat analyzer. Its core idea is to quickly analyze and hunt threats in real time!

## Demo

Here is a preview of `teler` with conditions of use as:

| **Buffer-streams**  | **Incremental**   |
|-------------------- |-----------------  |
| <a href="https://asciinema.org/a/367616" alt="teler"><img src="https://asciinema.org/a/367616.svg"></a> | <a href="https://asciinema.org/a/367610" alt="teler"><img src="https://asciinema.org/a/367610.svg"></a> |

## Documentation

All related documentation about installation, usage & configuration is on **[teler.app](https://teler.app)**.

## Supporting Materials

- [teler - Protect Your WebApp!](https://dw1.io/files/teler%20-%20Protect%20Your%20WebApp.pdf) Talks were brought to the **OWASP Jakarta: Virtual AppSec Indonesia 2020** event.
- [Tutorial: Cyber Threat Hunting - Useful Threat Hunting Tools (Part One)](https://youtu.be/0m54WOXO6Gc), Semi Yulianto gave a brief explanation and how to use **teler** in the video.
- [Empowering Teler HTTP Intrusion Detection as WAF with Fail2ban](https://link.medium.com/OXVZIMkZEeb).

## Contributors

[![contributions](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/kitabisa/teler/issues)

This project exists thanks to all the people who contribute. To learn how to setup a development environment and for contribution guidelines, see [CONTRIBUTING.md](/.github/CONTRIBUTING.md).

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://dw1.io"><img src="https://avatars0.githubusercontent.com/u/25837540?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Dwi Siswanto</b></sub></a><br /><a href="https://github.com/kitabisa/teler/commits?author=dwisiswant0" title="Code">💻</a> <a href="https://github.com/kitabisa/teler/commits?author=dwisiswant0" title="Documentation">📖</a> <a href="https://github.com/kitabisa/teler/commits?author=dwisiswant0" title="Tests">⚠️</a> <a href="#ideas-dwisiswant0" title="Ideas, Planning, & Feedback">🤔</a></td>
    <td align="center"><a href="https://projectdiscovery.io/open-source"><img src="https://avatars1.githubusercontent.com/u/50994705?v=4?s=100" width="100px;" alt=""/><br /><sub><b>ProjectDiscovery</b></sub></a><br /><a href="#tool-projectdiscovery" title="Tools">🔧</a></td>
    <td align="center"><a href="https://twitter.com/satyrius"><img src="https://avatars2.githubusercontent.com/u/278630?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Anton Egorov</b></sub></a><br /><a href="#tool-satyrius" title="Tools">🔧</a></td>
    <td align="center"><a href="https://github.com/0ktavandi"><img src="https://avatars0.githubusercontent.com/u/26356781?v=4?s=100" width="100px;" alt=""/><br /><sub><b>0ktavandi</b></sub></a><br /><a href="#ideas-0ktavandi" title="Ideas, Planning, & Feedback">🤔</a></td>
    <td align="center"><a href="http:///instagram.com/fikcompany"><img src="https://avatars3.githubusercontent.com/u/73404079?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Fik</b></sub></a><br /><a href="#design-fikridhiyau" title="Design">🎨</a></td>
    <td align="center"><a href="https://github.com/fairyhunter13"><img src="https://avatars3.githubusercontent.com/u/12372147?v=4?s=100" width="100px;" alt=""/><br /><sub><b>fairyhunter13</b></sub></a><br /><a href="https://github.com/kitabisa/teler/commits?author=fairyhunter13" title="Tests">⚠️</a></td>
    <td align="center"><a href="http://zufardhiyaulhaq.com"><img src="https://avatars3.githubusercontent.com/u/11990726?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Zufar Dhiyaulhaq</b></sub></a><br /><a href="https://github.com/kitabisa/teler/commits?author=zufardhiyaulhaq" title="Code">💻</a></td>
  </tr>
  <tr>
    <td align="center"><a href="https://github.com/JustHumanz"><img src="https://avatars3.githubusercontent.com/u/43176061?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Aldin Setiawan</b></sub></a><br /><a href="https://github.com/kitabisa/teler/commits?author=JustHumanz" title="Code">💻</a> <a href="#a11y-JustHumanz" title="Accessibility">️️️️♿️</a></td>
    <td align="center"><a href="https://www.kirsle.net/"><img src="https://avatars2.githubusercontent.com/u/1663507?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Noah Petherbridge</b></sub></a><br /><a href="#tool-kirsle" title="Tools">🔧</a></td>
    <td align="center"><a href="https://github.com/zackijack"><img src="https://avatars3.githubusercontent.com/u/1515471?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Zackky Muhammad</b></sub></a><br /><a href="#infra-zackijack" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a></td>
    <td align="center"><a href="https://github.com/acarl005"><img src="https://avatars0.githubusercontent.com/u/8334252?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Andy</b></sub></a><br /><a href="#tool-acarl005" title="Tools">🔧</a></td>
    <td align="center"><a href="https://victoriametrics.com"><img src="https://avatars0.githubusercontent.com/u/283442?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Aliaksandr Valialkin</b></sub></a><br /><a href="#tool-valyala" title="Tools">🔧</a></td>
    <td align="center"><a href="https://ma.rkus.io"><img src="https://avatars2.githubusercontent.com/u/1903284?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Markus Tenghamn</b></sub></a><br /><a href="https://github.com/kitabisa/teler/issues?q=author%3Amarkustenghamn" title="Bug reports">🐛</a></td>
    <td align="center"><a href="https://github.com/brownchow"><img src="https://avatars0.githubusercontent.com/u/8622915?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Rick</b></sub></a><br /><a href="#maintenance-brownchow" title="Maintenance">🚧</a> <a href="https://github.com/kitabisa/teler/commits?author=brownchow" title="Code">💻</a></td>
  </tr>
  <tr>
    <td align="center"><a href="http://michael.bouvy.net/blog/"><img src="https://avatars.githubusercontent.com/u/1674029?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Michael BOUVY</b></sub></a><br /><a href="https://github.com/kitabisa/teler/commits?author=michael-bouvy" title="Documentation">📖</a></td>
    <td align="center"><a href="https://github.com/ossie-git"><img src="https://avatars.githubusercontent.com/u/25382296?v=4?s=100" width="100px;" alt=""/><br /><sub><b>oelnaggar</b></sub></a><br /><a href="https://github.com/kitabisa/teler/commits?author=ossie-git" title="Documentation">📖</a> <a href="https://github.com/kitabisa/teler/issues?q=author%3Aossie-git" title="Bug reports">🐛</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

### Resources

All external resources used in this teler are **NOT** provided by us. See all peoples who involved in this resources at [teler Resource Collections](https://github.com/kitabisa/teler-resources).

## Pronunciation

[`jv_id`](https://www.localeplanet.com/java/jv-ID/index.html) • **/télér/** — bagaimana bisa seorang pemuda itu teler hanya dengan meminum sloki ciu _(?)_

## Changes

For changes, see the [CHANGELOG.md](/CHANGELOG.md).

## License

This program is developed and maintained by members of Kitabisa Security Team, and this is not an officially supported Kitabisa product. This program is free software: you can redistribute it and/or modify it under the terms of the [Apache license](/LICENSE). Kitabisa teler and any contributions are copyright © by Dwi Siswanto 2020-2022.

[![Stargazers over time](https://starchart.cc/kitabisa/teler.svg)](https://starchart.cc/kitabisa/teler)
