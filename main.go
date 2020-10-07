package main

// @title FAQ RESTful Service API
// @version 1.0
// @description This is a MVP for Questions and Answers for https://www.example.com/ page.
// @description Install MongoDB to test this service or use Docker executing Compose up in docker-compose.yml in the project folder check the CLI for reference in https://docs.docker.com/compose/reference/up/
// @description If executing this service by VS Code, just click on right button on docker-compose.yml and Compose Up
// @termsOfService https://www.example.com/
// @contact.name example API Support
// @contact.email flavio.costa@ecore.com.br
// @license.name Apache 2.0
// @license.url https://www.example.com/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	a := App{}
	a.Initialize("localhost", "db")
	a.Run(":3000")
}
