package vrf

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/log/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/vrf_coordinator_v2"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
)

func Test_BatchFulfillments_AddRun(t *testing.T) {
	batchLimit := uint32(2500)
	bfs := newBatchFulfillments(batchLimit)
	fromAddress := testutils.NewAddress()
	for i := 0; i < 4; i++ {
		bfs.addRun(vrfPipelineResult{
			gasLimit: 500,
			req: pendingRequest{
				req: &vrf_coordinator_v2.VRFCoordinatorV2RandomWordsRequested{
					RequestId: big.NewInt(1),
					Raw: types.Log{
						TxHash: common.HexToHash("0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65"),
					},
				},
				lb: mocks.NewBroadcast(t),
			},
			run: pipeline.NewRun(pipeline.Spec{}, pipeline.Vars{}),
		}, fromAddress)
		require.Len(t, bfs.fulfillments, 1)
	}

	require.Equal(t, uint32(2000), bfs.fulfillments[0].totalGasLimit)

	// This addition should create and add a new batch
	bfs.addRun(vrfPipelineResult{
		gasLimit: 500,
		req: pendingRequest{
			req: &vrf_coordinator_v2.VRFCoordinatorV2RandomWordsRequested{
				RequestId: big.NewInt(1),
				Raw: types.Log{
					TxHash: common.HexToHash("0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65"),
				},
			},
			lb: mocks.NewBroadcast(t),
		},
		run: pipeline.NewRun(pipeline.Spec{}, pipeline.Vars{}),
	}, fromAddress)
	require.Len(t, bfs.fulfillments, 2)
}
