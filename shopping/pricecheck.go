package shopping

import (
    "shopping/db"
)

// type Item struct {
//     Price float64
// }

func PriceCheck(itemId int) (float64, bool) {
    item := db.LoadItem(itemId)
    if item == nil {
        return 0, false
    }
    return item.Price, true
}