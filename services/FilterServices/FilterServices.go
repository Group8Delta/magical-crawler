package FilterServices

import (
	"errors"
	"magical-crwler/database"
	"magical-crwler/models"
	"magical-crwler/models/Dtos"
)

type FilterServices struct {
	repository database.IRepository
}

func NewFilterServices(repository database.IRepository) *FilterServices {
	return &FilterServices{repository: repository}
}
func (s FilterServices) CreateFilter(filter Dtos.FilterDto) (models.Filter, []error) {
	f := models.Filter{}
	Errors := make([]error, 0)
	if filter.CreationTimeRangeFrom.Compare(filter.CreationTimeRangeTo) > 0 {
		Errors = append(Errors, errors.New("Invalid Creation Time Range"))
	}
	if filter.FloorRange.Min > filter.FloorRange.Max {
		Errors = append(Errors, errors.New("Invalid Floor Range"))
	}
	if filter.PriceRange.Min > filter.PriceRange.Max {
		Errors = append(Errors, errors.New("Invalid Price Range"))

	}
	if filter.SizeRange.Min > filter.SizeRange.Max {
		Errors = append(Errors, errors.New("Invalid Size Range"))
	}
	if filter.RentPriceRange.Min > filter.RentPriceRange.Max {
		Errors = append(Errors, errors.New("Invalid Rent Price"))
	}
	if filter.BedroomRange.Min > filter.BedroomRange.Max {
		Errors = append(Errors, errors.New("Invalid Bed Room Range"))
	}
	if filter.AgeRange.Min > filter.AgeRange.Max {
		Errors = append(Errors, errors.New("Invalid Age Range"))
	}
	if len(Errors) > 0 {
		return f, Errors
	}
	res := s.repository.CreateFilter(filter)
	return res, nil
}
func (s FilterServices) UpdateFilter(filter Dtos.FilterDto) (models.Filter, []error) {
	f := models.Filter{}
	Errors := make([]error, 0)
	if filter.CreationTimeRangeFrom.Compare(filter.CreationTimeRangeTo) > 0 {
		Errors = append(Errors, errors.New("Invalid Creation Time Range"))
	}
	if filter.FloorRange.Min > filter.FloorRange.Max {
		Errors = append(Errors, errors.New("Invalid Floor Range"))
	}
	if filter.PriceRange.Min > filter.PriceRange.Max {
		Errors = append(Errors, errors.New("Invalid Price Range"))
	}
	if filter.SizeRange.Min > filter.SizeRange.Max {
		Errors = append(Errors, errors.New("Invalid Size Range"))
	}
	if filter.RentPriceRange.Min > filter.RentPriceRange.Max {
		Errors = append(Errors, errors.New("Invalid Rent Price"))
	}
	if filter.BedroomRange.Min > filter.BedroomRange.Max {
		Errors = append(Errors, errors.New("Invalid Bed Room Range"))
	}
	if filter.AgeRange.Min > filter.AgeRange.Max {
		Errors = append(Errors, errors.New("Invalid Age Range"))
	}
	if len(Errors) > 0 {
		return f, Errors
	}
	res, err := s.repository.UpdateFilter(filter)
	Errors = append(Errors, err)
	return res, Errors
}
func (s FilterServices) GetFilterById(id int) (models.Filter, error) {
	return s.repository.GetFilterById(id)
}
