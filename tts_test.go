package tts_test

import (
	"github.com/ErfanMomeniii/tts"
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
	// english US text
	err := tts.SaveToFile(sampleEnglishUsText, tts.EnglishUs, "./test")
	assert.NoError(t, err)

	// english UK text
	err = tts.SaveToFile(sampleEnglishUkText, tts.EnglishUk, "./test")
	assert.NoError(t, err)

	// arabic text
	err = tts.SaveToFile(sampleArabicText, tts.Arabic, "./test")
	assert.NoError(t, err)

	_ = os.RemoveAll("./test")
}
