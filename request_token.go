package main

import (
	"net/http"
	"net/url"
	"strings"
)

func requestToken(client_id string, client_secret string) tokenStruct {
	accessToken := tokenStruct{}
	data := url.Values{}
	data.Set("grant-type", "client_credentials")
	data.Set("client_id", client_id)
	data.Set("client_secret", client_secret)

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest()

	return accessToken
}
