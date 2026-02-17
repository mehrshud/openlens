# Contributing to OpenLens
Thank you for considering contributing to OpenLens, an open-source alternative to proprietary data analytics platforms. We appreciate your help in making OpenLens a better tool for everyone.

## Setup Steps
To get started with contributing to OpenLens, follow these steps:

1. **Fork the repository**: Create a fork of the OpenLens repository on GitHub.
2. **Clone the repository**: Clone the forked repository to your local machine using `git clone`.
3. **Install dependencies**: Install the required dependencies using `go mod tidy` and `npm install`.
4. **Setup your environment**: Configure your environment by setting up the necessary databases and services.

## Branch Naming
We use the following branch naming conventions:

* `feat/` for new features
* `fix/` for bug fixes
* `docs/` for documentation changes

## Conventional Commits
We use Conventional Commits to format our commit messages. Commit messages should be in the following format:
type(scope): brief description

body
Where `type` is one of `feat`, `fix`, `docs`, etc.

## PR Checklist
Before submitting a pull request, make sure to:

* [ ] Run `go fmt` and `go vet` to ensure code is formatted correctly
* [ ] Run `npm run lint` to ensure code is styled correctly
* [ ] Run `go test` and `npm run test` to ensure tests pass
* [ ] Make sure your commit messages follow the Conventional Commits format
* [ ] Ensure your code is formatted according to our code style guidelines (see below)

## Code Style
We follow the following code style guidelines:

* **Go**: Use `go fmt` to format Go code. Follow the standard Go coding conventions.
* **GraphQL**: Use `graphql-tag` to format GraphQL code. Follow the standard GraphQL coding conventions.
* **React**: Use `npm run lint` to format React code. Follow the standard React coding conventions.

## Running Tests
To run tests, use the following commands:

* `go test` for Go tests
* `npm run test` for React tests

## Reporting Bugs
To report a bug, please create an issue on GitHub with the following information:

* **Description**: A clear and concise description of the bug
* **Steps to reproduce**: Step-by-step instructions to reproduce the bug
* **Expected behavior**: What you expect to happen
* **Actual behavior**: What actually happens
* **Version**: The version of OpenLens you are using

## Additional Information
For more information on contributing to OpenLens, please see our [README](README.md) and [CODE_OF_CONDUCT](CODE_OF_CONDUCT.md). Thank you again for considering contributing to OpenLens!