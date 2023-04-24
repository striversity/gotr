package model

import "encoding/json"

type Track struct {
	Id     string `json:"id,omitempty"`
	Artist string `json:"artist,omitempty"`
	Title  string `json:"title,omitempty"`
}

type Playlists map[string][]Track

func (t Track) ToJSONBytes() []byte {
	b, _ := json.Marshal(t)
	return b
}
