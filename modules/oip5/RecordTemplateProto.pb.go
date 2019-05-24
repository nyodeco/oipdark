// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RecordTemplateProto.proto

package oip5

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RecordTemplateProto struct {
	// Human readable name to quickly identify type (non-unique)
	FriendlyName string `protobuf:"bytes,1,opt,name=friendlyName" json:"friendlyName,omitempty"`
	// Description of the purpose behind this new type
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Compiled descriptor including dependencies; Defines fields
	DescriptorSetProto []byte `protobuf:"bytes,4,opt,name=DescriptorSetProto,proto3" json:"DescriptorSetProto,omitempty"`
	// List of unique template identifiers required for use with this template
	Extends []int64 `protobuf:"fixed64,12,rep,packed,name=extends" json:"extends,omitempty"`
	// Populated by oipd with the unique identifier for this type
	Identifier int64 `protobuf:"fixed64,10,opt,name=identifier" json:"identifier,omitempty"`
}

func (m *RecordTemplateProto) Reset()                    { *m = RecordTemplateProto{} }
func (m *RecordTemplateProto) String() string            { return proto.CompactTextString(m) }
func (*RecordTemplateProto) ProtoMessage()               {}
func (*RecordTemplateProto) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *RecordTemplateProto) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *RecordTemplateProto) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RecordTemplateProto) GetDescriptorSetProto() []byte {
	if m != nil {
		return m.DescriptorSetProto
	}
	return nil
}

func (m *RecordTemplateProto) GetExtends() []int64 {
	if m != nil {
		return m.Extends
	}
	return nil
}

func (m *RecordTemplateProto) GetIdentifier() int64 {
	if m != nil {
		return m.Identifier
	}
	return 0
}

func init() {
	proto.RegisterType((*RecordTemplateProto)(nil), "oipProto.RecordTemplateProto")
}

func init() { proto.RegisterFile("RecordTemplateProto.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x0c, 0x4a, 0x4d, 0xce,
	0x2f, 0x4a, 0x09, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0x0d, 0x28, 0xca, 0x2f, 0xc9, 0xd7,
	0x2b, 0x00, 0x91, 0x42, 0x1c, 0xf9, 0x99, 0x05, 0x60, 0xbe, 0xd2, 0x51, 0x46, 0x2e, 0x61, 0x2c,
	0xea, 0x84, 0x94, 0xb8, 0x78, 0xd2, 0x8a, 0x32, 0x53, 0xf3, 0x52, 0x72, 0x2a, 0xfd, 0x12, 0x73,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x50, 0xc4, 0x84, 0x14, 0xb8, 0xb8, 0x53, 0x52,
	0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x24, 0x98, 0xc0, 0x4a, 0x90, 0x85, 0x84,
	0xf4, 0xb8, 0x84, 0x5c, 0xa0, 0xdc, 0xfc, 0xa2, 0xe0, 0xd4, 0x12, 0xb0, 0xd9, 0x12, 0x2c, 0x0a,
	0x8c, 0x1a, 0x3c, 0x41, 0x58, 0x64, 0x84, 0x24, 0xb8, 0xd8, 0x53, 0x2b, 0x4a, 0x52, 0xf3, 0x52,
	0x8a, 0x25, 0x78, 0x14, 0x98, 0x35, 0x04, 0x82, 0x60, 0x5c, 0x21, 0x39, 0x2e, 0xae, 0xcc, 0x94,
	0xd4, 0xbc, 0x92, 0xcc, 0xb4, 0xcc, 0xd4, 0x22, 0x09, 0x2e, 0x05, 0x46, 0x0d, 0x81, 0x20, 0x24,
	0x11, 0x27, 0xb6, 0x28, 0x96, 0xfc, 0xcc, 0x02, 0xd3, 0x24, 0x36, 0xb0, 0x07, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xc1, 0x11, 0x5d, 0xd5, 0xfd, 0x00, 0x00, 0x00,
}
