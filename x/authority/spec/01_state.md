# State

## Owner

The owner field is of type string, specifically a Noble address, stored via an [`collections.Item`][item].
It is used to store the current owner of this module.

```go
var OwnerKey = []byte("owner")
```

It is updated by the following messages:

- [`noble.authority.v1.AcceptOwnership`](./02_messages.md#accept-ownership)

## Pending Owner

The pending owner field is of type string, specifically a Noble address, stored via an [`collections.Item`][item].
It is used to store the current pending owner of this module.

```go
var PendingOwnerKey = []byte("pending_owner")
```

It is updated by the following messages:

- [`noble.authority.v1.TransferOwnership`](./02_messages.md#transfer-ownerhsip)
- [`noble.authority.v1.AcceptOwnership`](./02_messages.md#accept-ownership)

[item]: https://docs.cosmos.network/main/build/packages/collections#item
