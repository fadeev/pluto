package keeper

import (
	"github.com/fadeev/pluto/x/pluto/types"
)

var _ types.QueryServer = Keeper{}
