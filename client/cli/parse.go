package cli

import (
	"github.com/spf13/pflag"

	"cosmossdk.io/x/upgrade/types"
)

// parsePlan was copied from SDK becuase it was not exported
// https://github.com/cosmos/cosmos-sdk/blob/x/upgrade/v0.1.4/x/upgrade/client/cli/parse.go
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
