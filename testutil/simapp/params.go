package simapp

import (
	"crypto/ed25519"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

const (
	usernamePrefix = "user"
)

// TestUser is simapp user type for testing
type TestUser struct {
	PrvKey  secp256k1.PrivKey
	Address sdk.AccAddress
	Balance int64
}

// TestValidator is simapp validator type for testing
type TestValidator struct {
	PubKey      types.PubKey
	Address     sdk.ValAddress
	ConsAddress sdk.ConsAddress
	Power       sdkmath.Int
}

var (
	// TestParamUsers represents the map of simapp users
	TestParamUsers = make(map[string]TestUser)

	// TestParamValidatorAddresses represents the map of test validators
	TestParamValidatorAddresses = make(map[string]TestValidator)

	// TestOVMPublicKeys represents test public keys needed for ovm
	TestOVMPublicKeys []ed25519.PublicKey

	// TestOVMPrivateKeys represents test private keys needed for ovm
	TestOVMPrivateKeys []ed25519.PrivateKey
)

// GenesisModuleAccountsBalances holds the test genesis module accounts balance map.
var GenesisModuleAccountsBalances = map[string]int64{}
