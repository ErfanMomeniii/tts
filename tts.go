package tts

type TTS interface {
	Play(text string, language string) error
	Save(text string, language string, path string) error
}

func Speak(text string, language string, isMale ...bool) error {
	return nil
}

func SaveToFile(text string, language string, path string, isMale ...bool) error {
	return nil
}
