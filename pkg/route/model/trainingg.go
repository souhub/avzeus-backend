package model

// gorrila/scheme に必要
type TrainingData struct {
	ID                      int       `json:"id"`
	RecommendedActressesIDs []int     `json:"recommended_actresses_ids"`
	States                  []float32 `json:"states"`
	Epsilons                []float32 `json:"epsilons"`
	SelectedID              int       `json:"selected_id"`
}
