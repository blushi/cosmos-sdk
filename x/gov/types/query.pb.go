// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/gov/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryProposalRequest is the request type for the Query/Proposal RPC method
type QueryAllProposalsRequest struct {
	// status of the proposals.
	ProposalStatus int32 `protobuf:"varint,1,opt,name=proposal_status,json=proposalStatus,proto3" json:"proposal_status,omitempty"`
	// Voter address for the proposals.
	Voter github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=voter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"voter,omitempty"`
	// Deposit addresses from the proposals.
	Depositor github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=depositor,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"depositor,omitempty"`
	Req       *query.PageRequest                            `protobuf:"bytes,4,opt,name=req,proto3" json:"req,omitempty"`
}

func (m *QueryAllProposalsRequest) Reset()         { *m = QueryAllProposalsRequest{} }
func (m *QueryAllProposalsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllProposalsRequest) ProtoMessage()    {}
func (*QueryAllProposalsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6efb1c1bc2595eda, []int{0}
}
func (m *QueryAllProposalsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllProposalsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllProposalsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllProposalsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllProposalsRequest.Merge(m, src)
}
func (m *QueryAllProposalsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllProposalsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllProposalsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllProposalsRequest proto.InternalMessageInfo

func (m *QueryAllProposalsRequest) GetProposalStatus() int32 {
	if m != nil {
		return m.ProposalStatus
	}
	return 0
}

func (m *QueryAllProposalsRequest) GetVoter() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *QueryAllProposalsRequest) GetDepositor() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Depositor
	}
	return nil
}

func (m *QueryAllProposalsRequest) GetReq() *query.PageRequest {
	if m != nil {
		return m.Req
	}
	return nil
}

type QueryAllProposalsResponse struct {
	Proposals []byte              `protobuf:"bytes,1,opt,name=proposals,proto3,castrepeated=github.com/cosmos/cosmos-sdk/x/gov/types.Proposal" json:"proposals,omitempty"`
	Res       *query.PageResponse `protobuf:"bytes,2,opt,name=res,proto3" json:"res,omitempty"`
}

func (m *QueryAllProposalsResponse) Reset()         { *m = QueryAllProposalsResponse{} }
func (m *QueryAllProposalsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllProposalsResponse) ProtoMessage()    {}
func (*QueryAllProposalsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6efb1c1bc2595eda, []int{1}
}
func (m *QueryAllProposalsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllProposalsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllProposalsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllProposalsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllProposalsResponse.Merge(m, src)
}
func (m *QueryAllProposalsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllProposalsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllProposalsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllProposalsResponse proto.InternalMessageInfo

func (m *QueryAllProposalsResponse) GetProposals() []byte {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *QueryAllProposalsResponse) GetRes() *query.PageResponse {
	if m != nil {
		return m.Res
	}
	return nil
}

type QueryVotesRequest struct {
	// unique id of the proposal
	ProposalId uint64             `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Req        *query.PageRequest `protobuf:"bytes,2,opt,name=req,proto3" json:"req,omitempty"`
}

func (m *QueryVotesRequest) Reset()         { *m = QueryVotesRequest{} }
func (m *QueryVotesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryVotesRequest) ProtoMessage()    {}
func (*QueryVotesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6efb1c1bc2595eda, []int{2}
}
func (m *QueryVotesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVotesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVotesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVotesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVotesRequest.Merge(m, src)
}
func (m *QueryVotesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryVotesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVotesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVotesRequest proto.InternalMessageInfo

func (m *QueryVotesRequest) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *QueryVotesRequest) GetReq() *query.PageRequest {
	if m != nil {
		return m.Req
	}
	return nil
}

type QueryVotesResponse struct {
	Votes []byte              `protobuf:"bytes,1,opt,name=votes,proto3,castrepeated=github.com/cosmos/cosmos-sdk/x/gov/types.Vote" json:"votes,omitempty"`
	Res   *query.PageResponse `protobuf:"bytes,2,opt,name=res,proto3" json:"res,omitempty"`
}

func (m *QueryVotesResponse) Reset()         { *m = QueryVotesResponse{} }
func (m *QueryVotesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryVotesResponse) ProtoMessage()    {}
func (*QueryVotesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6efb1c1bc2595eda, []int{3}
}
func (m *QueryVotesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVotesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVotesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVotesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVotesResponse.Merge(m, src)
}
func (m *QueryVotesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryVotesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVotesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVotesResponse proto.InternalMessageInfo

func (m *QueryVotesResponse) GetVotes() []byte {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *QueryVotesResponse) GetRes() *query.PageResponse {
	if m != nil {
		return m.Res
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAllProposalsRequest)(nil), "cosmos.gov.QueryAllProposalsRequest")
	proto.RegisterType((*QueryAllProposalsResponse)(nil), "cosmos.gov.QueryAllProposalsResponse")
	proto.RegisterType((*QueryVotesRequest)(nil), "cosmos.gov.QueryVotesRequest")
	proto.RegisterType((*QueryVotesResponse)(nil), "cosmos.gov.QueryVotesResponse")
}

func init() { proto.RegisterFile("cosmos/gov/query.proto", fileDescriptor_6efb1c1bc2595eda) }

var fileDescriptor_6efb1c1bc2595eda = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0xce, 0xa4, 0x35, 0x52, 0x5f, 0x22, 0x10, 0x23, 0x84, 0x5c, 0x4b, 0x75, 0x22, 0x0b, 0x84,
	0x25, 0xa8, 0xad, 0x04, 0x71, 0x80, 0x78, 0x03, 0xac, 0x28, 0xae, 0xc4, 0x02, 0x09, 0x21, 0xd7,
	0x1e, 0x19, 0x8b, 0x34, 0xcf, 0xf1, 0x1b, 0x47, 0xf4, 0x0a, 0xb0, 0xe1, 0x04, 0x1c, 0xa0, 0x2b,
	0x8e, 0xc1, 0xb2, 0x4b, 0x56, 0x14, 0x25, 0xb7, 0x60, 0x85, 0x3c, 0x63, 0x37, 0x96, 0x68, 0x69,
	0x60, 0xe5, 0xd1, 0x37, 0x6f, 0xbe, 0xf9, 0x7e, 0xc6, 0x70, 0x37, 0x46, 0x3a, 0x46, 0xf2, 0x53,
	0x5c, 0xf8, 0xf3, 0x52, 0x14, 0x27, 0x5e, 0x5e, 0xa0, 0x44, 0x0e, 0x1a, 0xf7, 0x52, 0x5c, 0x58,
	0x7b, 0xf5, 0x8c, 0xda, 0xf7, 0xf3, 0x28, 0xcd, 0x66, 0x91, 0xcc, 0x70, 0xa6, 0x47, 0xad, 0x3b,
	0x29, 0xa6, 0xa8, 0x96, 0x7e, 0xb5, 0xd2, 0xa8, 0xf3, 0xb1, 0x0b, 0xe6, 0xcb, 0xea, 0xc0, 0x64,
	0x3a, 0x3d, 0x28, 0x30, 0x47, 0x8a, 0xa6, 0x14, 0x8a, 0x79, 0x29, 0x48, 0xf2, 0x07, 0x70, 0x2b,
	0xaf, 0xb1, 0xb7, 0x24, 0x23, 0x59, 0x92, 0xc9, 0x86, 0xcc, 0x35, 0xc2, 0x9b, 0x0d, 0x7c, 0xa8,
	0x50, 0xfe, 0x14, 0x8c, 0x05, 0x4a, 0x51, 0x98, 0xdd, 0x21, 0x73, 0xfb, 0xc1, 0xe8, 0xd7, 0x8f,
	0xc1, 0x7e, 0x9a, 0xc9, 0x77, 0xe5, 0x91, 0x17, 0xe3, 0xb1, 0x5f, 0x0b, 0xd3, 0x9f, 0x7d, 0x4a,
	0xde, 0xfb, 0xf2, 0x24, 0x17, 0xe4, 0x4d, 0xe2, 0x78, 0x92, 0x24, 0x85, 0x20, 0x0a, 0xf5, 0x79,
	0xfe, 0x02, 0x76, 0x12, 0x91, 0x23, 0x65, 0x12, 0x0b, 0x73, 0xeb, 0x7f, 0xc9, 0xd6, 0x1c, 0xfc,
	0x21, 0x6c, 0x15, 0x62, 0x6e, 0x6e, 0x0f, 0x99, 0xdb, 0x1b, 0xef, 0x7a, 0x75, 0x5c, 0x3a, 0xc2,
	0x83, 0x28, 0x15, 0xb5, 0xd5, 0xb0, 0x9a, 0x72, 0xbe, 0x30, 0xd8, 0xbd, 0x24, 0x0c, 0xca, 0x71,
	0x46, 0x82, 0x1f, 0xc2, 0x4e, 0x63, 0x5b, 0xe7, 0xd0, 0x0f, 0x9e, 0x9c, 0x9e, 0x0f, 0x46, 0x7f,
	0xd5, 0xf6, 0x41, 0x55, 0xa6, 0x15, 0x36, 0x94, 0xe1, 0x9a, 0x87, 0x3f, 0xaa, 0xf4, 0x91, 0xca,
	0xad, 0x37, 0xb6, 0x2e, 0xd3, 0xa7, 0x6f, 0xaf, 0x04, 0x92, 0x13, 0xc1, 0x6d, 0xa5, 0xef, 0x15,
	0x4a, 0x71, 0xd1, 0xd2, 0x00, 0x7a, 0x17, 0x2d, 0x65, 0x89, 0x52, 0xb6, 0x1d, 0x42, 0x03, 0x3d,
	0x4f, 0x9a, 0x0c, 0xba, 0x1b, 0x65, 0xf0, 0x89, 0x01, 0x6f, 0xdf, 0x51, 0x9b, 0xaf, 0x1b, 0x6e,
	0x8c, 0x8f, 0x4e, 0xcf, 0xaf, 0x29, 0xa5, 0x6d, 0xbc, 0xa2, 0xd2, 0x0d, 0xff, 0xa3, 0xe1, 0xf1,
	0x57, 0x06, 0x86, 0x52, 0xc3, 0xdf, 0x40, 0xbf, 0xdd, 0x0a, 0xbf, 0xe7, 0xad, 0x9f, 0xbe, 0x77,
	0xd5, 0x0b, 0xb6, 0xee, 0x5f, 0x33, 0xa5, 0xef, 0x72, 0x3a, 0xfc, 0x19, 0x18, 0xca, 0x30, 0xdf,
	0xfb, 0xe3, 0x44, 0x3b, 0x6c, 0xcb, 0xbe, 0x6a, 0xbb, 0x61, 0x0a, 0x82, 0x6f, 0x4b, 0x9b, 0x9d,
	0x2d, 0x6d, 0xf6, 0x73, 0x69, 0xb3, 0xcf, 0x2b, 0xbb, 0x73, 0xb6, 0xb2, 0x3b, 0xdf, 0x57, 0x76,
	0xe7, 0xb5, 0xbb, 0x69, 0x60, 0x47, 0x37, 0xd4, 0xcf, 0xf9, 0xf8, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x24, 0xed, 0xc7, 0x0e, 0xf7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Proposals queries all proposals.
	AllProposals(ctx context.Context, in *QueryAllProposalsRequest, opts ...grpc.CallOption) (*QueryAllProposalsResponse, error)
	// Votes queries votes of a given proposal
	Votes(ctx context.Context, in *QueryVotesRequest, opts ...grpc.CallOption) (*QueryVotesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) AllProposals(ctx context.Context, in *QueryAllProposalsRequest, opts ...grpc.CallOption) (*QueryAllProposalsResponse, error) {
	out := new(QueryAllProposalsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.Query/AllProposals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Votes(ctx context.Context, in *QueryVotesRequest, opts ...grpc.CallOption) (*QueryVotesResponse, error) {
	out := new(QueryVotesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.Query/Votes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Proposals queries all proposals.
	AllProposals(context.Context, *QueryAllProposalsRequest) (*QueryAllProposalsResponse, error)
	// Votes queries votes of a given proposal
	Votes(context.Context, *QueryVotesRequest) (*QueryVotesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) AllProposals(ctx context.Context, req *QueryAllProposalsRequest) (*QueryAllProposalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllProposals not implemented")
}
func (*UnimplementedQueryServer) Votes(ctx context.Context, req *QueryVotesRequest) (*QueryVotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Votes not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_AllProposals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllProposalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllProposals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.Query/AllProposals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllProposals(ctx, req.(*QueryAllProposalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Votes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Votes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.Query/Votes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Votes(ctx, req.(*QueryVotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.gov.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllProposals",
			Handler:    _Query_AllProposals_Handler,
		},
		{
			MethodName: "Votes",
			Handler:    _Query_Votes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/gov/query.proto",
}

func (m *QueryAllProposalsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllProposalsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllProposalsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Req != nil {
		{
			size, err := m.Req.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalStatus != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalStatus))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllProposalsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllProposalsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllProposalsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Res != nil {
		{
			size, err := m.Res.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Proposals) > 0 {
		i -= len(m.Proposals)
		copy(dAtA[i:], m.Proposals)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Proposals)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryVotesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVotesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVotesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Req != nil {
		{
			size, err := m.Req.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryVotesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVotesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVotesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Res != nil {
		{
			size, err := m.Res.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Votes) > 0 {
		i -= len(m.Votes)
		copy(dAtA[i:], m.Votes)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Votes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAllProposalsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalStatus != 0 {
		n += 1 + sovQuery(uint64(m.ProposalStatus))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Req != nil {
		l = m.Req.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllProposalsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proposals)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Res != nil {
		l = m.Res.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryVotesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovQuery(uint64(m.ProposalId))
	}
	if m.Req != nil {
		l = m.Req.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryVotesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Votes)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Res != nil {
		l = m.Res.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAllProposalsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllProposalsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllProposalsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalStatus", wireType)
			}
			m.ProposalStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalStatus |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = append(m.Voter[:0], dAtA[iNdEx:postIndex]...)
			if m.Voter == nil {
				m.Voter = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = append(m.Depositor[:0], dAtA[iNdEx:postIndex]...)
			if m.Depositor == nil {
				m.Depositor = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Req", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Req == nil {
				m.Req = &query.PageRequest{}
			}
			if err := m.Req.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllProposalsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllProposalsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllProposalsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposals = append(m.Proposals[:0], dAtA[iNdEx:postIndex]...)
			if m.Proposals == nil {
				m.Proposals = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Res", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Res == nil {
				m.Res = &query.PageResponse{}
			}
			if err := m.Res.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryVotesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryVotesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVotesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Req", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Req == nil {
				m.Req = &query.PageRequest{}
			}
			if err := m.Req.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryVotesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryVotesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVotesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes[:0], dAtA[iNdEx:postIndex]...)
			if m.Votes == nil {
				m.Votes = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Res", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Res == nil {
				m.Res = &query.PageResponse{}
			}
			if err := m.Res.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)