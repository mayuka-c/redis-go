package model

type Student struct {
	Name    string   `json:"name"`
	Age     string   `json:"age"`
	Hobbies []string `json:"hobbies"`
}
