# Contributing

By participating to this project, you agree to abide our [code of conduct](/.github/CODE_OF_CONDUCT.md).

## Development

For small things like fixing typos in documentation, you can [make edits through GitHub](https://help.github.com/articles/editing-files-in-another-user-s-repository/), which will handle forking and making a pull request (PR) for you. For anything bigger or more complex, you'll probably want to set up a development environment on your machine, a quick procedure for which is as folows:


### Setup your machine

`teler` is written in [Go](https://golang.org/).

Prerequisites:

- `make`
- [Go 1.18+](https://golang.org/doc/install)

Fork and clone **[teler](https://github.com/kitabisa/teler)** repository.

A good way of making sure everything is all right is running the following:

```bash
▶ make build
▶ ./bin/teler -v
```

### Test your change

When you are satisfied with the changes, we suggest you run:

```bash
▶ make test
```

Which runs all the linters and cross-compability checks.

### Create and/ update configuration for documentations

Add your new or updated configuration to `teler.example.yaml` so they will be shown in the documentations.

### Submit a pull request

As you are ready with your code contribution, push your branch to your `teler` fork and open a pull request against the **master** branch.

Please also update the [CHANGELOG.md](/CHANGELOG.md) to note what you've added or fixed.

### Pull request checks

By submitting PR to this project, you are accept to our [CLA](/.github/CONTRIBUTION_LICENSE_AGREEMENT.md).

Also, we run a few checks in CI by using GitHub actions, you can see them [here](/.github/workflows/pr.yaml).