package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func url_get(url string) []byte {
	res, _ := http.Get(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}

type Client struct {
	CLIENT_ID     string
	CLIENT_SECRET string
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func main() {
	cre_url := "http://localhost:8000/credentials"
	cred := url_get(cre_url)
	var m Client
	json.Unmarshal(cred, &m)
	fmt.Println(string(cred))
	fmt.Println(m.CLIENT_ID)

	tok_url := "http://localhost:8000/token?grant_type=client_credentials&client_id=" +
		m.CLIENT_ID +
		"&client_secret=" +
		m.CLIENT_SECRET +
		"&scope=all"

	access := url_get(tok_url)
	fmt.Println(string(access))

	s := "{\"access_token\":\"LL9B2D0BNSQ120AKPKKPMQ\",\"expires_in\":7200,\"scope\":\"all\",\"token_type\":\"Bearer\"}"
	var t Token
	json.Unmarshal([]byte(s), &t)
	fmt.Println(s, "access token:", t.AccessToken)

	test_url := "http://localhost:8000/protected?access_token=" + t.AccessToken
	res := url_get(test_url)
	fmt.Println(string(res))
}
