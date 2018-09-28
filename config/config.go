package config

import (
	"encoding/json"
	"errors"

	consul "github.com/hashicorp/consul/api"
)

const appKey = "mapped/config"

type Configuration struct {
	MongoDB string `json:"mongo_db"`
}

func FromConsul(client *consul.Client) (*Configuration, error) {
	var cfg = new(Configuration)

	pair, _, err := client.KV().Get(appKey, nil)
	if err != nil {
		return nil, err
	}

	if pair == nil {
		return nil, errors.New("configuration from consul is empty")
	}

	return cfg, json.Unmarshal(pair.Value, cfg)
}
