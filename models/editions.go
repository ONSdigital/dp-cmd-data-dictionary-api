package models

// Editions represents the editions response model
type Editions struct {
	Items           []Edition `json:"editions"`
	NumberOfResults int       `json:"number_of_results"`
}

// Edition represents a single edition response model
type Edition struct {
	ID          string       `json:"id"`
	Edition     string       `json:"edition"`
	Label       string       `json:"label"`
	ReleaseDate string       `json:"release_date"`
	Links       EditionLinks `json:"links"`
}

// EditionLinks reprsents the links returned for a specific edition
type EditionLinks struct {
	Self     Link `json:"self"`
	Editions Link `json:"editions"`
	Codes    Link `json:"codes"`
}
