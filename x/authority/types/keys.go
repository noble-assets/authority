package types

import authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

const ModuleName = "authority"

var ModuleAddress = authtypes.NewModuleAddress(ModuleName)

var AuthorityKey = []byte("authority")
