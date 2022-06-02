package main

import "music-player-backend/route"

func main() {
	r := route.Router()
	r.Run(":8080")
}
