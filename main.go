package main

import "go-gin-template/src/routers"

// use godot package to load/read the .env file and
// return the value of the key

func main() {
	routers.HandleRouter()
}
