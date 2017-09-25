package main

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_GenerateServerCfg(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		cfg     *Config
		args    args
		wantErr bool
	}{
		{
			"required",
			&Config{
				Announce:   &[]bool{true}[0],
				Hostname:   &[]string{"Test"}[0],
				MaxPlayers: &[]int{32}[0],
				Port:       &[]int{8080}[0],
				RCON:       &[]bool{true}[0],
				Language:   &[]string{"English"}[0],
				Gamemodes: []string{
					"rivershell",
					"baserace",
				},
				RCONPassword: &[]string{"test"}[0],
			},
			args{"./testspace"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.GenerateServerCfg(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Config.GenerateServerCfg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_GenerateJSON(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		config  *Config
		want    []byte
		args    args
		wantErr bool
	}{
		{
			"minimal",
			&Config{
				Gamemodes: []string{
					"rivershell",
					"baserace",
				},
				RCONPassword: &[]string{"test"}[0],
				Port:         &[]int{8080}[0],
				Hostname:     &[]string{"Test"}[0],
				MaxPlayers:   &[]int{32}[0],
				Language:     &[]string{"English"}[0],
				Announce:     &[]bool{true}[0],
				RCON:         &[]bool{true}[0],
			},
			[]byte(`{
	"gamemodes": [
		"rivershell",
		"baserace"
	],
	"rcon_password": "test",
	"port": 8080,
	"hostname": "Test",
	"maxplayers": 32,
	"language": "English",
	"announce": true,
	"rcon": true
}`),
			args{"./testspace"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.GenerateJSON(tt.args.dir)
			assert.NoError(t, err)

			contents, err := ioutil.ReadFile(filepath.Join(tt.args.dir, "samp.json"))
			assert.NoError(t, err)

			assert.Equal(t, string(tt.want), string(contents))
		})
	}
}
