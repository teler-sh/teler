# Changelog

All notable changes to this project should be documented in this file.

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