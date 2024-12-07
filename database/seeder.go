package database

import (
	"dashboard-ecommerce-team2/models"
	"fmt"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedAll(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		seeds := dataSeeds()
		for i := range seeds {
			err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(seeds[i]).Error
			if nil != err {
				name := reflect.TypeOf(seeds[i]).String()
				errorMessage := err.Error()
				return fmt.Errorf("%s seeder fail with %s", name, errorMessage)
			}
		}
		return nil
	})
}

func dataSeeds() []interface{} {
	return []interface{}{
		models.UserSeed(),
		// models.CategorySeed(),
		// models.ProductSeed(),
		// models.StockHistorySeed(),
		// models.BannerSeed(),
		// models.OrderSeed(),
		// models.OrderItemSeed(),
		// models.PromotionSeed(),
	}
}
