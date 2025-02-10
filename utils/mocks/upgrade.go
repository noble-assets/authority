// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2025, NASD Inc. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN "AS IS" BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

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
