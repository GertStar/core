package hard

import (
	"github.com/Gert-Chain/core/x/hard/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker updates interest rates
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.ApplyInterestRateUpdates(ctx)
}