# Messages

## Execute

`noble.authority.v1.MsgExecute`

A message to execute arbitrary messages on behalf of the `x/authority` module.
If other module authorities are configured to this module, this can be used to
execute admin messages such as updating params, software ugprades, etc.

```shell
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /noble.authority.v1.MsgExecute
    messages: [...]
    signer: noble1owner
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
```

### Arguments

- `messages` — An array of encoded messages to execute.

### Requirements

- Signer must be the current [`owner`](./01_state.md#owner).

### State Changes

This message doesn't affect any `x/authority` state.

However, the executed messages may contain state changes.

### Events Emitted

This message emits no events.

## Transfer Ownership

`noble.authority.v1.MsgTransferOwnership`

A message that initiates an ownership transfer of this module.

```shell
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /noble.authority.v1.MsgTransferOwnership
    new_owner: noble1owner
    signer: noble1signer
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
```

### Arguments

- `new_owner` — The Noble address to transfer ownership to.

### Requirements

- Signer must be the current [`owner`](./01_state.md#owner).

### State Changes

- [`pending_owner`](./01_state.md#pending-owner)

### Events Emitted

- [`noble.authority.v1.OwnershipTransferStarted`](./03_events.md#ownershiptransferstarted)

## Accept Ownership

`noble.authority.v1.MsgAcceptOwnership`

A message that finalizes an ownership transfer of this module.

```shell
auth_info:
  fee:
    amount: []
    gas_limit: "200000"
    granter: ""
    payer: ""
  signer_infos: []
  tip: null
body:
  extension_options: []
  memo: ""
  messages:
  - '@type': /noble.authority.v1.MsgAcceptOwnership
    signer: noble1signer
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
```

### Arguments

This message takes no arguments.

### Requirements

- Signer must be the current [`pending_owner`](./01_state.md#pending-owner).

### State Changes

- [`owner`](./01_state.md#owner)
- [`pending_authority`](./01_state.md#pending-owner)

### Events Emitted

- [`noble.authority.v1.OwnershipTransferred`](./03_events.md#ownershiptransferred)
