package model

type Disease struct {
	ID         int             `db:"id"`
	Name       string          `db:"name"`
	CategoryID int             `db:"category_id"`
	Category   DiseaseCategory `db:""`
}

type DiseaseCategory struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
