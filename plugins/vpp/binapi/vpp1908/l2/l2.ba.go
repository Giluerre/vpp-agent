// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: /usr/share/vpp/api/core/l2.api.json

/*
 Package l2 is a generated from VPP binary API module 'l2'.

 It contains following objects:
	 25 services
	  3 enums
	  3 aliases
	  7 types
	  1 union
	 51 messages
*/
package l2

import api "git.fd.io/govpp.git/api"
import struc "github.com/lunixbochs/struc"
import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// Services represents VPP binary API services:
type Services interface {
	DumpBdIPMac(*BdIPMacDump) ([]*BdIPMacDetails, error)
	DumpBridgeDomain(*BridgeDomainDump) ([]*BridgeDomainDetails, error)
	DumpL2FibTable(*L2FibTableDump) ([]*L2FibTableDetails, error)
	DumpL2Xconnect(*L2XconnectDump) ([]*L2XconnectDetails, error)
	BdIPMacAddDel(*BdIPMacAddDel) (*BdIPMacAddDelReply, error)
	BdIPMacFlush(*BdIPMacFlush) (*BdIPMacFlushReply, error)
	BridgeDomainAddDel(*BridgeDomainAddDel) (*BridgeDomainAddDelReply, error)
	BridgeDomainSetMacAge(*BridgeDomainSetMacAge) (*BridgeDomainSetMacAgeReply, error)
	BridgeFlags(*BridgeFlags) (*BridgeFlagsReply, error)
	BviCreate(*BviCreate) (*BviCreateReply, error)
	BviDelete(*BviDelete) (*BviDeleteReply, error)
	L2FibClearTable(*L2FibClearTable) (*L2FibClearTableReply, error)
	L2Flags(*L2Flags) (*L2FlagsReply, error)
	L2InterfaceEfpFilter(*L2InterfaceEfpFilter) (*L2InterfaceEfpFilterReply, error)
	L2InterfacePbbTagRewrite(*L2InterfacePbbTagRewrite) (*L2InterfacePbbTagRewriteReply, error)
	L2InterfaceVlanTagRewrite(*L2InterfaceVlanTagRewrite) (*L2InterfaceVlanTagRewriteReply, error)
	L2PatchAddDel(*L2PatchAddDel) (*L2PatchAddDelReply, error)
	L2fibAddDel(*L2fibAddDel) (*L2fibAddDelReply, error)
	L2fibFlushAll(*L2fibFlushAll) (*L2fibFlushAllReply, error)
	L2fibFlushBd(*L2fibFlushBd) (*L2fibFlushBdReply, error)
	L2fibFlushInt(*L2fibFlushInt) (*L2fibFlushIntReply, error)
	SwInterfaceSetL2Bridge(*SwInterfaceSetL2Bridge) (*SwInterfaceSetL2BridgeReply, error)
	SwInterfaceSetL2Xconnect(*SwInterfaceSetL2Xconnect) (*SwInterfaceSetL2XconnectReply, error)
	SwInterfaceSetVpath(*SwInterfaceSetVpath) (*SwInterfaceSetVpathReply, error)
	WantL2MacsEvents(*WantL2MacsEvents) (*WantL2MacsEventsReply, error)
}

/* Enums */

// AddressFamily represents VPP binary API enum 'address_family':
type AddressFamily uint32

const (
	ADDRESS_IP4 AddressFamily = 0
	ADDRESS_IP6 AddressFamily = 1
)

// BdFlags represents VPP binary API enum 'bd_flags':
type BdFlags uint32

const (
	BRIDGE_API_FLAG_NONE     BdFlags = 0
	BRIDGE_API_FLAG_LEARN    BdFlags = 1
	BRIDGE_API_FLAG_FWD      BdFlags = 2
	BRIDGE_API_FLAG_FLOOD    BdFlags = 4
	BRIDGE_API_FLAG_UU_FLOOD BdFlags = 8
	BRIDGE_API_FLAG_ARP_TERM BdFlags = 16
	BRIDGE_API_FLAG_ARP_UFWD BdFlags = 32
)

// L2PortType represents VPP binary API enum 'l2_port_type':
type L2PortType uint32

const (
	L2_API_PORT_TYPE_NORMAL L2PortType = 0
	L2_API_PORT_TYPE_BVI    L2PortType = 1
	L2_API_PORT_TYPE_UU_FWD L2PortType = 2
)

/* Aliases */

// IP4Address represents VPP binary API alias 'ip4_address':
type IP4Address [4]uint8

// IP6Address represents VPP binary API alias 'ip6_address':
type IP6Address [16]uint8

// MacAddress represents VPP binary API alias 'mac_address':
type MacAddress [6]uint8

/* Types */

// Address represents VPP binary API type 'address':
type Address struct {
	Af AddressFamily
	Un AddressUnion
}

func (*Address) GetTypeName() string {
	return "address"
}
func (*Address) GetCrcString() string {
	return "09f11671"
}

// BridgeDomainSwIf represents VPP binary API type 'bridge_domain_sw_if':
type BridgeDomainSwIf struct {
	Context   uint32
	SwIfIndex uint32
	Shg       uint8
}

func (*BridgeDomainSwIf) GetTypeName() string {
	return "bridge_domain_sw_if"
}
func (*BridgeDomainSwIf) GetCrcString() string {
	return "a06dd426"
}

// IP4Prefix represents VPP binary API type 'ip4_prefix':
type IP4Prefix struct {
	Prefix IP4Address
	Len    uint8
}

func (*IP4Prefix) GetTypeName() string {
	return "ip4_prefix"
}
func (*IP4Prefix) GetCrcString() string {
	return "ea8dc11d"
}

// IP6Prefix represents VPP binary API type 'ip6_prefix':
type IP6Prefix struct {
	Prefix IP6Address
	Len    uint8
}

func (*IP6Prefix) GetTypeName() string {
	return "ip6_prefix"
}
func (*IP6Prefix) GetCrcString() string {
	return "779fd64f"
}

// MacEntry represents VPP binary API type 'mac_entry':
type MacEntry struct {
	SwIfIndex uint32
	MacAddr   []byte `struc:"[6]byte"`
	Action    uint8
	Flags     uint8
}

func (*MacEntry) GetTypeName() string {
	return "mac_entry"
}
func (*MacEntry) GetCrcString() string {
	return "971135b8"
}

// Mprefix represents VPP binary API type 'mprefix':
type Mprefix struct {
	Af               AddressFamily
	GrpAddressLength uint16
	GrpAddress       AddressUnion
	SrcAddress       AddressUnion
}

func (*Mprefix) GetTypeName() string {
	return "mprefix"
}
func (*Mprefix) GetCrcString() string {
	return "1c4cba05"
}

// Prefix represents VPP binary API type 'prefix':
type Prefix struct {
	Address       Address
	AddressLength uint8
}

func (*Prefix) GetTypeName() string {
	return "prefix"
}
func (*Prefix) GetCrcString() string {
	return "0403aebc"
}

/* Unions */

// AddressUnion represents VPP binary API union 'address_union':
type AddressUnion struct {
	Union_data [16]byte
}

func (*AddressUnion) GetTypeName() string {
	return "address_union"
}
func (*AddressUnion) GetCrcString() string {
	return "d68a2fb4"
}

func AddressUnionIP4(a IP4Address) (u AddressUnion) {
	u.SetIP4(a)
	return
}
func (u *AddressUnion) SetIP4(a IP4Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.Union_data[:], b.Bytes())
}
func (u *AddressUnion) GetIP4() (a IP4Address) {
	var b = bytes.NewReader(u.Union_data[:])
	struc.Unpack(b, &a)
	return
}

func AddressUnionIP6(a IP6Address) (u AddressUnion) {
	u.SetIP6(a)
	return
}
func (u *AddressUnion) SetIP6(a IP6Address) {
	var b = new(bytes.Buffer)
	if err := struc.Pack(b, &a); err != nil {
		return
	}
	copy(u.Union_data[:], b.Bytes())
}
func (u *AddressUnion) GetIP6() (a IP6Address) {
	var b = bytes.NewReader(u.Union_data[:])
	struc.Unpack(b, &a)
	return
}

/* Messages */

// BdIPMacAddDel represents VPP binary API message 'bd_ip_mac_add_del':
type BdIPMacAddDel struct {
	BdID  uint32
	IsAdd uint8
	IP    Address
	Mac   MacAddress
}

func (*BdIPMacAddDel) GetMessageName() string {
	return "bd_ip_mac_add_del"
}
func (*BdIPMacAddDel) GetCrcString() string {
	return "fb873902"
}
func (*BdIPMacAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BdIPMacAddDelReply represents VPP binary API message 'bd_ip_mac_add_del_reply':
type BdIPMacAddDelReply struct {
	Retval int32
}

func (*BdIPMacAddDelReply) GetMessageName() string {
	return "bd_ip_mac_add_del_reply"
}
func (*BdIPMacAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BdIPMacAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BdIPMacDetails represents VPP binary API message 'bd_ip_mac_details':
type BdIPMacDetails struct {
	BdID       uint32
	IsIPv6     uint8
	IPAddress  []byte `struc:"[16]byte"`
	MacAddress MacAddress
}

func (*BdIPMacDetails) GetMessageName() string {
	return "bd_ip_mac_details"
}
func (*BdIPMacDetails) GetCrcString() string {
	return "c05c27de"
}
func (*BdIPMacDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BdIPMacDump represents VPP binary API message 'bd_ip_mac_dump':
type BdIPMacDump struct {
	BdID uint32
}

func (*BdIPMacDump) GetMessageName() string {
	return "bd_ip_mac_dump"
}
func (*BdIPMacDump) GetCrcString() string {
	return "c25fdce6"
}
func (*BdIPMacDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BdIPMacFlush represents VPP binary API message 'bd_ip_mac_flush':
type BdIPMacFlush struct {
	BdID uint32
}

func (*BdIPMacFlush) GetMessageName() string {
	return "bd_ip_mac_flush"
}
func (*BdIPMacFlush) GetCrcString() string {
	return "c25fdce6"
}
func (*BdIPMacFlush) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BdIPMacFlushReply represents VPP binary API message 'bd_ip_mac_flush_reply':
type BdIPMacFlushReply struct {
	Retval int32
}

func (*BdIPMacFlushReply) GetMessageName() string {
	return "bd_ip_mac_flush_reply"
}
func (*BdIPMacFlushReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BdIPMacFlushReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BridgeDomainAddDel represents VPP binary API message 'bridge_domain_add_del':
type BridgeDomainAddDel struct {
	BdID    uint32
	Flood   uint8
	UuFlood uint8
	Forward uint8
	Learn   uint8
	ArpTerm uint8
	ArpUfwd uint8
	MacAge  uint8
	BdTag   []byte `struc:"[64]byte"`
	IsAdd   uint8
}

func (*BridgeDomainAddDel) GetMessageName() string {
	return "bridge_domain_add_del"
}
func (*BridgeDomainAddDel) GetCrcString() string {
	return "c6360720"
}
func (*BridgeDomainAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BridgeDomainAddDelReply represents VPP binary API message 'bridge_domain_add_del_reply':
type BridgeDomainAddDelReply struct {
	Retval int32
}

func (*BridgeDomainAddDelReply) GetMessageName() string {
	return "bridge_domain_add_del_reply"
}
func (*BridgeDomainAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BridgeDomainAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BridgeDomainDetails represents VPP binary API message 'bridge_domain_details':
type BridgeDomainDetails struct {
	BdID           uint32
	Flood          uint8
	UuFlood        uint8
	Forward        uint8
	Learn          uint8
	ArpTerm        uint8
	ArpUfwd        uint8
	MacAge         uint8
	BdTag          []byte `struc:"[64]byte"`
	BviSwIfIndex   uint32
	UuFwdSwIfIndex uint32
	NSwIfs         uint32 `struc:"sizeof=SwIfDetails"`
	SwIfDetails    []BridgeDomainSwIf
}

func (*BridgeDomainDetails) GetMessageName() string {
	return "bridge_domain_details"
}
func (*BridgeDomainDetails) GetCrcString() string {
	return "53fd2ff5"
}
func (*BridgeDomainDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BridgeDomainDump represents VPP binary API message 'bridge_domain_dump':
type BridgeDomainDump struct {
	BdID uint32
}

func (*BridgeDomainDump) GetMessageName() string {
	return "bridge_domain_dump"
}
func (*BridgeDomainDump) GetCrcString() string {
	return "c25fdce6"
}
func (*BridgeDomainDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BridgeDomainSetMacAge represents VPP binary API message 'bridge_domain_set_mac_age':
type BridgeDomainSetMacAge struct {
	BdID   uint32
	MacAge uint8
}

func (*BridgeDomainSetMacAge) GetMessageName() string {
	return "bridge_domain_set_mac_age"
}
func (*BridgeDomainSetMacAge) GetCrcString() string {
	return "b537ad7b"
}
func (*BridgeDomainSetMacAge) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BridgeDomainSetMacAgeReply represents VPP binary API message 'bridge_domain_set_mac_age_reply':
type BridgeDomainSetMacAgeReply struct {
	Retval int32
}

func (*BridgeDomainSetMacAgeReply) GetMessageName() string {
	return "bridge_domain_set_mac_age_reply"
}
func (*BridgeDomainSetMacAgeReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BridgeDomainSetMacAgeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BridgeFlags represents VPP binary API message 'bridge_flags':
type BridgeFlags struct {
	BdID  uint32
	IsSet uint8
	Flags BdFlags
}

func (*BridgeFlags) GetMessageName() string {
	return "bridge_flags"
}
func (*BridgeFlags) GetCrcString() string {
	return "8563d406"
}
func (*BridgeFlags) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BridgeFlagsReply represents VPP binary API message 'bridge_flags_reply':
type BridgeFlagsReply struct {
	Retval                 int32
	ResultingFeatureBitmap uint32
}

func (*BridgeFlagsReply) GetMessageName() string {
	return "bridge_flags_reply"
}
func (*BridgeFlagsReply) GetCrcString() string {
	return "29b2a2b3"
}
func (*BridgeFlagsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BviCreate represents VPP binary API message 'bvi_create':
type BviCreate struct {
	Mac          MacAddress
	UserInstance uint32
}

func (*BviCreate) GetMessageName() string {
	return "bvi_create"
}
func (*BviCreate) GetCrcString() string {
	return "27a79e9e"
}
func (*BviCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BviCreateReply represents VPP binary API message 'bvi_create_reply':
type BviCreateReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*BviCreateReply) GetMessageName() string {
	return "bvi_create_reply"
}
func (*BviCreateReply) GetCrcString() string {
	return "fda5941f"
}
func (*BviCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// BviDelete represents VPP binary API message 'bvi_delete':
type BviDelete struct {
	SwIfIndex uint32
}

func (*BviDelete) GetMessageName() string {
	return "bvi_delete"
}
func (*BviDelete) GetCrcString() string {
	return "529cb13f"
}
func (*BviDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// BviDeleteReply represents VPP binary API message 'bvi_delete_reply':
type BviDeleteReply struct {
	Retval int32
}

func (*BviDeleteReply) GetMessageName() string {
	return "bvi_delete_reply"
}
func (*BviDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*BviDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2FibClearTable represents VPP binary API message 'l2_fib_clear_table':
type L2FibClearTable struct{}

func (*L2FibClearTable) GetMessageName() string {
	return "l2_fib_clear_table"
}
func (*L2FibClearTable) GetCrcString() string {
	return "51077d14"
}
func (*L2FibClearTable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2FibClearTableReply represents VPP binary API message 'l2_fib_clear_table_reply':
type L2FibClearTableReply struct {
	Retval int32
}

func (*L2FibClearTableReply) GetMessageName() string {
	return "l2_fib_clear_table_reply"
}
func (*L2FibClearTableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*L2FibClearTableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2FibTableDetails represents VPP binary API message 'l2_fib_table_details':
type L2FibTableDetails struct {
	BdID      uint32
	Mac       []byte `struc:"[6]byte"`
	SwIfIndex uint32
	StaticMac uint8
	FilterMac uint8
	BviMac    uint8
}

func (*L2FibTableDetails) GetMessageName() string {
	return "l2_fib_table_details"
}
func (*L2FibTableDetails) GetCrcString() string {
	return "c7392706"
}
func (*L2FibTableDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2FibTableDump represents VPP binary API message 'l2_fib_table_dump':
type L2FibTableDump struct {
	BdID uint32
}

func (*L2FibTableDump) GetMessageName() string {
	return "l2_fib_table_dump"
}
func (*L2FibTableDump) GetCrcString() string {
	return "c25fdce6"
}
func (*L2FibTableDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2Flags represents VPP binary API message 'l2_flags':
type L2Flags struct {
	SwIfIndex     uint32
	IsSet         uint8
	FeatureBitmap uint32
}

func (*L2Flags) GetMessageName() string {
	return "l2_flags"
}
func (*L2Flags) GetCrcString() string {
	return "0e889fb9"
}
func (*L2Flags) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2FlagsReply represents VPP binary API message 'l2_flags_reply':
type L2FlagsReply struct {
	Retval                 int32
	ResultingFeatureBitmap uint32
}

func (*L2FlagsReply) GetMessageName() string {
	return "l2_flags_reply"
}
func (*L2FlagsReply) GetCrcString() string {
	return "29b2a2b3"
}
func (*L2FlagsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2InterfaceEfpFilter represents VPP binary API message 'l2_interface_efp_filter':
type L2InterfaceEfpFilter struct {
	SwIfIndex     uint32
	EnableDisable uint8
}

func (*L2InterfaceEfpFilter) GetMessageName() string {
	return "l2_interface_efp_filter"
}
func (*L2InterfaceEfpFilter) GetCrcString() string {
	return "69d24598"
}
func (*L2InterfaceEfpFilter) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2InterfaceEfpFilterReply represents VPP binary API message 'l2_interface_efp_filter_reply':
type L2InterfaceEfpFilterReply struct {
	Retval int32
}

func (*L2InterfaceEfpFilterReply) GetMessageName() string {
	return "l2_interface_efp_filter_reply"
}
func (*L2InterfaceEfpFilterReply) GetCrcString() string {
	return "e8d4e804"
}
func (*L2InterfaceEfpFilterReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2InterfacePbbTagRewrite represents VPP binary API message 'l2_interface_pbb_tag_rewrite':
type L2InterfacePbbTagRewrite struct {
	SwIfIndex uint32
	VtrOp     uint32
	OuterTag  uint16
	BDmac     []byte `struc:"[6]byte"`
	BSmac     []byte `struc:"[6]byte"`
	BVlanid   uint16
	ISid      uint32
}

func (*L2InterfacePbbTagRewrite) GetMessageName() string {
	return "l2_interface_pbb_tag_rewrite"
}
func (*L2InterfacePbbTagRewrite) GetCrcString() string {
	return "6cf815f9"
}
func (*L2InterfacePbbTagRewrite) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2InterfacePbbTagRewriteReply represents VPP binary API message 'l2_interface_pbb_tag_rewrite_reply':
type L2InterfacePbbTagRewriteReply struct {
	Retval int32
}

func (*L2InterfacePbbTagRewriteReply) GetMessageName() string {
	return "l2_interface_pbb_tag_rewrite_reply"
}
func (*L2InterfacePbbTagRewriteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*L2InterfacePbbTagRewriteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2InterfaceVlanTagRewrite represents VPP binary API message 'l2_interface_vlan_tag_rewrite':
type L2InterfaceVlanTagRewrite struct {
	SwIfIndex uint32
	VtrOp     uint32
	PushDot1q uint32
	Tag1      uint32
	Tag2      uint32
}

func (*L2InterfaceVlanTagRewrite) GetMessageName() string {
	return "l2_interface_vlan_tag_rewrite"
}
func (*L2InterfaceVlanTagRewrite) GetCrcString() string {
	return "b90be6b4"
}
func (*L2InterfaceVlanTagRewrite) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2InterfaceVlanTagRewriteReply represents VPP binary API message 'l2_interface_vlan_tag_rewrite_reply':
type L2InterfaceVlanTagRewriteReply struct {
	Retval int32
}

func (*L2InterfaceVlanTagRewriteReply) GetMessageName() string {
	return "l2_interface_vlan_tag_rewrite_reply"
}
func (*L2InterfaceVlanTagRewriteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*L2InterfaceVlanTagRewriteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2MacsEvent represents VPP binary API message 'l2_macs_event':
type L2MacsEvent struct {
	PID   uint32
	NMacs uint32 `struc:"sizeof=Mac"`
	Mac   []MacEntry
}

func (*L2MacsEvent) GetMessageName() string {
	return "l2_macs_event"
}
func (*L2MacsEvent) GetCrcString() string {
	return "4e5ab0c8"
}
func (*L2MacsEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

// L2PatchAddDel represents VPP binary API message 'l2_patch_add_del':
type L2PatchAddDel struct {
	RxSwIfIndex uint32
	TxSwIfIndex uint32
	IsAdd       uint8
}

func (*L2PatchAddDel) GetMessageName() string {
	return "l2_patch_add_del"
}
func (*L2PatchAddDel) GetCrcString() string {
	return "62506e63"
}
func (*L2PatchAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2PatchAddDelReply represents VPP binary API message 'l2_patch_add_del_reply':
type L2PatchAddDelReply struct {
	Retval int32
}

func (*L2PatchAddDelReply) GetMessageName() string {
	return "l2_patch_add_del_reply"
}
func (*L2PatchAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*L2PatchAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2XconnectDetails represents VPP binary API message 'l2_xconnect_details':
type L2XconnectDetails struct {
	RxSwIfIndex uint32
	TxSwIfIndex uint32
}

func (*L2XconnectDetails) GetMessageName() string {
	return "l2_xconnect_details"
}
func (*L2XconnectDetails) GetCrcString() string {
	return "722e2378"
}
func (*L2XconnectDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2XconnectDump represents VPP binary API message 'l2_xconnect_dump':
type L2XconnectDump struct{}

func (*L2XconnectDump) GetMessageName() string {
	return "l2_xconnect_dump"
}
func (*L2XconnectDump) GetCrcString() string {
	return "51077d14"
}
func (*L2XconnectDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2fibAddDel represents VPP binary API message 'l2fib_add_del':
type L2fibAddDel struct {
	Mac       []byte `struc:"[6]byte"`
	BdID      uint32
	SwIfIndex uint32
	IsAdd     uint8
	StaticMac uint8
	FilterMac uint8
	BviMac    uint8
}

func (*L2fibAddDel) GetMessageName() string {
	return "l2fib_add_del"
}
func (*L2fibAddDel) GetCrcString() string {
	return "34ced3eb"
}
func (*L2fibAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2fibAddDelReply represents VPP binary API message 'l2fib_add_del_reply':
type L2fibAddDelReply struct {
	Retval int32
}

func (*L2fibAddDelReply) GetMessageName() string {
	return "l2fib_add_del_reply"
}
func (*L2fibAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*L2fibAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2fibFlushAll represents VPP binary API message 'l2fib_flush_all':
type L2fibFlushAll struct{}

func (*L2fibFlushAll) GetMessageName() string {
	return "l2fib_flush_all"
}
func (*L2fibFlushAll) GetCrcString() string {
	return "51077d14"
}
func (*L2fibFlushAll) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2fibFlushAllReply represents VPP binary API message 'l2fib_flush_all_reply':
type L2fibFlushAllReply struct {
	Retval int32
}

func (*L2fibFlushAllReply) GetMessageName() string {
	return "l2fib_flush_all_reply"
}
func (*L2fibFlushAllReply) GetCrcString() string {
	return "e8d4e804"
}
func (*L2fibFlushAllReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2fibFlushBd represents VPP binary API message 'l2fib_flush_bd':
type L2fibFlushBd struct {
	BdID uint32
}

func (*L2fibFlushBd) GetMessageName() string {
	return "l2fib_flush_bd"
}
func (*L2fibFlushBd) GetCrcString() string {
	return "c25fdce6"
}
func (*L2fibFlushBd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2fibFlushBdReply represents VPP binary API message 'l2fib_flush_bd_reply':
type L2fibFlushBdReply struct {
	Retval int32
}

func (*L2fibFlushBdReply) GetMessageName() string {
	return "l2fib_flush_bd_reply"
}
func (*L2fibFlushBdReply) GetCrcString() string {
	return "e8d4e804"
}
func (*L2fibFlushBdReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// L2fibFlushInt represents VPP binary API message 'l2fib_flush_int':
type L2fibFlushInt struct {
	SwIfIndex uint32
}

func (*L2fibFlushInt) GetMessageName() string {
	return "l2fib_flush_int"
}
func (*L2fibFlushInt) GetCrcString() string {
	return "529cb13f"
}
func (*L2fibFlushInt) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// L2fibFlushIntReply represents VPP binary API message 'l2fib_flush_int_reply':
type L2fibFlushIntReply struct {
	Retval int32
}

func (*L2fibFlushIntReply) GetMessageName() string {
	return "l2fib_flush_int_reply"
}
func (*L2fibFlushIntReply) GetCrcString() string {
	return "e8d4e804"
}
func (*L2fibFlushIntReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceSetL2Bridge represents VPP binary API message 'sw_interface_set_l2_bridge':
type SwInterfaceSetL2Bridge struct {
	RxSwIfIndex uint32
	BdID        uint32
	PortType    L2PortType
	Shg         uint8
	Enable      uint8
}

func (*SwInterfaceSetL2Bridge) GetMessageName() string {
	return "sw_interface_set_l2_bridge"
}
func (*SwInterfaceSetL2Bridge) GetCrcString() string {
	return "2af7795e"
}
func (*SwInterfaceSetL2Bridge) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SwInterfaceSetL2BridgeReply represents VPP binary API message 'sw_interface_set_l2_bridge_reply':
type SwInterfaceSetL2BridgeReply struct {
	Retval int32
}

func (*SwInterfaceSetL2BridgeReply) GetMessageName() string {
	return "sw_interface_set_l2_bridge_reply"
}
func (*SwInterfaceSetL2BridgeReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SwInterfaceSetL2BridgeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceSetL2Xconnect represents VPP binary API message 'sw_interface_set_l2_xconnect':
type SwInterfaceSetL2Xconnect struct {
	RxSwIfIndex uint32
	TxSwIfIndex uint32
	Enable      uint8
}

func (*SwInterfaceSetL2Xconnect) GetMessageName() string {
	return "sw_interface_set_l2_xconnect"
}
func (*SwInterfaceSetL2Xconnect) GetCrcString() string {
	return "95de3988"
}
func (*SwInterfaceSetL2Xconnect) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SwInterfaceSetL2XconnectReply represents VPP binary API message 'sw_interface_set_l2_xconnect_reply':
type SwInterfaceSetL2XconnectReply struct {
	Retval int32
}

func (*SwInterfaceSetL2XconnectReply) GetMessageName() string {
	return "sw_interface_set_l2_xconnect_reply"
}
func (*SwInterfaceSetL2XconnectReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SwInterfaceSetL2XconnectReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceSetVpath represents VPP binary API message 'sw_interface_set_vpath':
type SwInterfaceSetVpath struct {
	SwIfIndex uint32
	Enable    uint8
}

func (*SwInterfaceSetVpath) GetMessageName() string {
	return "sw_interface_set_vpath"
}
func (*SwInterfaceSetVpath) GetCrcString() string {
	return "a36fadc0"
}
func (*SwInterfaceSetVpath) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SwInterfaceSetVpathReply represents VPP binary API message 'sw_interface_set_vpath_reply':
type SwInterfaceSetVpathReply struct {
	Retval int32
}

func (*SwInterfaceSetVpathReply) GetMessageName() string {
	return "sw_interface_set_vpath_reply"
}
func (*SwInterfaceSetVpathReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SwInterfaceSetVpathReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// WantL2MacsEvents represents VPP binary API message 'want_l2_macs_events':
type WantL2MacsEvents struct {
	LearnLimit     uint32
	ScanDelay      uint8
	MaxMacsInEvent uint8
	EnableDisable  uint8
	PID            uint32
}

func (*WantL2MacsEvents) GetMessageName() string {
	return "want_l2_macs_events"
}
func (*WantL2MacsEvents) GetCrcString() string {
	return "94e63394"
}
func (*WantL2MacsEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// WantL2MacsEventsReply represents VPP binary API message 'want_l2_macs_events_reply':
type WantL2MacsEventsReply struct {
	Retval int32
}

func (*WantL2MacsEventsReply) GetMessageName() string {
	return "want_l2_macs_events_reply"
}
func (*WantL2MacsEventsReply) GetCrcString() string {
	return "e8d4e804"
}
func (*WantL2MacsEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*BdIPMacAddDel)(nil), "l2.BdIPMacAddDel")
	api.RegisterMessage((*BdIPMacAddDelReply)(nil), "l2.BdIPMacAddDelReply")
	api.RegisterMessage((*BdIPMacDetails)(nil), "l2.BdIPMacDetails")
	api.RegisterMessage((*BdIPMacDump)(nil), "l2.BdIPMacDump")
	api.RegisterMessage((*BdIPMacFlush)(nil), "l2.BdIPMacFlush")
	api.RegisterMessage((*BdIPMacFlushReply)(nil), "l2.BdIPMacFlushReply")
	api.RegisterMessage((*BridgeDomainAddDel)(nil), "l2.BridgeDomainAddDel")
	api.RegisterMessage((*BridgeDomainAddDelReply)(nil), "l2.BridgeDomainAddDelReply")
	api.RegisterMessage((*BridgeDomainDetails)(nil), "l2.BridgeDomainDetails")
	api.RegisterMessage((*BridgeDomainDump)(nil), "l2.BridgeDomainDump")
	api.RegisterMessage((*BridgeDomainSetMacAge)(nil), "l2.BridgeDomainSetMacAge")
	api.RegisterMessage((*BridgeDomainSetMacAgeReply)(nil), "l2.BridgeDomainSetMacAgeReply")
	api.RegisterMessage((*BridgeFlags)(nil), "l2.BridgeFlags")
	api.RegisterMessage((*BridgeFlagsReply)(nil), "l2.BridgeFlagsReply")
	api.RegisterMessage((*BviCreate)(nil), "l2.BviCreate")
	api.RegisterMessage((*BviCreateReply)(nil), "l2.BviCreateReply")
	api.RegisterMessage((*BviDelete)(nil), "l2.BviDelete")
	api.RegisterMessage((*BviDeleteReply)(nil), "l2.BviDeleteReply")
	api.RegisterMessage((*L2FibClearTable)(nil), "l2.L2FibClearTable")
	api.RegisterMessage((*L2FibClearTableReply)(nil), "l2.L2FibClearTableReply")
	api.RegisterMessage((*L2FibTableDetails)(nil), "l2.L2FibTableDetails")
	api.RegisterMessage((*L2FibTableDump)(nil), "l2.L2FibTableDump")
	api.RegisterMessage((*L2Flags)(nil), "l2.L2Flags")
	api.RegisterMessage((*L2FlagsReply)(nil), "l2.L2FlagsReply")
	api.RegisterMessage((*L2InterfaceEfpFilter)(nil), "l2.L2InterfaceEfpFilter")
	api.RegisterMessage((*L2InterfaceEfpFilterReply)(nil), "l2.L2InterfaceEfpFilterReply")
	api.RegisterMessage((*L2InterfacePbbTagRewrite)(nil), "l2.L2InterfacePbbTagRewrite")
	api.RegisterMessage((*L2InterfacePbbTagRewriteReply)(nil), "l2.L2InterfacePbbTagRewriteReply")
	api.RegisterMessage((*L2InterfaceVlanTagRewrite)(nil), "l2.L2InterfaceVlanTagRewrite")
	api.RegisterMessage((*L2InterfaceVlanTagRewriteReply)(nil), "l2.L2InterfaceVlanTagRewriteReply")
	api.RegisterMessage((*L2MacsEvent)(nil), "l2.L2MacsEvent")
	api.RegisterMessage((*L2PatchAddDel)(nil), "l2.L2PatchAddDel")
	api.RegisterMessage((*L2PatchAddDelReply)(nil), "l2.L2PatchAddDelReply")
	api.RegisterMessage((*L2XconnectDetails)(nil), "l2.L2XconnectDetails")
	api.RegisterMessage((*L2XconnectDump)(nil), "l2.L2XconnectDump")
	api.RegisterMessage((*L2fibAddDel)(nil), "l2.L2fibAddDel")
	api.RegisterMessage((*L2fibAddDelReply)(nil), "l2.L2fibAddDelReply")
	api.RegisterMessage((*L2fibFlushAll)(nil), "l2.L2fibFlushAll")
	api.RegisterMessage((*L2fibFlushAllReply)(nil), "l2.L2fibFlushAllReply")
	api.RegisterMessage((*L2fibFlushBd)(nil), "l2.L2fibFlushBd")
	api.RegisterMessage((*L2fibFlushBdReply)(nil), "l2.L2fibFlushBdReply")
	api.RegisterMessage((*L2fibFlushInt)(nil), "l2.L2fibFlushInt")
	api.RegisterMessage((*L2fibFlushIntReply)(nil), "l2.L2fibFlushIntReply")
	api.RegisterMessage((*SwInterfaceSetL2Bridge)(nil), "l2.SwInterfaceSetL2Bridge")
	api.RegisterMessage((*SwInterfaceSetL2BridgeReply)(nil), "l2.SwInterfaceSetL2BridgeReply")
	api.RegisterMessage((*SwInterfaceSetL2Xconnect)(nil), "l2.SwInterfaceSetL2Xconnect")
	api.RegisterMessage((*SwInterfaceSetL2XconnectReply)(nil), "l2.SwInterfaceSetL2XconnectReply")
	api.RegisterMessage((*SwInterfaceSetVpath)(nil), "l2.SwInterfaceSetVpath")
	api.RegisterMessage((*SwInterfaceSetVpathReply)(nil), "l2.SwInterfaceSetVpathReply")
	api.RegisterMessage((*WantL2MacsEvents)(nil), "l2.WantL2MacsEvents")
	api.RegisterMessage((*WantL2MacsEventsReply)(nil), "l2.WantL2MacsEventsReply")
}

var Messages = []api.Message{
	(*BdIPMacAddDel)(nil),
	(*BdIPMacAddDelReply)(nil),
	(*BdIPMacDetails)(nil),
	(*BdIPMacDump)(nil),
	(*BdIPMacFlush)(nil),
	(*BdIPMacFlushReply)(nil),
	(*BridgeDomainAddDel)(nil),
	(*BridgeDomainAddDelReply)(nil),
	(*BridgeDomainDetails)(nil),
	(*BridgeDomainDump)(nil),
	(*BridgeDomainSetMacAge)(nil),
	(*BridgeDomainSetMacAgeReply)(nil),
	(*BridgeFlags)(nil),
	(*BridgeFlagsReply)(nil),
	(*BviCreate)(nil),
	(*BviCreateReply)(nil),
	(*BviDelete)(nil),
	(*BviDeleteReply)(nil),
	(*L2FibClearTable)(nil),
	(*L2FibClearTableReply)(nil),
	(*L2FibTableDetails)(nil),
	(*L2FibTableDump)(nil),
	(*L2Flags)(nil),
	(*L2FlagsReply)(nil),
	(*L2InterfaceEfpFilter)(nil),
	(*L2InterfaceEfpFilterReply)(nil),
	(*L2InterfacePbbTagRewrite)(nil),
	(*L2InterfacePbbTagRewriteReply)(nil),
	(*L2InterfaceVlanTagRewrite)(nil),
	(*L2InterfaceVlanTagRewriteReply)(nil),
	(*L2MacsEvent)(nil),
	(*L2PatchAddDel)(nil),
	(*L2PatchAddDelReply)(nil),
	(*L2XconnectDetails)(nil),
	(*L2XconnectDump)(nil),
	(*L2fibAddDel)(nil),
	(*L2fibAddDelReply)(nil),
	(*L2fibFlushAll)(nil),
	(*L2fibFlushAllReply)(nil),
	(*L2fibFlushBd)(nil),
	(*L2fibFlushBdReply)(nil),
	(*L2fibFlushInt)(nil),
	(*L2fibFlushIntReply)(nil),
	(*SwInterfaceSetL2Bridge)(nil),
	(*SwInterfaceSetL2BridgeReply)(nil),
	(*SwInterfaceSetL2Xconnect)(nil),
	(*SwInterfaceSetL2XconnectReply)(nil),
	(*SwInterfaceSetVpath)(nil),
	(*SwInterfaceSetVpathReply)(nil),
	(*WantL2MacsEvents)(nil),
	(*WantL2MacsEventsReply)(nil),
}
