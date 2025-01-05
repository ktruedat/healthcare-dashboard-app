package entities

type DiseaseCategory struct {
	ID         int                  `json:"id"`
	Name       string               `json:"name"`
	Diseases   []Disease            `json:"diseases"`
	YearlyData []CategoryYearlyData `json:"yearly_data"`
}

type Disease struct {
	ID         int                 `json:"id"`
	Name       string              `json:"name"`
	CategoryID int                 `json:"category_id"`
	Category   DiseaseCategory     `json:"category"`
	YearlyData []DiseaseYearlyData `json:"yearly_data"`
}

type CategoryYearlyData struct {
	ID            int                     `json:"id"`
	Year          int                     `json:"year"`
	Prevalence    *float64                `json:"prevalence"`
	Incidence     *float64                `json:"incidence"`
	CategoryID    int                     `json:"category_id"`
	Category      DiseaseCategory         `json:"category"`
	QuarterlyData []CategoryQuarterlyData `json:"quarterly_data"`
}

type CategoryQuarterlyData struct {
	ID         int                `json:"id"`
	YearID     int                `json:"year_id"`
	Quarter    string             `json:"quarter"`
	Cases      *int               `json:"cases"`
	YearlyData CategoryYearlyData `json:"yearly_data"`
}

type DiseaseYearlyData struct {
	ID        int     `json:"id"`
	Year      int     `json:"year"`
	DiseaseID int     `json:"disease_id"`
	Disease   Disease `json:"disease"`
	Q1Cases   *int    `json:"q1_cases"`
	Q2Cases   *int    `json:"q2_cases"`
	Q3Cases   *int    `json:"q3_cases"`
	Q4Cases   *int    `json:"q4_cases"`
}
