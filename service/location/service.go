package location

import (
	"context"
	"github.com/guneyin/locator/dto"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Add(ctx context.Context, loc *dto.LocationDto) (dto.LocationResponseDto, error) {
	return dto.LocationResponseDto{
		Id:          "1",
		LocationDto: *loc,
	}, nil
}

func (s *Service) List(ctx context.Context) (dto.LocationListResponseDto, error) {
	return dto.LocationListResponseDto{}, nil
}

func (s *Service) Detail(ctx context.Context, id string) (dto.LocationResponseDto, error) {
	return dto.LocationResponseDto{
		Id:          "1",
		LocationDto: dto.LocationDto{},
	}, nil
}

func (s *Service) Edit(ctx context.Context, id string, loc *dto.LocationDto) (dto.LocationResponseDto, error) {
	return dto.LocationResponseDto{
		Id:          "1",
		LocationDto: *loc,
	}, nil
}

func (s *Service) Route(ctx context.Context, loc *dto.LocationDto) (dto.LocationListResponseDto, error) {
	return dto.LocationListResponseDto{}, nil
}
