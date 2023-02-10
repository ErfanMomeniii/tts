package freetts

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const Url string = "https://freetts.com"

type FreeTts struct {
	IsMale bool
}

func (f *FreeTts) TextToSpeak(text string, language string) (error, []byte) {
	speaker, err := selectSpeaker(language, f.IsMale)
	if err != nil {
		return err, nil
	}

	u := fmt.Sprintf("%s/Home/PlayAudio?Language=%s&Voice=%s&TextMessage=%s&id=%s&type=1",
		Url, language, speaker, url.QueryEscape(text), strings.Split(speaker, "_")[0])
	fmt.Println(u)
	resp, err := http.Get(u)
	if err != nil {
		return err, nil
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var response struct {
		Message string `json:"msg"`
		Id      string `json:"id"`
		Counts  int    `json:"counts"`
	}

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}

	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return err, nil
	}

	u2 := fmt.Sprintf("%s/audio/%s", Url, response.Id)
	result, err := http.Get(u2)
	if err != nil {
		return err, nil
	}
	defer func() {
		_ = result.Body.Close()
	}()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return err, nil
	}
	return err, body
}

func selectSpeaker(language string, isMale bool) (string, error) {

	switch language {
	case "en-US":
		if isMale {
			return UsMaleVoice, nil
		} else {
			return UsFemaleVoice, nil
		}
	case "en-GB":
		if isMale {
			return UkMaleVoice, nil
		} else {
			return UkFemaleVoice, nil
		}
	case "ar-XA":
		if isMale {
			return ArMaleVoice, nil
		} else {
			return ArFemaleVoice, nil
		}
	}

	return "", errors.New("not found")
}

func New(isMale bool) *FreeTts {
	return &FreeTts{
		IsMale: isMale,
	}
}
