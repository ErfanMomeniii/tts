package main

import (
	"fmt"
	"github.com/ErfanMomeniii/tts"
)

func main() {
	_ = tts.Speak("hello my name is Erfan", tts.EnglishUs, true)

	fileName, _ := tts.SaveToFile("hello my friend", tts.EnglishUs, "./example", false)
	fmt.Println(fileName)
}
