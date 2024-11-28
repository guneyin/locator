package location

import (
	"context"
	"github.com/guneyin/locator/dto"
	"github.com/guneyin/locator/repository/location"
	"gorm.io/gorm"
)

type Service struct {
	repo *location.Repository
}

func New(db *gorm.DB) *Service {
	return &Service{repo: location.NewRepository(db)}
}

func (s *Service) Add(ctx context.Context, loc *dto.LocationDto) (*dto.LocationResponseDto, error) {
	added, err := s.repo.Add(ctx, loc.ToEntity())
	if err != nil {
		return nil, err
	}

	return dto.NewLocationResponseDto(added)
}

func (s *Service) List(ctx context.Context) (*dto.LocationListResponseDto, error) {
	locList, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return dto.NewLocationListResponseDto(locList)
}

func (s *Service) Detail(ctx context.Context, id uint) (*dto.LocationResponseDto, error) {
	loc, err := s.repo.Detail(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.NewLocationResponseDto(loc)
}

func (s *Service) Edit(ctx context.Context, id uint, loc *dto.LocationDto) (*dto.LocationResponseDto, error) {
	edited, err := s.repo.Edit(ctx, id, loc.ToEntity())
	if err != nil {
		return nil, err
	}

	return dto.NewLocationResponseDto(edited)
}

func (s *Service) Route(ctx context.Context, loc *dto.LocationDto) (*dto.LocationListResponseDto, error) {
	locList, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return dto.NewLocationListResponseDto(locList)
}
