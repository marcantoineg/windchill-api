package main

import "gin-api-template/api"

func main() {
	r := api.SetupRouter()

	r.Run(":1313")
}
