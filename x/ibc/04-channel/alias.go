package channel

// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/04-channel/keeper

import (
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/keeper"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"
)

const (
	AttributeKeyConnectionID       = types.AttributeKeyConnectionID
	AttributeKeyPortID             = types.AttributeKeyPortID
	AttributeKeyChannelID          = types.AttributeKeyChannelID
	AttributeCounterpartyPortID    = types.AttributeCounterpartyPortID
	AttributeCounterpartyChannelID = types.AttributeCounterpartyChannelID
	EventTypeSendPacket            = types.EventTypeSendPacket
	EventTypeRecvPacket            = types.EventTypeRecvPacket
	EventTypeAcknowledgePacket     = types.EventTypeAcknowledgePacket
	EventTypeTimeoutPacket         = types.EventTypeTimeoutPacket
	AttributeKeyData               = types.AttributeKeyData
	AttributeKeyAck                = types.AttributeKeyAck
	AttributeKeyTimeoutHeight      = types.AttributeKeyTimeoutHeight
	AttributeKeyTimeoutTimestamp   = types.AttributeKeyTimeoutTimestamp
	AttributeKeySequence           = types.AttributeKeySequence
	AttributeKeySrcPort            = types.AttributeKeySrcPort
	AttributeKeySrcChannel         = types.AttributeKeySrcChannel
	AttributeKeyDstPort            = types.AttributeKeyDstPort
	AttributeKeyDstChannel         = types.AttributeKeyDstChannel
	SubModuleName                  = types.SubModuleName
	StoreKey                       = types.StoreKey
	RouterKey                      = types.RouterKey
	QuerierRoute                   = types.QuerierRoute
	UNINITIALIZED                  = types.UNINITIALIZED
	INIT                           = types.INIT
	TRYOPEN                        = types.TRYOPEN
	OPEN                           = types.OPEN
	CLOSED                         = types.CLOSED
	NONE                           = types.NONE
	UNORDERED                      = types.UNORDERED
	ORDERED                        = types.ORDERED
)

var (
	// functions aliases
	NewChannel                = types.NewChannel
	NewCounterparty           = types.NewCounterparty
	NewIdentifiedChannel      = types.NewIdentifiedChannel
	RegisterCodec             = types.RegisterCodec
	RegisterInterfaces        = types.RegisterInterfaces
	NewPacketAckCommitment    = types.NewPacketAckCommitment
	NewPacketSequence         = types.NewPacketSequence
	NewGenesisState           = types.NewGenesisState
	DefaultGenesisState       = types.DefaultGenesisState
	NewMsgChannelOpenInit     = types.NewMsgChannelOpenInit
	NewMsgChannelOpenTry      = types.NewMsgChannelOpenTry
	NewMsgChannelOpenAck      = types.NewMsgChannelOpenAck
	NewMsgChannelOpenConfirm  = types.NewMsgChannelOpenConfirm
	NewMsgChannelCloseInit    = types.NewMsgChannelCloseInit
	NewMsgChannelCloseConfirm = types.NewMsgChannelCloseConfirm
	NewMsgPacket              = types.NewMsgPacket
	NewMsgTimeout             = types.NewMsgTimeout
	NewMsgAcknowledgement     = types.NewMsgAcknowledgement
	CommitPacket              = types.CommitPacket
	CommitAcknowledgement     = types.CommitAcknowledgement
	NewPacket                 = types.NewPacket
	NewKeeper                 = keeper.NewKeeper

	// variable aliases
	SubModuleCdc                 = types.SubModuleCdc
	ErrChannelExists             = types.ErrChannelExists
	ErrChannelNotFound           = types.ErrChannelNotFound
	ErrInvalidChannel            = types.ErrInvalidChannel
	ErrInvalidChannelState       = types.ErrInvalidChannelState
	ErrInvalidChannelOrdering    = types.ErrInvalidChannelOrdering
	ErrInvalidCounterparty       = types.ErrInvalidCounterparty
	ErrInvalidChannelCapability  = types.ErrInvalidChannelCapability
	ErrChannelCapabilityNotFound = types.ErrChannelCapabilityNotFound
	ErrSequenceSendNotFound      = types.ErrSequenceSendNotFound
	ErrSequenceReceiveNotFound   = types.ErrSequenceReceiveNotFound
	ErrSequenceAckNotFound       = types.ErrSequenceAckNotFound
	ErrInvalidPacket             = types.ErrInvalidPacket
	ErrPacketTimeout             = types.ErrPacketTimeout
	ErrTooManyConnectionHops     = types.ErrTooManyConnectionHops
	ErrAcknowledgementTooLong    = types.ErrAcknowledgementTooLong
	EventTypeChannelOpenInit     = types.EventTypeChannelOpenInit
	EventTypeChannelOpenTry      = types.EventTypeChannelOpenTry
	EventTypeChannelOpenAck      = types.EventTypeChannelOpenAck
	EventTypeChannelOpenConfirm  = types.EventTypeChannelOpenConfirm
	EventTypeChannelCloseInit    = types.EventTypeChannelCloseInit
	EventTypeChannelCloseConfirm = types.EventTypeChannelCloseConfirm
	AttributeValueCategory       = types.AttributeValueCategory
)

type (
	IdentifiedChannel      = types.IdentifiedChannel
	ClientKeeper           = types.ClientKeeper
	ConnectionKeeper       = types.ConnectionKeeper
	PortKeeper             = types.PortKeeper
	PacketAckCommitment    = types.PacketAckCommitment
	PacketSequence         = types.PacketSequence
	GenesisState           = types.GenesisState
	State                  = types.State
	Order                  = types.Order
	MsgChannelOpenInit     = types.MsgChannelOpenInit
	MsgChannelOpenTry      = types.MsgChannelOpenTry
	MsgChannelOpenAck      = types.MsgChannelOpenAck
	MsgChannelOpenConfirm  = types.MsgChannelOpenConfirm
	MsgChannelCloseInit    = types.MsgChannelCloseInit
	MsgChannelCloseConfirm = types.MsgChannelCloseConfirm
	MsgPacket              = types.MsgPacket
	MsgTimeout             = types.MsgTimeout
	MsgAcknowledgement     = types.MsgAcknowledgement
	Channel                = types.Channel
	Counterparty           = types.Counterparty
	Packet                 = types.Packet
	Keeper                 = keeper.Keeper
)
