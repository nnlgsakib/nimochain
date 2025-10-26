package tokenfactory

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"nimo-chain/testutil/sample"
	tokenfactorysimulation "nimo-chain/x/tokenfactory/simulation"
	"nimo-chain/x/tokenfactory/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	tokenfactoryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DenomMap: []types.Denom{{Owner: sample.AccAddress(),
			Denom: "0",
		}, {Owner: sample.AccAddress(),
			Denom: "1",
		}}}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&tokenfactoryGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	
	// CreateDenom operation
	const (
		opWeightMsgCreateDenom          = "op_weight_msg_create_denom"
		defaultWeightMsgCreateDenom int = 100
	)
	var weightMsgCreateDenom int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateDenom, &weightMsgCreateDenom, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDenom = defaultWeightMsgCreateDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDenom,
		tokenfactorysimulation.SimulateMsgCreateDenom(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	// UpdateDenom operation
	const (
		opWeightMsgUpdateDenom          = "op_weight_msg_update_denom"
		defaultWeightMsgUpdateDenom int = 100
	)
	var weightMsgUpdateDenom int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateDenom, &weightMsgUpdateDenom, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDenom = defaultWeightMsgUpdateDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDenom,
		tokenfactorysimulation.SimulateMsgUpdateDenom(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	// DeleteDenom operation
	const (
		opWeightMsgDeleteDenom          = "op_weight_msg_delete_denom"
		defaultWeightMsgDeleteDenom int = 100
	)
	var weightMsgDeleteDenom int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteDenom, &weightMsgDeleteDenom, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDenom = defaultWeightMsgDeleteDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDenom,
		tokenfactorysimulation.SimulateMsgDeleteDenom(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	// MintAndSendTokens operation
	const (
		opWeightMsgMintAndSendTokens          = "op_weight_msg_mint_and_send_tokens"
		defaultWeightMsgMintAndSendTokens int = 100
	)
	var weightMsgMintAndSendTokens int
	simState.AppParams.GetOrGenerate(opWeightMsgMintAndSendTokens, &weightMsgMintAndSendTokens, nil,
		func(_ *rand.Rand) {
			weightMsgMintAndSendTokens = defaultWeightMsgMintAndSendTokens
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintAndSendTokens,
		tokenfactorysimulation.SimulateMsgMintAndSendTokens(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	// UpdateOwner operation
	const (
		opWeightMsgUpdateOwner          = "op_weight_msg_update_owner"
		defaultWeightMsgUpdateOwner int = 100
	)
	var weightMsgUpdateOwner int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateOwner, &weightMsgUpdateOwner, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateOwner = defaultWeightMsgUpdateOwner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateOwner,
		tokenfactorysimulation.SimulateMsgUpdateOwner(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
