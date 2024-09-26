// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package e2e

import (
	_ "embed"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed upgrade.json
var Upgrade []byte

// TestScheduleUpgrade tests the module's ability to schedule an upgrade on-chain.
func TestScheduleUpgrade(t *testing.T) {
	t.Parallel()

	var wrapper Wrapper
	ctx, _, _ := Suite(t, &wrapper, false)
	validator := wrapper.chain.Validators[0]

	EnsureUpgrade(t, wrapper, ctx, "", 0)

	require.NoError(t, validator.WriteFile(ctx, Upgrade, "upgrade.json"))
	_, err := validator.ExecTx(
		ctx, wrapper.owner.KeyName(),
		"authority", "execute", path.Join(validator.HomeDir(), "upgrade.json"),
	)
	require.NoError(t, err)

	EnsureUpgrade(t, wrapper, ctx, "v2", 50)
}
