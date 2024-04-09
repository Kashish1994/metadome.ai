package configs

import (
	"fmt"
	"github.com/eduhub/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	config, err := bootDB()
	if err != nil {
		return nil, err
	}
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	fmt.Println(err)
	err = autoMigrate(db)
	if err != nil {
		return nil, err
	}
	initData(db)
	return db, nil
}

func initData(db *gorm.DB) {
}
func autoMigrate(db *gorm.DB) error {
	db.Debug()
	err := db.Table("users").AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	return nil
}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func bootDB() (*Config, error) {

	env := "local" // Change this to "prod" for production
	configFileName := fmt.Sprintf("config.%s"+".yaml", env)

	viper.AddConfigPath("./configs")
	fmt.Println(configFileName)
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		err = fmt.Errorf("error reading configs file: %s", err)
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)

	if err != nil {
		return nil, err
	}
	return &config, nil
}
