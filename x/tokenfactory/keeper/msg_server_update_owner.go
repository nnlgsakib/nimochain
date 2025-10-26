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

func (k msgServer) UpdateOwner(ctx context.Context, msg *types.MsgUpdateOwner) (*types.MsgUpdateOwnerResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator address: %s", err))
	}

	// Validate new owner address
	if _, err := k.addressCodec.StringToBytes(msg.NewOwner); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid new owner address: %s", err))
	}

	// Get the denom to check ownership
	denom, err := k.Denom.Get(ctx, msg.Denom)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "denom not found")
		}
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Check if the creator is the current owner of the denom
	if msg.Creator != denom.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the current owner can update ownership")
	}

	// Update the owner
	denom.Owner = msg.NewOwner
	if err := k.Denom.Set(ctx, msg.Denom, denom); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update denom owner")
	}

	return &types.MsgUpdateOwnerResponse{}, nil
}
