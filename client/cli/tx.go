// Copyright 2024 NASD Inc.
//
// Use of this source code is governed by a BSL-style
// license that can be found in the LICENSE file or at
// https://mariadb.com/bsl11.

package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/noble-assets/authority/types"
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: fmt.Sprintf("Transactions commands for the %s module", types.ModuleName),
	}

	cmd.AddCommand(NewCmdExecute())

	return cmd
}

func NewCmdExecute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute [path to underlying tx file]",
		Short: "Execute arbitrary messages as authority module",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			underlying, err := authclient.ReadTxFromFile(clientCtx, args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgExecute(clientCtx.FromAddress.String(), underlying.GetMsgs())

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
