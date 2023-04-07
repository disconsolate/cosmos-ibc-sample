package keeper

import (
	"github.com/disconsolate/cosmos-ibc-sample/x/cosmosibcsample/types"
)

var _ types.QueryServer = Keeper{}
