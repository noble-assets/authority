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
    signer: noble1authority
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
```

### Arguments

- `messages` — An array of encoded messages to execute.

### Requirements

- Signer must be the current underlying [`authority`](./01_state.md#authority).

### State Changes

This message doesn't affect any `x/authority` state.

However, the executed messages may contain state changes.

### Events Emitted

This message emits no events.

## Update Authority

`noble.authority.v1.MsgUpdateAuthority`

A message that updates the underlying authority of this module.

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
  - '@type': /noble.authority.v1.MsgUpdateAuthority
    new_authority: noble1demo
    signer: noble1authority
  non_critical_extension_options: []
  timeout_height: "0"
signatures: []
```

### Arguments

- `new_authority` — The Noble address to transfer underlying authority to.

### Requirements

- Signer must be the current underlying [`authority`](./01_state.md#authority).

### State Changes

- [`authority`](./01_state.md#authority)

### Events Emitted

This message emits no events.