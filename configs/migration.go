package configs

import (
	CategoryModel "batch11_golang_pos/models/category"
	UserModel "batch11_golang_pos/models/user"
)

func initMigrate() {
	DB.AutoMigrate(&UserModel.User{})
	DB.AutoMigrate(&CategoryModel.Category{})
}
