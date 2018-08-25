// Code generated by protoc-gen-go. DO NOT EDIT.
// source: applicationInformation.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestSoftwareInformation_RequestInformationType int32

const (
	RequestSoftwareInformation_FIRMWARE_VERSION RequestSoftwareInformation_RequestInformationType = 1
)

var RequestSoftwareInformation_RequestInformationType_name = map[int32]string{
	1: "FIRMWARE_VERSION",
}
var RequestSoftwareInformation_RequestInformationType_value = map[string]int32{
	"FIRMWARE_VERSION": 1,
}

func (x RequestSoftwareInformation_RequestInformationType) Enum() *RequestSoftwareInformation_RequestInformationType {
	p := new(RequestSoftwareInformation_RequestInformationType)
	*p = x
	return p
}
func (x RequestSoftwareInformation_RequestInformationType) String() string {
	return proto.EnumName(RequestSoftwareInformation_RequestInformationType_name, int32(x))
}
func (x *RequestSoftwareInformation_RequestInformationType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RequestSoftwareInformation_RequestInformationType_value, data, "RequestSoftwareInformation_RequestInformationType")
	if err != nil {
		return err
	}
	*x = RequestSoftwareInformation_RequestInformationType(value)
	return nil
}
func (RequestSoftwareInformation_RequestInformationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_applicationInformation_21f2f7cd76d92e97, []int{0, 0}
}

type RequestSoftwareInformation struct {
	Information          *RequestSoftwareInformation_RequestInformationType `protobuf:"varint,1,req,name=information,enum=protobuf.RequestSoftwareInformation_RequestInformationType,def=1" json:"information,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_unrecognized     []byte                                             `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *RequestSoftwareInformation) Reset()         { *m = RequestSoftwareInformation{} }
func (m *RequestSoftwareInformation) String() string { return proto.CompactTextString(m) }
func (*RequestSoftwareInformation) ProtoMessage()    {}
func (*RequestSoftwareInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationInformation_21f2f7cd76d92e97, []int{0}
}
func (m *RequestSoftwareInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestSoftwareInformation.Unmarshal(m, b)
}
func (m *RequestSoftwareInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestSoftwareInformation.Marshal(b, m, deterministic)
}
func (dst *RequestSoftwareInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestSoftwareInformation.Merge(dst, src)
}
func (m *RequestSoftwareInformation) XXX_Size() int {
	return xxx_messageInfo_RequestSoftwareInformation.Size(m)
}
func (m *RequestSoftwareInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestSoftwareInformation.DiscardUnknown(m)
}

var xxx_messageInfo_RequestSoftwareInformation proto.InternalMessageInfo

const Default_RequestSoftwareInformation_Information RequestSoftwareInformation_RequestInformationType = RequestSoftwareInformation_FIRMWARE_VERSION

func (m *RequestSoftwareInformation) GetInformation() RequestSoftwareInformation_RequestInformationType {
	if m != nil && m.Information != nil {
		return *m.Information
	}
	return Default_RequestSoftwareInformation_Information
}

type FirmwareVersion struct {
	Major                *uint32  `protobuf:"varint,1,req,name=major" json:"major,omitempty"`
	Minor                *uint32  `protobuf:"varint,2,req,name=minor" json:"minor,omitempty"`
	Patch                *uint32  `protobuf:"varint,3,req,name=patch" json:"patch,omitempty"`
	Hotfix               *uint32  `protobuf:"varint,4,req,name=hotfix" json:"hotfix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FirmwareVersion) Reset()         { *m = FirmwareVersion{} }
func (m *FirmwareVersion) String() string { return proto.CompactTextString(m) }
func (*FirmwareVersion) ProtoMessage()    {}
func (*FirmwareVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationInformation_21f2f7cd76d92e97, []int{1}
}
func (m *FirmwareVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FirmwareVersion.Unmarshal(m, b)
}
func (m *FirmwareVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FirmwareVersion.Marshal(b, m, deterministic)
}
func (dst *FirmwareVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FirmwareVersion.Merge(dst, src)
}
func (m *FirmwareVersion) XXX_Size() int {
	return xxx_messageInfo_FirmwareVersion.Size(m)
}
func (m *FirmwareVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_FirmwareVersion.DiscardUnknown(m)
}

var xxx_messageInfo_FirmwareVersion proto.InternalMessageInfo

func (m *FirmwareVersion) GetMajor() uint32 {
	if m != nil && m.Major != nil {
		return *m.Major
	}
	return 0
}

func (m *FirmwareVersion) GetMinor() uint32 {
	if m != nil && m.Minor != nil {
		return *m.Minor
	}
	return 0
}

func (m *FirmwareVersion) GetPatch() uint32 {
	if m != nil && m.Patch != nil {
		return *m.Patch
	}
	return 0
}

func (m *FirmwareVersion) GetHotfix() uint32 {
	if m != nil && m.Hotfix != nil {
		return *m.Hotfix
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestSoftwareInformation)(nil), "protobuf.requestSoftwareInformation")
	proto.RegisterType((*FirmwareVersion)(nil), "protobuf.firmwareVersion")
	proto.RegisterEnum("protobuf.RequestSoftwareInformation_RequestInformationType", RequestSoftwareInformation_RequestInformationType_name, RequestSoftwareInformation_RequestInformationType_value)
}

func init() {
	proto.RegisterFile("applicationInformation.proto", fileDescriptor_applicationInformation_21f2f7cd76d92e97)
}

var fileDescriptor_applicationInformation_21f2f7cd76d92e97 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x2c, 0x28, 0xc8,
	0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xf3, 0xcc, 0x4b, 0xcb, 0x2f, 0xca, 0x05, 0x33, 0xf5,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0xc0, 0x54, 0x52, 0x69, 0x9a, 0xd2, 0x5e, 0x46, 0x2e,
	0xa9, 0xa2, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x92, 0xe0, 0xfc, 0xb4, 0x92, 0xf2, 0xc4, 0xa2, 0x54,
	0x24, 0xe5, 0x42, 0xf9, 0x5c, 0xdc, 0x99, 0x08, 0xae, 0x04, 0xa3, 0x02, 0x93, 0x06, 0x9f, 0x91,
	0xb5, 0x1e, 0x4c, 0xbb, 0x1e, 0x6e, 0xad, 0x7a, 0x41, 0x10, 0x29, 0x24, 0xa1, 0x90, 0xca, 0x82,
	0x54, 0x2b, 0x01, 0x37, 0xcf, 0x20, 0xdf, 0x70, 0xc7, 0x20, 0xd7, 0xf8, 0x30, 0xd7, 0xa0, 0x60,
	0x4f, 0x7f, 0xbf, 0x20, 0x64, 0x1b, 0x94, 0xf4, 0xb8, 0xc4, 0xb0, 0x6b, 0x14, 0x12, 0xe1, 0xc2,
	0xd0, 0x2a, 0xc0, 0xa8, 0x94, 0xcd, 0xc5, 0x9f, 0x96, 0x59, 0x94, 0x0b, 0xb2, 0x3c, 0x2c, 0xb5,
	0xa8, 0x18, 0xe4, 0x66, 0x11, 0x2e, 0xd6, 0xdc, 0xc4, 0xac, 0xfc, 0x22, 0xb0, 0x6b, 0x79, 0x83,
	0x20, 0x1c, 0xb0, 0x68, 0x66, 0x5e, 0x7e, 0x91, 0x04, 0x13, 0x54, 0x14, 0xc4, 0x01, 0x89, 0x16,
	0x24, 0x96, 0x24, 0x67, 0x48, 0x30, 0x43, 0x44, 0xc1, 0x1c, 0x21, 0x31, 0x2e, 0xb6, 0x8c, 0xfc,
	0x92, 0xb4, 0xcc, 0x0a, 0x09, 0x16, 0xb0, 0x30, 0x94, 0x07, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb9,
	0x36, 0x5c, 0xc9, 0x55, 0x01, 0x00, 0x00,
}