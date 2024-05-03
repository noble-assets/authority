package types

import "cosmossdk.io/errors"

var (
	ErrInvalidAuthority = errors.Register(ModuleName, 1, "signer is not authority")
	ErrSameAuthority    = errors.Register(ModuleName, 2, "provided authority is the current authority")
	ErrInvalidMessage   = errors.Register(ModuleName, 3, "message is invalid")
)
