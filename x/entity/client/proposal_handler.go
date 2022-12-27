package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/furya-official/fury-blockchain/x/entity/client/cli"
	"github.com/furya-official/fury-blockchain/x/entity/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdUpdateEntityParamsProposal, rest.ProposalRESTHandler)
