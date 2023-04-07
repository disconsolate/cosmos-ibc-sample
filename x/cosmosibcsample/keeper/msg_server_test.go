package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/disconsolate/cosmos-ibc-sample/testutil/keeper"
	"github.com/disconsolate/cosmos-ibc-sample/x/cosmosibcsample/keeper"
	"github.com/disconsolate/cosmos-ibc-sample/x/cosmosibcsample/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmosibcsampleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
