package model

type Woman struct {
	ID        int    `json:"id"`
	ImagePath string `json:"image_path"`
	Vector    string `json:"vector"`
}

type Wemen []Woman
