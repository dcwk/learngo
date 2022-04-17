package main

import "github.com/dcwk/learngo/src/interfaces/gadget"

type Player interface {
	Play(string)
	Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	playList(player, mixtape)

	recorder := gadget.TapeRecorder{}
	playList(recorder, mixtape)
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}

	player, ok := device.(gadget.TapeRecorder)
	if ok {
		player.Record()
	}

	device.Stop()
}
