package main

type Mailbox struct {
	name        string
	hasChildren bool
	isSent      bool
	isSpam      bool
	isTrash     bool
}
