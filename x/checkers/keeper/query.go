package keeper

import (
	"github.com/alice/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
