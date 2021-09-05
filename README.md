# Cookiecutter Golang CLI Tool

> Cookiecutter template to create the skeleton for a Golang CLI tool with test, build and release CircleCI pipeline.

The CircleCI pipeline includes:
- Testing with CodeClimate integration
- GitHub release with semantic-release
- Compiling of binaries for multiple operating systems (with upload to the GitHub release)
- Creation and update of a Homebrew formula.

## CircleCI env var setup

The CircleCI job is requiring the following environment variables:
- `CC_TEST_REPORTER_ID` - Required to upload the coverage results to CodeClimate (CodeClimate Repo settings -> Test Coverage)
- `GITHUB_TOKEN` - Required to create GitHub releases and to upload the binaries and Homebrew formula (must have repo permissions, write:packages and delete:packages)

## Usage
```
pip3 install cookiecutter
cookiecutter gh:janritter/cookiecutter-golang-cli-tool
```

## License and Author

Author: Jan Ritter

License: MIT
