package entities

type Urls struct {
	ShortCode    string `json:"short_code"`
	FullUrl      string `json:"full_url"`
	Expiry       string `json:"expiry"`
	NumberOfHits string `json:"number_of_hits"`
}
