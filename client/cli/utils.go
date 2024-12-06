package cli

import (
	"os"
	"path/filepath"

	"github.com/spf13/pflag"

	"cosmossdk.io/x/upgrade/types"
)

// parsePlan is copied from the Cosmos SDK because it is not exported.
//
// https://github.com/cosmos/cosmos-sdk/blob/x/upgrade/v0.1.4/x/upgrade/client/cli/parse.go#L9-L21
func parsePlan(fs *pflag.FlagSet, name string) (types.Plan, error) {
	height, err := fs.GetInt64(FlagUpgradeHeight)
	if err != nil {
		return types.Plan{}, err
	}

	info, err := fs.GetString(FlagUpgradeInfo)
	if err != nil {
		return types.Plan{}, err
	}

	return types.Plan{Name: name, Height: height, Info: info}, nil
}

// getDefaultDaemonName is copied from the Cosmos SDK because it is not exported.
//
// https://github.com/cosmos/cosmos-sdk/blob/x/upgrade/v0.1.4/x/upgrade/client/cli/tx.go#L184-L194
func getDefaultDaemonName() string {
	// DAEMON_NAME is specifically used here to correspond with the Cosmovisor setup env vars.
	name := os.Getenv("DAEMON_NAME")
	if len(name) == 0 {
		_, name = filepath.Split(os.Args[0])
	}
	return name
}
