package main

import (
	"github.com/ErfanMomeniii/tts"
)

func main() {
	_ = tts.Speak("hello my name is Erfan", tts.EnglishUs, true)

	_ = tts.SaveToFile("hello my friend", tts.EnglishUs, "./example", false)
}
