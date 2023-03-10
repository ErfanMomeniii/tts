package tts

import "github.com/ErfanMomeniii/tts/pkg/freetts"

type TTS interface {
	Play(text string, language string) error
	Save(text string, language string, path string) (string, error)
}

// Speak generates voice from input text in the input language
// with chosen male or female sound
func Speak(text string, language string, isMale ...bool) error {
	m := true

	if len(isMale) > 0 {
		m = isMale[0]
	}

	f := freetts.New(m)

	if err := f.Play(text, language); err != nil {
		return err
	}

	return nil
}

// SaveToFile generates file of speaking in the input path and returns
// generation file exact name
func SaveToFile(text string, language string, path string, isMale ...bool) (string, error) {
	m := true

	if len(isMale) > 0 {
		m = isMale[0]
	}

	f := freetts.New(m)

	name, err := f.Save(text, language, path)
	if err != nil {
		return "", err
	}

	return name, nil
}
