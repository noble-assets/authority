// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package types

func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

func (gs *GenesisState) Validate() error {
	return nil
}
