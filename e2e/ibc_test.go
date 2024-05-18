package e2e

import "testing"

// TestIBCClientSubstitution tests the module's ability to substitute an expired
// IBC client. This is currently skipped due to simapp not supporting IBC.
func TestIBCClientSubstitution(t *testing.T) {
	// TODO: Implement once IBC is added to the simapp.
	t.Skip()
	t.Parallel()
}
