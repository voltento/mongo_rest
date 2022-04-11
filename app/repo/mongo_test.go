package repo

import (
	"github.com/stretchr/testify/assert"
	"github.com/voltento/mongo_rest/app/config"
	"github.com/voltento/mongo_rest/app/dto"
	"testing"
)

func Test_buildMongoHost(t *testing.T) {
	tests := []struct {
		name    string
		cfg     config.Config
		want    string
		wantErr bool
	}{
		{
			name: "All params selected",
			cfg: config.Config{
				MongoDBHost:  "localhost:27017",
				MongoDBCreds: "user:root",
			},
			want: "mongodb://user:root@localhost:27017",
		},
		{
			name: "Creds are empty",
			cfg: config.Config{
				MongoDBHost: "localhost:27017",
			},
			want: "mongodb://localhost:27017",
		},
		{
			name: "Host is not provided",
			cfg: config.Config{
				MongoDBHost: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildMongoHost(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildMongoHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("buildMongoHost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildFindParams_set_limit(t *testing.T) {
	opts, _ := buildFindParams(&dto.Filters{Limit: 1})
	assert.Equal(t, int64(1), *opts.Limit)
}

func Test_buildFindParams_set_found(t *testing.T) {
	_, filters := buildFindParams(&dto.Filters{Found: &[]bool{true}[0]})
	assert.Equal(t, true, filters["found"])
}

func Test_buildFindParams_set_number(t *testing.T) {
	_, filters := buildFindParams(&dto.Filters{Number: &[]float64{13}[0]})
	assert.Equal(t, float64(13), filters["number"])
}

func Test_buildFindParams_set_type(t *testing.T) {
	_, filters := buildFindParams(&dto.Filters{Type: &[]string{"test"}[0]})
	assert.Equal(t, "test", filters["type"])
}
