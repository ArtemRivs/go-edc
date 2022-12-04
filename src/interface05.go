package main

import (
	"fmt"
)

type Player interface{
	Play(string)
	Stop()
}

type TapePlayer struct {
	paramTP string
}
type CDPlayer struct {
	paramCD string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing song (TP): ", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stop playing")
}
func (t TapePlayer) Record() {
	fmt.Println("Recording...")
}
func (t CDPlayer) Play(song string) {
	fmt.Println("Playing song (CD): ", song)
}
func (t CDPlayer) Stop() {
	fmt.Println("Stop playing")
}

func runPlayList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	
	mix := []string{"Song 2", "Smile", "Nevermind"}
	device := TapePlayer{}
	runPlayList(device, mix)

	device2 := CDPlayer{}
	runPlayList(device2, mix)

	
	
}
