package main

type Message struct {
	Id          int
	GuildId     int
	Contents    string
	Timestamp   string
	Attachments string
}

type Channel struct {
	Id       int `json:"id,string"`
	Name     string
	Type     string
	Messages []Message
	Guild    Guild
}

type Guild struct {
	Id   int `json:"id,string"`
	Name string
}
