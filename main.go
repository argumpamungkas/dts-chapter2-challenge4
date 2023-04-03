package main

import (
	"DTS/Chapter-2/chapter2-challenge-sesi-4/repo"
	"DTS/Chapter-2/chapter2-challenge-sesi-4/routers"
)

func main() {
	repo.StartDB()

	routers.StartServer().Run(":8080")
}
