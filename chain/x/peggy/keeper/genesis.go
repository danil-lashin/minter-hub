package keeper

import (
	"github.com/MinterTeam/mhub/chain/x/peggy/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis starts a chain from a genesis state
func InitGenesis(ctx sdk.Context, k Keeper, data types.GenesisState) {
	k.SetParams(ctx, *data.Params)
	// reset valsets in state
	for _, vs := range data.Valsets {
		// TODO: block height?
		k.StoreValsetUnsafe(ctx, vs)
	}

	// reset valset confirmations in state
	for _, conf := range data.ValsetConfirms {
		k.SetValsetConfirm(ctx, *conf)
	}

	// reset batches in state
	for _, batch := range data.Batches {
		// TODO: block height?
		k.StoreBatchUnsafe(ctx, batch)
	}

	// reset batch confirmations in state
	for _, conf := range data.BatchConfirms {
		k.SetBatchConfirm(ctx, &conf)
	}

	// reset attestations in state
	for _, att := range data.Attestations {
		// TODO: block height?
		k.SetAttestationUnsafe(ctx, &att)
	}
}

// ExportGenesis exports all the state needed to restart the chain
// from the current state of the chain
func ExportGenesis(ctx sdk.Context, k Keeper) types.GenesisState {
	var (
		p = k.GetParams(ctx)
	)

	return types.GenesisState{
		Params: &p,
	}
}
