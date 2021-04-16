package model

import "github.com/dmmlabo/dmm-go-sdk/api"

type Actress struct {
	ID        int                    `json:"id"`
	Name      string                 `json:"name"`
	ImagePath string                 `json:"image_path"`
	ListURL   api.ActressProductList `json:"list_url"`
}

type Actresses []Actress
