package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	dto "dumbmerch/dto/result"
	tripdto "dumbmerch/dto/trip"
	"dumbmerch/models"
	"dumbmerch/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type tripHandler struct {
	TripRepository repositories.TripRepository
}

func NewTripHandler(tripRepository repositories.TripRepository) *tripHandler {
	return &tripHandler{tripRepository}

}

var path_file = "http://localhost:5000/uploads/"

func (h *tripHandler) FindTrip(c echo.Context) error {
	trips, err := h.TripRepository.FindTrip()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})

	}

	for i, p := range trips {
		trips[i].Image = path_file + p.Image
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: trips})

}

func (h *tripHandler) GetTripByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	trip, err := h.TripRepository.GetTripByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	trip.Image = path_file + trip.Image

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(trip)})

}

func (h *tripHandler) CreateTrip(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this data file", dataFile)

	country_id, _ := strconv.Atoi(c.FormValue("country_id"))
	day, _ := strconv.Atoi(c.FormValue("day"))
	night, _ := strconv.Atoi(c.FormValue("night"))
	price, _ := strconv.Atoi(c.FormValue("price"))
	quota, _ := strconv.Atoi(c.FormValue("quota"))

	request := tripdto.CreateTripRequest{
		Title:          c.FormValue("title"),
		CountryID:      country_id,
		Accommodation:  c.FormValue("accommodation"),
		Transportation: c.FormValue("transportation"),
		Eat:            c.FormValue("eat"),
		Day:            day,
		Night:          night,
		DateTrip:       c.FormValue("date_trip"),
		Price:          price,
		Quota:          quota,
		Description:    c.FormValue("description"),
		Image:          dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	country, _ := h.TripRepository.GetCountryByID(request.CountryID)

	trip := models.Trip{
		Title:          request.Title,
		Country:        country,
		CountryID:      request.CountryID,
		Accommodation:  request.Accommodation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		Day:            request.Day,
		Night:          request.Night,
		DateTrip:       request.DateTrip,
		Price:          request.Price,
		Quota:          request.Quota,
		Description:    request.Description,
		Image:          request.Image,
		UserID:         int(userId),
	}

	data, err := h.TripRepository.CreateTrip(trip)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)})
}

func (h *tripHandler) UpdateTrip(c echo.Context) error {
	request := new(tripdto.UpdateTripRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	trip, err := h.TripRepository.GetTripByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	country, _ := h.TripRepository.GetCountryByID(request.CountryID)
	fmt.Println(country)

	if request.Title != "" {
		trip.Title = request.Title
	}

	// if request.Country != modelCountry {
	// 	trip.Country = request.Country
	// }

	trip.Country = country
	// if request.Country != 0 {
	// 	trip.Country = request.Country
	// }

	// trip.CountryID = request.CountryID

	if request.CountryID != 0 {
		trip.CountryID = request.CountryID
	}

	if request.Accommodation != "" {
		trip.Accommodation = request.Accommodation
	}

	if request.Transportation != "" {
		trip.Transportation = request.Transportation
	}

	if request.Eat != "" {
		trip.Eat = request.Eat
	}

	if request.Day != 0 {
		trip.Day = request.Day
	}

	if request.Night != 0 {
		trip.Night = request.Night
	}

	if request.DateTrip != "" {
		trip.DateTrip = request.DateTrip
	}

	if request.Price != 0 {
		trip.Price = request.Price
	}

	if request.Quota != 0 {
		trip.Quota = request.Quota
	}

	if request.Description != "" {
		trip.Description = request.Description
	}

	if request.Image != "" {
		trip.Eat = request.Image
	}

	data, err := h.TripRepository.UpdateTrip(trip)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *tripHandler) DeleteTrip(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	trip, err := h.TripRepository.GetTripByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TripRepository.DeleteTrip(trip, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)})
}

func convertResponseTrip(u models.Trip) tripdto.TripResponse {
	return tripdto.TripResponse{
		ID:             u.ID,
		Title:          u.Title,
		CountryID:      u.CountryID,
		Country:        models.Country(u.Country),
		Accommodation:  u.Accommodation,
		Transportation: u.Transportation,
		Eat:            u.Eat,
		Day:            u.Day,
		Night:          u.Night,
		DateTrip:       u.DateTrip,
		Price:          u.Price,
		Quota:          u.Quota,
		Description:    u.Description,
		Image:          u.Image,
	}
}

// func convertUpdateTrip(u models.Trip) tripdto.UpdateTripRequest {
// 	return tripdto.UpdateTripRequest{
// 		Title:          u.Title,
// 		CountryID:      u.CountryID,
// 		Accommodation:  u.Accommodation,
// 		Transportation: u.Transportation,
// 		Eat:            u.Eat,
// 		Day:            u.Day,
// 		Night:          u.Night,
// 		DateTrip:       u.DateTrip,
// 		Price:          u.Price,
// 		Quota:          u.Quota,
// 		Description:    u.Description,
// 		Image:          u.Image,
// 	}
// }
