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

package e2e

import (
	_ "embed"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed params.json
var Params []byte

// TestParameterUpdate tests the module's ability to modify module parameters.
// This test uses the "execute" command to test broadcasting arbitrary messages.
func TestParameterUpdate(t *testing.T) {
	t.Parallel()

	var wrapper Wrapper
	ctx, _, _, _ := Suite(t, &wrapper, false)
	validator := wrapper.chain.Validators[0]

	EnsureParams(t, wrapper, ctx, 0)

	require.NoError(t, validator.WriteFile(ctx, Params, "params.json"))
	_, err := validator.ExecTx(
		ctx, wrapper.owner.KeyName(),
		"authority", "execute", path.Join(validator.HomeDir(), "params.json"),
	)
	require.NoError(t, err)

	EnsureParams(t, wrapper, ctx, 50)
}
