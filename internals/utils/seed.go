package utils

import (
	fake "github.com/brianvoe/gofakeit/v7"
)

func SeedDb() error {
	for i := 0; i < 100; i++ {
		err := CreateSecret(fake.Name(), fake.Username(), fake.Password(true, true, true, true, false, 8))
		if err != nil {
			return err
		}
	}
	return nil
}

