package dbutil

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

// type Token struct {
// 	Token      string    `gorm:"column:TOKEN"`
// 	UserName   string    `gorm:"column:USER_NAME"`
// 	Password   string    `gorm:"column:PASSWORD"`
// 	LoginDate  time.Time `gorm:"column:LOGIN_DATE"`
// 	ExpireDate time.Time `gorm:"column:EXPIRE_DATE"`
// 	ExpireTime time.Time `gorm:"column:EXPIRE_TIME"`
// 	IPAddress  string    `gorm:"column:IP_ADDRESS"`
// }

func ConnectDB() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	url := oracle.BuildUrl("118.69.35.119", 1521, "hhm", "MiniMDM10", "MiniMDM10", nil)
	db, err := gorm.Open(oracle.Open(url), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	fmt.Println("connected to the database:", db)
	return db, nil
}

// 	// Query the Tokens table
// 	tokens, err := queryTokens(db)
// 	if err != nil {
// 		fmt.Println("failed to query Tokens:", err)
// 		return
// 	}

// 	// Print the results
// 	fmt.Println("Tokens:")
// 	for _, token := range tokens {
// 		fmt.Printf("%+v\n", token)
// 	}
// }

//	func queryTokens(db *gorm.DB) ([]Token, error) {
//		var tokens []Token
//		result := db.Table("Tokens").Find(&tokens)
//		if result.Error != nil {
//			return nil, result.Error
//		}
//		return tokens, nil
