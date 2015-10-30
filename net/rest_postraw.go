package main

import (
	"encoding/json"
	"github.com/rendon/anaconda"
	"log"
	"menteslibres.net/gosexy/rest"
	"time"
)

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Tickets API types
type Credential struct {
	Owner       string `bson:"owner"          json:"owner"`
	Server      string `bson:"server"         json:"server"`
	Tokens      string `bson:"tokens"         json:"tokens"`
	Counter     int    `bson:"counter"        json:"counter"`
	ActivatedAt int64  `bson:"activated_at"   json:"activated_at"`
}

type CredentialRequestData struct {
	Server  string `json:"server"`
	Tickets int    `json:"tickets"`
}

type CredentialResponseData struct {
	Tokens  string `json:"tokens"`
	Tickets int    `json:"tickets"`
}

type CredentialResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    CredentialResponseData `json:"data,omitempty"`
}

type User struct {
	Id          int64         `json:"id"            bson:"id"`
	Status      string        `json:"status"        bson:"status"`
	RetrievedAt time.Duration `json:"retrieved_at"  bson:"retrieved_at"`
	User        anaconda.User `json:"user"          bson:"user"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data,omitempty"`
}

type Followers struct {
	Id          int64         `json:"id"           bson:"id"`
	Status      string        `json:"status"       bson:"status"`
	RetrievedAt time.Duration `json:"retrieved_at" bson:"retrieved_at"`
	Ids         []int64       `json:"ids"          bson:"ids"`
}

type FollowersResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    Followers `json:"data,omitempty"`
}

type Friends struct {
	Id          int64         `json:"id"            bson:"id"`
	Status      string        `json:"status"        bson:"status"`
	RetrievedAt time.Duration `json:"retrieved_at"  bson:"retrieved_at"`
	Ids         []int64       `json:"ids"           bson:"ids"`
}

type FriendsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Friends `json:"data,omitempty"`
}

type Timeline struct {
	Id          int64            `json:"id"            bson:"id"`
	Status      string           `json:"status"        bson:"status"`
	RetrievedAt time.Duration    `json:"retrieved_at"  bson:"retrieved_at"`
	Tweets      []anaconda.Tweet `json:"tweets"        bson:"tweets"`
}

type TimelineResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Timeline `json:"data,omitempty"`
}

func main() {
	client, err := rest.New("http://auth-server:10000/")
	if err != nil {
		panic(err)
	}

	var c = CredentialRequestData{
		Server:  "twitter",
		Tickets: 1,
	}
	req, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	//var resp CredentialResponse
	var rb = []byte{}
	client.Header.Set("Content-type", "application/json")
	err = client.PostRaw(&rb, "/credentials/request", req)
	if err != nil {
		panic(err)
	}

	log.Printf("Request body: %s\n", req)
	log.Printf("Response body: %s\n", rb)
}
