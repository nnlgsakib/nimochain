package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"nimo-chain/x/tokenfactory/types"
)

func (k msgServer) MintAndSendTokens(ctx context.Context, msg *types.MsgMintAndSendTokens) (*types.MsgMintAndSendTokensResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator address: %s", err))
	}

	// Get the denom to check ownership and update supply
	denom, err := k.Denom.Get(ctx, msg.Denom)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "denom not found")
		}
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Check if the creator is the owner of the denom
	if msg.Creator != denom.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the owner can mint tokens")
	}

	// Check if minting would exceed max supply
	newSupply := denom.Supply + msg.Amount
	if newSupply > denom.MaxSupply {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "minting would exceed max supply")
	}

	// Create coins to mint
	coins := sdk.NewCoins(sdk.NewInt64Coin(msg.Denom, msg.Amount))

	// Mint coins to the module account
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, fmt.Sprintf("failed to mint coins: %s", err))
	}

	// Send coins from module to recipient
	recipientAddr, err := k.addressCodec.StringToBytes(msg.Recipient)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid recipient address: %s", err))
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipientAddr, coins); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, fmt.Sprintf("failed to send coins: %s", err))
	}

	// Update the denom supply
	denom.Supply = newSupply
	if err := k.Denom.Set(ctx, msg.Denom, denom); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update denom supply")
	}

	return &types.MsgMintAndSendTokensResponse{}, nil
}
