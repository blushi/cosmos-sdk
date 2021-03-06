package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	clientkeeper "github.com/cosmos/cosmos-sdk/x/ibc/02-client/keeper"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
	connectionkeeper "github.com/cosmos/cosmos-sdk/x/ibc/03-connection/keeper"
	channelkeeper "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/keeper"
	port "github.com/cosmos/cosmos-sdk/x/ibc/05-port"
)

// Keeper defines each ICS keeper for IBC
type Keeper struct {
	aminoCdc *codec.Codec
	cdc      codec.Marshaler

	ClientKeeper     clientkeeper.Keeper
	ConnectionKeeper connectionkeeper.Keeper
	ChannelKeeper    channelkeeper.Keeper
	PortKeeper       port.Keeper
	Router           *port.Router
}

// NewKeeper creates a new ibc Keeper
func NewKeeper(
	aminoCdc *codec.Codec, cdc codec.Marshaler, key sdk.StoreKey, stakingKeeper clienttypes.StakingKeeper, scopedKeeper capabilitykeeper.ScopedKeeper,
) *Keeper {
	clientKeeper := clientkeeper.NewKeeper(aminoCdc, key, stakingKeeper)
	connectionKeeper := connectionkeeper.NewKeeper(aminoCdc, cdc, key, clientKeeper)
	portKeeper := port.NewKeeper(scopedKeeper)
	channelKeeper := channelkeeper.NewKeeper(cdc, key, clientKeeper, connectionKeeper, portKeeper, scopedKeeper)

	return &Keeper{
		aminoCdc:         aminoCdc,
		cdc:              cdc,
		ClientKeeper:     clientKeeper,
		ConnectionKeeper: connectionKeeper,
		ChannelKeeper:    channelKeeper,
		PortKeeper:       portKeeper,
	}
}

// Codecs returns the IBC module codec.
func (k Keeper) Codecs() (codec.Marshaler, *codec.Codec) {
	return k.cdc, k.aminoCdc
}

// SetRouter sets the Router in IBC Keeper and seals it. The method panics if
// there is an existing router that's already sealed.
func (k *Keeper) SetRouter(rtr *port.Router) {
	if k.Router != nil && k.Router.Sealed() {
		panic("cannot reset a sealed router")
	}
	k.Router = rtr
	k.Router.Seal()
}
