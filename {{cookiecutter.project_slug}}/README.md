# {{ cookiecutter.readme_title }}

[![CircleCI](https://circleci.com/gh/{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}/tree/main.svg?style=svg)](https://circleci.com/gh/{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}/tree/main)

> {{ cookiecutter.readme_short_description }}

## Prerequisites
- TODO

## Installation via go

### Clone git repo
```bash
git clone git@github.com:{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}.git
```

### Open project directory
```bash
cd {{cookiecutter.project_slug}}
```

### Install via go
```bash
go install
```

## Installation via Homebrew (For Mac / Linux)

### Get the formula
```bash
brew tap {{cookiecutter.github_username}}/{{cookiecutter.project_slug}} https://github.com/{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}
```

### Install formula
```bash
brew install {{cookiecutter.project_slug}}
```

## Usage

### Run
```bash
{{cookiecutter.project_slug}}
```

## License and Author

Author: {{ cookiecutter.author }}

License: {{ cookiecutter.license }}
