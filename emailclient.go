package main

import (
	"fmt"
)

func main() {
	client, err := NewImapClient("imap.yandex.kz:993")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	client.SendLoginCommand()
	client.SendListCommand()
	client.SelectBoxCommand()
	client.FetchMailCommand()
}
