package checkers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sidrisov/checkers/x/checkers/keeper"
	"github.com/sidrisov/checkers/x/checkers/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetSystemInfo(ctx, genState.SystemInfo)
	// Set all the storedGame
	for _, elem := range genState.StoredGameList {
		k.SetStoredGame(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all systemInfo
	systemInfo, found := k.GetSystemInfo(ctx)
	if found {
		genesis.SystemInfo = systemInfo
	}
	genesis.StoredGameList = k.GetAllStoredGame(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
