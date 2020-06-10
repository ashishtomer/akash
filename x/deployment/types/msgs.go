package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	msgTypeCreateDeployment = "create-deployment"
	msgTypeUpdateDeployment = "update-deployment"
	msgTypeCloseDeployment  = "close-deployment"
)

// MsgCreateDeployment defines an SDK message for creating deployment
type MsgCreateDeployment struct {
	ID DeploymentID `json:"id"`
	// Version []byte      `json:"version"`
	Groups []GroupSpec `json:"groups"`
}

// Route implements the sdk.Msg interface
func (msg MsgCreateDeployment) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgCreateDeployment) Type() string { return msgTypeCreateDeployment }

// GetSignBytes encodes the message for signing
func (msg MsgCreateDeployment) GetSignBytes() []byte {
	return sdk.MustSortJSON(cdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateDeployment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ID.Owner}
}

// ValidateBasic does basic validation like check owner and groups length
func (msg MsgCreateDeployment) ValidateBasic() error {
	if err := msg.ID.Validate(); err != nil {
		return err
	}
	if len(msg.Groups) == 0 {
		return ErrEmptyGroups
	}
	// TODO: version
	return nil
}

// MsgUpdateDeployment defines an SDK message for updating deployment
type MsgUpdateDeployment struct {
	ID      DeploymentID
	Version sdk.AccAddress
}

// Route implements the sdk.Msg interface
func (msg MsgUpdateDeployment) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgUpdateDeployment) Type() string { return msgTypeUpdateDeployment }

// ValidateBasic does basic validation
func (msg MsgUpdateDeployment) ValidateBasic() error {
	if err := msg.ID.Validate(); err != nil {
		return err
	}

	if msg.Version.Empty() {
		return ErrEmptyVersion
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUpdateDeployment) GetSignBytes() []byte {
	return sdk.MustSortJSON(cdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUpdateDeployment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ID.Owner}
}

// MsgCloseDeployment defines an SDK message for closing deployment
type MsgCloseDeployment struct {
	ID DeploymentID
}

// Route implements the sdk.Msg interface
func (msg MsgCloseDeployment) Route() string { return RouterKey }

// Type implements the sdk.Msg interface
func (msg MsgCloseDeployment) Type() string { return msgTypeCloseDeployment }

// ValidateBasic does basic validation with deployment details
func (msg MsgCloseDeployment) ValidateBasic() error {
	if err := msg.ID.Validate(); err != nil {
		return err
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCloseDeployment) GetSignBytes() []byte {
	return sdk.MustSortJSON(cdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCloseDeployment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ID.Owner}
}
