// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2025, NASD Inc. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN "AS IS" BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package cli

import (
	"fmt"
	"strconv"

	"cosmossdk.io/math"
	"cosmossdk.io/x/upgrade/plan"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	ratelimittypes "github.com/cosmos/ibc-apps/modules/rate-limiting/v8/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	"github.com/noble-assets/authority/types"
	"github.com/spf13/cobra"
)

const (
	FlagUpgradeHeight      = "upgrade-height"
	FlagUpgradeInfo        = "upgrade-info"
	FlagNoValidate         = "no-validate"
	FlagNoChecksumRequired = "no-checksum-required"
	FlagDaemonName         = "daemon-name"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: fmt.Sprintf("Transactions commands for the %s module", types.ModuleName),
	}

	cmd.AddCommand(NewCmdExecute())
	cmd.AddCommand(NewCmdSoftwareUpgrade())
	cmd.AddCommand(NewCmdCancelSoftwareUpgrade())
	cmd.AddCommand(NewCmdRecoverClient())
	cmd.AddCommand(NewCmdAddRateLimit())
	cmd.AddCommand(NewCmdUpdateRateLimit())
	cmd.AddCommand(NewCmdRemoveRateLimit())
	cmd.AddCommand(NewCmdResetRateLimit())

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

/////////////////////////////
//    Software Upgrades    //
/////////////////////////////

// NewCmdSoftwareUpgrade is a helper for scheduling a software upgrade.
//
// This command has been adapted from the Cosmos SDK implementation.
// https://github.com/cosmos/cosmos-sdk/blob/x/upgrade/v0.1.4/x/upgrade/client/cli/tx.go#L46-L133
func NewCmdSoftwareUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "software-upgrade <name> [--upgrade-height <height>] [--upgrade-info <info>] [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "Helper for scheduling a software upgrade",
		Long: "Helper for scheduling a software upgrade.\n\n" +
			"You can additionally include upgrade info via a flag to reference pre-built binaries, documentation, etc.\n" +
			"https://docs.cosmos.network/main/build/tooling/cosmovisor",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			p, err := parsePlan(cmd.Flags(), name)
			if err != nil {
				return err
			}

			noValidate, err := cmd.Flags().GetBool(FlagNoValidate)
			if err != nil {
				return err
			}

			if !noValidate {
				daemonName, err := cmd.Flags().GetString(FlagDaemonName)
				if err != nil {
					return err
				}

				noChecksum, err := cmd.Flags().GetBool(FlagNoChecksumRequired)
				if err != nil {
					return err
				}

				var planInfo *plan.Info
				if planInfo, err = plan.ParseInfo(p.Info, plan.ParseOptionEnforceChecksum(!noChecksum)); err != nil {
					return err
				}

				if err = planInfo.ValidateFull(daemonName); err != nil {
					return err
				}
			}

			msgExecute := types.NewMsgExecute(
				clientCtx.FromAddress.String(),
				[]sdk.Msg{
					&upgradetypes.MsgSoftwareUpgrade{
						Authority: types.ModuleAddress.String(),
						Plan:      p,
					},
				})

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgExecute)
		},
	}

	cmd.Flags().Int64(FlagUpgradeHeight, 0, "The height at which the upgrade must happen")
	cmd.Flags().String(FlagUpgradeInfo, "", "Info for the upgrade plan such as new version download urls, etc.")
	cmd.Flags().Bool(FlagNoValidate, false, "Skip validation of the upgrade info (dangerous!)")
	cmd.Flags().Bool(FlagNoChecksumRequired, false, "Skip requirement of checksums for binaries in the upgrade info")
	cmd.Flags().String(FlagDaemonName, getDefaultDaemonName(), "The name of the executable being upgraded (for upgrade-info validation). Default is the DAEMON_NAME env var if set, or else this executable")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCmdCancelSoftwareUpgrade is a helper for canceling the currently scheduled software upgrade.
//
// This command has been adapted from the Cosmos SDK implementation.
// https://github.com/cosmos/cosmos-sdk/blob/x/upgrade/v0.1.4/x/upgrade/client/cli/tx.go#L135-L182
func NewCmdCancelSoftwareUpgrade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-software-upgrade [flags]",
		Args:  cobra.NoArgs,
		Short: "Helper for canceling the currently scheduled software upgrade",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msgExecute := types.NewMsgExecute(
				clientCtx.FromAddress.String(),
				[]sdk.Msg{
					&upgradetypes.MsgCancelUpgrade{
						Authority: types.ModuleAddress.String(),
					},
				})

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgExecute)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

///////////////////////////////
//    IBC Client Recovery    //
///////////////////////////////

// NewCmdRecoverClient is a helper for recovering an expired client.
//
// This command has been adapted from the IBC-Go implementation.
// https://github.com/cosmos/ibc-go/blob/v8.5.2/modules/core/02-client/client/cli/tx.go#L248-L303
func NewCmdRecoverClient() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recover-client [subject-client-id] [substitute-client-id] [flags]",
		Args:  cobra.ExactArgs(2),
		Short: "Helper for recovering an expired client",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subjectClientID, substituteClientID := args[0], args[1]

			authority := types.ModuleAddress.String()

			msg := clienttypes.NewMsgRecoverClient(authority, subjectClientID, substituteClientID)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("error validating %T: %w", clienttypes.MsgRecoverClient{}, err)
			}

			msgExecute := types.NewMsgExecute(clientCtx.FromAddress.String(), []sdk.Msg{msg})

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgExecute)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

///////////////////////////
//    IBC Rate Limits    //
///////////////////////////

// NewCmdAddRateLimit is a helper for adding a new rate limit to a denom.
func NewCmdAddRateLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add-rate-limit [denom] [channel-id] [max-percent-send] [max-percent-receive] [duration-hours] [flags]",
		Short:   "Add a new rate limit to a denom.",
		Args:    cobra.ExactArgs(5),
		Example: "simd tx authority add-rate-limit utoken channel-0 10 10 24",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom, channelID, maxPercentSend, maxPercentReceive, duration := args[0], args[1], args[2], args[3], args[4]

			maxPercentSendInt, ok := math.NewIntFromString(maxPercentSend)
			if !ok {
				return fmt.Errorf("invalid max percent send: %s", maxPercentSend)
			}

			maxPercentReceiveInt, ok := math.NewIntFromString(maxPercentReceive)
			if !ok {
				return fmt.Errorf("invalid max percent receive: %s", maxPercentReceive)
			}

			durationInt, err := strconv.ParseUint(duration, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid duration: %s, err : %w", duration, err)
			}

			msg := ratelimittypes.NewMsgAddRateLimit(denom, channelID, maxPercentSendInt, maxPercentReceiveInt, durationInt)
			msg.Authority = types.ModuleAddress.String()

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("error validating %T: %w", msg, err)
			}

			msgExecute := types.NewMsgExecute(clientCtx.FromAddress.String(), []sdk.Msg{msg})

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgExecute)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCmdUpdateRateLimit is a helper to update an already rate limited to a denom.
func NewCmdUpdateRateLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update-rate-limit [denom] [channel-id] [max-percent-send] [max-percent-receive] [duration-hours] [flags]",
		Short:   "Update an already rate limited denom.",
		Args:    cobra.ExactArgs(5),
		Example: "simd tx authority update-rate-limit utoken channel-0 10 10 24",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom, channelID, maxPercentSend, maxPercentReceive, duration := args[0], args[1], args[2], args[3], args[4]

			maxPercentSendInt, ok := math.NewIntFromString(maxPercentSend)
			if !ok {
				return fmt.Errorf("invalid max percent send: %s", maxPercentSend)
			}

			maxPercentReceiveInt, ok := math.NewIntFromString(maxPercentReceive)
			if !ok {
				return fmt.Errorf("invalid max percent receive: %s", maxPercentReceive)
			}

			durationInt, err := strconv.ParseUint(duration, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid duration: %s, err : %w", duration, err)
			}

			msg := ratelimittypes.NewMsgUpdateRateLimit(denom, channelID, maxPercentSendInt, maxPercentReceiveInt, durationInt)
			msg.Authority = types.ModuleAddress.String()

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("error validating %T: %w", msg, err)
			}

			msgExecute := types.NewMsgExecute(clientCtx.FromAddress.String(), []sdk.Msg{msg})

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgExecute)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCmdRemoveRateLimit is a helper to remove the rate limit from a denom.
func NewCmdRemoveRateLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove-rate-limit [denom] [channel-id] [flags]",
		Short:   "Remove a rate limit from a denom.",
		Args:    cobra.ExactArgs(2),
		Example: "simd tx authority remove-rate-limit utoken channel-0",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom, channelID := args[0], args[1]

			msg := ratelimittypes.NewMsgRemoveRateLimit(denom, channelID)
			msg.Authority = types.ModuleAddress.String()

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("error validating %T: %w", msg, err)
			}

			msgExecute := types.NewMsgExecute(clientCtx.FromAddress.String(), []sdk.Msg{msg})

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgExecute)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCmdResetRateLimit is a helper to reset the rate limit for a denom.
func NewCmdResetRateLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "reset-rate-limit [denom] [channel-id] [flags]",
		Short:   "Reset a rate limit for a denom.",
		Args:    cobra.ExactArgs(2),
		Example: "simd tx authority reset-rate-limit utoken channel-0",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom, channelID := args[0], args[1]

			msg := ratelimittypes.NewMsgResetRateLimit(denom, channelID)
			msg.Authority = types.ModuleAddress.String()

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("error validating %T: %w", msg, err)
			}

			msgExecute := types.NewMsgExecute(clientCtx.FromAddress.String(), []sdk.Msg{msg})

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgExecute)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
