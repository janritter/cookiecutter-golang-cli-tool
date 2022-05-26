# Cookiecutter Golang CLI Tool

> Cookiecutter template to create the skeleton for a Golang CLI tool with test, build and release CircleCI pipeline.

The CircleCI pipeline includes:
- Testing with CodeClimate integration
- GitHub release with semantic-release
- Compiling of binaries for multiple operating systems with upload to the GitHub release and signing with cosign
- Creation and update of a Homebrew formula.

## Usage
```
pip3 install cookiecutter
cookiecutter gh:janritter/cookiecutter-golang-cli-tool
```

## Config in the newly created repo

### Cosign

1. Generate new cosign key pair with `cosign generate-key-pair`
2. Commit the `cosign.pub` file to the newly created repo
3. Set the password you used in the cosign generation as an environment variable named `COSIGN_PWD` in CircleCI
4. Set the base64 of the cosign key as an environment variable named `COSIGN_KEY` in CircleCI, you can generate the base64 with the following command
    ```bash
    cat cosign.key | base64 -w 0
    ```

### Additional CircleCI env vars

The CircleCI job is requiring the following environment variables next to the ones configured in the previous step:
- `CC_TEST_REPORTER_ID` - Required to upload the coverage results to CodeClimate (CodeClimate Repo settings -> Test Coverage)
- `GITHUB_TOKEN` - Required to create GitHub releases and to upload the binaries and Homebrew formula (must have repo permissions, write:packages and delete:packages)

## License and Author

Author: Jan Ritter

License: MIT
