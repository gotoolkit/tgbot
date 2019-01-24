package tgbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

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

func (b *Bot) sendMessage(id string, text string) (*Message, error) {
	params := map[string]string{
		"chat_id": id,
		"text":    text,
	}
	messageBts, err := b.Request("sendMessage", params)
	if err != nil {
		return nil, err
	}
	var messageInfo struct {
		Ok          bool
		Result      *Message
		Description string
	}
	err = json.Unmarshal(messageBts, &messageInfo)
	if err != nil {
		return nil, err
	}
	if !messageInfo.Ok {
		return nil, errors.New(fmt.Sprintf("api err: %s", messageInfo.Description))
	}

	return messageInfo.Result, nil
}

func (b *Bot) deleteMessage(id string, messageID int) error {
	params := map[string]string{
		"chat_id":    id,
		"message_id": fmt.Sprintf("%d", messageID),
	}
	messageBts, err := b.Request("deleteMessage", params)
	if err != nil {
		return err
	}
	var messageInfo struct {
		Ok          bool
		Description string
	}
	err = json.Unmarshal(messageBts, &messageInfo)
	if err != nil {
		return err
	}

	if !messageInfo.Ok {
		return errors.New(fmt.Sprintf("api err: %s", messageInfo.Description))
	}

	return nil
}

func (b *Bot) getUpdates(offset int, timeout time.Duration) ([]Update, error) {
	params := map[string]string{
		"offset":  strconv.Itoa(offset),
		"timeout": strconv.Itoa(int(timeout / time.Second)),
	}
	updateBts, err := b.Request("getUpdates", params)
	if err != nil {
		return nil, err
	}
	var updatesInfo struct {
		Ok          bool
		Result      []Update
		Description string
	}

	err = json.Unmarshal(updateBts, &updatesInfo)
	if err != nil {
		return nil, err
	}
	if !updatesInfo.Ok {
		return nil, errors.New(fmt.Sprintf("api err: %s", updatesInfo.Description))
	}

	return updatesInfo.Result, nil
}

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
