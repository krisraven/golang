package db

// type Item struct {
//     Price float64
// }

import (
    "shopping/models"
)

// if the name of the type of function starts with
// an uppercase letter, then it's visible outside of a package
func LoadItem(id int) *models.Item {
    return &models.Item{
        Price: 9.001,
    }
}