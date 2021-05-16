package main

import (
	"bufio"
	"fmt"
	"strings"
)

type CommandHandler struct {
	client  *ImapClient
	reader  *bufio.Reader
	counter int
}

func (client *ImapClient) SendLoginCommand() {
	command := `00000001 LOGIN "" ""` + "\r\n"
	fmt.Fprintf(client.connection, command)

	reader := bufio.NewReader(client.connection)
	for {
		message, err := reader.ReadString('\r')
		message = strings.Trim(message, "\r\n")
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		if strings.HasPrefix(message, "00000001") {
			tokens := strings.Fields(message)
			if tokens[1] == "OK" {
				fmt.Println("Login successful!")
			} else {
				fmt.Println("Logon error occured")
			}
			break
		}
	}
}

func (client *ImapClient) SendListCommand() {
	command := `00000002 LIST "" "*"` + "\r\n"
	fmt.Fprintf(client.connection, command)

	reader := bufio.NewReader(client.connection)
	for {
		message, err := reader.ReadString('\r')
		message = strings.Trim(message, "\r\n")
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		if strings.HasPrefix(message, "* LIST") {
			tokens := strings.Fields(message)
			fmt.Println(strings.Trim(tokens[len(tokens)-1], `"`))
		}

		if strings.HasPrefix(message, "00000002") {
			fmt.Println("Fetched all mailboxes")
			break
		}
	}
}

func (client *ImapClient) SelectBoxCommand() {
	command := `00000003 SELECT "INBOX"` + "\r\n"
	fmt.Fprintf(client.connection, command)

	reader := bufio.NewReader(client.connection)
	for {
		message, err := reader.ReadString('\r')
		message = strings.Trim(message, "\r\n")
		fmt.Println(message)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		if strings.HasPrefix(message, "00000003") {
			fmt.Println("SELECTED INBOX")
			break
		}
	}
}

func (client *ImapClient) FetchMailCommand() {
	command := `00000004 fetch 1:4 (BODY[])` + "\r\n"
	fmt.Fprintf(client.connection, command)

	reader := bufio.NewReader(client.connection)
	for {
		message, err := reader.ReadString('\r')
		message = strings.Trim(message, "\r\n")
		fmt.Println(message)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		if strings.HasPrefix(message, "00000004") {
			fmt.Println("Fetched all mails")
			break
		}
	}
}
