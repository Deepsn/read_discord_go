package main

import "encoding/json"

func ParseMessages(messagesBytes []byte) ([]Message, error) {
	var msgs []Message

	err := json.Unmarshal(messagesBytes, &msgs)

	if err != nil && len(err.Error()) != 0 {
		return nil, err
	}

	return msgs, nil
}
