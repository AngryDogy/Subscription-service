package service

import (
	"context"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
)

type CountryService struct {
	countryRepository repository.CountryRepository
	protogen.UnimplementedCountryServiceServer
}

func NewCountryService(repository repository.Repository) *CountryService {
	return &CountryService{countryRepository: repository}
}

func (s *CountryService) GetCountryByName(ctx context.Context, countryName *protogen.CountryName) (*protogen.Country, error) {
	country, err := s.countryRepository.FindCountryByName(ctx, countryName.Name)
	if err != nil {
		return nil, err
	}

	return &protogen.Country{Id: country.Id, Name: country.Name}, nil
}
