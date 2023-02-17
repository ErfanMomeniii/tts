package freetts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const Url string = "https://freetts.com"

type FreeTts struct {
	IsMale bool
}

func textToSpeak(text string, language string, isMale bool) ([]byte, string, error) {
	speaker, err := selectSpeaker(language, isMale)
	if err != nil {
		return nil, "", err
	}

	u := fmt.Sprintf("%s/Home/PlayAudio?Language=%s&Voice=%s&TextMessage=%s&id=%s&type=1",
		Url, language, speaker, url.QueryEscape(text), strings.Split(speaker, "_")[0])
	fmt.Println(u)
	resp, err := http.Get(u)
	if err != nil {
		return nil, "", err
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
		return nil, "", err
	}

	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, "", err
	}

	u2 := fmt.Sprintf("%s/audio/%s", Url, response.Id)
	result, err := http.Get(u2)
	if err != nil {
		return nil, "", err
	}
	defer func() {
		_ = result.Body.Close()
	}()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, "", err
	}
	return body, response.Id, err
}

func (f *FreeTts) Play(text string, language string) error {
	result, _, err := textToSpeak(text, language, f.IsMale)
	if err != nil {
		return err
	}

	bytesReader := bytes.NewReader(result)
	decodedMp3, err := mp3.NewDecoder(bytesReader)
	if err != nil {
		return err
	}

	numOfChannels := 2
	audioBitDepth := 2

	otoCtx, readyChan, err := oto.NewContext(decodedMp3.SampleRate(), numOfChannels, audioBitDepth)
	if err != nil {
		return err
	}

	<-readyChan

	player := otoCtx.NewPlayer(decodedMp3)
	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	err = player.Close()
	if err != nil {
		return err
	}

	return nil
}

func (f *FreeTts) Save(text string, language string, path string) (string, error) {
	result, id, err := textToSpeak(text, language, f.IsMale)
	if err != nil {
		return "", err
	}

	_ = os.Mkdir(path, 0777)

	if err = os.WriteFile(path+"/"+id, result, 0777); err != nil {
		return "", err
	}

	return id, nil
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
