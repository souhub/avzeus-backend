package model

type TrainingData struct {
	ID           int    `scheme:"id"`
	ActressesIDs string `schema:"actresses_ids"` // gorrila/scheme に必要
	States       string `scheme:"states"`
	Epsilons     string `scheme:"epsions"`
	SelectedID   int    `scheme:"selected_id"`
}
