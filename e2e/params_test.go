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
func TestParameterUpdate(t *testing.T) {
	t.Parallel()

	var wrapper Wrapper
	ctx := Suite(t, &wrapper)
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
