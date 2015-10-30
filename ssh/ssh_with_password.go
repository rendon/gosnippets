package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {
	fmt.Printf("Enter host name: ")
	var host string
	if _, err := fmt.Scanf("%s", &host); err != nil {
		log.Fatalf("Failed to read host name: %s", err)
	}

	fmt.Printf("Enter your user: ")
	var user string
	if _, err := fmt.Scanf("%s", &user); err != nil {
		log.Fatalf("Failed to read user: %s", err)
	}

	fmt.Printf("Enter your password: ")
	var pass string
	if _, err := fmt.Scanf("%s", &pass); err != nil {
		log.Fatalf("Failed to read password: ", err)
	}

	var config = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		log.Fatalf("Failed to connect to server: %s", err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("ls"); err != nil {
		log.Fatalf("Failed to run: %s", err)
	}
	fmt.Printf(">> %s", b.String())
}
