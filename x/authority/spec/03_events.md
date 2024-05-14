# Events

## OwnershipTransferStarted

This event is emitted when an ownership transfer is started.

```shell
- attributes:
  - index: true
    key: previous_owner
    value: '"noble1signer"'
  - index: true
    key: new_owner
    value: '"noble1owner"'
  - index: true
    key: msg_index
    value: "0"
  type: noble.authority.v1.OwnershipTransferStarted
```

This event is emitted by the following transactions:

- [`noble.authority.v1.MsgTransferOwnership`](./02_messages.md#transfer-ownership)

## OwnershipTransferred

This event is emitted when an ownership transfer is finalized.

```shell
- attributes:
  - index: true
    key: previous_owner
    value: '"noble1signer"'
  - index: true
    key: new_owner
    value: '"noble1owner"'
  - index: true
    key: msg_index
    value: "0"
  type: noble.authority.v1.OwnershipTransferStarted
```

This event is emitted by the following transactions:

- [`noble.authority.v1.MsgAcceptOwnership`](./02_messages.md#accept-ownership)
