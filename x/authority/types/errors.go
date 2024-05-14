package types

import "cosmossdk.io/errors"

var (
	ErrInvalidOwner        = errors.Register(ModuleName, 1, "signer is not authority")
	ErrSameOwner           = errors.Register(ModuleName, 2, "provided owner is the current owner")
	ErrNoPendingOwner      = errors.Register(ModuleName, 3, "no pending owner found")
	ErrInvalidPendingOwner = errors.Register(ModuleName, 4, "signer is not pending owner")
	ErrInvalidMessage      = errors.Register(ModuleName, 5, "message is invalid")
)
