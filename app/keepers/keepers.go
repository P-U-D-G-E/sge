package keepers

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/group"
	groupkeeper "github.com/cosmos/cosmos-sdk/x/group/keeper"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icahost "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host"
	ibc "github.com/cosmos/ibc-go/v5/modules/core"

	// ibc-go

	ica "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts"
	icacontroller "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/types"
	icahostkeeper "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/host/types"
	ibcfee "github.com/cosmos/ibc-go/v5/modules/apps/29-fee"
	ibcfeekeeper "github.com/cosmos/ibc-go/v5/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v5/modules/apps/29-fee/types"
	"github.com/cosmos/ibc-go/v5/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v5/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v5/modules/core/02-client"
	ibcclienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v5/modules/core/keeper"
	mintkeeper "github.com/sge-network/sge/x/mint/keeper"
	minttypes "github.com/sge-network/sge/x/mint/types"

	// sge

	betmodule "github.com/sge-network/sge/x/bet"
	betmodulekeeper "github.com/sge-network/sge/x/bet/keeper"
	betmoduletypes "github.com/sge-network/sge/x/bet/types"

	marketmodule "github.com/sge-network/sge/x/market"
	marketmodulekeeper "github.com/sge-network/sge/x/market/keeper"
	marketmoduletypes "github.com/sge-network/sge/x/market/types"

	ovmmodule "github.com/sge-network/sge/x/ovm"
	ovmmodulekeeper "github.com/sge-network/sge/x/ovm/keeper"
	ovmmoduletypes "github.com/sge-network/sge/x/ovm/types"

	housemodule "github.com/sge-network/sge/x/house"
	housemodulekeeper "github.com/sge-network/sge/x/house/keeper"
	housemoduletypes "github.com/sge-network/sge/x/house/types"

	orderbookmodule "github.com/sge-network/sge/x/orderbook"
	orderbookmodulekeeper "github.com/sge-network/sge/x/orderbook/keeper"
	orderbookmoduletypes "github.com/sge-network/sge/x/orderbook/types"

	// unnamed import of statik for swagger UI support
	_ "github.com/cosmos/cosmos-sdk/client/docs/statik"
)

type AppKeepers struct {
	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// SDK keepers
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper
	FeeGrantKeeper   feegrantkeeper.Keeper
	AuthzKeeper      authzkeeper.Keeper
	GroupKeeper      groupkeeper.Keeper

	//// SGE keepers \\\\
	BetKeeper       *betmodulekeeper.Keeper
	MarketKeeper    *marketmodulekeeper.Keeper
	MintKeeper      mintkeeper.Keeper
	HouseKeeper     *housemodulekeeper.Keeper
	OrderbookKeeper *orderbookmodulekeeper.Keeper
	OVMKeeper       *ovmmodulekeeper.Keeper

	//// SGE modules \\\\
	BetModule       betmodule.AppModule
	MarketModule    marketmodule.AppModule
	HouseModule     housemodule.AppModule
	OrderbookModule orderbookmodule.AppModule
	OVMModule       ovmmodule.AppModule

	// make scoped keepers public for test purposes
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper      capabilitykeeper.ScopedKeeper
	ScopedIBCFeeKeeper        capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper

	// IBC Keepers
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	IBCFeeKeeper        ibcfeekeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper

	// IBC Modules
	IBCModule      ibc.AppModule
	ICAModule      ica.AppModule
	TransferModule transfer.AppModule
	IBCFeeModule   ibcfee.AppModule
}

func NewAppKeeper(
	appCodec codec.Codec,
	bApp *baseapp.BaseApp,
	cdc *codec.LegacyAmino,
	maccPerms map[string][]string,
	moduleAccAddress map[string]bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	_ servertypes.AppOptions,
) AppKeepers {
	appKeepers := AppKeepers{}
	// Set keys KVStoreKey, TransientStoreKey, MemoryStoreKey
	appKeepers.GenerateKeys()

	appKeepers.ParamsKeeper = initParamsKeeper(
		appCodec,
		cdc,
		appKeepers.keys[paramstypes.StoreKey],
		appKeepers.tkeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(
		appKeepers.ParamsKeeper.Subspace(baseapp.Paramspace).
			WithKeyTable(paramstypes.ConsensusParamsKeyTable()),
	)

	// add capability keeper and ScopeToModule for ibc module
	appKeepers.CapabilityKeeper = capabilitykeeper.NewKeeper(
		appCodec,
		appKeepers.keys[capabilitytypes.StoreKey],
		appKeepers.memKeys[capabilitytypes.MemStoreKey],
	)

	// grant capabilities for the ibc and ibc-transfer modules
	appKeepers.ScopedIBCKeeper = appKeepers.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	appKeepers.ScopedTransferKeeper = appKeepers.CapabilityKeeper.ScopeToModule(
		ibctransfertypes.ModuleName,
	)
	appKeepers.ScopedICAControllerKeeper = appKeepers.CapabilityKeeper.ScopeToModule(
		icacontrollertypes.SubModuleName,
	)
	appKeepers.ScopedICAHostKeeper = appKeepers.CapabilityKeeper.ScopeToModule(
		icahosttypes.SubModuleName,
	)

	appKeepers.CapabilityKeeper.Seal()

	// add keepers
	appKeepers.CrisisKeeper = crisiskeeper.NewKeeper(
		appKeepers.GetSubspace(crisistypes.ModuleName),
		invCheckPeriod,
		appKeepers.BankKeeper,
		authtypes.FeeCollectorName,
	)

	appKeepers.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		appKeepers.keys[authtypes.StoreKey],
		appKeepers.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
	)
	appKeepers.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		appKeepers.keys[banktypes.StoreKey],
		appKeepers.AccountKeeper,
		appKeepers.GetSubspace(banktypes.ModuleName),
		moduleAccAddress,
	)
	appKeepers.AuthzKeeper = authzkeeper.NewKeeper(
		appKeepers.keys[authzkeeper.StoreKey],
		appCodec,
		bApp.MsgServiceRouter(),
		appKeepers.AccountKeeper,
	)

	groupConfig := group.DefaultConfig()
	appKeepers.GroupKeeper = groupkeeper.NewKeeper(
		appKeepers.keys[group.StoreKey],
		appCodec,
		bApp.MsgServiceRouter(),
		appKeepers.AccountKeeper, groupConfig,
	)

	appKeepers.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[feegrant.StoreKey],
		appKeepers.AccountKeeper,
	)

	appKeepers.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[stakingtypes.StoreKey],
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.GetSubspace(stakingtypes.ModuleName),
	)

	appKeepers.MintKeeper = *mintkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[minttypes.StoreKey],
		appKeepers.GetSubspace(minttypes.ModuleName),
		appKeepers.AccountKeeper,
		mintkeeper.ExpectedKeepers{
			StakingKeeper: appKeepers.StakingKeeper,
			BankKeeper:    appKeepers.BankKeeper,
		},
		authtypes.FeeCollectorName,
	)

	appKeepers.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[distrtypes.StoreKey],
		appKeepers.GetSubspace(distrtypes.ModuleName),
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.StakingKeeper,
		authtypes.FeeCollectorName,
	)
	appKeepers.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[slashingtypes.StoreKey],
		appKeepers.StakingKeeper,
		appKeepers.GetSubspace(slashingtypes.ModuleName),
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	appKeepers.StakingKeeper = *appKeepers.StakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(
			appKeepers.DistrKeeper.Hooks(),
			appKeepers.SlashingKeeper.Hooks(),
		),
	)

	// UpgradeKeeper must be created before IBCKeeper
	appKeepers.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		appKeepers.keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		bApp,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	// Create IBC Keeper
	appKeepers.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		appKeepers.keys[ibchost.StoreKey],
		appKeepers.GetSubspace(ibchost.ModuleName),
		appKeepers.StakingKeeper,
		appKeepers.UpgradeKeeper,
		appKeepers.ScopedIBCKeeper,
	)

	// register the proposal types
	govRouter := govv1beta1.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(appKeepers.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(appKeepers.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(appKeepers.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(appKeepers.IBCKeeper.ClientKeeper))

	govConfig := govtypes.DefaultConfig()

	appKeepers.GovKeeper = govkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[govtypes.StoreKey],
		appKeepers.GetSubspace(govtypes.ModuleName),
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.StakingKeeper,
		govRouter,
		bApp.MsgServiceRouter(),
		govConfig,
	)

	// IBC Fee Module keeper
	appKeepers.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[ibcfeetypes.StoreKey],
		appKeepers.GetSubspace(ibcfeetypes.ModuleName),
		appKeepers.IBCKeeper.ChannelKeeper, // more middlewares can be added in future
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.IBCKeeper.PortKeeper, appKeepers.AccountKeeper, appKeepers.BankKeeper,
	)

	appKeepers.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[ibctransfertypes.StoreKey],
		appKeepers.GetSubspace(ibctransfertypes.ModuleName),
		appKeepers.IBCFeeKeeper, // ISC4 Wrapper: fee IBC middleware
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.IBCKeeper.PortKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.ScopedTransferKeeper,
	)

	appKeepers.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		appCodec,
		appKeepers.keys[icacontrollertypes.StoreKey],
		appKeepers.GetSubspace(icacontrollertypes.SubModuleName),
		appKeepers.IBCFeeKeeper,
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.IBCKeeper.PortKeeper,
		appKeepers.ScopedICAControllerKeeper,
		bApp.MsgServiceRouter(),
	)

	appKeepers.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec, appKeepers.keys[icahosttypes.StoreKey],
		appKeepers.GetSubspace(icahosttypes.SubModuleName),
		appKeepers.IBCFeeKeeper,
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.IBCKeeper.PortKeeper,
		appKeepers.AccountKeeper,
		appKeepers.ScopedICAHostKeeper,
		bApp.MsgServiceRouter(),
	)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	appKeepers.EvidenceKeeper = *evidencekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[evidencetypes.StoreKey],
		&appKeepers.StakingKeeper,
		appKeepers.SlashingKeeper,
	)

	//// SGE keepers \\\\

	appKeepers.OrderbookKeeper = orderbookmodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[orderbookmoduletypes.StoreKey],
		appKeepers.GetSubspace(orderbookmoduletypes.ModuleName),
		orderbookmodulekeeper.SdkExpectedKeepers{
			BankKeeper:     appKeepers.BankKeeper,
			AccountKeeper:  appKeepers.AccountKeeper,
			FeeGrantKeeper: appKeepers.FeeGrantKeeper,
		},
	)

	appKeepers.OVMKeeper = ovmmodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[ovmmoduletypes.StoreKey],
		appKeepers.keys[ovmmoduletypes.MemStoreKey],
		appKeepers.GetSubspace(ovmmoduletypes.ModuleName),
	)

	appKeepers.MarketKeeper = marketmodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[marketmoduletypes.StoreKey],
		appKeepers.keys[marketmoduletypes.MemStoreKey],
		appKeepers.GetSubspace(marketmoduletypes.ModuleName),
	)
	appKeepers.MarketKeeper.SetOVMKeeper(appKeepers.OVMKeeper)
	appKeepers.MarketKeeper.SetOrderbookKeeper(appKeepers.OrderbookKeeper)

	appKeepers.BetKeeper = betmodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[betmoduletypes.StoreKey],
		appKeepers.keys[betmoduletypes.MemStoreKey],
		appKeepers.GetSubspace(betmoduletypes.ModuleName),
	)
	appKeepers.BetKeeper.SetMarketKeeper(appKeepers.MarketKeeper)
	appKeepers.BetKeeper.SetOrderbookKeeper(appKeepers.OrderbookKeeper)
	appKeepers.BetKeeper.SetOVMKeeper(appKeepers.OVMKeeper)

	appKeepers.OrderbookKeeper.SetBetKeeper(appKeepers.BetKeeper)
	appKeepers.OrderbookKeeper.SetMarketKeeper(appKeepers.MarketKeeper)
	appKeepers.OrderbookKeeper.SetOVMKeeper(appKeepers.OVMKeeper)

	appKeepers.HouseKeeper = housemodulekeeper.NewKeeper(
		appCodec,
		appKeepers.keys[housemoduletypes.StoreKey],
		appKeepers.OrderbookKeeper,
		appKeepers.OVMKeeper,
		appKeepers.GetSubspace(housemoduletypes.ModuleName),
		housemodulekeeper.SdkExpectedKeepers{
			AuthzKeeper: appKeepers.AuthzKeeper,
		},
	)
	appKeepers.OrderbookKeeper.SetHouseKeeper(appKeepers.HouseKeeper)

	//// SGE modules \\\\

	appKeepers.BetModule = betmodule.NewAppModule(
		appCodec,
		*appKeepers.BetKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.MarketKeeper,
		appKeepers.OrderbookKeeper,
		appKeepers.OVMKeeper,
	)
	appKeepers.MarketModule = marketmodule.NewAppModule(
		appCodec,
		*appKeepers.MarketKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.OVMKeeper,
	)
	appKeepers.HouseModule = housemodule.NewAppModule(
		appCodec,
		*appKeepers.HouseKeeper,
	)
	appKeepers.OrderbookModule = orderbookmodule.NewAppModule(
		appCodec,
		*appKeepers.OrderbookKeeper,
	)
	appKeepers.OVMModule = ovmmodule.NewAppModule(
		appCodec,
		*appKeepers.OVMKeeper,
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
	)

	//// IBC modules \\\\
	appKeepers.IBCModule = ibc.NewAppModule(appKeepers.IBCKeeper)
	appKeepers.IBCFeeModule = ibcfee.NewAppModule(appKeepers.IBCFeeKeeper)
	appKeepers.TransferModule = transfer.NewAppModule(appKeepers.TransferKeeper)
	appKeepers.ICAModule = ica.NewAppModule(&appKeepers.ICAControllerKeeper, &appKeepers.ICAHostKeeper)

	// IBC stacks \\\
	var transferStack ibcporttypes.IBCModule
	transferStack = transfer.NewIBCModule(appKeepers.TransferKeeper)
	transferStack = ibcfee.NewIBCMiddleware(transferStack, appKeepers.IBCFeeKeeper)

	var icaControllerStack ibcporttypes.IBCModule
	icaControllerStack = icacontroller.NewIBCMiddleware(icaControllerStack, appKeepers.ICAControllerKeeper)
	icaControllerStack = ibcfee.NewIBCMiddleware(icaControllerStack, appKeepers.IBCFeeKeeper)

	var icaHostStack ibcporttypes.IBCModule
	icaHostStack = icahost.NewIBCModule(appKeepers.ICAHostKeeper)
	icaHostStack = ibcfee.NewIBCMiddleware(icaHostStack, appKeepers.IBCFeeKeeper)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(icacontrollertypes.SubModuleName, icaControllerStack)
	ibcRouter.AddRoute(icahosttypes.SubModuleName, icaHostStack)
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferStack)

	appKeepers.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/
	return appKeepers
}

// GetSubspace returns a param subspace for a given module name.
func (appKeepers *AppKeepers) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := appKeepers.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec,
	legacyAmino *codec.LegacyAmino,
	key, tkey storetypes.StoreKey,
) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govv1.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	paramsKeeper.Subspace(betmoduletypes.ModuleName)
	paramsKeeper.Subspace(marketmoduletypes.ModuleName)
	paramsKeeper.Subspace(orderbookmoduletypes.ModuleName)
	paramsKeeper.Subspace(ovmmoduletypes.ModuleName)
	paramsKeeper.Subspace(housemoduletypes.ModuleName)

	return paramsKeeper
}
