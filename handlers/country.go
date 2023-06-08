package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	countrydto "dumbmerch/dto/country"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
)

type countryHandler struct {
	CountryRepository repositories.CountryRepository
}

func NewCountryHandler(countryRepository repositories.CountryRepository) *countryHandler {
	return &countryHandler{countryRepository}
}

func (h *countryHandler) FindCountry(c echo.Context) error {
	countries, err := h.CountryRepository.FindCountry()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: countries})
}

func (h *countryHandler) GetCountry(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	country, err := h.CountryRepository.GetCountry(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry(country)})

}

// func (h *countryHandler) CreateCountry(c echo.Context) error {
// 	request := new(countrydto.CreateCountryRequest)
// 	if err := c.Bind(&request); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})

// 	}

// 	country := models.Country{
// 		Name: request.Name,
// 	}

// 	data, err := h.CountryRepository.CreateCountry(country)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{
// 		Code: http.StatusInternalServerError,
// 		Data: convertResponseCountry(data),
// 	})
// }

func (h *countryHandler) CreateCountry(c echo.Context) error {
	request := new(countrydto.CreateCountryRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// Get country data from REST Countries API
	countryData, err := h.fetchCountryData(request.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// Create country model
	country := models.Country{
		Name: countryData.Name,
		// Include other fields as needed
	}

	// Save country to repository
	data, err := h.CountryRepository.CreateCountry(country)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponseCountry(data),
	})
}

func (h *countryHandler) fetchCountryData(countryName string) (*CountryAPIResponse, error) {
	// Construct the REST Countries API URL
	url := fmt.Sprintf("https://restcountries.com/v2/name/%s", strings.ReplaceAll(countryName, " ", "%20"))

	// Send GET request to the API
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response JSON
	var countryData []CountryAPIResponse
	err = json.Unmarshal(body, &countryData)
	if err != nil {
		return nil, err
	}

	if len(countryData) == 0 {
		return nil, fmt.Errorf("country not found")

	}

	return &countryData[0], nil
}

type CountryAPIResponse struct {
	Name string `json:"name"`
	// Add other fields as per the API response structure
}

func (h *countryHandler) DeleteCountry(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	country, err := h.CountryRepository.GetCountry(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CountryRepository.DeleteCountry(country, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: converDeleteCountry(data)})

}

func (h *countryHandler) UpdateCountry(c echo.Context) error {
	request := new(countrydto.UpdateCountryRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	country, err := h.CountryRepository.GetCountry(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		country.Name = request.Name
	}

	data, err := h.CountryRepository.UpdateCountry(country)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertUpdateCountry(data)})
}

func convertResponseCountry(u models.Country) countrydto.CountryResponse {
	return countrydto.CountryResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}

func convertUpdateCountry(u models.Country) countrydto.UpdateCountryRequest {
	return countrydto.UpdateCountryRequest{
		Name: u.Name,
	}
}

func converDeleteCountry(u models.Country) countrydto.DeleteCountryResponse {
	return countrydto.DeleteCountryResponse{
		ID: u.ID,
	}
}
