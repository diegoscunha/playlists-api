package configs

import (
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	ServerPort    string `mapstructure:"SERVER_PORT"`
	MongoUser     string `mapstructure:"MONGO_USER"`
	MongoPassword string `mapstructure:"MONGO_PASSWORD"`
	MongoHost     string `mapstructure:"MONGO_HOST"`
	MongoPort     string `mapstructure:"MONGO_PORT"`
	MongoDbName   string `mapstructure:"MONGO_DB_NAME"`
}

var (
	logger *Logger
	db     *mongo.Client
	config *Config
)

const (
	ERRO_INICIALIZACAO_VAR_AMBIENTE string = "Erro inicialização variaveis de ambiente: %v"
	ERRO_INICIALIZACAO_MONGODB      string = "Erro inicialização mongodb: %v"
)

func Init() error {
	var err error
	err = initConfig()
	if err != nil {
		logger.Errorf(ERRO_INICIALIZACAO_VAR_AMBIENTE, err)
		return fmt.Errorf(ERRO_INICIALIZACAO_VAR_AMBIENTE, err)
	}

	db, err = InitMongodb()
	if err != nil {
		logger.Errorf(ERRO_INICIALIZACAO_MONGODB, err)
		return fmt.Errorf(ERRO_INICIALIZACAO_MONGODB, err)
	}
	return nil

}

func initConfig() (err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}
	return
}

func GetMongodb() *mongo.Client {
	return db
}

func GetLogger(prefix string) *Logger {
	return NewLogger(prefix)
}

func GetConfig() *Config {
	return config
}
