package config

import (
	"github.com/BurntSushi/toml"
)

type Mode string

type HTTP struct {
	Address  string `toml:"address"`
	BasePath string `toml:"base_path"`
}

type AWS struct {
	Region          string `toml:"region"`
	AccessKeyID     string `toml:"access_key_id"`
	SecretAccessKey string `toml:"secret_access_key"`
}

type Store struct {
	DynamoDBTable string `toml:"dynamodb_table"`
}

type Cathedral struct {
	Mode  Mode   `toml:"mode"`
	HTTP  *HTTP  `toml:"http"`
	AWS   *AWS   `toml:"AWS"`
	Store *Store `toml:"Store"`
}

func LoadConfig(path string) (*Cathedral, error) {
	cathedral := &Cathedral{}
	_, err := toml.DecodeFile(path, cathedral)
	if err != nil {
		return nil, err
	}
	return cathedral, nil
}
