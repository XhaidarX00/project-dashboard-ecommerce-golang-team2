package database

import (
	"dashboard-ecommerce-team2/models"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		seeds := dataSeeds()

		for _, seedFunc := range seeds {
			// Each seed function should check for existing data internally
			if err := seedFunc.(func(*gorm.DB) error)(tx); err != nil {
				seedName := reflect.TypeOf(seedFunc).Name()
				return fmt.Errorf("seeding failed for %s: %v", seedName, err)
			}
		}

		return nil
	})
}

func dataSeeds() []interface{} {
	return []interface{}{
		models.UserSeed(),
		models.CategorySeed(),
		models.ProductSeed(),
		models.StockHistorySeed(),
		models.BannerSeed(),
		models.OrderSeed(),
		models.OrderItemSeed(),
		models.PromotionSeed(),
	}
}
