package internal

import (
	"bytes"
	base64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Auth struct {
	c     AConfig
	uri   string
	cType string
	Token Token
}

type AConfig struct {
	Id  string
	Sec string
}

type Token struct {
	TokenRes
	Created time.Time
	Expires time.Time
}

type TokenReq struct {
	GrantType string `json:"grant_type"`
}

type TokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Type        string `json:"token_type"`
}

var CLIENT_ID = "SPOT_BU_CLI_CLIENT_ID"
var SECRET = "SPOT_BU_CLI_SECRET"

/*
Auth.GetConfig() - Gets required environment variables for Spotify authentication
*/
func (auth *Auth) getConfig() {
	auth.c.Id = os.Getenv(CLIENT_ID)
	auth.c.Sec = os.Getenv(SECRET)
	auth.uri = "https://accounts.spotify.com/api/token"
	auth.cType = "application/x-www-form-urlencoded"

	if auth.c.Id == "" {
		panic(`No envvar for Client ID detected. Please define an environment variable called $SPOT_BU_CLI_CLIENT_ID`)
	}

	if auth.c.Sec == "" {
		panic(`No envvar for Client Secret detected. Please define an environment variable called $SPOT_BU_CLI_SECRET`)
	}
}

/*
Auth.GetToken() - Requests
*/
func (auth *Auth) Init() {
	fmt.Print("Fetching credentials from Spotify...")
	auth.getConfig()

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	r, err := http.NewRequest("POST", auth.uri, bytes.NewBufferString(data.Encode()))

	if err != nil {
		panic("Error: problem fetching token.")
	}

	cEnc := base64.StdEncoding.EncodeToString([]byte(auth.c.Id + ":" + auth.c.Sec))
	aVal := "Basic " + cEnc
	r.Header.Add("Content-Type", auth.cType)
	r.Header.Add("Authorization", aVal)

	c := &http.Client{}
	res, err := c.Do(r)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	rBody, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var b TokenRes
	err = json.Unmarshal(rBody, &b)

	if err != nil {
		panic(err)
	}

	auth.Token.AccessToken = b.AccessToken
	auth.Token.ExpiresIn = b.ExpiresIn
	auth.Token.Type = b.Type

	auth.setCreated()
	auth.setExpiry()
}

/*
Auth.SetCreated() - setter function that sets the `Token.Created` field
*/
func (auth *Auth) setCreated() {
	auth.Token.Created = time.Now()
}

/*
Auth.SetExpiry() - setter function that sets the `Token.Expires` field
*/
func (auth *Auth) setExpiry() {
	t := time.Now().Add(time.Duration(auth.Token.ExpiresIn) * time.Second)

	auth.Token.Expires = t
}

/*
Auth.IsTokenExpired() - Checks if the cached token is expired
*/
func (auth *Auth) IsTokenExpired() bool {
	if auth.Token.AccessToken == "" {
		return true
	}

	/**
	Compares if `Token.Expires` is before or after the current time
	0 == Same
	-1 == Not expired
	+1 == Expired
	*/
	tComp := auth.Token.Expires.Compare(time.Now())

	return tComp >= 0
}
