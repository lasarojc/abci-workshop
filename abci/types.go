package abci

import (
	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fatal-fruit/cosmapp/mempool"
	"github.com/fatal-fruit/cosmapp/provider"
)

type PrepareProposalHandler struct {
	logger      log.Logger
	txConfig    client.TxConfig
	cdc         codec.Codec
	mempool     *mempool.ThresholdMempool
	txProvider  provider.TxProvider
	keyname     string
	runProvider bool
}

type ProcessProposalHandler struct {
	TxConfig client.TxConfig
	Codec    codec.Codec
	Logger   log.Logger
}

type VoteExtHandler struct {
	logger            log.Logger
	currentBlock      int64
	mempool           *mempool.ThresholdMempool
	cdc               codec.Codec
	extendSubscribers map[string]func(sdk.Context, *abci.RequestExtendVote) ([]byte, error)
	verifySubscribers map[string]func(sdk.Context, []byte) (abci.ResponseVerifyVoteExtension_VerifyStatus, error)
}

type InjectedVoteExt struct {
	VoteExtSigner []byte
	Bids          [][]byte
}

type InjectedVotes struct {
	Votes []InjectedVoteExt
}

type AppVoteExtension struct {
	Height    int64
	Bids      [][]byte
	ExtraInfo []byte
}

type SpecialTransaction struct {
	Height int
	Bids   [][]byte
}
