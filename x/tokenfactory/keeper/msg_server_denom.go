package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"nimo-chain/x/tokenfactory/types"
)

func (k msgServer) CreateDenom(ctx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Owner); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value already exists
	_, err := k.Denom.Get(ctx, msg.Denom)
	if err == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        msg.Description,
		Ticker:             msg.Ticker,
		Precision:          msg.Precision,
		Url:                msg.Url,
		MaxSupply:          msg.MaxSupply,
		Supply:             0, // Initial supply is 0
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	if err := k.Denom.Set(ctx, msg.Denom, denom); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set denom")
	}

	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(ctx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Owner); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Denom.Get(ctx, msg.Denom)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Owner != val.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        msg.Description,
		Ticker:             val.Ticker,    // Keep original ticker
		Precision:          val.Precision, // Keep original precision
		Url:                msg.Url,
		MaxSupply:          msg.MaxSupply,
		Supply:             val.Supply, // Keep current supply
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	if err := k.Denom.Set(ctx, msg.Denom, denom); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update denom")
	}

	return &types.MsgUpdateDenomResponse{}, nil
}

func (k msgServer) DeleteDenom(ctx context.Context, msg *types.MsgDeleteDenom) (*types.MsgDeleteDenomResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator address: %s", err))
	}

	// Get the denom to check ownership
	denom, err := k.Denom.Get(ctx, msg.Denom)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "denom not found")
		}
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Check if the creator is the owner of the denom
	if msg.Creator != denom.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the owner can delete the denom")
	}

	// Check if there are any tokens in circulation
	if denom.Supply > 0 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "cannot delete denom with tokens in circulation")
	}

	// Delete the denom
	if err := k.Denom.Remove(ctx, msg.Denom); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to delete denom")
	}

	return &types.MsgDeleteDenomResponse{}, nil
}
