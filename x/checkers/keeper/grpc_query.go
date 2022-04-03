package keeper

import (
	"github.com/gupnik/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
