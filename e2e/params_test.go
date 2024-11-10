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
