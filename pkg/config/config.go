package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	sandBoxUserStoreURI = "https://sandbox.evernote.com/edam/user"
	// Get it here https://sandbox.evernote.com/api/DeveloperToken.action
	sandBoxNoteStoreURI = "https://sandbox.evernote.com/shard/s1/notestore"
)

type ConfigKey string

const (
	ApiKey       ConfigKey = "apiKey"
	UserStoreURI ConfigKey = "userStoreURI"
	NoteStoreURI ConfigKey = "noteStoreURI"
)

func GetApiKey() string {
	return viper.GetString(string(ApiKey))
}

func GetUserStoreURI() string {
	value := viper.GetString(string(UserStoreURI))
	if value != "" {
		return value
	}
	return sandBoxUserStoreURI
}

func GetNoteStoreURI() string {
	value := viper.GetString(string(NoteStoreURI))
	if value != "" {
		return value
	}
	return sandBoxNoteStoreURI
}

func SetConfigAndSave(key ConfigKey, value interface{}) {
	viper.Set(string(key), value)
	err := viper.WriteConfig()
	if err != nil {
		log.Printf("Warning: can not write config, error: %s", err)
	}
}
