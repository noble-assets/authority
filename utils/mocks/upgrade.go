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
