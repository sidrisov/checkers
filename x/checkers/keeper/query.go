package keeper

import (
	"github.com/sidrisov/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
