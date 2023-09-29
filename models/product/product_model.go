package ProductModel

import CategoryModel "batch11_golang_pos/models/category"

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Stock int    `json:"stock" validate:"required"`

	Category_id int                    `json:"category_id" validate:"required"`
	Category    CategoryModel.Category `gorm:"foreignKey:Category_id"`
}
