package main

import (
	"crypto/tls"
	"crypto/x509"
	"net"
)

type ImapClient struct {
	connection net.Conn
}

func NewImapClient(url string) (ImapClient, error) {
	roots, err := x509.SystemCertPool()
	if err != nil {
		panic("failed to parse system CAs")
	}
	conn, err := tls.Dial("tcp", url, &tls.Config{
		RootCAs: roots,
	})
	client := ImapClient{}

	if err != nil {
		return client, err
	}

	client.connection = conn
	return client, nil
}
