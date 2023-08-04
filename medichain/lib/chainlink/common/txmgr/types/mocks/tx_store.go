// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	feetypes "github.com/smartcontractkit/chainlink/v2/common/fee/types"
	mock "github.com/stretchr/testify/mock"

	pg "github.com/smartcontractkit/chainlink/v2/core/services/pg"

	time "time"

	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"

	types "github.com/smartcontractkit/chainlink/v2/common/types"

	uuid "github.com/google/uuid"
)

// TxStore is an autogenerated mock type for the TxStore type
type TxStore[ADDR types.Hashable, CHAIN_ID types.ID, TX_HASH types.Hashable, BLOCK_HASH types.Hashable, R txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH], SEQ types.Sequence, FEE feetypes.Fee] struct {
	mock.Mock
}

// Abandon provides a mock function with given fields: id, addr
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Abandon(id CHAIN_ID, addr ADDR) error {
	ret := _m.Called(id, addr)

	var r0 error
	if rf, ok := ret.Get(0).(func(CHAIN_ID, ADDR) error); ok {
		r0 = rf(id, addr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckTxQueueCapacity provides a mock function with given fields: fromAddress, maxQueuedTransactions, chainID, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CheckTxQueueCapacity(fromAddress ADDR, maxQueuedTransactions uint64, chainID CHAIN_ID, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fromAddress, maxQueuedTransactions, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(ADDR, uint64, CHAIN_ID, ...pg.QOpt) error); ok {
		r0 = rf(fromAddress, maxQueuedTransactions, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) Close() {
	_m.Called()
}

// CountUnconfirmedTransactions provides a mock function with given fields: fromAddress, chainID, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CountUnconfirmedTransactions(fromAddress ADDR, chainID CHAIN_ID, qopts ...pg.QOpt) (uint32, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fromAddress, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, CHAIN_ID, ...pg.QOpt) (uint32, error)); ok {
		return rf(fromAddress, chainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(ADDR, CHAIN_ID, ...pg.QOpt) uint32); ok {
		r0 = rf(fromAddress, chainID, qopts...)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(ADDR, CHAIN_ID, ...pg.QOpt) error); ok {
		r1 = rf(fromAddress, chainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountUnstartedTransactions provides a mock function with given fields: fromAddress, chainID, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CountUnstartedTransactions(fromAddress ADDR, chainID CHAIN_ID, qopts ...pg.QOpt) (uint32, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fromAddress, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, CHAIN_ID, ...pg.QOpt) (uint32, error)); ok {
		return rf(fromAddress, chainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(ADDR, CHAIN_ID, ...pg.QOpt) uint32); ok {
		r0 = rf(fromAddress, chainID, qopts...)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(ADDR, CHAIN_ID, ...pg.QOpt) error); ok {
		r1 = rf(fromAddress, chainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTransaction provides a mock function with given fields: txRequest, chainID, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) CreateTransaction(txRequest txmgrtypes.TxRequest[ADDR, TX_HASH], chainID CHAIN_ID, qopts ...pg.QOpt) (txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, txRequest, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(txmgrtypes.TxRequest[ADDR, TX_HASH], CHAIN_ID, ...pg.QOpt) (txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(txRequest, chainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(txmgrtypes.TxRequest[ADDR, TX_HASH], CHAIN_ID, ...pg.QOpt) txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(txRequest, chainID, qopts...)
	} else {
		r0 = ret.Get(0).(txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
	}

	if rf, ok := ret.Get(1).(func(txmgrtypes.TxRequest[ADDR, TX_HASH], CHAIN_ID, ...pg.QOpt) error); ok {
		r1 = rf(txRequest, chainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteInProgressAttempt provides a mock function with given fields: ctx, attempt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) DeleteInProgressAttempt(ctx context.Context, attempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	ret := _m.Called(ctx, attempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error); ok {
		r0 = rf(ctx, attempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindNextUnstartedTransactionFromAddress provides a mock function with given fields: etx, fromAddress, chainID, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindNextUnstartedTransactionFromAddress(etx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], fromAddress ADDR, chainID CHAIN_ID, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etx, fromAddress, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], ADDR, CHAIN_ID, ...pg.QOpt) error); ok {
		r0 = rf(etx, fromAddress, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindReceiptsPendingConfirmation provides a mock function with given fields: ctx, blockNum, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindReceiptsPendingConfirmation(ctx context.Context, blockNum int64, chainID CHAIN_ID) ([]txmgrtypes.ReceiptPlus[R], error) {
	ret := _m.Called(ctx, blockNum, chainID)

	var r0 []txmgrtypes.ReceiptPlus[R]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, CHAIN_ID) ([]txmgrtypes.ReceiptPlus[R], error)); ok {
		return rf(ctx, blockNum, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, CHAIN_ID) []txmgrtypes.ReceiptPlus[R]); ok {
		r0 = rf(ctx, blockNum, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgrtypes.ReceiptPlus[R])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, CHAIN_ID) error); ok {
		r1 = rf(ctx, blockNum, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTransactionsConfirmedInBlockRange provides a mock function with given fields: highBlockNumber, lowBlockNumber, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTransactionsConfirmedInBlockRange(highBlockNumber int64, lowBlockNumber int64, chainID CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(highBlockNumber, lowBlockNumber, chainID)

	var r0 []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64, CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(highBlockNumber, lowBlockNumber, chainID)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, CHAIN_ID) []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(highBlockNumber, lowBlockNumber, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(int64, int64, CHAIN_ID) error); ok {
		r1 = rf(highBlockNumber, lowBlockNumber, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxAttemptsConfirmedMissingReceipt provides a mock function with given fields: chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxAttemptsConfirmedMissingReceipt(chainID CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(chainID)

	var r0 []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(chainID)
	}
	if rf, ok := ret.Get(0).(func(CHAIN_ID) []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(CHAIN_ID) error); ok {
		r1 = rf(chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxAttemptsRequiringReceiptFetch provides a mock function with given fields: chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxAttemptsRequiringReceiptFetch(chainID CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(chainID)

	var r0 []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(chainID)
	}
	if rf, ok := ret.Get(0).(func(CHAIN_ID) []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(CHAIN_ID) error); ok {
		r1 = rf(chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxAttemptsRequiringResend provides a mock function with given fields: olderThan, maxInFlightTransactions, chainID, address
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxAttemptsRequiringResend(olderThan time.Time, maxInFlightTransactions uint32, chainID CHAIN_ID, address ADDR) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(olderThan, maxInFlightTransactions, chainID, address)

	var r0 []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, uint32, CHAIN_ID, ADDR) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(olderThan, maxInFlightTransactions, chainID, address)
	}
	if rf, ok := ret.Get(0).(func(time.Time, uint32, CHAIN_ID, ADDR) []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(olderThan, maxInFlightTransactions, chainID, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time, uint32, CHAIN_ID, ADDR) error); ok {
		r1 = rf(olderThan, maxInFlightTransactions, chainID, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxWithSequence provides a mock function with given fields: fromAddress, seq
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxWithSequence(fromAddress ADDR, seq SEQ) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(fromAddress, seq)

	var r0 *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, SEQ) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(fromAddress, seq)
	}
	if rf, ok := ret.Get(0).(func(ADDR, SEQ) *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(fromAddress, seq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(ADDR, SEQ) error); ok {
		r1 = rf(fromAddress, seq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxsRequiringGasBump provides a mock function with given fields: ctx, address, blockNum, gasBumpThreshold, depth, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxsRequiringGasBump(ctx context.Context, address ADDR, blockNum int64, gasBumpThreshold int64, depth int64, chainID CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, address, blockNum, gasBumpThreshold, depth, chainID)

	var r0 []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, int64, int64, int64, CHAIN_ID) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, address, blockNum, gasBumpThreshold, depth, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, int64, int64, int64, CHAIN_ID) []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, address, blockNum, gasBumpThreshold, depth, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, int64, int64, int64, CHAIN_ID) error); ok {
		r1 = rf(ctx, address, blockNum, gasBumpThreshold, depth, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTxsRequiringResubmissionDueToInsufficientFunds provides a mock function with given fields: address, chainID, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) FindTxsRequiringResubmissionDueToInsufficientFunds(address ADDR, chainID CHAIN_ID, qopts ...pg.QOpt) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, CHAIN_ID, ...pg.QOpt) ([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(address, chainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(ADDR, CHAIN_ID, ...pg.QOpt) []*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(address, chainID, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(ADDR, CHAIN_ID, ...pg.QOpt) error); ok {
		r1 = rf(address, chainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInProgressTxAttempts provides a mock function with given fields: ctx, address, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) GetInProgressTxAttempts(ctx context.Context, address ADDR, chainID CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	ret := _m.Called(ctx, address, chainID)

	var r0 []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) ([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(ctx, address, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ADDR, CHAIN_ID) []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(ctx, address, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ADDR, CHAIN_ID) error); ok {
		r1 = rf(ctx, address, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxInProgress provides a mock function with given fields: fromAddress, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) GetTxInProgress(fromAddress ADDR, qopts ...pg.QOpt) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, fromAddress)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, ...pg.QOpt) (*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], error)); ok {
		return rf(fromAddress, qopts...)
	}
	if rf, ok := ret.Get(0).(func(ADDR, ...pg.QOpt) *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]); ok {
		r0 = rf(fromAddress, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE])
		}
	}

	if rf, ok := ret.Get(1).(func(ADDR, ...pg.QOpt) error); ok {
		r1 = rf(fromAddress, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasInProgressTransaction provides a mock function with given fields: account, chainID, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) HasInProgressTransaction(account ADDR, chainID CHAIN_ID, qopts ...pg.QOpt) (bool, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, account, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(ADDR, CHAIN_ID, ...pg.QOpt) (bool, error)); ok {
		return rf(account, chainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(ADDR, CHAIN_ID, ...pg.QOpt) bool); ok {
		r0 = rf(account, chainID, qopts...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(ADDR, CHAIN_ID, ...pg.QOpt) error); ok {
		r1 = rf(account, chainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadTxAttempts provides a mock function with given fields: etx, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) LoadTxAttempts(etx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], ...pg.QOpt) error); ok {
		r0 = rf(etx, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkAllConfirmedMissingReceipt provides a mock function with given fields: chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) MarkAllConfirmedMissingReceipt(chainID CHAIN_ID) error {
	ret := _m.Called(chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(CHAIN_ID) error); ok {
		r0 = rf(chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkOldTxesMissingReceiptAsErrored provides a mock function with given fields: blockNum, finalityDepth, chainID, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) MarkOldTxesMissingReceiptAsErrored(blockNum int64, finalityDepth uint32, chainID CHAIN_ID, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, blockNum, finalityDepth, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, uint32, CHAIN_ID, ...pg.QOpt) error); ok {
		r0 = rf(blockNum, finalityDepth, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PreloadTxes provides a mock function with given fields: attempts, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) PreloadTxes(attempts []txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, attempts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func([]txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], ...pg.QOpt) error); ok {
		r0 = rf(attempts, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PruneUnstartedTxQueue provides a mock function with given fields: queueSize, subject, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) PruneUnstartedTxQueue(queueSize uint32, subject uuid.UUID, qopts ...pg.QOpt) (int64, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, queueSize, subject)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(uint32, uuid.UUID, ...pg.QOpt) (int64, error)); ok {
		return rf(queueSize, subject, qopts...)
	}
	if rf, ok := ret.Get(0).(func(uint32, uuid.UUID, ...pg.QOpt) int64); ok {
		r0 = rf(queueSize, subject, qopts...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(uint32, uuid.UUID, ...pg.QOpt) error); ok {
		r1 = rf(queueSize, subject, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReapTxHistory provides a mock function with given fields: minBlockNumberToKeep, timeThreshold, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) ReapTxHistory(minBlockNumberToKeep int64, timeThreshold time.Time, chainID CHAIN_ID) error {
	ret := _m.Called(minBlockNumberToKeep, timeThreshold, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, time.Time, CHAIN_ID) error); ok {
		r0 = rf(minBlockNumberToKeep, timeThreshold, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveConfirmedMissingReceiptAttempt provides a mock function with given fields: ctx, timeout, attempt, broadcastAt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveConfirmedMissingReceiptAttempt(ctx context.Context, timeout time.Duration, attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], broadcastAt time.Time) error {
	ret := _m.Called(ctx, timeout, attempt, broadcastAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration, *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], time.Time) error); ok {
		r0 = rf(ctx, timeout, attempt, broadcastAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveFetchedReceipts provides a mock function with given fields: receipts, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveFetchedReceipts(receipts []R, chainID CHAIN_ID) error {
	ret := _m.Called(receipts, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func([]R, CHAIN_ID) error); ok {
		r0 = rf(receipts, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveInProgressAttempt provides a mock function with given fields: attempt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveInProgressAttempt(attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	ret := _m.Called(attempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error); ok {
		r0 = rf(attempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveInsufficientFundsAttempt provides a mock function with given fields: timeout, attempt, broadcastAt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveInsufficientFundsAttempt(timeout time.Duration, attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], broadcastAt time.Time) error {
	ret := _m.Called(timeout, attempt, broadcastAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], time.Time) error); ok {
		r0 = rf(timeout, attempt, broadcastAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveReplacementInProgressAttempt provides a mock function with given fields: oldAttempt, replacementAttempt, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveReplacementInProgressAttempt(oldAttempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], replacementAttempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, oldAttempt, replacementAttempt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], ...pg.QOpt) error); ok {
		r0 = rf(oldAttempt, replacementAttempt, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveSentAttempt provides a mock function with given fields: timeout, attempt, broadcastAt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SaveSentAttempt(timeout time.Duration, attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], broadcastAt time.Time) error {
	ret := _m.Called(timeout, attempt, broadcastAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], time.Time) error); ok {
		r0 = rf(timeout, attempt, broadcastAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetBroadcastBeforeBlockNum provides a mock function with given fields: blockNum, chainID
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) SetBroadcastBeforeBlockNum(blockNum int64, chainID CHAIN_ID) error {
	ret := _m.Called(blockNum, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, CHAIN_ID) error); ok {
		r0 = rf(blockNum, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBroadcastAts provides a mock function with given fields: now, etxIDs
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateBroadcastAts(now time.Time, etxIDs []int64) error {
	ret := _m.Called(now, etxIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time, []int64) error); ok {
		r0 = rf(now, etxIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateKeyNextSequence provides a mock function with given fields: newNextSequence, currentNextSequence, address, chainID, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateKeyNextSequence(newNextSequence SEQ, currentNextSequence SEQ, address ADDR, chainID CHAIN_ID, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, newNextSequence, currentNextSequence, address, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(SEQ, SEQ, ADDR, CHAIN_ID, ...pg.QOpt) error); ok {
		r0 = rf(newNextSequence, currentNextSequence, address, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxAttemptInProgressToBroadcast provides a mock function with given fields: etx, attempt, NewAttemptState, incrNextSequenceCallback, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxAttemptInProgressToBroadcast(etx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], attempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], NewAttemptState txmgrtypes.TxAttemptState, incrNextSequenceCallback func(pg.Queryer) error, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etx, attempt, NewAttemptState, incrNextSequenceCallback)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttemptState, func(pg.Queryer) error, ...pg.QOpt) error); ok {
		r0 = rf(etx, attempt, NewAttemptState, incrNextSequenceCallback, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxFatalError provides a mock function with given fields: etx, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxFatalError(etx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], ...pg.QOpt) error); ok {
		r0 = rf(etx, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxForRebroadcast provides a mock function with given fields: etx, etxAttempt
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxForRebroadcast(etx txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], etxAttempt txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error {
	ret := _m.Called(etx, etxAttempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE]) error); ok {
		r0 = rf(etx, etxAttempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxUnstartedToInProgress provides a mock function with given fields: etx, attempt, qopts
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxUnstartedToInProgress(etx *txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], attempt *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etx, attempt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgrtypes.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], *txmgrtypes.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, SEQ, FEE], ...pg.QOpt) error); ok {
		r0 = rf(etx, attempt, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTxsUnconfirmed provides a mock function with given fields: ids
func (_m *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]) UpdateTxsUnconfirmed(ids []int64) error {
	ret := _m.Called(ids)

	var r0 error
	if rf, ok := ret.Get(0).(func([]int64) error); ok {
		r0 = rf(ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTxStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewTxStore creates a new instance of TxStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTxStore[ADDR types.Hashable, CHAIN_ID types.ID, TX_HASH types.Hashable, BLOCK_HASH types.Hashable, R txmgrtypes.ChainReceipt[TX_HASH, BLOCK_HASH], SEQ types.Sequence, FEE feetypes.Fee](t mockConstructorTestingTNewTxStore) *TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE] {
	mock := &TxStore[ADDR, CHAIN_ID, TX_HASH, BLOCK_HASH, R, SEQ, FEE]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}