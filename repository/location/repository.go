package location

import (
	"context"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	r := &Repository{db: db}
	return r.Migrate()
}

func (r *Repository) Migrate() *Repository {
	err := r.db.AutoMigrate(Location{})
	if err != nil {
		panic(err)
	}

	return r
}

func (r *Repository) Add(ctx context.Context, loc *Location) (*Location, error) {
	err := r.db.WithContext(ctx).Create(&loc).Error
	if err != nil {
		return nil, err
	}

	return loc, nil
}

func (r *Repository) List(ctx context.Context) (LocationList, error) {
	locList := make(LocationList, 0)
	err := r.db.WithContext(ctx).Find(&locList).Error
	if err != nil {
		return nil, err
	}

	return locList, nil
}

func (r *Repository) Detail(ctx context.Context, id uint) (*Location, error) {
	loc := new(Location)
	loc.ID = id
	err := r.db.WithContext(ctx).First(&loc).Error
	if err != nil {
		return nil, err
	}

	return loc, nil
}

func (r *Repository) Edit(ctx context.Context, id uint, loc *Location) (*Location, error) {
	loc.ID = id
	err := r.db.WithContext(ctx).Updates(&loc).Error
	if err != nil {
		return nil, err
	}

	return r.Detail(ctx, id)
}

func (r *Repository) Delete(ctx context.Context, idList []uint) error {
	return r.db.WithContext(ctx).Delete(&Location{}, idList).Error
}
