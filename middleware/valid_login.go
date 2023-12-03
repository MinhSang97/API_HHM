package middleware

import "gorm.io/gorm"

func valid_login(db *gorm.DB){
	result := db.execute("Tokens").Find(&tokens)

}