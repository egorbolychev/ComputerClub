package models

type Client struct {
	IsInside bool
	TableNum int
}

func NewClient() *Client {
	return &Client{}
}

// func (с *Client) SatDown() {
// 	return nil
// }
