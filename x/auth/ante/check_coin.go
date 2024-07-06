package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// CheckTURAFeeDecorator checks that all fees are paid in TURA.
type CheckTURAFeeDecorator struct{}

// NewCheckTURAFeeDecorator creates a new CheckTURAFeeDecorator.
func NewCheckTURAFeeDecorator() CheckTURAFeeDecorator {
	return CheckTURAFeeDecorator{}
}

// AnteHandle checks that all fees are paid in TURA.
func (ctfd CheckTURAFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	fee := feeTx.GetFee()

	// Check that all fees are paid in TURA
	for _, coin := range fee {
		if coin.Denom != "TURA" {
			return ctx, errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "fees must be paid in TURA")
		}
	}

	// Call next AnteHandler if fees are valid
	return next(ctx, tx, simulate)
}
