package types

import "cosmossdk.io/errors"

var (
	ErrInvalidAuthority        = errors.Register(ModuleName, 1, "signer is not authority")
	ErrSameAuthority           = errors.Register(ModuleName, 2, "provided authority is the current authority")
	ErrInvalidPendingAuthority = errors.Register(ModuleName, 3, "signer is not pending authority")
	ErrInvalidMessage          = errors.Register(ModuleName, 4, "message is invalid")
)
