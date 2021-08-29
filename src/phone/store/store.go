package store

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	db gorm.DB
}

const dsn = "host=localhost dbname=phone port=5432 sslmode=disable TimeZone=America/New_York"

func NewStore() (*Store, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Contact{})
	return &Store{db: *db}, nil
}

func (s Store) AddNumber(number string) error {
	res := s.db.Create(&Contact{Number: number})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (s Store) GetNumbers() ([]string, error) {
	var contacts []Contact
	res := s.db.Find(&contacts)
	if res.Error != nil {
		return nil, res.Error
	}
	numbers := make([]string, 0)
	for _, contact := range contacts {
		numbers = append(numbers, contact.Number)
	}
	return numbers, nil
}

func (s Store) Update(oldnumber, newnumber string) error {
	var contact Contact
	res := s.db.Where(&Contact{Number: oldnumber}).Find(&contact)
	if res.Error != nil {
		return res.Error
	}
	contact.Number = newnumber
	res = s.db.Save(&contact)
	if res.Error != nil {
		return res.Error
	}

	err := s.dedupe(newnumber)
	if err != nil {
		return err
	}

	return nil
}

func (s Store) Clear() error {
	res := s.db.Where("1 = 1").Delete(&Contact{})
	return res.Error
}

func (s Store) dedupe(num string) error {
	var contacts []Contact

	res := s.db.Where(&Contact{Number: num}).Find(&contacts)
	if res.RowsAffected <= 1 {
		return nil
	}

	// Delete all but the last
	for i, c := range contacts {
		if i < len(contacts)-1 {
			res := s.db.Delete(&c)
			if res.Error != nil {
				return res.Error
			}
		}
	}

	return nil
}
