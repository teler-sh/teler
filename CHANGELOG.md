# Changelog

All notable changes to this project should be documented in this file.

### v1.1.0

- Upgrade dependencies.

### v1.0.3

- Remove unused package

### v1.0.2

- Refactor versioning package

### v1.0.2-dev

- Inject build version (#95)
- Update `Makefile` script
- Update documentations

### v1.0.1

- Improves whitelist for Common Web Attack & CVE threats
- Matching status code & request method to reduce false-positives CVE (#72)

### v1.0.0

- Fix mismatch breaking logic for detecting CVE (#71)

### v1.0.0-rc

- Add CVE resource
- Add metrics for CVEs
- Fix Bad Referrer

### v0.0.5-dev

- Justifying informations

### v0.0.4

- Fix the Slack token validation not accepting some tokens (#57)

### v0.0.4-dev

- Add spinner while getting resources

### v0.0.3

- Fix the _wrong catch_ for `3xx` status codes to be ignored in Directory Bruteforce (#56)

### v0.0.2

- Fix threat metrics

### v0.0.1

- Add Discord notification alert

### v0.0.1-dev5.1

- Add prints JSON format options
- Update Dockerfile
- Fix SIGSEGV of `syscall` within Docker container

### v0.0.1-rc2.1

- Add remove cached resources flag
- Add purge caches function
- Fix fails to get cached resources
- Add cache option in configuration file
- Remove unused files

### v0.0.1-rc2

- Add caching supports for resources (#45)

### v0.0.1-dev5

- Refactor teler configurations
- Update Prometheus package & runner
- Add Exporters to Prometheus (#42)

### v0.0.1-dev.4.3

- Add errors.Abort for supporting cross-platform

### v0.0.1-rc1.3

- Update Dockerfile

### v0.0.1-beta4

- Refactoring configurations
- Add Telegram notification alert (#38)
- Add partially unit test (#34)

### v0.0.1-beta3.2

- Delete element in Slack alert
- Move Slack alert parts

### v0.0.1-dev.4.2

- Add remote IP addr part on Slack alert

### v0.0.1-beta3.1

- Update documentations

### v0.0.1-dev.4.1

- Removes default attachments alert
- Fix ignored whitelists on query parameters (#28)

### v0.0.1-beta2

- Remove trailing newlines
- Convert output threats to JSON
- Merging `out` mapping into `log`

### v0.0.1-dev1

- Fix invalid URL escape by adding error handle
- Fix justifies common web attack category

### v0.0.1-rc1.2

- Fix mismatch pattern of referrer
- Add whitelist for logs (#16)

### v0.0.1-rc1

- Fix redeclared & unused functions and variables
- Remove `active` part in example config file
- Add `whitelists` documentations for configuration
- Add `whitelists` in example config file
- Replace configs to resource
- Refactor `teler.Analyze`
- Add trailing newlines in analyzer

### v0.0.1-beta1.2

- Fix Common Web Attack threat

### v0.0.1-beta1

- Initial beta release