package models

//RecipeDBO model
type RecipeDBO struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
}
