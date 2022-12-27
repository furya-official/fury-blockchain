package cli

import (
	"encoding/json"

	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/furya-official/fury-blockchain/x/entity/types"
	"github.com/spf13/cobra"
)

func NewTxCmd() *cobra.Command {
	entityTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "entity transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	entityTxCmd.AddCommand(
		NewCmdCreateEntity(),
	)

	return entityTxCmd
}

// NewCmdSubmitUpgradeProposal implements a command handler for submitting a software upgrade proposal transaction.
func NewCmdUpdateEntityParamsProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-entity-params [nft_contract_code] [nft_minter_address] [flags]",
		Args:  cobra.ExactArgs(2),
		Short: "Submit a proposal to update entity params",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			codeId, err := strconv.ParseUint(args[0], 0, 64)
			if err != nil {
				return err
			}

			content := types.NewInitializeNftContract(codeId, args[1])

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(&content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")

	return cmd
}

func NewCmdCreateEntity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-entity [entity-iid]",
		Short: "Create a new EntityDoc",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			var msg types.MsgCreateEntity
			err := json.Unmarshal([]byte(args[0]), &msg)
			if err != nil {
				return err
			}

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg.OwnerAddress = clientCtx.GetFromAddress().String()

			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
			if err != nil {
				return err
			}

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// func NewCmdUpdateProjectStatus() *cobra.Command {
// cmd := &cobra.Command{
// 	Use:   "update-project-status [sender-did] [status] [fury-did]",
// 	Short: "Update the status of a project signed by the furyDid of the project",
// 	Args:  cobra.ExactArgs(3),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		senderDid := args[0]
// 		status := args[1]
// 		furyDid, err := didtypes.UnmarshalFuryDid(args[2])
// 		if err != nil {
// 			return err
// 		}

// 		projectStatus := types.ProjectStatus(status)
// 		if projectStatus != types.CreatedProject &&
// 			projectStatus != types.PendingStatus &&
// 			projectStatus != types.FundedStatus &&
// 			projectStatus != types.StartedStatus &&
// 			projectStatus != types.StoppedStatus &&
// 			projectStatus != types.PaidoutStatus {
// 			return errors.New("The status must be one of 'CREATED', " +
// 				"'PENDING', 'FUNDED', 'STARTED', 'STOPPED' or 'PAIDOUT'")
// 		}

// 		updateProjectStatusDoc := types.NewUpdateProjectStatusDoc(
// 			projectStatus, "")

// 		clientCtx, err := client.GetClientTxContext(cmd)
// 		if err != nil {
// 			return err
// 		}
// 		clientCtx = clientCtx.WithFromAddress(furyDid.Address())

// 		msg := types.NewMsgUpdateProjectStatus(senderDid, updateProjectStatusDoc, furyDid.Did)
// 		err = msg.ValidateBasic()
// 		if err != nil {
// 			return err
// 		}

// 		return furytypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), furyDid, msg)
// 	},
// }

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

// func NewCmdCreateAgent() *cobra.Command {
// cmd := &cobra.Command{
// 	Use: "create-agent [tx-hash] [sender-did] [agent-did] " +
// 		"[role] [project-did]",
// 	Short: "Create a new agent on a project signed by the furyDid of the project",
// 	Args:  cobra.ExactArgs(5),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		txHash := args[0]
// 		senderDid := args[1]
// 		agentDid := args[2]
// 		role := args[3]
// 		if role != "SA" && role != "EA" && role != "IA" {
// 			return errors.New("The role must be one of 'SA', 'EA' or 'IA'")
// 		}

// 		createAgentDoc := types.NewCreateAgentDoc(agentDid, role)

// 		furyDid, err := didtypes.UnmarshalFuryDid(args[4])
// 		if err != nil {
// 			return err
// 		}

// 		clientCtx, err := client.GetClientTxContext(cmd)
// 		if err != nil {
// 			return err
// 		}
// 		clientCtx = clientCtx.WithFromAddress(furyDid.Address())

// 		msg := types.NewMsgCreateAgent(txHash, senderDid, createAgentDoc, furyDid.Did)
// 		err = msg.ValidateBasic()
// 		if err != nil {
// 			return err
// 		}

// 		return furytypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), furyDid, msg)
// 	},
// }

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

// func NewCmdUpdateAgent() *cobra.Command {
// cmd := &cobra.Command{
// 	Use: "update-agent [tx-hash] [sender-did] [agent-did] " +
// 		"[status] [fury-did]",
// 	Short: "Update the status of an agent on a project signed by the furyDid of the project",
// 	Args:  cobra.ExactArgs(6),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		txHash := args[0]
// 		senderDid := args[1]
// 		agentDid := args[2]
// 		agentRole := args[4]
// 		agentStatus := types.AgentStatus(args[3])
// 		if agentStatus != types.PendingAgent && agentStatus != types.ApprovedAgent && agentStatus != types.RevokedAgent {
// 			return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Revoked)")
// 		}

// 		updateAgentDoc := types.NewUpdateAgentDoc(
// 			agentDid, agentStatus, agentRole)

// 		furyDid, err := didtypes.UnmarshalFuryDid(args[5])
// 		if err != nil {
// 			return err
// 		}

// 		clientCtx, err := client.GetClientTxContext(cmd)
// 		if err != nil {
// 			return err
// 		}
// 		clientCtx = clientCtx.WithFromAddress(furyDid.Address())

// 		msg := types.NewMsgUpdateAgent(txHash, senderDid, updateAgentDoc, furyDid.Did)
// 		err = msg.ValidateBasic()
// 		if err != nil {
// 			return err
// 		}

// 		return furytypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), furyDid, msg)
// 	},
// }

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

// func NewCmdCreateClaim() *cobra.Command {
// cmd := &cobra.Command{
// 	Use:   "create-claim [tx-hash] [sender-did] [claim-id] [claim-template-id] [fury-did]",
// 	Short: "Create a new claim on a project signed by the furyDid of the project",
// 	Args:  cobra.ExactArgs(5),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		txHash := args[0]
// 		senderDid := args[1]
// 		claimId := args[2]
// 		claimTemplateId := args[3]
// 		createClaimDoc := types.NewCreateClaimDoc(claimId, claimTemplateId)

// 		furyDid, err := didtypes.UnmarshalFuryDid(args[4])
// 		if err != nil {
// 			return err
// 		}

// 		clientCtx, err := client.GetClientTxContext(cmd)
// 		if err != nil {
// 			return err
// 		}
// 		clientCtx = clientCtx.WithFromAddress(furyDid.Address())

// 		msg := types.NewMsgCreateClaim(txHash, senderDid, createClaimDoc, furyDid.Did)
// 		err = msg.ValidateBasic()
// 		if err != nil {
// 			return err
// 		}

// 		return furytypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), furyDid, msg)
// 	},
// }

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

// func NewCmdCreateEvaluation() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use: "create-evaluation [tx-hash] [sender-did] [claim-id] " +
// 			"[status] [fury-did]",
// 		Short: "Create a new claim evaluation on a project signed by the furyDid of the project",
// 		Args:  cobra.ExactArgs(5),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			txHash := args[0]
// 			senderDid := args[1]
// 			claimId := args[2]
// 			claimStatus := types.ClaimStatus(args[3])
// 			if claimStatus != types.PendingClaim && claimStatus != types.ApprovedClaim && claimStatus != types.RejectedClaim {
// 				return errors.New("The status must be one of '0' (Pending), '1' (Approved) or '2' (Rejected)")
// 			}

// 			createEvaluationDoc := types.NewCreateEvaluationDoc(
// 				claimId, claimStatus)

// 			furyDid, err := didtypes.UnmarshalFuryDid(args[4])
// 			if err != nil {
// 				return err
// 			}

// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}
// 			clientCtx = clientCtx.WithFromAddress(furyDid.Address())

// 			msg := types.NewMsgCreateEvaluation(txHash, senderDid, createEvaluationDoc, furyDid.Did)
// 			err = msg.ValidateBasic()
// 			if err != nil {
// 				return err
// 			}

// 			return furytypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), furyDid, msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

// func NewCmdWithdrawFunds() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "withdraw-funds [sender-did] [iid]",
// 		Short: "Withdraw funds.",
// 		Args:  cobra.ExactArgs(2),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			furyDid, err := didtypes.UnmarshalFuryDid(args[0])
// 			if err != nil {
// 				return err
// 			}

// 			var data types.WithdrawFundsDoc
// 			err = json.Unmarshal([]byte(args[1]), &data)
// 			if err != nil {
// 				return err
// 			}

// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}
// 			clientCtx = clientCtx.WithFromAddress(furyDid.Address())

// 			msg := types.NewMsgWithdrawFunds(furyDid.Did, data)
// 			err = msg.ValidateBasic()
// 			if err != nil {
// 				return err
// 			}

// 			return furytypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), furyDid, msg)
// 		},
// 	}

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }

// func NewCmdUpdateProjectDoc() *cobra.Command {
// cmd := &cobra.Command{
// 	Use:   "update-project-doc [sender-did] [project-iid-json] [fury-did]",
// 	Short: "Update a project's iid signed by the furyDid of the project",
// 	Args:  cobra.ExactArgs(3),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		senderDid := args[0]
// 		projectDataStr := args[1]
// 		furyDid, err := didtypes.UnmarshalFuryDid(args[2])
// 		if err != nil {
// 			return err
// 		}

// 		clientCtx, err := client.GetClientTxContext(cmd)
// 		if err != nil {
// 			return err
// 		}
// 		clientCtx = clientCtx.WithFromAddress(furyDid.Address())

// 		msg := types.NewMsgUpdateProjectDoc(senderDid, json.RawMessage(projectDataStr), furyDid.Did)
// 		err = msg.ValidateBasic()
// 		if err != nil {
// 			return err
// 		}

// 		return furytypes.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), furyDid, msg)
// 	},
// }

// 	flags.AddTxFlagsToCmd(cmd)
// 	return cmd
// }
