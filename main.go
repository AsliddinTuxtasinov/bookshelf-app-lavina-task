package main

import (
	"bookshelf-app/handlers"
	"bookshelf-app/initializers"
)

func init() {
	initializers.ConnectToDB()
}
func main() {
	handlers.Run()
}
