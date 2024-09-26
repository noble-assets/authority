// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package types

import (
	"fmt"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramstypes.ParamSet = &Params{}

var AuthorityKey = []byte("authority")

// Params is the legacy ParamAuthority interface.
// https://github.com/strangelove-ventures/paramauthority/blob/v1.1.0/x/params/types/proposal/genesis.pb.go#L71-L74
type Params struct {
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty" yaml:"authority"`
}

// Deprecated: ParamKeyTable returns the legacy ParamAuthority key table.
// https://github.com/strangelove-ventures/paramauthority/blob/v1.1.0/x/params/types/proposal/params.go#L11-L14
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface.
// https://github.com/strangelove-ventures/paramauthority/blob/v1.1.0/x/params/types/proposal/params.go#L33-L38
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(AuthorityKey, p.Authority, validateAuthority),
	}
}

// https://github.com/strangelove-ventures/paramauthority/blob/v1.1.0/x/params/types/proposal/params.go#L40-L51
func validateAuthority(i interface{}) error {
	a, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if _, err := sdk.AccAddressFromBech32(a); err != nil {
		return errors.Wrap(err, "authority")
	}

	return nil
}
