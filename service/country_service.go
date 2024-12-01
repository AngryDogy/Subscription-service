package service

import (
	"context"
	protogen "dev/master/protogen/proto/api/v1"
	"dev/master/repository"
	"google.golang.org/protobuf/types/known/emptypb"
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
		return nil, nil
	}

	return &protogen.Country{Id: country.Id, Name: country.Name}, nil
}

func (s *CountryService) GetAllCountries(ctx context.Context, request *emptypb.Empty) (*protogen.Countries, error) {
	countries, err := s.countryRepository.GetAllCountries(ctx)
	if err != nil {
		return nil, nil
	}

	var protogenCountries []*protogen.Country
	for _, country := range countries {
		protogenCountries = append(protogenCountries, &protogen.Country{Id: country.Id, Name: country.Name})
	}

	return &protogen.Countries{Countries: protogenCountries}, nil
}

func (s *CountryService) CreateCountry(ctx context.Context, countryName *protogen.CountryName) (*protogen.Country, error) {
	country, err := s.countryRepository.CreateCountry(ctx, countryName.Name)
	if err != nil {
		return nil, nil
	}

	return &protogen.Country{Id: country.Id, Name: country.Name}, nil

}
