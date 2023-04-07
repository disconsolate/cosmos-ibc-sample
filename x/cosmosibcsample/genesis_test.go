package cosmosibcsample_test

import (
	"testing"

	keepertest "github.com/disconsolate/cosmos-ibc-sample/testutil/keeper"
	"github.com/disconsolate/cosmos-ibc-sample/testutil/nullify"
	"github.com/disconsolate/cosmos-ibc-sample/x/cosmosibcsample"
	"github.com/disconsolate/cosmos-ibc-sample/x/cosmosibcsample/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmosibcsampleKeeper(t)
	cosmosibcsample.InitGenesis(ctx, *k, genesisState)
	got := cosmosibcsample.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
