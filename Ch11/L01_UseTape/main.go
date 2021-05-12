package main

import (
	"fmt"
	"github.com/TimLeary/headFirstGo/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}

	device.Stop()
}

func TryOut(player Player, mixtape []string)  {
	playList(player, mixtape)
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}

func main()  {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	TryOut(player, mixtape)

	player = gadget.TapeRecorder{}
	TryOut(player, mixtape)
}
