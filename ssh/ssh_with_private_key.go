package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
)

func getPublicKeys(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(key), nil
}

func main() {
	fmt.Printf("Enter your user: ")
	var user string
	if _, err := fmt.Scanf("%s", &user); err != nil {
		log.Fatalf("Failed to read user: %s", err)
	}

	fmt.Printf("Enter host name: ")
	var host string
	if _, err := fmt.Scanf("%s", &host); err != nil {
		log.Fatalf("Failed to read host name: %s", err)
	}

	fmt.Printf("Enter your private key file path: ")
	var file string
	if _, err := fmt.Scanf("%s", &file); err != nil {
		log.Fatalf("Failed to read file name: ", err)
	}

	auth, err := getPublicKeys(file)
	if err != nil {
		log.Fatalf("Failed to get public keys: %s", err)
	}

	var config = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{auth},
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
