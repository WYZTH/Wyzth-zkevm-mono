package proof

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/wyzth_zkevmxyz/wyzth_zkevm-mono/packages/relayer"
)

type blocker interface {
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
}
type Prover struct {
	blocker blocker
}

func New(blocker blocker) (*Prover, error) {
	if blocker == nil {
		return nil, relayer.ErrNoEthClient
	}

	return &Prover{
		blocker: blocker,
	}, nil
}
