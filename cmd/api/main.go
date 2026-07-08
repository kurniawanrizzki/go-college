package main

import "go-college/internal/app"

// @title			Go-College API
// @version		1.0.0
// @description	RESTful API for managing colleges, courses, and enrollments.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host			localhost:8181
// @schemes		http https
func main() {
	app.Run()
}
