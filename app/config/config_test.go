package config

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name string
		want Config
		envs map[string]string
	}{
		{
			name: "Default values",
			want: Config{ServiceHost: "localhost:8080",
				MongoDBHost:         "localhost:27017",
				MongoDBName:         "wave",
				MongoCollectionName: "numbers",
				User:                "user",
				Password:            "password",
			},
		},
		{
			name: "Get service host port",
			want: cfg(Config{ServiceHost: "localhost:9977"}),
			envs: map[string]string{"SERVICE_HOST": "localhost:9977"},
		},
		{
			name: "Get mongoDB name from envs",
			want: cfg(Config{MongoDBName: "database_name"}),
			envs: map[string]string{"MONGO_DB_NAME": "database_name"},
		},
		{
			name: "Get mongoDB collection name",
			want: cfg(Config{MongoCollectionName: "collection_name"}),
			envs: map[string]string{"MONGO_COLLECTION_NAME": "collection_name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envs {
				err := os.Setenv(k, v)
				if err != nil {
					log.Fatal(err)
				}
			}

			if got := GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}

			for k := range tt.envs {
				err := os.Unsetenv(k)
				if err != nil {
					log.Fatal(err)
				}
			}
		})
	}
}

// Build default config with custom values
func cfg(customConfig Config) Config {
	defaultConfig := GetConfig()
	if len(customConfig.ServiceHost) > 0 {
		defaultConfig.ServiceHost = customConfig.ServiceHost
	}

	if len(customConfig.MongoDBHost) > 0 {
		defaultConfig.MongoDBHost = customConfig.MongoDBHost
	}

	if len(customConfig.MongoDBName) > 0 {
		defaultConfig.MongoDBName = customConfig.MongoDBName
	}

	if len(customConfig.MongoDBCreds) > 0 {
		defaultConfig.MongoDBCreds = customConfig.MongoDBCreds
	}

	if len(customConfig.ServiceHost) > 0 {
		defaultConfig.ServiceHost = customConfig.ServiceHost
	}

	if len(customConfig.MongoCollectionName) > 0 {
		defaultConfig.MongoCollectionName = customConfig.MongoCollectionName
	}

	return defaultConfig
}
