package slashing

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/params"
)

// Default parameter namespace
const (
	DefaultParamSpace = "Slashing"
)

// nolint - Key generators for parameter access
func MaxEvidenceAgeKey() params.Key           { return params.NewKey("MaxEvidenceAge") }
func SignedBlocksWindowKey() params.Key       { return params.NewKey("SignedBlocksWindow") }
func MinSignedPerWindowKey() params.Key       { return params.NewKey("MinSignedPerWindow") }
func DoubleSignUnbondDurationKey() params.Key { return params.NewKey("DoubleSignUnbondDuration") }
func DowntimeUnbondDurationKey() params.Key   { return params.NewKey("DowntimeUnbondDuration") }
func SlashFractionDoubleSignKey() params.Key  { return params.NewKey("SlashFractionDoubleSign") }
func SlashFractionDowntimeKey() params.Key    { return params.NewKey("SlashFractionDowntime") }

// Cached parameter keys
var (
	maxEvidenceAgeKey           = MaxEvidenceAgeKey()
	signedBlocksWindowKey       = SignedBlocksWindowKey()
	minSignedPerWindowKey       = MinSignedPerWindowKey()
	doubleSignUnbondDurationKey = DoubleSignUnbondDurationKey()
	downtimeUnbondDurationKey   = DowntimeUnbondDurationKey()
	slashFractionDoubleSignKey  = SlashFractionDoubleSignKey()
	slashFractionDowntimeKey    = SlashFractionDowntimeKey()
)

// DefaultParams - used for initializing default parameter for slashing at genesis
type DefaultParams struct {
	MaxEvidenceAge           int64
	SignedBlocksWindow       int64
	MinSignedPerWindow       sdk.Rat
	DoubleSignUnbondDuration int64
	DowntimeUnbondDuration   int64
	SlashFractionDoubleSign  sdk.Rat
	SlashFractionDowntime    sdk.Rat
}

// Default parameters used by Cosmos Hub
func HubDefaultParams() DefaultParams {
	return DefaultParams{
		// defaultMaxEvidenceAge = 60 * 60 * 24 * 7 * 3
		// TODO Temporarily set to 2 minutes for testnets.
		MaxEvidenceAge: 60 * 2,

		// TODO Temporarily set to five minutes for testnets
		DoubleSignUnbondDuration: 60 * 5,

		// TODO Temporarily set to 100 blocks for testnets
		SignedBlocksWindow: 100,

		// TODO Temporarily set to 10 minutes for testnets
		DowntimeUnbondDuration: 60 * 10,

		MinSignedPerWindow: sdk.NewRat(1, 2),

		SlashFractionDoubleSign: sdk.NewRat(1).Quo(sdk.NewRat(20)),

		SlashFractionDowntime: sdk.NewRat(1).Quo(sdk.NewRat(100)),
	}
}

// MaxEvidenceAge - Max age for evidence - 21 days (3 weeks)
// MaxEvidenceAge = 60 * 60 * 24 * 7 * 3
func (k Keeper) MaxEvidenceAge(ctx sdk.Context) time.Duration {
	var t int64
	k.paramstore.Get(ctx, maxEvidenceAgeKey, &t)
	return time.Duration(t) * time.Second
}

// SignedBlocksWindow - sliding window for downtime slashing
func (k Keeper) SignedBlocksWindow(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, SignedBlocksWindowKey(), &res)
	return
}

// Downtime slashing thershold - default 50% of the SignedBlocksWindow
func (k Keeper) MinSignedPerWindow(ctx sdk.Context) int64 {
	var minSignedPerWindow sdk.Rat
	k.paramstore.Get(ctx, MinSignedPerWindowKey(), &minSignedPerWindow)
	signedBlocksWindow := k.SignedBlocksWindow(ctx)
	return sdk.NewRat(signedBlocksWindow).Mul(minSignedPerWindow).RoundInt64()
}

// Double-sign unbond duration
func (k Keeper) DoubleSignUnbondDuration(ctx sdk.Context) time.Duration {
	var t int64
	k.params.Get(ctx, doubleSignUnbondDurationKey, &t)
	return time.Duration(t) * time.Second
}

// Downtime unbond duration
func (k Keeper) DowntimeUnbondDuration(ctx sdk.Context) time.Duration {
	var t int64
	k.params.Get(ctx, downtimeUnbondDurationKey, &t)
	return time.Duration(t) * time.Second
}

// SlashFractionDoubleSign - currently default 5%
func (k Keeper) SlashFractionDoubleSign(ctx sdk.Context) (res sdk.Rat) {
	k.paramstore.Get(ctx, SlashFractionDoubleSignKey(), &res)
	return
}

// SlashFractionDowntime - currently default 1%
func (k Keeper) SlashFractionDowntime(ctx sdk.Context) (res sdk.Rat) {
	k.paramstore.Get(ctx, SlashFractionDowntimeKey(), &res)
	return
}
