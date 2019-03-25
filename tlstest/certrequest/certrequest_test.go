package certrequest

import (
	"context"
	"testing"

	"github.com/dgraph-io/dgo/protos/api"
	"github.com/dgraph-io/dgraph/z"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestAccessOverPlaintext(t *testing.T) {
	dg := z.DgraphClient(":9180")
	err := dg.Alter(context.Background(), &api.Operation{DropAll: true})
	require.Error(t, err, "The authentication handshake should have failed")
}

func TestAccessWithCaCert(t *testing.T) {
	conf := viper.New()
	conf.Set("tls_cacert", "../tls/ca.crt")
	conf.Set("tls_server_name", "node")

	dg, err := z.DgraphClientWithCerts(":9180", conf)
	require.NoError(t, err, "Unable to get dgraph client: %v", err)
	err = dg.Alter(context.Background(), &api.Operation{DropAll: true})
	require.NoError(t, err, "Unable to perform dropall: %v", err)
}
