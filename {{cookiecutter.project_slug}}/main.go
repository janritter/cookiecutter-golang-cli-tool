package main

import "fmt"

func returnGreeting(name string) string {
	return "Hello World from " + name
}

func main() {
	name := "{{cookiecutter.project_slug}}"
	fmt.Println(returnGreeting(name))
}
