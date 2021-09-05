prepare:
	go mod download

build: prepare
	go build -o ./bin/{{cookiecutter.project_slug}} -v -ldflags "-X github.com/{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}/cmd.gitSha=`git rev-parse HEAD` -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}/cmd.buildTime=`date +'%Y-%m-%d_%T'` -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.project_slug}}/cmd.version=LOCAL_BUILD"

tests:
	go test ./... -v  --cover

run:
	go run main.go
