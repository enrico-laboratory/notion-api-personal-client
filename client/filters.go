package client

type Or struct {
	Property string `json:"property"`
	Select   struct {
		Equals string `json:"equals"`
	} `json:"select"`
}

type FilterOrSelect struct {
	Filter struct {
		Or []Or `json:"or"`
	} `json:"filter"`
}
