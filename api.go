package tgbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func (b *Bot) getMe() (*User, error) {
	meBts, err := b.Request("getMe", nil)
	if err != nil {
		return nil, err
	}
	var botInfo struct {
		Ok          bool
		Result      *User
		Description string
	}

	err = json.Unmarshal(meBts, &botInfo)
	if err != nil {
		return nil, err
	}
	if !botInfo.Ok {
		return nil, errors.New(fmt.Sprintf("api err: %s", botInfo.Description))
	}
	return botInfo.Result, nil
}

func (b *Bot) Request(method string, payload interface{}) ([]byte, error) {
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
