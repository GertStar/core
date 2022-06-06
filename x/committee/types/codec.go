package types

import (
	gertdisttypes "github.com/Gert-Chain/core/x/gertdist/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	proposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/committee module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/committee and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	// amino is not sealed so that other modules can register their own pubproposal and/or permission types.

	// Register external module pubproposal types. Ideally these would be registered within the modules' types pkg init function.
	// However registration happens here as a work-around.
	RegisterProposalTypeCodec(distrtypes.CommunityPoolSpendProposal{}, "cosmos-sdk/CommunityPoolSpendProposal")
	RegisterProposalTypeCodec(proposaltypes.ParameterChangeProposal{}, "cosmos-sdk/ParameterChangeProposal")
	RegisterProposalTypeCodec(govtypes.TextProposal{}, "cosmos-sdk/TextProposal")
	RegisterProposalTypeCodec(upgradetypes.SoftwareUpgradeProposal{}, "cosmos-sdk/SoftwareUpgradeProposal")
	RegisterProposalTypeCodec(upgradetypes.CancelSoftwareUpgradeProposal{}, "cosmos-sdk/CancelSoftwareUpgradeProposal")
}

// RegisterLegacyAminoCodec registers all the necessary types and interfaces for the module.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// Proposals
	cdc.RegisterInterface((*PubProposal)(nil), nil)
	cdc.RegisterConcrete(CommitteeChangeProposal{}, "gert/CommitteeChangeProposal", nil)
	cdc.RegisterConcrete(CommitteeDeleteProposal{}, "gert/CommitteeDeleteProposal", nil)

	// Committees
	cdc.RegisterInterface((*Committee)(nil), nil)
	cdc.RegisterConcrete(BaseCommittee{}, "gert/BaseCommittee", nil)
	cdc.RegisterConcrete(MemberCommittee{}, "gert/MemberCommittee", nil)
	cdc.RegisterConcrete(TokenCommittee{}, "gert/TokenCommittee", nil)

	// Permissions
	cdc.RegisterInterface((*Permission)(nil), nil)
	cdc.RegisterConcrete(GodPermission{}, "gert/GodPermission", nil)
	cdc.RegisterConcrete(TextPermission{}, "gert/TextPermission", nil)
	cdc.RegisterConcrete(SoftwareUpgradePermission{}, "gert/SoftwareUpgradePermission", nil)
	cdc.RegisterConcrete(ParamsChangePermission{}, "gert/ParamsChangePermission", nil)

	// Msgs
	cdc.RegisterConcrete(&MsgSubmitProposal{}, "gert/MsgSubmitProposal", nil)
	cdc.RegisterConcrete(&MsgVote{}, "gert/MsgVote", nil)
}

// RegisterProposalTypeCodec allows external modules to register their own pubproposal types on the
// internal ModuleCdc. This allows the MsgSubmitProposal to be correctly Amino encoded and decoded.
func RegisterProposalTypeCodec(o interface{}, name string) {
	ModuleCdc.RegisterConcrete(o, name, nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitProposal{},
		&MsgVote{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	registry.RegisterInterface(
		"gert.committee.v1beta1.Committee",
		(*Committee)(nil),
		&BaseCommittee{},
		&TokenCommittee{},
		&MemberCommittee{},
	)

	registry.RegisterInterface(
		"gert.committee.v1beta1.Permission",
		(*Permission)(nil),
		&GodPermission{},
		&TextPermission{},
		&SoftwareUpgradePermission{},
		&ParamsChangePermission{},
	)

	// Need to register PubProposal here since we use this as alias for the x/gov Content interface for all the proposal implementations used in this module.
	// Note that all proposals supported by x/committee needed to be registered here, including the proposals from x/gov.
	registry.RegisterInterface(
		"gert.committee.v1beta1.PubProposal",
		(*PubProposal)(nil),
		&Proposal{},
		&distrtypes.CommunityPoolSpendProposal{},
		&govtypes.TextProposal{},
		&gertdisttypes.CommunityPoolMultiSpendProposal{},
		&proposaltypes.ParameterChangeProposal{},
		&upgradetypes.SoftwareUpgradeProposal{},
		&upgradetypes.CancelSoftwareUpgradeProposal{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&CommitteeChangeProposal{},
		&CommitteeDeleteProposal{},
	)
}
