// +build integration

package config

import (
	"encoding/json"
	"os"
	"testing"

	consul "github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/require"
)

const (
	CONSUL_PORT = "MAPPED_CONSUL_PORT"
)

func addr() string {
	return os.Getenv(CONSUL_PORT)
}

func TestFromConsul(t *testing.T) {

	sample := &Configuration{
		MongoDB: "oops",
	}

	buf, err := json.Marshal(sample)
	require.NoError(t, err)

	consulCfg := consul.DefaultConfig()
	consulCfg.Address = addr()

	client, err := consul.NewClient(consulCfg)
	require.NoError(t, err)

	_, err = client.KV().Put(&consul.KVPair{
		Key:   appKey,
		Value: buf,
	}, nil)
	require.NoError(t, err)

	cfg, err := FromConsul(client)
	require.NoError(t, err)

	require.Equal(t, sample, cfg)
}
