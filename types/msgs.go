// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

func NewMsgExecute(signer string, msgs []sdk.Msg) *MsgExecute {
	rawMsgs, err := tx.SetMsgs(msgs)
	if err != nil {
		panic(err)
	}

	return &MsgExecute{
		Signer:   signer,
		Messages: rawMsgs,
	}
}

func (exec *MsgExecute) GetMessages(cdc codec.Codec) ([]sdk.Msg, error) {
	var msgs []sdk.Msg

	for _, msgAny := range exec.Messages {
		var msg sdk.Msg
		err := cdc.UnpackAny(msgAny, &msg)
		if err != nil {
			return nil, err
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}
