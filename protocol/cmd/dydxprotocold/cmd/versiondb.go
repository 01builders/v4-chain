//go:build rocksdb
// +build rocksdb

package cmd

import (
	"sort"

	versiondbclient "github.com/crypto-org-chain/cronos/versiondb/client"
	"github.com/dydxprotocol/v4-chain/protocol/app"
	"github.com/dydxprotocol/v4-chain/protocol/cmd/dydxprotocold/opendb"
	"github.com/spf13/cobra"
)

func ChangeSetCmd() *cobra.Command {
	keys, _, _ := app.StoreKeys()
	storeNames := make([]string, 0, len(keys))
	for name := range keys {
		storeNames = append(storeNames, name)
	}
	sort.Strings(storeNames)

	return versiondbclient.ChangeSetGroupCmd(versiondbclient.Options{
		DefaultStores:     storeNames,
		OpenReadOnlyDB:    opendb.OpenReadOnlyDB,
		AppRocksDBOptions: opendb.NewRocksdbOptions,
	})
}
