package keeper

import (
	"encoding/binary"
	"github.com/MinterTeam/mhub/chain/x/minter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func InitGenesis(ctx sdk.Context, keeper Keeper, data types.GenesisState) {
	keeper.setParams(ctx, data.Params)
	ctx.KVStore(keeper.storeKey).Set(types.MinterNonce, sdk.Uint64ToBigEndian(data.StartMinterNonce))
}

func ExportGenesis(ctx sdk.Context, k Keeper) types.GenesisState {
	p := k.GetParams(ctx)

	bz := ctx.KVStore(k.storeKey).Get(types.MinterNonce)
	var startMinterNonce uint64 = 1
	if bz != nil {
		startMinterNonce = binary.BigEndian.Uint64(bz)
	}

	var addresses []*types.MsgSetMinterAddress

	k.StakingKeeper.IterateValidators(ctx, func(index int64, validator stakingtypes.ValidatorI) (stop bool) {
		addresses = append(addresses, &types.MsgSetMinterAddress{
			Address:   k.GetMinterAddress(ctx, sdk.AccAddress(validator.GetOperator())),
			Validator: validator.GetOperator().String(),
		})
		return false
	})

	return types.GenesisState{
		Params:           &p,
		StartMinterNonce: startMinterNonce,
		MinterAddresses:  addresses,
	}
}
