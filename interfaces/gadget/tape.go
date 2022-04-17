package gadget

import "fmt"

type TapePlayer struct {
	Batteriess string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing by player: ", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphones string
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing by recorder: ", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
