package data

import (
	"fmt"

	"github.com/icrowley/fake"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("../config/")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func SeedDockerdb() {
	viper.WatchConfig()
	// db, err := gorm.Open("postgres", "host=postgres"+ /* viper.GetString("dbHost")+*/ " user=relay"+ /* viper.GetString("dbUser")+*/ " dbname=relay_development"+ /* viper.GetString("dbName")+*/ " sslmode=disable password=noahheil1" /*+ viper.GetString("dbPassword")*/)
	db, _ := gorm.Open("postgres", "host=postgres user=relay dbname=relay_development sslmode=disable password=noahheil1")
	// if err != nil { // Handle errors reading the config file
	// 	panic(fmt.Errorf("Fatal error connecting to database: %s \n", err))
	// }
	defer db.Close()

	db.AutoMigrate(&User{})

	for index := 0; index < 10; index++ {
		user := User{FirstName: fake.FirstName(), LastName: fake.LastName(), Email: fake.EmailAddress()}
		db.Create(&user)
	}
}
