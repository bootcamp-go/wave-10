package domain

type Client struct {
	id      int    `json:"id,omitempty"`
	name    string `json:"name,omitempty"`
	address string `json:"address,omitempty"`
}
