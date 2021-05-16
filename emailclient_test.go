package main

import "testing"

func TestConnection(t *testing.T) {
	client, err := NewImapClient("imap.yandex.com:993")
	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	client.SendLoginCommand()
}
