package tts

import "github.com/ErfanMomeniii/tts/pkg/freetts"

type TTS interface {
	Play(text string, language string) error
	Save(text string, language string, path string) error
}

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

func SaveToFile(text string, language string, path string, isMale ...bool) error {
	m := true

	if len(isMale) > 0 {
		m = isMale[0]
	}

	f := freetts.New(m)

	if err := f.Save(text, language, path); err != nil {
		return err
	}

	return nil
}
