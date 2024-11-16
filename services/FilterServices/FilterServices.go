package FilterServices

import (
	"errors"
	"fmt"
	"magical-crwler/database"
	"magical-crwler/models"
	"magical-crwler/models/Dtos"
	"magical-crwler/services/Logger"
	"magical-crwler/utils"
)

type FilterServices struct {
	repository database.IRepository
	logger     *Logger.Logger
}

func NewFilterServices(repository database.IRepository) *FilterServices {
	return &FilterServices{
		repository: repository,
		logger:     Logger.NewLogger(repository),
	}
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

func (s FilterServices) ApplyFilters() error {
	filters, err := s.repository.GetAllFilters()
	if err != nil {
		return err
	}

	for index := range filters {
		filter := filters[index]
		ads, err := s.repository.SearchAds(filter, "id") // just fetch ads id filed
		if err != nil {
			s.logger.Error(fmt.Sprintf("Error in search filters :%s", err.Error()))
			continue
		}

		if len(ads) == 0 {
			continue
		}
		ids := make([]int, 0)
		for index := range ads {
			ids = append(ids, int(ads[index].ID))
		}
		user, err := s.repository.GetAFilterOwner(filter)
		if err != nil {
			s.logger.Error(fmt.Sprintf("Error in fetching user :%s", err.Error()))
			continue
		}
		existingAdIds, err := s.repository.GetExistingFiltersAds(filter)
		if err != nil {
			s.logger.Error(fmt.Sprintf("Error in search filter_ad :%s", err.Error()))
			continue
		}
		var diff []int
		if len(existingAdIds) == 0 {
			diff = existingAdIds
		} else {
			diff = utils.Difference(ids, existingAdIds)
		}

		if len(diff) == 0 {
			continue
		}
		newAds, err := s.repository.GetAdsByIDs(diff)
		if err != nil {
			s.logger.Error(fmt.Sprintf("Error in fetching ads :%s", err.Error()))
			continue
		}
		messageContent := utils.GenerateFilterMessage(newAds)
		s.repository.SaveFilterAds(diff, user.ID, filter.ID)
		//TODO: send message with telegram bot
		fmt.Printf("sending message to %s content %s\n", user.FirstName, messageContent)
	}
	return nil
}
