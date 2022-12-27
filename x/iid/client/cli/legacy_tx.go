package cli

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	libfury "github.com/furya-official/fury-blockchain/lib/fury"
	"github.com/furya-official/fury-blockchain/lib/legacydid"
	"github.com/furya-official/fury-blockchain/x/iid/types"
	"github.com/spf13/cobra"
)

func NewCreateIidDocumentFormLegacyDidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-iid-from-legacy-did [did]",
		Short:   "create decentralized did (did) document from legacy did",
		Example: "creates a did document for users",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			furyDid, err := legacydid.UnmarshalFuryDid(args[0])
			if err != nil {
				return err
			}

			// chaincmd.Flags().GetString(flags.FlagChainID)
			// if err != nil {
			// 	return err
			// }
			// did
			did := types.DID(furyDid.Did)
			// verification
			// signer := clientCtx.GetFromAddress()
			// pubkey

			pubKey := furyDid.VerifyKey

			clientCtx = clientCtx.WithFromAddress(furyDid.Address())

			// understand the vmType

			auth := types.NewVerification(
				types.NewVerificationMethod(
					furyDid.Did,
					did,
					types.NewPublicKeyMultibase(base58.Decode(pubKey), types.DIDVMethodTypeEd25519VerificationKey2018),
				),
				[]string{types.Authentication},
				nil,
			)
			// create the message
			msg := types.NewMsgCreateIidDocument(
				did.String(),
				types.Verifications{auth},
				types.Services{},
				types.AccordedRights{},
				types.LinkedResources{},
				types.LinkedEntities{},
				furyDid.Address().String(),
				types.Contexts{},
			)
			// validate
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// execute
			return libfury.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), furyDid, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
