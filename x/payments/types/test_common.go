package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// --------------------------------------------- TestPeriod

var _ Period = &TestPeriod{}

// TestPeriod Is identical to BlockPeriod but does
// not take into consideration the context in periodEnded() and periodStarted()

func NewTestPeriod(periodLength, periodStartBlock int64) TestPeriod {
	return TestPeriod{
		PeriodLength:     periodLength,
		PeriodStartBlock: periodStartBlock,
	}
}

func (p TestPeriod) periodEndBlock() int64 {
	return p.PeriodStartBlock + p.PeriodLength
}

func (p TestPeriod) GetPeriodUnit() string {
	return BlockPeriodUnit
}

func (p TestPeriod) Validate() error {
	// Validate period-related values
	if p.PeriodStartBlock > p.periodEndBlock() {
		return sdkerrors.Wrap(ErrInvalidPeriod, "start time is after end time")
	} else if p.PeriodLength <= 0 {
		return sdkerrors.Wrap(ErrInvalidPeriod, "period length must be greater than zero")
	}

	return nil
}

func (p TestPeriod) periodStarted(ctx sdk.Context) bool {
	return true
}

func (p TestPeriod) periodEnded(ctx sdk.Context) bool {
	return true
}

func (p TestPeriod) nextPeriod() Period {
	p.PeriodStartBlock = p.periodEndBlock()
	return &p
}
