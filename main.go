package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := flag.String("input", "package.zip", "a package.zip file")

	flag.Parse()

	if _, err := os.Stat(*input); os.IsNotExist(err) {
		println("Error: FIle \"" + *input + "\" doesn't exist")
		return
	}

	zipReader, err := zip.OpenReader(*input)
	if err != nil {
		panic(err)
	}

	defer zipReader.Close()

	messages := map[int][]Message{}
	channels := map[int]Channel{}

	for _, file := range zipReader.File {
		fileHandle, err := file.Open()
		if err != nil {
			continue
		}

		defer fileHandle.Close()

		if file.FileInfo().IsDir() {
			continue
		}

		data := strings.Split(file.Name, "/")
		content, err := io.ReadAll(fileHandle)

		if err != nil || len(content) == 0 {
			continue
		}

		dataLength := len(data)
		collection := data[0]

		var id string
		var dataType string

		if dataLength > 2 {
			id = data[1]
			dataType = strings.Replace(data[2], ".json", "", 1)
		} else if dataLength == 2 {
			dataType = strings.Replace(data[1], ".json", "", 1)
		}

		switch {
		case collection == "messages" && dataType == "channel":
			{
				parsedChannel, err := ParseChannel(content)

				if err != nil {
					continue
				}

				if parsedChannel.Name == "" {
					parsedChannel.Name = "unknown-channel-" + strconv.Itoa(parsedChannel.Id)
				}

				channels[parsedChannel.Id] = parsedChannel
			}
		case collection == "messages" && dataType == "messages":
			{
				channelId, err := strconv.Atoi(strings.Replace(id, "c", "", 1))

				if err != nil {
					println("Failed to parse channel id", id)
					continue
				}

				parsedMessages, err := ParseMessages(content)

				if err != nil {
					continue
				}

				for _, message := range parsedMessages {
					if _, ok := messages[channelId]; !ok {
						messages[channelId] = []Message{}
					}

					messages[channelId] = append(messages[channelId], message)
				}
			}
		}

	}

	for channelId, msgs := range messages {
		channel, ok := channels[channelId]

		if !ok {
			continue
		}

		sort.Slice(msgs, func(i, j int) bool {
			return msgs[i].Id < msgs[j].Id
		})

		channel.Messages = append(channel.Messages, msgs...)
		channels[channelId] = channel
	}

	if err := CreateFolder("out"); err != nil {
		return
	}

	os.Chdir("out")

	for _, channel := range channels {
		fileName := SanitizeFileName(channel.Name)
		file, err := os.Create(fileName + ".txt")

		if err != nil {
			continue
		}

		defer file.Close()

		file.WriteString(fmt.Sprintf("Channel: %s\nChannel Id: %d\n\nMessages:\n", channel.Name, channel.Id))

		for _, msg := range channel.Messages {
			file.WriteString(fmt.Sprintf("[%d] %s\n", msg.Id, msg.Contents))
		}

		file.Sync()
	}

	os.Chdir("..")

	println("Channels", len(channels))
	println("Message channels", len(messages))
}
