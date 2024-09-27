package main

import "encoding/json"

func ParseChannelsMap(channelsMapBytes []byte) ([]Channel, error) {
	channels := []Channel{}
	channelsMap := map[int]string{}

	err := json.Unmarshal(channelsMapBytes, &channelsMap)

	if err != nil && len(err.Error()) != 0 {
		println(err.Error(), len(channelsMapBytes))
		println(string(channelsMapBytes))

		return nil, err
	}

	for id, name := range channelsMap {
		channels = append(channels, Channel{
			Id:   id,
			Name: name,
		})
	}

	return channels, nil
}

func ParseChannel(channelBytes []byte) (channel Channel, err error) {
	err = json.Unmarshal(channelBytes, &channel)

	return channel, err
}
