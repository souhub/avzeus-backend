package model

type Actress struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
	Vector    string `json:"vector"`
}

type Actresses []Actress
