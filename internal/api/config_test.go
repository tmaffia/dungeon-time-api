package api

import (
	"os"
	"reflect"
	"testing"
)

func Test_newConfig(t *testing.T) {
	tests := []struct {
		name string
		want *config
	}{
		{name: "Config Envs", want: &config{
			databaseUrl: os.Getenv("DUNGEON_TIME_API_DATABASE_URL"),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
