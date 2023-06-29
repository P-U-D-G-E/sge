package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	batchSettlementCount  = 1000
	maxBetByUIDQueryCount = 10
)

var (
	defaultMinBetAmount = sdk.NewInt(1000000)
	defaultBetFee       = sdk.NewInt(100)
)

// parameter store keys
var (
	// keyBatchSettlementCount is the batch settlement
	// count of bets
	keyBatchSettlementCount = []byte("BatchSettlementCount")

	// keyMaxBetByUIDQueryCount is the max count of
	// the queryable bets by UID list.
	keyMaxBetByUIDQueryCount = []byte("MaxBetByUidQueryCount")

	// keyPlacementConstraints is the default bet placement
	// constraints.
	keyPlacementConstraints = []byte("PlacementConstraints")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		BatchSettlementCount:  batchSettlementCount,
		MaxBetByUidQueryCount: maxBetByUIDQueryCount,
		PlacementConstraints: PlacementConstraints{
			MinAmount: defaultMinBetAmount,
			BetFee:    defaultBetFee,
		},
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			keyBatchSettlementCount,
			&p.BatchSettlementCount,
			validateBatchSettlementCount,
		),
		paramtypes.NewParamSetPair(
			keyMaxBetByUIDQueryCount,
			&p.MaxBetByUidQueryCount,
			validateMaxBetByUIDQueryCount,
		),
		paramtypes.NewParamSetPair(
			keyPlacementConstraints,
			&p.PlacementConstraints,
			validatePlacementConstraints,
		),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateBatchSettlementCount(p.BatchSettlementCount); err != nil {
		return err
	}

	if err := validateMaxBetByUIDQueryCount(p.MaxBetByUidQueryCount); err != nil {
		return err
	}

	if err := validatePlacementConstraints(p.PlacementConstraints); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func validateBatchSettlementCount(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("%s: %T", ErrTextInvalidParamType, i)
	}

	if v <= 0 {
		return fmt.Errorf("%s: %d", ErrTextBatchSettlementCountMustBePositive, v)
	}

	return nil
}

func validateMaxBetByUIDQueryCount(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("%s: %T", ErrTextInvalidParamType, i)
	}

	if v <= 0 {
		return fmt.Errorf("%s: %d", ErrTextMaxBetUIDQueryCountMustBePositive, v)
	}

	return nil
}

func validatePlacementConstraints(i interface{}) error {
	v, ok := i.(PlacementConstraints)
	if !ok {
		return fmt.Errorf("%s: %T", ErrTextInvalidParamType, i)
	}

	if v.MinAmount.LTE(sdk.OneInt()) {
		return fmt.Errorf("minimum bet amount must be more than one: %d", v.MinAmount.Int64())
	}

	if v.BetFee.LT(sdk.ZeroInt()) {
		return fmt.Errorf("minimum bet fee must be positive: %d", v.BetFee.Int64())
	}

	return nil
}
