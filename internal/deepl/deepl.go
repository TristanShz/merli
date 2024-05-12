package deepl

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	FREE_MODE = "free"
	PRO_MODE  = "pro"
	FREE_URL  = "https://api-free.deepl.com/v2/translate"
	PRO_URL   = "https://api.deepl.com/v2/translate"
)

type Translator struct {
	authKey string
	mode    string
}

func NewTranslator(authKey string, mode string) *Translator {
	if mode != FREE_MODE && mode != PRO_MODE {
		return nil
	}

	return &Translator{authKey: authKey, mode: mode}
}

func (t *Translator) getURL() string {
	if t.mode == FREE_MODE {
		return FREE_URL
	}

	return PRO_URL
}

func (t *Translator) createBody(text []string, targetLang string) *bytes.Buffer {
	bodyData := struct {
		Text       []string `json:"text"`
		TargetLang string   `json:"target_lang"`
	}{
		Text:       text,
		TargetLang: targetLang,
	}

	bodyJSON, _ := json.Marshal(bodyData)
	return bytes.NewBuffer(bodyJSON)
}

type Response struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

func (t *Translator) Translate(text []string, targetLang string) (Response, error) {
	c := http.Client{Timeout: time.Duration(5 * time.Second)}

	req, err := http.NewRequest("POST", t.getURL(), t.createBody(text, targetLang))
	if err != nil {
		return Response{}, err
	}

	req.Header.Add("Authorization", "DeepL-Auth-Key "+t.authKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	rawResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r Response
	if err := json.Unmarshal(rawResponse, &r); err != nil {
		return Response{}, err
	}

	return r, nil
}
