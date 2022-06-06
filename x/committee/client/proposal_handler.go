package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/Gert-Chain/core/x/committee/client/cli"
	"github.com/Gert-Chain/core/x/committee/client/rest"
)

// ProposalHandler is a struct containing handler funcs for submiting CommitteeChange/Delete proposal txs to the gov module through the cli or rest.
var ProposalHandler = govclient.NewProposalHandler(cli.GetGovCmdSubmitProposal, rest.ProposalRESTHandler)
