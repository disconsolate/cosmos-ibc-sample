package keeper_test

import (
	"testing"

	testkeeper "github.com/disconsolate/cosmos-ibc-sample/testutil/keeper"
	"github.com/disconsolate/cosmos-ibc-sample/x/cosmosibcsample/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmosibcsampleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
