package main

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Site     string `json:"site"`
	Slot     string `json:"slot"`
	Force    bool   `json:"force"`
	Commit   bool   `json:"commit"`
}
