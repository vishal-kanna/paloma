package client

import (
	govclient "cosmossdk.io/x/gov/client"
	// govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/palomachain/paloma/x/evm/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.CmdEvmChainProposalHandler)
