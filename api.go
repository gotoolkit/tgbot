package tgbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (b *Bot) getMe() ([]byte, error) {
	meJson, err := b.SendRequest("getMe", nil)
	if err != nil {
		return nil, err
	}
	return meJson, nil
}

func (b *Bot) SendRequest(method string, payload interface{}) ([]byte, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.Token, method)

	bts, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := b.client.Post(url, "application/json", bytes.NewReader(bts))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	js, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return js, nil

}
