package model

type TrainingData struct {
	ID                      int       `json:"id"`
	States                  []float64 `json:"states"`
	Epsilons                []float64 `json:"epsilons"`
	Result                  Result    `json:"result"`
	RecommendedActressesIDs []int     `json:"recommended_actresses_ids"`
}
type Vector struct {
	ID         int     `json:"id"`
	Val        float64 `json:"val"`
	TrainingID int     `json:"training_id"`
}

type Result struct {
	ID         int `json:"id"`
	Val        int `json:"val"`
	TrainingID int `json:"training_id"`
}

type Vectors []Vector
