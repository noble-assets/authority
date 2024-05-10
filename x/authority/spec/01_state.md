# State

## Authority

The authority field is of type bytes, specifically a Noble address, stored via an [`collections.Item`][item].
It is used to store the current underlying authority address of this module.

```go
var AuthorityKey = []byte("authority")
```

It is updated by the following messages:

- [`noble.authority.v1.AcceptAuthority`](./02_messages.md#accept-authority)

## Pending Authority

The pending authority field is of type bytes, specifically a Noble address, stored via an [`collections.Item`][item].
It is used to store the current pending underlying authority address of this module.

```go
var PendingAuthorityKey = []byte("pending_authority")
```

It is updated by the following messages:

- [`noble.authority.v1.TransferAuthority`](./02_messages.md#transfer-authority)
- [`noble.authority.v1.AcceptAuthority`](./02_messages.md#accept-authority)

[item]: https://docs.cosmos.network/main/build/packages/collections#item
