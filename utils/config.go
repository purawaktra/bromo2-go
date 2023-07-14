package utils

import (
	"github.com/spf13/viper"
	"os"
)

var (
	AppName        string
	AppHost        string
	AppPort        string
	AppEnvironment string
	AppLogLevel    string

	AppAuthUsername string
	AppAuthPassword string

	MongoDBHost     string
	MongoDBPort     string
	MongoDBDatabase string
	MongoDBUser     string
	MongoDBPassword string

	CollectionFiles string
)

func InitConfig() {
	// check app environment on env
	env := os.Getenv("APP_ENV")

	// check for value
	if env == "" {
		env = "development"
	}

	if env == "development" {
		// check for config.json
		viper.SetConfigFile(`config.json`)

		// read file
		err := viper.ReadInConfig()
		if err != nil {
			Fatal(err, "InitConfig", "")
			panic(err)
		}

		// get variable for app
		AppName = viper.GetString("app.name")
		AppHost = viper.GetString("app.host")
		AppPort = viper.GetString("app.port")
		AppEnvironment = viper.GetString("app.environment")
		AppLogLevel = viper.GetString("app.log.level")
		AppAuthUsername = viper.GetString("app.auth.username")
		AppAuthPassword = viper.GetString("app.auth.password")

		// get variable for db
		MongoDBHost = viper.GetString("database.mongodb.host")
		MongoDBPort = viper.GetString("database.mongodb.port")
		MongoDBDatabase = viper.GetString("database.mongodb.database")
		MongoDBUser = viper.GetString("database.mongodb.user")
		MongoDBPassword = viper.GetString("database.mongodb.password")

		// get mongodb collection name
		CollectionFiles = viper.GetString("database.mongodb.collection.files")

		// create return
		return
	}

	if env == "staging" || env == "production" {
		// get variable for app
		AppName = os.Getenv("APP_NAME")
		AppPort = os.Getenv("APP_PORT")
		AppEnvironment = os.Getenv("APP_ENV")
		AppLogLevel = os.Getenv("APP_LOG_LEVEL")
		AppAuthUsername = os.Getenv("APP_AUTH_USERNAME")
		AppAuthPassword = os.Getenv("APP_AUTH_PASSWORD")

		// get variable for db
		MongoDBHost = os.Getenv("MONGODB_HOST")
		MongoDBPort = os.Getenv("MONGODB_PORT")
		MongoDBDatabase = os.Getenv("MONGODB_DATABASE")
		MongoDBUser = os.Getenv("MONGODB_USER")
		MongoDBPassword = os.Getenv("MONGODB_PASSWORD")
	}
}
