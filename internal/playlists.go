package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Playlist struct {
	Id string
}

type Owner struct {
	Id string `json:"id"`
}

type Tracks struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type SimplePlaylist struct {
	Description string `json:"description"`
	Owner       Owner  `json:"owner"`
	Name        string `json:"name"`
	SnapshotId  string `json:"snapshot_id"`
	Tracks      Tracks `json:"tracks"`
}

type PlaylistReq struct {
}

type PlaylistRes struct {
	Href  string           `json:"href"`
	Limit int              `json:"limit"`
	Next  string           `json:"next"`
	Total int              `json:"total"`
	Items []SimplePlaylist `json:"items"`
}

var GET_PLAYLISTS_URL = "https://api.spotify.com/v1/users/wijohnst/playlists"

func GetPlaylists(auth *Auth) []Playlist {
	token := auth.Token.AccessToken

	r, err := http.NewRequest("GET", GET_PLAYLISTS_URL, nil)

	if err != nil {
		panic(err)
	}

	r.Header.Add("Authorization", "Bearer "+token)

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

	var b PlaylistRes
	err = json.Unmarshal(rBody, &b)

	if err != nil {
		panic(err)
	}

	for _, i := range b.Items {
		fmt.Print(i.Name + "\n")
	}

	return []Playlist{}
}
