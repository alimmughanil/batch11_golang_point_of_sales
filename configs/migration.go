package configs

import UserModel "batch11_golang_pos/models/user"

func initMigrate() {
	DB.AutoMigrate(&UserModel.User{})
}
