package main

import "example/crud_api/route"

func main() {
	r := route.SetupRoutes()
	r.Run(":8080")
}
