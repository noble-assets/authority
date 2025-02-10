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

package e2e

import (
	"context"
	"encoding/json"
	"testing"

	upgradetypes "cosmossdk.io/x/upgrade/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/noble-assets/authority/types"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/relayer/rly"
	"github.com/strangelove-ventures/interchaintest/v8/testreporter"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

type Wrapper struct {
	chain *cosmos.CosmosChain
	gaia  *cosmos.CosmosChain

	owner        ibc.Wallet
	pendingOwner ibc.Wallet
}

func Suite(t *testing.T, wrapper *Wrapper, ibcEnabled bool) (ctx context.Context, logger *zap.Logger, execReporter *testreporter.RelayerExecReporter, relayer *rly.CosmosRelayer) {
	ctx = context.Background()
	logger = zaptest.NewLogger(t)
	reporter := testreporter.NewNopReporter()
	execReporter = reporter.RelayerExecReporter(t)
	client, network := interchaintest.DockerSetup(t)

	numValidators, numFullNodes := 1, 0

	specs := []*interchaintest.ChainSpec{
		{
			Name:          "authority",
			Version:       "local",
			NumValidators: &numValidators,
			NumFullNodes:  &numFullNodes,
			ChainConfig: ibc.ChainConfig{
				Type:    "cosmos",
				Name:    "authority",
				ChainID: "authority-1",
				Images: []ibc.DockerImage{
					{
						Repository: "noble-authority-simd",
						Version:    "local",
						UidGid:     "1025:1025",
					},
				},
				Bin:            "simd",
				Bech32Prefix:   "noble",
				Denom:          "uusdc",
				GasPrices:      "0.1uusdc",
				GasAdjustment:  5,
				TrustingPeriod: "504h",
				NoHostMount:    false,
				PreGenesis: func(cfg ibc.Chain) (err error) {
					validator := wrapper.chain.Validators[0]
					coins := sdk.NewCoins(sdk.NewInt64Coin("uusdc", 2_500_000))

					wrapper.owner, err = wrapper.chain.BuildWallet(ctx, "owner", "")
					if err != nil {
						return err
					}
					err = validator.AddGenesisAccount(ctx, wrapper.owner.FormattedAddress(), coins)
					if err != nil {
						return err
					}

					wrapper.pendingOwner, err = wrapper.chain.BuildWallet(ctx, "pending-owner", "")
					if err != nil {
						return err
					}
					err = validator.AddGenesisAccount(ctx, wrapper.pendingOwner.FormattedAddress(), coins)
					if err != nil {
						return err
					}

					return nil
				},
				ModifyGenesis: func(cfg ibc.ChainConfig, bz []byte) ([]byte, error) {
					changes := []cosmos.GenesisKV{
						cosmos.NewGenesisKV("app_state.authority.owner", wrapper.owner.FormattedAddress()),
					}

					return cosmos.ModifyGenesis(changes)(cfg, bz)
				},
			},
		},
	}
	if ibcEnabled {
		specs = append(specs, &interchaintest.ChainSpec{
			Name:          "gaia",
			Version:       "v16.0.0",
			NumValidators: &numValidators,
			NumFullNodes:  &numFullNodes,
			ChainConfig: ibc.ChainConfig{
				ChainID: "cosmoshub-4",
			},
		})
	}
	factory := interchaintest.NewBuiltinChainFactory(logger, specs)

	chains, err := factory.Chains(t.Name())
	require.NoError(t, err)

	noble := chains[0].(*cosmos.CosmosChain)
	wrapper.chain = noble

	interchain := interchaintest.NewInterchain().AddChain(noble)
	if ibcEnabled {
		relayer = interchaintest.NewBuiltinRelayerFactory(
			ibc.CosmosRly,
			logger,
		).Build(t, client, network).(*rly.CosmosRelayer)

		gaia := chains[1].(*cosmos.CosmosChain)
		wrapper.gaia = gaia

		interchain = interchain.
			AddChain(gaia).
			AddRelayer(relayer, "relayer").
			AddLink(interchaintest.InterchainLink{
				Chain1:  noble,
				Chain2:  gaia,
				Relayer: relayer,
				Path:    "transfer",
			})
	}
	require.NoError(t, interchain.Build(ctx, execReporter, interchaintest.InterchainBuildOptions{
		TestName:         t.Name(),
		Client:           client,
		NetworkID:        network,
		SkipPathCreation: true,
	}))

	t.Cleanup(func() {
		_ = interchain.Close()
	})

	return
}

//

func EnsureOwner(t *testing.T, wrapper Wrapper, ctx context.Context, owner string) {
	validator := wrapper.chain.Validators[0]

	raw, _, err := validator.ExecQuery(ctx, "authority", "owner")
	require.NoError(t, err)

	var res types.QueryOwnerResponse
	require.NoError(t, json.Unmarshal(raw, &res))

	require.Equal(t, owner, res.Owner)
}

func EnsurePendingOwner(t *testing.T, wrapper Wrapper, ctx context.Context, pendingOwner string) {
	validator := wrapper.chain.Validators[0]

	raw, _, err := validator.ExecQuery(ctx, "authority", "pending-owner")
	require.NoError(t, err)

	var res types.QueryPendingOwnerResponse
	require.NoError(t, json.Unmarshal(raw, &res))

	require.Equal(t, pendingOwner, res.PendingOwner)
}

func EnsureParams(t *testing.T, wrapper Wrapper, ctx context.Context, height int64) {
	validator := wrapper.chain.Validators[0]

	raw, _, err := validator.ExecQuery(ctx, "consensus", "params")
	require.NoError(t, err)

	var res consensustypes.QueryParamsResponse
	require.NoError(t, jsonpb.UnmarshalString(string(raw), &res))

	require.NotNil(t, res.Params)
	require.NotNil(t, res.Params.Abci)
	require.Equal(t, height, res.Params.Abci.VoteExtensionsEnableHeight)
}

func EnsureUpgrade(t *testing.T, wrapper Wrapper, ctx context.Context, name string, height int64) {
	validator := wrapper.chain.Validators[0]

	raw, _, err := validator.ExecQuery(ctx, "upgrade", "plan")
	require.NoError(t, err)

	var res upgradetypes.QueryCurrentPlanResponse
	require.NoError(t, jsonpb.UnmarshalString(string(raw), &res))

	if name == "" {
		require.Nil(t, res.Plan)
	} else {
		require.NotNil(t, res.Plan)
		require.Equal(t, name, res.Plan.Name)
		require.Equal(t, height, res.Plan.Height)
	}
}
