package data

import (
	"github.com/jinzhu/gorm"
)

// Model structs
type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string
}

// Data accessors
func GetUser(id string) *User {
	// db, err := gorm.Open("postgres", "host="+
	// 	viper.GetString("dbHost")+
	// 	" user="+
	// 	viper.GetString("dbUser")+
	// 	" dbname="+
	// 	viper.GetString("dbName")+
	// 	" sslmode=disable password="+
	// 	viper.GetString("dbPassword"))
	// if err != nil { // Handle errors reading the config file
	// 	panic(fmt.Errorf("Fatal error connecting to database: %s \n", err))
	// }
	db, _ := gorm.Open("postgres", "host=postgres user=relay dbname=relay_development sslmode=disable password=noahheil1")
	defer db.Close()

	var user User

	db.First(&user, id)

	return &user
}
