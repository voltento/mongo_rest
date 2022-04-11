package service

import (
	"github.com/voltento/mongo_rest/app/dto"
	"github.com/voltento/mongo_rest/app/repo"
)

type Search struct {
	repo *repo.Mongo
}

func NewSearch(r *repo.Mongo) *Search {
	return &Search{repo: r}
}

func (s *Search) FindRecords(f *dto.Filters) []dto.Record {
	return s.repo.FindRecords(f)
}
