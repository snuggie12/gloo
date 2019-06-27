// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/transformation/parameters.proto

package transformation

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Parameters struct {
	// headers that will be used to extract data for processing output templates
	// Gloo will search for parameters by their name in header value strings, enclosed in single
	// curly braces
	// Example:
	//   extensions:
	//     parameters:
	//         headers:
	//           x-user-id: { userId }
	Headers map[string]string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// part of the (or the entire) path that will be used extract data for processing output templates
	// Gloo will search for parameters by their name in header value strings, enclosed in single
	// curly braces
	// Example:
	//   extensions:
	//     parameters:
	//         path: /users/{ userId }
	Path                 *types.StringValue `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Parameters) Reset()         { *m = Parameters{} }
func (m *Parameters) String() string { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()    {}
func (*Parameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_2165d9f00bb6c750, []int{0}
}
func (m *Parameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameters.Unmarshal(m, b)
}
func (m *Parameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameters.Marshal(b, m, deterministic)
}
func (m *Parameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameters.Merge(m, src)
}
func (m *Parameters) XXX_Size() int {
	return xxx_messageInfo_Parameters.Size(m)
}
func (m *Parameters) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameters.DiscardUnknown(m)
}

var xxx_messageInfo_Parameters proto.InternalMessageInfo

func (m *Parameters) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Parameters) GetPath() *types.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "transformation.plugins.gloo.solo.io.Parameters")
	proto.RegisterMapType((map[string]string)(nil), "transformation.plugins.gloo.solo.io.Parameters.HeadersEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/transformation/parameters.proto", fileDescriptor_2165d9f00bb6c750)
}

var fileDescriptor_2165d9f00bb6c750 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xd9, 0xd6, 0x3f, 0x74, 0xeb, 0x41, 0x42, 0x0f, 0x25, 0x48, 0x09, 0x7a, 0xc9, 0xc5,
	0x59, 0xad, 0x17, 0x29, 0x9e, 0x04, 0xd1, 0x8b, 0x20, 0x51, 0x7a, 0xf0, 0xb6, 0xa9, 0xdb, 0xcd,
	0xda, 0x24, 0xb3, 0xcc, 0x6e, 0x2a, 0x7d, 0x23, 0x9f, 0xc7, 0x47, 0xf0, 0x49, 0x24, 0x89, 0xad,
	0x7a, 0x91, 0xde, 0xbe, 0x19, 0xe6, 0xfb, 0x7d, 0x1f, 0x0c, 0x7f, 0xd2, 0xc6, 0x67, 0x55, 0x0a,
	0x33, 0x2c, 0x84, 0xc3, 0x1c, 0x4f, 0x0d, 0x0a, 0x9d, 0x23, 0x0a, 0x4b, 0xf8, 0xaa, 0x66, 0xde,
	0xb5, 0x93, 0xb4, 0x46, 0x2c, 0xcf, 0x85, 0xcd, 0x2b, 0x6d, 0x4a, 0x27, 0x3c, 0xc9, 0xd2, 0xcd,
	0x91, 0x0a, 0xe9, 0x0d, 0x96, 0xc2, 0x4a, 0x92, 0x85, 0xf2, 0x8a, 0x1c, 0x58, 0x42, 0x8f, 0xc1,
	0xc9, 0xdf, 0x03, 0xf8, 0xf6, 0x41, 0xcd, 0x82, 0x3a, 0x06, 0x0c, 0x86, 0x23, 0x8d, 0xa8, 0x73,
	0x25, 0x1a, 0x4b, 0x5a, 0xcd, 0xc5, 0x1b, 0x49, 0x6b, 0x37, 0x90, 0x70, 0xa0, 0x51, 0x63, 0x23,
	0x45, 0xad, 0xda, 0xed, 0xf1, 0x07, 0xe3, 0xfc, 0x61, 0x93, 0x17, 0x4c, 0xf9, 0x7e, 0xa6, 0xe4,
	0x8b, 0x22, 0x37, 0x64, 0x51, 0x37, 0xee, 0x8f, 0xaf, 0x60, 0x8b, 0x6c, 0xf8, 0x21, 0xc0, 0x5d,
	0x6b, 0xbf, 0x29, 0x3d, 0xad, 0x92, 0x35, 0x2c, 0x38, 0xe3, 0x3b, 0x56, 0xfa, 0x6c, 0xd8, 0x89,
	0x58, 0xdc, 0x1f, 0x1f, 0x41, 0xdb, 0x15, 0xd6, 0x5d, 0xe1, 0xd1, 0x93, 0x29, 0xf5, 0x54, 0xe6,
	0x95, 0x4a, 0x9a, 0xcb, 0x70, 0xc2, 0x0f, 0x7e, 0xa3, 0x82, 0x43, 0xde, 0x5d, 0xa8, 0xd5, 0x90,
	0x45, 0x2c, 0xee, 0x25, 0xb5, 0x0c, 0x06, 0x7c, 0x77, 0x59, 0x1b, 0x1a, 0x68, 0x2f, 0x69, 0x87,
	0x49, 0xe7, 0x92, 0x5d, 0xdf, 0xbf, 0x7f, 0x8e, 0xd8, 0xf3, 0xed, 0x76, 0xbf, 0xb0, 0x0b, 0xfd,
	0xff, 0x3f, 0xd2, 0xbd, 0xa6, 0xe6, 0xc5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xd4, 0x02,
	0x49, 0xdd, 0x01, 0x00, 0x00,
}

func (this *Parameters) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Parameters)
	if !ok {
		that2, ok := that.(Parameters)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if this.Headers[i] != that1.Headers[i] {
			return false
		}
	}
	if !this.Path.Equal(that1.Path) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
