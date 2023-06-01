package countrydto

type CreateCountryRequest struct {
	Name string `json:"name" forsm:"name" validate:"required"`
}

type UpdateCountryRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}
