package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ValidateBasic performs basic validation for MsgCreateDenom
func (msg *MsgCreateDenom) ValidateBasic() error {
	if msg == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "message cannot be nil")
	}
	
	if msg.Owner == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "owner cannot be empty")
	}
	
	if _, err := sdk.AccAddressFromBech32(msg.Owner); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid owner address: %s", err))
	}
	
	if msg.Denom == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "denom cannot be empty")
	}
	
	if msg.MaxSupply <= 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "max supply must be positive")
	}
	
	return nil
}

// ValidateBasic performs basic validation for MsgUpdateDenom
func (msg *MsgUpdateDenom) ValidateBasic() error {
	if msg == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "message cannot be nil")
	}
	
	if msg.Owner == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "owner cannot be empty")
	}
	
	if _, err := sdk.AccAddressFromBech32(msg.Owner); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid owner address: %s", err))
	}
	
	if msg.Denom == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "denom cannot be empty")
	}
	
	if msg.MaxSupply <= 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "max supply must be positive")
	}
	
	return nil
}

// ValidateBasic performs basic validation for MsgDeleteDenom
func (msg *MsgDeleteDenom) ValidateBasic() error {
	if msg == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "message cannot be nil")
	}
	
	if msg.Creator == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "creator cannot be empty")
	}
	
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator address: %s", err))
	}
	
	if msg.Denom == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "denom cannot be empty")
	}
	
	return nil
}

// ValidateBasic performs basic validation for MsgMintAndSendTokens
func (msg *MsgMintAndSendTokens) ValidateBasic() error {
	if msg == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "message cannot be nil")
	}
	
	if msg.Creator == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "creator cannot be empty")
	}
	
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator address: %s", err))
	}
	
	if msg.Recipient == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "recipient cannot be empty")
	}
	
	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid recipient address: %s", err))
	}
	
	if msg.Denom == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "denom cannot be empty")
	}
	
	if msg.Amount <= 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount must be positive")
	}
	
	return nil
}

// ValidateBasic performs basic validation for MsgUpdateOwner
func (msg *MsgUpdateOwner) ValidateBasic() error {
	if msg == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "message cannot be nil")
	}
	
	if msg.Creator == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "creator cannot be empty")
	}
	
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator address: %s", err))
	}
	
	if msg.NewOwner == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "new owner cannot be empty")
	}
	
	if _, err := sdk.AccAddressFromBech32(msg.NewOwner); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid new owner address: %s", err))
	}
	
	if msg.Denom == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "denom cannot be empty")
	}
	
	return nil
}

// ValidateBasic performs basic validation for MsgUpdateParams
func (msg *MsgUpdateParams) ValidateBasic() error {
	if msg == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "message cannot be nil")
	}
	
	if msg.Authority == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "authority cannot be empty")
	}
	
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid authority address: %s", err))
	}
	
	return msg.Params.Validate()
}