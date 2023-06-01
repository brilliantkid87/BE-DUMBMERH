package countrydto

type CountryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type DeleteCountryResponse struct {
	ID int `json:"id"`
}
