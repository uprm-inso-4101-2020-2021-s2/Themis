// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Themis/poll.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Poll struct {
	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Group       string   `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Title       string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Options     []string `protobuf:"bytes,5,rep,name=options,proto3" json:"options,omitempty"`
	Deadline    int64    `protobuf:"varint,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (m *Poll) Reset()         { *m = Poll{} }
func (m *Poll) String() string { return proto.CompactTextString(m) }
func (*Poll) ProtoMessage()    {}
func (*Poll) Descriptor() ([]byte, []int) {
	return fileDescriptor_9809cdc3832974ea, []int{0}
}
func (m *Poll) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Poll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Poll.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Poll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Poll.Merge(m, src)
}
func (m *Poll) XXX_Size() int {
	return m.Size()
}
func (m *Poll) XXX_DiscardUnknown() {
	xxx_messageInfo_Poll.DiscardUnknown(m)
}

var xxx_messageInfo_Poll proto.InternalMessageInfo

func (m *Poll) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Poll) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Poll) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Poll) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Poll) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Poll) GetDeadline() int64 {
	if m != nil {
		return m.Deadline
	}
	return 0
}

type MsgCreatePoll struct {
	Creator     string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Group       string   `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Title       string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Options     []string `protobuf:"bytes,5,rep,name=options,proto3" json:"options,omitempty"`
	Deadline    int64    `protobuf:"varint,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (m *MsgCreatePoll) Reset()         { *m = MsgCreatePoll{} }
func (m *MsgCreatePoll) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePoll) ProtoMessage()    {}
func (*MsgCreatePoll) Descriptor() ([]byte, []int) {
	return fileDescriptor_9809cdc3832974ea, []int{1}
}
func (m *MsgCreatePoll) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePoll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePoll.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePoll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePoll.Merge(m, src)
}
func (m *MsgCreatePoll) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePoll) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePoll.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePoll proto.InternalMessageInfo

func (m *MsgCreatePoll) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreatePoll) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *MsgCreatePoll) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgCreatePoll) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgCreatePoll) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *MsgCreatePoll) GetDeadline() int64 {
	if m != nil {
		return m.Deadline
	}
	return 0
}

type MsgSetPollDesc struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id          string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *MsgSetPollDesc) Reset()         { *m = MsgSetPollDesc{} }
func (m *MsgSetPollDesc) String() string { return proto.CompactTextString(m) }
func (*MsgSetPollDesc) ProtoMessage()    {}
func (*MsgSetPollDesc) Descriptor() ([]byte, []int) {
	return fileDescriptor_9809cdc3832974ea, []int{2}
}
func (m *MsgSetPollDesc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetPollDesc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetPollDesc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetPollDesc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetPollDesc.Merge(m, src)
}
func (m *MsgSetPollDesc) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetPollDesc) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetPollDesc.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetPollDesc proto.InternalMessageInfo

func (m *MsgSetPollDesc) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSetPollDesc) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgSetPollDesc) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type MsgExtendPollDeadline struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Deadline int64  `protobuf:"varint,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (m *MsgExtendPollDeadline) Reset()         { *m = MsgExtendPollDeadline{} }
func (m *MsgExtendPollDeadline) String() string { return proto.CompactTextString(m) }
func (*MsgExtendPollDeadline) ProtoMessage()    {}
func (*MsgExtendPollDeadline) Descriptor() ([]byte, []int) {
	return fileDescriptor_9809cdc3832974ea, []int{3}
}
func (m *MsgExtendPollDeadline) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExtendPollDeadline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExtendPollDeadline.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExtendPollDeadline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExtendPollDeadline.Merge(m, src)
}
func (m *MsgExtendPollDeadline) XXX_Size() int {
	return m.Size()
}
func (m *MsgExtendPollDeadline) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExtendPollDeadline.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExtendPollDeadline proto.InternalMessageInfo

func (m *MsgExtendPollDeadline) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgExtendPollDeadline) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgExtendPollDeadline) GetDeadline() int64 {
	if m != nil {
		return m.Deadline
	}
	return 0
}

func init() {
	proto.RegisterType((*Poll)(nil), "uprminso410120202021s2.Themis.Themis.Poll")
	proto.RegisterType((*MsgCreatePoll)(nil), "uprminso410120202021s2.Themis.Themis.MsgCreatePoll")
	proto.RegisterType((*MsgSetPollDesc)(nil), "uprminso410120202021s2.Themis.Themis.MsgSetPollDesc")
	proto.RegisterType((*MsgExtendPollDeadline)(nil), "uprminso410120202021s2.Themis.Themis.MsgExtendPollDeadline")
}

func init() { proto.RegisterFile("Themis/poll.proto", fileDescriptor_9809cdc3832974ea) }

var fileDescriptor_9809cdc3832974ea = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x97, 0x76, 0x7f, 0x5c, 0xc4, 0x81, 0x65, 0x42, 0xd8, 0xa1, 0x94, 0xe1, 0x61, 0x97,
	0xae, 0xeb, 0xf4, 0xe8, 0xc9, 0x3f, 0xc7, 0x81, 0x4c, 0xf1, 0x20, 0x7a, 0xd8, 0xda, 0x90, 0x05,
	0xba, 0xbe, 0x25, 0xc9, 0x60, 0x7e, 0x0b, 0xcf, 0x7e, 0x07, 0xbf, 0x87, 0xc7, 0x1d, 0x3d, 0xca,
	0xf6, 0x45, 0x24, 0xc9, 0x36, 0xa6, 0x82, 0xe0, 0x4d, 0x0a, 0xcd, 0xfb, 0x3c, 0xc9, 0xcb, 0xfb,
	0xe3, 0xe5, 0xc1, 0x87, 0xb7, 0x13, 0x3a, 0xe5, 0x32, 0x2a, 0x20, 0xcb, 0xba, 0x85, 0x00, 0x05,
	0xde, 0xf1, 0xac, 0x10, 0x53, 0x9e, 0x4b, 0x38, 0x8d, 0x7b, 0x71, 0xbf, 0x67, 0xbe, 0x58, 0xf6,
	0xbb, 0xf6, 0xe5, 0xfa, 0x68, 0x35, 0x19, 0x30, 0x30, 0x0d, 0x91, 0xae, 0x6c, 0x6f, 0xfb, 0x05,
	0xe1, 0xf2, 0x35, 0x64, 0x99, 0xd7, 0xc0, 0x0e, 0x4f, 0x09, 0x0a, 0x50, 0xa7, 0x3e, 0x74, 0x78,
	0xea, 0x35, 0x71, 0x85, 0x09, 0x98, 0x15, 0xc4, 0x31, 0x96, 0x15, 0xda, 0x55, 0x5c, 0x65, 0x94,
	0xb8, 0xd6, 0x35, 0xc2, 0x0b, 0xf0, 0x7e, 0x4a, 0x65, 0x22, 0x78, 0xa1, 0x38, 0xe4, 0xa4, 0x6c,
	0xee, 0x76, 0x2d, 0x8f, 0xe0, 0x1a, 0x98, 0x4a, 0x92, 0x4a, 0xe0, 0x76, 0xea, 0xc3, 0x8d, 0xf4,
	0x5a, 0x78, 0x2f, 0xa5, 0xa3, 0x34, 0xe3, 0x39, 0x25, 0xd5, 0x00, 0x75, 0xdc, 0xe1, 0x56, 0xb7,
	0x5f, 0x11, 0x3e, 0x18, 0x48, 0x76, 0x21, 0xe8, 0x48, 0x51, 0x43, 0x49, 0x70, 0x2d, 0xd1, 0x0a,
	0xc4, 0x1a, 0x75, 0x23, 0xff, 0x05, 0xef, 0x03, 0x6e, 0x0c, 0x24, 0xbb, 0xa1, 0x4a, 0xb3, 0x5e,
	0x52, 0x99, 0xfc, 0xc2, 0x6b, 0xf7, 0xed, 0x6c, 0xf7, 0xfd, 0x8d, 0xc9, 0xfd, 0xc1, 0xd4, 0x7e,
	0xc4, 0x47, 0x03, 0xc9, 0xae, 0xe6, 0x8a, 0xe6, 0xa9, 0x1d, 0x60, 0xc7, 0xfe, 0x61, 0xc8, 0x2e,
	0xbc, 0xfb, 0x15, 0xfe, 0xfc, 0xee, 0x6d, 0xe9, 0xa3, 0xc5, 0xd2, 0x47, 0x1f, 0x4b, 0x1f, 0x3d,
	0xaf, 0xfc, 0xd2, 0x62, 0xe5, 0x97, 0xde, 0x57, 0x7e, 0xe9, 0xfe, 0x8c, 0x71, 0x35, 0x99, 0x8d,
	0xbb, 0x09, 0x4c, 0x23, 0x1d, 0xb5, 0x50, 0x67, 0x2d, 0xd4, 0x61, 0x0b, 0x75, 0xd6, 0xf4, 0x2f,
	0x0e, 0x65, 0x3f, 0x5a, 0x27, 0x73, 0xbe, 0x29, 0xd4, 0x53, 0x41, 0xe5, 0xb8, 0x6a, 0x82, 0x76,
	0xf2, 0x19, 0x00, 0x00, 0xff, 0xff, 0x60, 0x8d, 0x3e, 0x7a, 0xb9, 0x02, 0x00, 0x00,
}

func (m *Poll) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Poll) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Poll) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Deadline != 0 {
		i = encodeVarintPoll(dAtA, i, uint64(m.Deadline))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Options) > 0 {
		for iNdEx := len(m.Options) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Options[iNdEx])
			copy(dAtA[i:], m.Options[iNdEx])
			i = encodeVarintPoll(dAtA, i, uint64(len(m.Options[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Group) > 0 {
		i -= len(m.Group)
		copy(dAtA[i:], m.Group)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Group)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePoll) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePoll) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePoll) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Deadline != 0 {
		i = encodeVarintPoll(dAtA, i, uint64(m.Deadline))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Options) > 0 {
		for iNdEx := len(m.Options) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Options[iNdEx])
			copy(dAtA[i:], m.Options[iNdEx])
			i = encodeVarintPoll(dAtA, i, uint64(len(m.Options[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Group) > 0 {
		i -= len(m.Group)
		copy(dAtA[i:], m.Group)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Group)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetPollDesc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetPollDesc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetPollDesc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgExtendPollDeadline) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExtendPollDeadline) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExtendPollDeadline) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Deadline != 0 {
		i = encodeVarintPoll(dAtA, i, uint64(m.Deadline))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPoll(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoll(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoll(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Poll) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.Group)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	if len(m.Options) > 0 {
		for _, s := range m.Options {
			l = len(s)
			n += 1 + l + sovPoll(uint64(l))
		}
	}
	if m.Deadline != 0 {
		n += 1 + sovPoll(uint64(m.Deadline))
	}
	return n
}

func (m *MsgCreatePoll) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.Group)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	if len(m.Options) > 0 {
		for _, s := range m.Options {
			l = len(s)
			n += 1 + l + sovPoll(uint64(l))
		}
	}
	if m.Deadline != 0 {
		n += 1 + sovPoll(uint64(m.Deadline))
	}
	return n
}

func (m *MsgSetPollDesc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	return n
}

func (m *MsgExtendPollDeadline) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPoll(uint64(l))
	}
	if m.Deadline != 0 {
		n += 1 + sovPoll(uint64(m.Deadline))
	}
	return n
}

func sovPoll(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoll(x uint64) (n int) {
	return sovPoll(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Poll) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoll
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
			return fmt.Errorf("proto: Poll: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Poll: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = append(m.Options, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			m.Deadline = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Deadline |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPoll(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoll
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
func (m *MsgCreatePoll) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoll
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
			return fmt.Errorf("proto: MsgCreatePoll: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePoll: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = append(m.Options, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			m.Deadline = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Deadline |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPoll(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoll
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
func (m *MsgSetPollDesc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoll
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
			return fmt.Errorf("proto: MsgSetPollDesc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetPollDesc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPoll(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoll
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
func (m *MsgExtendPollDeadline) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoll
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
			return fmt.Errorf("proto: MsgExtendPollDeadline: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExtendPollDeadline: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPoll
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPoll
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			m.Deadline = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoll
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Deadline |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPoll(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoll
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
func skipPoll(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoll
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
					return 0, ErrIntOverflowPoll
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
					return 0, ErrIntOverflowPoll
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
				return 0, ErrInvalidLengthPoll
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoll
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoll
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoll        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoll          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoll = fmt.Errorf("proto: unexpected end of group")
)
