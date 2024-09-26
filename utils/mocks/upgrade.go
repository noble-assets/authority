// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package mocks

import (
	"context"
	"errors"

	upgradetypes "cosmossdk.io/x/upgrade/types"
)

var _ upgradetypes.MsgServer = UpgradeKeeper{}

type UpgradeKeeper struct{}

func (UpgradeKeeper) SoftwareUpgrade(_ context.Context, _ *upgradetypes.MsgSoftwareUpgrade) (*upgradetypes.MsgSoftwareUpgradeResponse, error) {
	return &upgradetypes.MsgSoftwareUpgradeResponse{}, nil
}

func (UpgradeKeeper) CancelUpgrade(_ context.Context, _ *upgradetypes.MsgCancelUpgrade) (*upgradetypes.MsgCancelUpgradeResponse, error) {
	return &upgradetypes.MsgCancelUpgradeResponse{}, errors.New("intentional error")
}
