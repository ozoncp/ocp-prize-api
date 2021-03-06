// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/ocp-prize-api/ocp-prize-api.proto

package ocp_prize_api

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ListPrizeV1Request struct {
	Limit                uint64   `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPrizeV1Request) Reset()         { *m = ListPrizeV1Request{} }
func (m *ListPrizeV1Request) String() string { return proto.CompactTextString(m) }
func (*ListPrizeV1Request) ProtoMessage()    {}
func (*ListPrizeV1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{0}
}

func (m *ListPrizeV1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPrizeV1Request.Unmarshal(m, b)
}
func (m *ListPrizeV1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPrizeV1Request.Marshal(b, m, deterministic)
}
func (m *ListPrizeV1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPrizeV1Request.Merge(m, src)
}
func (m *ListPrizeV1Request) XXX_Size() int {
	return xxx_messageInfo_ListPrizeV1Request.Size(m)
}
func (m *ListPrizeV1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPrizeV1Request.DiscardUnknown(m)
}

var xxx_messageInfo_ListPrizeV1Request proto.InternalMessageInfo

func (m *ListPrizeV1Request) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListPrizeV1Request) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListPrizeV1Response struct {
	Prizes               []*Prize `protobuf:"bytes,1,rep,name=prizes,proto3" json:"prizes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPrizeV1Response) Reset()         { *m = ListPrizeV1Response{} }
func (m *ListPrizeV1Response) String() string { return proto.CompactTextString(m) }
func (*ListPrizeV1Response) ProtoMessage()    {}
func (*ListPrizeV1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{1}
}

func (m *ListPrizeV1Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPrizeV1Response.Unmarshal(m, b)
}
func (m *ListPrizeV1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPrizeV1Response.Marshal(b, m, deterministic)
}
func (m *ListPrizeV1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPrizeV1Response.Merge(m, src)
}
func (m *ListPrizeV1Response) XXX_Size() int {
	return xxx_messageInfo_ListPrizeV1Response.Size(m)
}
func (m *ListPrizeV1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPrizeV1Response.DiscardUnknown(m)
}

var xxx_messageInfo_ListPrizeV1Response proto.InternalMessageInfo

func (m *ListPrizeV1Response) GetPrizes() []*Prize {
	if m != nil {
		return m.Prizes
	}
	return nil
}

type CreatePrizeV1Request struct {
	Link                 string   `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	IssueId              uint64   `protobuf:"varint,2,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePrizeV1Request) Reset()         { *m = CreatePrizeV1Request{} }
func (m *CreatePrizeV1Request) String() string { return proto.CompactTextString(m) }
func (*CreatePrizeV1Request) ProtoMessage()    {}
func (*CreatePrizeV1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{2}
}

func (m *CreatePrizeV1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePrizeV1Request.Unmarshal(m, b)
}
func (m *CreatePrizeV1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePrizeV1Request.Marshal(b, m, deterministic)
}
func (m *CreatePrizeV1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePrizeV1Request.Merge(m, src)
}
func (m *CreatePrizeV1Request) XXX_Size() int {
	return xxx_messageInfo_CreatePrizeV1Request.Size(m)
}
func (m *CreatePrizeV1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePrizeV1Request.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePrizeV1Request proto.InternalMessageInfo

func (m *CreatePrizeV1Request) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *CreatePrizeV1Request) GetIssueId() uint64 {
	if m != nil {
		return m.IssueId
	}
	return 0
}

type CreatePrizeV1Response struct {
	PrizeId              uint64   `protobuf:"varint,1,opt,name=prize_id,json=prizeId,proto3" json:"prize_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePrizeV1Response) Reset()         { *m = CreatePrizeV1Response{} }
func (m *CreatePrizeV1Response) String() string { return proto.CompactTextString(m) }
func (*CreatePrizeV1Response) ProtoMessage()    {}
func (*CreatePrizeV1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{3}
}

func (m *CreatePrizeV1Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePrizeV1Response.Unmarshal(m, b)
}
func (m *CreatePrizeV1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePrizeV1Response.Marshal(b, m, deterministic)
}
func (m *CreatePrizeV1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePrizeV1Response.Merge(m, src)
}
func (m *CreatePrizeV1Response) XXX_Size() int {
	return xxx_messageInfo_CreatePrizeV1Response.Size(m)
}
func (m *CreatePrizeV1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePrizeV1Response.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePrizeV1Response proto.InternalMessageInfo

func (m *CreatePrizeV1Response) GetPrizeId() uint64 {
	if m != nil {
		return m.PrizeId
	}
	return 0
}

type MultiCreatePrizeV1Request struct {
	Prizes               []*Prize `protobuf:"bytes,1,rep,name=prizes,proto3" json:"prizes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiCreatePrizeV1Request) Reset()         { *m = MultiCreatePrizeV1Request{} }
func (m *MultiCreatePrizeV1Request) String() string { return proto.CompactTextString(m) }
func (*MultiCreatePrizeV1Request) ProtoMessage()    {}
func (*MultiCreatePrizeV1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{4}
}

func (m *MultiCreatePrizeV1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiCreatePrizeV1Request.Unmarshal(m, b)
}
func (m *MultiCreatePrizeV1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiCreatePrizeV1Request.Marshal(b, m, deterministic)
}
func (m *MultiCreatePrizeV1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiCreatePrizeV1Request.Merge(m, src)
}
func (m *MultiCreatePrizeV1Request) XXX_Size() int {
	return xxx_messageInfo_MultiCreatePrizeV1Request.Size(m)
}
func (m *MultiCreatePrizeV1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiCreatePrizeV1Request.DiscardUnknown(m)
}

var xxx_messageInfo_MultiCreatePrizeV1Request proto.InternalMessageInfo

func (m *MultiCreatePrizeV1Request) GetPrizes() []*Prize {
	if m != nil {
		return m.Prizes
	}
	return nil
}

type MultiCreatePrizeV1Response struct {
	PrizeIds             []uint64 `protobuf:"varint,1,rep,packed,name=prize_ids,json=prizeIds,proto3" json:"prize_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiCreatePrizeV1Response) Reset()         { *m = MultiCreatePrizeV1Response{} }
func (m *MultiCreatePrizeV1Response) String() string { return proto.CompactTextString(m) }
func (*MultiCreatePrizeV1Response) ProtoMessage()    {}
func (*MultiCreatePrizeV1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{5}
}

func (m *MultiCreatePrizeV1Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiCreatePrizeV1Response.Unmarshal(m, b)
}
func (m *MultiCreatePrizeV1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiCreatePrizeV1Response.Marshal(b, m, deterministic)
}
func (m *MultiCreatePrizeV1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiCreatePrizeV1Response.Merge(m, src)
}
func (m *MultiCreatePrizeV1Response) XXX_Size() int {
	return xxx_messageInfo_MultiCreatePrizeV1Response.Size(m)
}
func (m *MultiCreatePrizeV1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiCreatePrizeV1Response.DiscardUnknown(m)
}

var xxx_messageInfo_MultiCreatePrizeV1Response proto.InternalMessageInfo

func (m *MultiCreatePrizeV1Response) GetPrizeIds() []uint64 {
	if m != nil {
		return m.PrizeIds
	}
	return nil
}

type UpdatePrizeV1Request struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Link                 string   `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	IssueId              uint64   `protobuf:"varint,3,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePrizeV1Request) Reset()         { *m = UpdatePrizeV1Request{} }
func (m *UpdatePrizeV1Request) String() string { return proto.CompactTextString(m) }
func (*UpdatePrizeV1Request) ProtoMessage()    {}
func (*UpdatePrizeV1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{6}
}

func (m *UpdatePrizeV1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePrizeV1Request.Unmarshal(m, b)
}
func (m *UpdatePrizeV1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePrizeV1Request.Marshal(b, m, deterministic)
}
func (m *UpdatePrizeV1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePrizeV1Request.Merge(m, src)
}
func (m *UpdatePrizeV1Request) XXX_Size() int {
	return xxx_messageInfo_UpdatePrizeV1Request.Size(m)
}
func (m *UpdatePrizeV1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePrizeV1Request.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePrizeV1Request proto.InternalMessageInfo

func (m *UpdatePrizeV1Request) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdatePrizeV1Request) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *UpdatePrizeV1Request) GetIssueId() uint64 {
	if m != nil {
		return m.IssueId
	}
	return 0
}

type UpdatePrizeV1Response struct {
	Succeed              bool     `protobuf:"varint,1,opt,name=succeed,proto3" json:"succeed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePrizeV1Response) Reset()         { *m = UpdatePrizeV1Response{} }
func (m *UpdatePrizeV1Response) String() string { return proto.CompactTextString(m) }
func (*UpdatePrizeV1Response) ProtoMessage()    {}
func (*UpdatePrizeV1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{7}
}

func (m *UpdatePrizeV1Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePrizeV1Response.Unmarshal(m, b)
}
func (m *UpdatePrizeV1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePrizeV1Response.Marshal(b, m, deterministic)
}
func (m *UpdatePrizeV1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePrizeV1Response.Merge(m, src)
}
func (m *UpdatePrizeV1Response) XXX_Size() int {
	return xxx_messageInfo_UpdatePrizeV1Response.Size(m)
}
func (m *UpdatePrizeV1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePrizeV1Response.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePrizeV1Response proto.InternalMessageInfo

func (m *UpdatePrizeV1Response) GetSucceed() bool {
	if m != nil {
		return m.Succeed
	}
	return false
}

type RemovePrizeV1Request struct {
	PrizeId              uint64   `protobuf:"varint,1,opt,name=prize_id,json=prizeId,proto3" json:"prize_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemovePrizeV1Request) Reset()         { *m = RemovePrizeV1Request{} }
func (m *RemovePrizeV1Request) String() string { return proto.CompactTextString(m) }
func (*RemovePrizeV1Request) ProtoMessage()    {}
func (*RemovePrizeV1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{8}
}

func (m *RemovePrizeV1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemovePrizeV1Request.Unmarshal(m, b)
}
func (m *RemovePrizeV1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemovePrizeV1Request.Marshal(b, m, deterministic)
}
func (m *RemovePrizeV1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemovePrizeV1Request.Merge(m, src)
}
func (m *RemovePrizeV1Request) XXX_Size() int {
	return xxx_messageInfo_RemovePrizeV1Request.Size(m)
}
func (m *RemovePrizeV1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_RemovePrizeV1Request.DiscardUnknown(m)
}

var xxx_messageInfo_RemovePrizeV1Request proto.InternalMessageInfo

func (m *RemovePrizeV1Request) GetPrizeId() uint64 {
	if m != nil {
		return m.PrizeId
	}
	return 0
}

type RemovePrizeV1Response struct {
	Found                bool     `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemovePrizeV1Response) Reset()         { *m = RemovePrizeV1Response{} }
func (m *RemovePrizeV1Response) String() string { return proto.CompactTextString(m) }
func (*RemovePrizeV1Response) ProtoMessage()    {}
func (*RemovePrizeV1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{9}
}

func (m *RemovePrizeV1Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemovePrizeV1Response.Unmarshal(m, b)
}
func (m *RemovePrizeV1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemovePrizeV1Response.Marshal(b, m, deterministic)
}
func (m *RemovePrizeV1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemovePrizeV1Response.Merge(m, src)
}
func (m *RemovePrizeV1Response) XXX_Size() int {
	return xxx_messageInfo_RemovePrizeV1Response.Size(m)
}
func (m *RemovePrizeV1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_RemovePrizeV1Response.DiscardUnknown(m)
}

var xxx_messageInfo_RemovePrizeV1Response proto.InternalMessageInfo

func (m *RemovePrizeV1Response) GetFound() bool {
	if m != nil {
		return m.Found
	}
	return false
}

type DescribePrizeV1Request struct {
	PrizeId              uint64   `protobuf:"varint,1,opt,name=prize_id,json=prizeId,proto3" json:"prize_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribePrizeV1Request) Reset()         { *m = DescribePrizeV1Request{} }
func (m *DescribePrizeV1Request) String() string { return proto.CompactTextString(m) }
func (*DescribePrizeV1Request) ProtoMessage()    {}
func (*DescribePrizeV1Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{10}
}

func (m *DescribePrizeV1Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribePrizeV1Request.Unmarshal(m, b)
}
func (m *DescribePrizeV1Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribePrizeV1Request.Marshal(b, m, deterministic)
}
func (m *DescribePrizeV1Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribePrizeV1Request.Merge(m, src)
}
func (m *DescribePrizeV1Request) XXX_Size() int {
	return xxx_messageInfo_DescribePrizeV1Request.Size(m)
}
func (m *DescribePrizeV1Request) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribePrizeV1Request.DiscardUnknown(m)
}

var xxx_messageInfo_DescribePrizeV1Request proto.InternalMessageInfo

func (m *DescribePrizeV1Request) GetPrizeId() uint64 {
	if m != nil {
		return m.PrizeId
	}
	return 0
}

type DescribePrizeV1Response struct {
	Prize                *Prize   `protobuf:"bytes,1,opt,name=prize,proto3" json:"prize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribePrizeV1Response) Reset()         { *m = DescribePrizeV1Response{} }
func (m *DescribePrizeV1Response) String() string { return proto.CompactTextString(m) }
func (*DescribePrizeV1Response) ProtoMessage()    {}
func (*DescribePrizeV1Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{11}
}

func (m *DescribePrizeV1Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribePrizeV1Response.Unmarshal(m, b)
}
func (m *DescribePrizeV1Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribePrizeV1Response.Marshal(b, m, deterministic)
}
func (m *DescribePrizeV1Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribePrizeV1Response.Merge(m, src)
}
func (m *DescribePrizeV1Response) XXX_Size() int {
	return xxx_messageInfo_DescribePrizeV1Response.Size(m)
}
func (m *DescribePrizeV1Response) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribePrizeV1Response.DiscardUnknown(m)
}

var xxx_messageInfo_DescribePrizeV1Response proto.InternalMessageInfo

func (m *DescribePrizeV1Response) GetPrize() *Prize {
	if m != nil {
		return m.Prize
	}
	return nil
}

// Prize description
type Prize struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IssueId              uint64   `protobuf:"varint,2,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	Link                 string   `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Prize) Reset()         { *m = Prize{} }
func (m *Prize) String() string { return proto.CompactTextString(m) }
func (*Prize) ProtoMessage()    {}
func (*Prize) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9f7ba6cff1c934, []int{12}
}

func (m *Prize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prize.Unmarshal(m, b)
}
func (m *Prize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prize.Marshal(b, m, deterministic)
}
func (m *Prize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prize.Merge(m, src)
}
func (m *Prize) XXX_Size() int {
	return xxx_messageInfo_Prize.Size(m)
}
func (m *Prize) XXX_DiscardUnknown() {
	xxx_messageInfo_Prize.DiscardUnknown(m)
}

var xxx_messageInfo_Prize proto.InternalMessageInfo

func (m *Prize) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Prize) GetIssueId() uint64 {
	if m != nil {
		return m.IssueId
	}
	return 0
}

func (m *Prize) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func init() {
	proto.RegisterType((*ListPrizeV1Request)(nil), "ocp.prize.api.ListPrizeV1Request")
	proto.RegisterType((*ListPrizeV1Response)(nil), "ocp.prize.api.ListPrizeV1Response")
	proto.RegisterType((*CreatePrizeV1Request)(nil), "ocp.prize.api.CreatePrizeV1Request")
	proto.RegisterType((*CreatePrizeV1Response)(nil), "ocp.prize.api.CreatePrizeV1Response")
	proto.RegisterType((*MultiCreatePrizeV1Request)(nil), "ocp.prize.api.MultiCreatePrizeV1Request")
	proto.RegisterType((*MultiCreatePrizeV1Response)(nil), "ocp.prize.api.MultiCreatePrizeV1Response")
	proto.RegisterType((*UpdatePrizeV1Request)(nil), "ocp.prize.api.UpdatePrizeV1Request")
	proto.RegisterType((*UpdatePrizeV1Response)(nil), "ocp.prize.api.UpdatePrizeV1Response")
	proto.RegisterType((*RemovePrizeV1Request)(nil), "ocp.prize.api.RemovePrizeV1Request")
	proto.RegisterType((*RemovePrizeV1Response)(nil), "ocp.prize.api.RemovePrizeV1Response")
	proto.RegisterType((*DescribePrizeV1Request)(nil), "ocp.prize.api.DescribePrizeV1Request")
	proto.RegisterType((*DescribePrizeV1Response)(nil), "ocp.prize.api.DescribePrizeV1Response")
	proto.RegisterType((*Prize)(nil), "ocp.prize.api.Prize")
}

func init() {
	proto.RegisterFile("api/ocp-prize-api/ocp-prize-api.proto", fileDescriptor_2f9f7ba6cff1c934)
}

var fileDescriptor_2f9f7ba6cff1c934 = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0xfd, 0x2d, 0x6b, 0xd7, 0xee, 0xee, 0x57, 0x90, 0x4c, 0xba, 0x75, 0x01, 0xa4, 0x61, 0x18,
	0x1a, 0x13, 0x4d, 0xb4, 0xf2, 0xc4, 0x1f, 0x09, 0xad, 0x63, 0x48, 0x93, 0x40, 0xa0, 0x48, 0xe3,
	0x81, 0x97, 0x29, 0x4d, 0xdc, 0xd4, 0x5a, 0x1b, 0x9b, 0xda, 0xa9, 0xe8, 0xd0, 0x84, 0xc4, 0x57,
	0xe0, 0x9b, 0xc1, 0x57, 0xe0, 0x53, 0xf0, 0x84, 0xea, 0xa4, 0xed, 0xe2, 0xa4, 0xe5, 0xcf, 0x5b,
	0x6f, 0x7a, 0x7c, 0xce, 0xb9, 0xf7, 0xfa, 0xc8, 0xb0, 0xeb, 0x71, 0xea, 0x30, 0x9f, 0x37, 0xf9,
	0x90, 0x5e, 0x90, 0x66, 0xae, 0xb2, 0xf9, 0x90, 0x49, 0x86, 0x6a, 0xcc, 0xe7, 0xb6, 0xfa, 0x68,
	0x7b, 0x9c, 0x5a, 0xb7, 0x42, 0xc6, 0xc2, 0x3e, 0x71, 0x26, 0x70, 0x2f, 0x8a, 0x98, 0xf4, 0x24,
	0x65, 0x91, 0x48, 0xc0, 0xd6, 0x61, 0x48, 0x65, 0x2f, 0xee, 0xd8, 0x3e, 0x1b, 0x38, 0x24, 0x1a,
	0xb1, 0x31, 0x1f, 0xb2, 0x8f, 0x63, 0x47, 0xfd, 0xe9, 0x37, 0x43, 0x12, 0x35, 0x47, 0x5e, 0x9f,
	0x06, 0x9e, 0x24, 0x4e, 0xee, 0x47, 0x42, 0x81, 0xdb, 0x80, 0x5e, 0x51, 0x21, 0xdf, 0x4e, 0x14,
	0xdf, 0x1d, 0xb8, 0xe4, 0x43, 0x4c, 0x84, 0x44, 0x26, 0x94, 0xfb, 0x74, 0x40, 0x65, 0x63, 0x65,
	0x67, 0x65, 0xaf, 0xe4, 0x26, 0x05, 0xda, 0x84, 0x35, 0xd6, 0xed, 0x0a, 0x22, 0x1b, 0x86, 0xfa,
	0x9c, 0x56, 0xf8, 0x08, 0x6e, 0x64, 0x38, 0x04, 0x67, 0x91, 0x20, 0xe8, 0x21, 0xac, 0xa9, 0x46,
	0x44, 0x63, 0x65, 0x67, 0x75, 0x6f, 0xa3, 0x65, 0xda, 0x99, 0xde, 0x6c, 0x85, 0x77, 0x53, 0x0c,
	0x3e, 0x06, 0xf3, 0x68, 0x48, 0x3c, 0x49, 0x34, 0x2b, 0x08, 0x4a, 0x7d, 0x1a, 0x9d, 0x2b, 0x27,
	0xeb, 0xae, 0xfa, 0x8d, 0xb6, 0xa1, 0x4a, 0x85, 0x88, 0xc9, 0x19, 0x0d, 0x52, 0x2b, 0x15, 0x55,
	0x9f, 0x04, 0xb8, 0x05, 0x75, 0x8d, 0x26, 0x75, 0xb3, 0x0d, 0x55, 0xa5, 0x34, 0x39, 0x93, 0x74,
	0x55, 0x51, 0xf5, 0x49, 0x80, 0x4f, 0x60, 0xfb, 0x75, 0xdc, 0x97, 0xb4, 0x50, 0xff, 0xef, 0xba,
	0x78, 0x0c, 0x56, 0x11, 0x55, 0xea, 0xe1, 0x26, 0xac, 0x4f, 0x3d, 0x24, 0x74, 0x25, 0xb7, 0x9a,
	0x9a, 0x10, 0xf8, 0x14, 0xcc, 0x53, 0x1e, 0xe4, 0x0d, 0x5c, 0x03, 0x63, 0x66, 0xd9, 0xa0, 0xc1,
	0x6c, 0x20, 0xc6, 0x82, 0x81, 0xac, 0x66, 0x07, 0x72, 0x00, 0x75, 0x8d, 0x36, 0x35, 0xd3, 0x80,
	0x8a, 0x88, 0x7d, 0x9f, 0x90, 0x84, 0xbc, 0xea, 0x4e, 0x4b, 0xfc, 0x04, 0x4c, 0x97, 0x0c, 0xd8,
	0x48, 0x77, 0x82, 0xf5, 0x11, 0xb6, 0x2b, 0x3f, 0xdb, 0xa5, 0x96, 0xb1, 0xf3, 0xdf, 0x7c, 0x96,
	0x4d, 0xa8, 0x6b, 0x67, 0x53, 0x39, 0x13, 0xca, 0x5d, 0x16, 0x47, 0x53, 0xb1, 0xa4, 0xc0, 0xcf,
	0x60, 0xf3, 0x05, 0x11, 0xfe, 0x90, 0x76, 0xfe, 0x45, 0xec, 0x18, 0xb6, 0x72, 0xa7, 0x53, 0xb9,
	0x7d, 0x28, 0x2b, 0x94, 0x3a, 0xbb, 0x68, 0x6b, 0x09, 0x04, 0xbf, 0x84, 0xb2, 0xaa, 0x73, 0xa3,
	0x5e, 0x7c, 0xcf, 0x66, 0x5b, 0x58, 0x9d, 0x6f, 0xa1, 0xf5, 0xad, 0x0c, 0x1b, 0x6f, 0x7c, 0xae,
	0xb8, 0x0e, 0x39, 0x45, 0x3d, 0xd8, 0xb8, 0x92, 0x0b, 0x74, 0x47, 0xf3, 0x90, 0xcf, 0x9d, 0x85,
	0x97, 0x41, 0x92, 0xce, 0x30, 0xfa, 0xf2, 0xfd, 0xc7, 0x57, 0xe3, 0x7f, 0x04, 0xce, 0xe8, 0xc0,
	0x49, 0xae, 0x1d, 0xfa, 0x0c, 0xd7, 0xb5, 0x41, 0xa0, 0x5d, 0x8d, 0xaa, 0x78, 0xcc, 0xd6, 0xfd,
	0xdf, 0xc1, 0x52, 0xd5, 0xdb, 0x4a, 0x75, 0x0b, 0xd5, 0xe7, 0xaa, 0xce, 0xa7, 0xe9, 0x82, 0x2e,
	0x51, 0x04, 0xb5, 0xcc, 0x95, 0x47, 0x77, 0x35, 0xde, 0xa2, 0x6c, 0x59, 0xf7, 0x96, 0x83, 0xb2,
	0x0d, 0xe3, 0xab, 0x0d, 0x5f, 0x02, 0xca, 0xe7, 0x0c, 0xed, 0x69, 0x7c, 0x0b, 0x53, 0x6d, 0x3d,
	0xf8, 0x03, 0xe4, 0x12, 0xf9, 0x08, 0x6a, 0x99, 0x50, 0xe5, 0xda, 0x2d, 0x4a, 0x72, 0xae, 0xdd,
	0xc2, 0x5c, 0x16, 0xea, 0x8d, 0xa1, 0x96, 0x49, 0x55, 0x4e, 0xaf, 0x28, 0xaf, 0x39, 0xbd, 0xc2,
	0x60, 0x4e, 0x37, 0xbb, 0x5f, 0xbc, 0xd9, 0xf6, 0xe1, 0xfb, 0xe7, 0x21, 0xed, 0xc9, 0xf4, 0x95,
	0x61, 0x17, 0x2c, 0xf2, 0xb9, 0xf6, 0x8e, 0xf1, 0xf3, 0x30, 0xfb, 0xe5, 0x29, 0xf3, 0xf9, 0x59,
	0x42, 0xe1, 0x71, 0xda, 0x59, 0x53, 0x4f, 0xcd, 0xa3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8b,
	0xde, 0x36, 0x2d, 0x03, 0x07, 0x00, 0x00,
}
