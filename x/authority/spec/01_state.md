# State

## Authority

The authority field is of type bytes, specifically a Noble address, stored via an [`collections.Item`][item].
It is used to store the current underlying authority address of this module.

```go
var AuthorityKey = []byte("authority")
```

It is updated by the following messages:

- [`noble.authority.v1.UpdateAuthority`](./02_messages.md#update-authority)

[item]: https://docs.cosmos.network/main/build/packages/collections#item
