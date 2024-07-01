package main

import (
	"wijohnst/spot/internal"
)

func main() {
	auth := &internal.Auth{}
	auth.Init()

	internal.GetPlaylists(auth)
}
