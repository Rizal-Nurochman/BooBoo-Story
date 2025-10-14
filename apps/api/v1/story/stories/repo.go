package stories

import "gorm.io/gorm"

type Repo interface{
	
}

type repo struct{
	db *gorm.DB
}


func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}