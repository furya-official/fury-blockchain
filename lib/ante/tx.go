package ante

// import (
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	iidkeeper "github.com/furya-official/fury-blockchain/x/iid/keeper"
// )

// type FeePayer struct {
// 	feePayerAccount  authtypes.AccountI
// 	recipientAddress sdk.AccAddress
// }

// func NewFeePayer(feePayerAccount authtypes.AccountI, recipientAddress sdk.AccAddress) FeePayer {
// 	return FeePayer{
// 		feePayerAccount:  feePayerAccount,
// 		recipientAddress: recipientAddress,
// 	}
// }

// func (fp *FeePayer) GetFeePayerAccount() authtypes.AccountI { return fp.feePayerAccount }
// func (fp *FeePayer) GetRecipientAddress() sdk.AccAddress    { return fp.recipientAddress }

// type FuryFeeTxMsg interface {
// 	FeePayerFromIid(ctx sdk.Context, accountKeeper authante.AccountKeeper, iidKeeper iidkeeper.Keeper) (FeePayer, error)
// }

// type FuryFeeTx struct {
// 	sdk.FeeTx
// }

// func (tx *FuryFeeTx) GetFeePayerMsgs() []FuryFeeTxMsg {
// 	var msgs []FuryFeeTxMsg

// 	for _, msg := range tx.GetMsgs() {
// 		if msg, ok := msg.(FuryFeeTxMsg); ok {
// 			msgs = append(msgs, msg)
// 		}
// 	}

// 	return msgs
// }
