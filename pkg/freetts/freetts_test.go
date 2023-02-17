package freetts_test

import (
	"github.com/ErfanMomeniii/tts"
	"github.com/ErfanMomeniii/tts/pkg/freetts"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	sampleEnglishUsText = "hello"
	sampleEnglishUkText = "hello"
	sampleArabicText    = "مرحبًا"
)

func TestSave(t *testing.T) {
	f := freetts.New(true)

	// english US text
	err := f.Save(sampleEnglishUsText, tts.EnglishUs, "./test")
	assert.NoError(t, err)

	// english UK text
	err = f.Save(sampleEnglishUkText, tts.EnglishUk, "./test")
	assert.NoError(t, err)

	// arabic text
	err = f.Save(sampleArabicText, tts.Arabic, "./test")
	assert.NoError(t, err)

	_ = os.RemoveAll("./test")
}
