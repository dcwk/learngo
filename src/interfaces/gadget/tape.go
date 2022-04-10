package gadget

import "fmt"

type TapeGadget struct {
	Batteriess string
}

func (t TapeGadget) Play(song string) {
	fmt.Println("Playing: ", song)
}

func (t TapeGadget) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphones string
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing: ", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
