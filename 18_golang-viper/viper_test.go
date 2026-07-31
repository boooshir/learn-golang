package golangviper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()

	assert.NotNil(t, config)
}

func TestJson(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	// read config
	err := config.ReadInConfig()

	assert.Nil(t, err)

	assert.Equal(t, "learn golang viper", config.GetString("app.name"))
	assert.Equal(t, "boshir", config.GetString("app.author"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.True(t, true, config.GetBool("database.show_sql"))
}

func TestYaml(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")

	// read config
	err := config.ReadInConfig()

	assert.Nil(t, err)

	assert.Equal(t, "learn golang viper", config.GetString("app.name"))
	assert.Equal(t, "boshir", config.GetString("app.author"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.True(t, true, config.GetBool("database.show_sql"))
}

func TestEnv(t *testing.T) {
	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")

	//read .env
	config.AutomaticEnv()
	// read config
	err := config.ReadInConfig()

	assert.Nil(t, err)

	assert.Equal(t, "learn-golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "boshir", config.GetString("APP_AUTHOR"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.True(t, true, config.GetBool("DATABASE_SHOW_SQL"))

	assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
}
