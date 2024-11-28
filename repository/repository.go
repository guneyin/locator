package repository

type IRepository interface {
	Migrate() IRepository
}
