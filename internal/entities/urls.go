package entities

type Urls struct {
	ShortCode    string `json:"short_code"`
	FullUrl      string `json:"full_url"`
	Expiry       int32  `json:"expiry"`
	NumberOfHits int32  `json:"number_of_hits"`
}
