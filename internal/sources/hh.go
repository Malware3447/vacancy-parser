package sources

import "vacancy-parser/internal/models"

func NewHh(baseUrl string) *models.ItemParams {
	return &models.ItemParams{
		BaseUrl: baseUrl,
		Query: models.ListParamsQuery{
			Items:      ".vacancy-serp-content .magritte-redesign",
			Link:       "h2 a",
			Title:      "h2 a",
			Salary:     ".magritte-text_typography-label-1-regular___pi3R-_3-0-41",
			Company:    "span[data-qa=vacancy-serp__vacancy-employer-text]",
			Location:   "",
			Experience: "",
		},
	}
}
