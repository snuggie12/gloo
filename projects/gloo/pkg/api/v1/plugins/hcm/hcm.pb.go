// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/hcm/hcm.proto

package hcm

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Contains various settings for Envoy's http connection manager.
// See here for more information: https://www.envoyproxy.io/docs/envoy/v1.9.0/configuration/http_conn_man/http_conn_man
type HttpConnectionManagerSettings struct {
	SkipXffAppend       bool               `protobuf:"varint,1,opt,name=skip_xff_append,json=skipXffAppend,proto3" json:"skip_xff_append,omitempty"`
	Via                 string             `protobuf:"bytes,2,opt,name=via,proto3" json:"via,omitempty"`
	XffNumTrustedHops   uint32             `protobuf:"varint,3,opt,name=xff_num_trusted_hops,json=xffNumTrustedHops,proto3" json:"xff_num_trusted_hops,omitempty"`
	UseRemoteAddress    *types.BoolValue   `protobuf:"bytes,4,opt,name=use_remote_address,json=useRemoteAddress,proto3" json:"use_remote_address,omitempty"`
	GenerateRequestId   *types.BoolValue   `protobuf:"bytes,5,opt,name=generate_request_id,json=generateRequestId,proto3" json:"generate_request_id,omitempty"`
	Proxy_100Continue   bool               `protobuf:"varint,6,opt,name=proxy_100_continue,json=proxy100Continue,proto3" json:"proxy_100_continue,omitempty"`
	StreamIdleTimeout   *time.Duration     `protobuf:"bytes,7,opt,name=stream_idle_timeout,json=streamIdleTimeout,proto3,stdduration" json:"stream_idle_timeout,omitempty"`
	IdleTimeout         *time.Duration     `protobuf:"bytes,8,opt,name=idle_timeout,json=idleTimeout,proto3,stdduration" json:"idle_timeout,omitempty"`
	MaxRequestHeadersKb *types.UInt32Value `protobuf:"bytes,9,opt,name=max_request_headers_kb,json=maxRequestHeadersKb,proto3" json:"max_request_headers_kb,omitempty"`
	RequestTimeout      *time.Duration     `protobuf:"bytes,10,opt,name=request_timeout,json=requestTimeout,proto3,stdduration" json:"request_timeout,omitempty"`
	DrainTimeout        *time.Duration     `protobuf:"bytes,12,opt,name=drain_timeout,json=drainTimeout,proto3,stdduration" json:"drain_timeout,omitempty"`
	DelayedCloseTimeout *time.Duration     `protobuf:"bytes,13,opt,name=delayed_close_timeout,json=delayedCloseTimeout,proto3,stdduration" json:"delayed_close_timeout,omitempty"`
	ServerName          string             `protobuf:"bytes,14,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	// For explanation of these settings see: https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/protocol.proto#envoy-api-msg-core-http1protocoloptions
	AcceptHttp_10         bool     `protobuf:"varint,15,opt,name=accept_http_10,json=acceptHttp10,proto3" json:"accept_http_10,omitempty"`
	DefaultHostForHttp_10 string   `protobuf:"bytes,16,opt,name=default_host_for_http_10,json=defaultHostForHttp10,proto3" json:"default_host_for_http_10,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *HttpConnectionManagerSettings) Reset()         { *m = HttpConnectionManagerSettings{} }
func (m *HttpConnectionManagerSettings) String() string { return proto.CompactTextString(m) }
func (*HttpConnectionManagerSettings) ProtoMessage()    {}
func (*HttpConnectionManagerSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c9393403d6dbb8c, []int{0}
}
func (m *HttpConnectionManagerSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpConnectionManagerSettings.Unmarshal(m, b)
}
func (m *HttpConnectionManagerSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpConnectionManagerSettings.Marshal(b, m, deterministic)
}
func (m *HttpConnectionManagerSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpConnectionManagerSettings.Merge(m, src)
}
func (m *HttpConnectionManagerSettings) XXX_Size() int {
	return xxx_messageInfo_HttpConnectionManagerSettings.Size(m)
}
func (m *HttpConnectionManagerSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpConnectionManagerSettings.DiscardUnknown(m)
}

var xxx_messageInfo_HttpConnectionManagerSettings proto.InternalMessageInfo

func (m *HttpConnectionManagerSettings) GetSkipXffAppend() bool {
	if m != nil {
		return m.SkipXffAppend
	}
	return false
}

func (m *HttpConnectionManagerSettings) GetVia() string {
	if m != nil {
		return m.Via
	}
	return ""
}

func (m *HttpConnectionManagerSettings) GetXffNumTrustedHops() uint32 {
	if m != nil {
		return m.XffNumTrustedHops
	}
	return 0
}

func (m *HttpConnectionManagerSettings) GetUseRemoteAddress() *types.BoolValue {
	if m != nil {
		return m.UseRemoteAddress
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetGenerateRequestId() *types.BoolValue {
	if m != nil {
		return m.GenerateRequestId
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetProxy_100Continue() bool {
	if m != nil {
		return m.Proxy_100Continue
	}
	return false
}

func (m *HttpConnectionManagerSettings) GetStreamIdleTimeout() *time.Duration {
	if m != nil {
		return m.StreamIdleTimeout
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetIdleTimeout() *time.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetMaxRequestHeadersKb() *types.UInt32Value {
	if m != nil {
		return m.MaxRequestHeadersKb
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetRequestTimeout() *time.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetDrainTimeout() *time.Duration {
	if m != nil {
		return m.DrainTimeout
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetDelayedCloseTimeout() *time.Duration {
	if m != nil {
		return m.DelayedCloseTimeout
	}
	return nil
}

func (m *HttpConnectionManagerSettings) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *HttpConnectionManagerSettings) GetAcceptHttp_10() bool {
	if m != nil {
		return m.AcceptHttp_10
	}
	return false
}

func (m *HttpConnectionManagerSettings) GetDefaultHostForHttp_10() string {
	if m != nil {
		return m.DefaultHostForHttp_10
	}
	return ""
}

func init() {
	proto.RegisterType((*HttpConnectionManagerSettings)(nil), "hcm.plugins.gloo.solo.io.HttpConnectionManagerSettings")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/hcm/hcm.proto", fileDescriptor_1c9393403d6dbb8c)
}

var fileDescriptor_1c9393403d6dbb8c = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdd, 0x6e, 0x13, 0x39,
	0x14, 0x80, 0x95, 0x6d, 0xb7, 0x3f, 0x6e, 0xd2, 0xa6, 0x93, 0xee, 0x6a, 0xb6, 0x5a, 0xda, 0x08,
	0x21, 0x94, 0x0b, 0x98, 0x49, 0x5a, 0x89, 0x2b, 0x6e, 0x9a, 0x56, 0x28, 0x05, 0x51, 0x44, 0x5a,
	0x10, 0xe2, 0xc6, 0x72, 0x66, 0xce, 0x4c, 0x4c, 0x67, 0x7c, 0x8c, 0x7f, 0x4a, 0xfa, 0x26, 0x88,
	0x27, 0xe0, 0xad, 0x90, 0x78, 0x12, 0x34, 0x76, 0xa6, 0x80, 0x0a, 0x22, 0x17, 0x91, 0xec, 0x73,
	0xce, 0xf7, 0xd9, 0x3e, 0xf1, 0x98, 0x0c, 0x73, 0x6e, 0xa6, 0x76, 0x12, 0x25, 0x58, 0xc6, 0x1a,
	0x0b, 0x7c, 0xc8, 0x31, 0xce, 0x0b, 0xc4, 0x58, 0x2a, 0x7c, 0x07, 0x89, 0xd1, 0x7e, 0xc6, 0x24,
	0x8f, 0xaf, 0x06, 0xb1, 0x2c, 0x6c, 0xce, 0x85, 0x8e, 0xa7, 0x49, 0x59, 0xfd, 0x22, 0xa9, 0xd0,
	0x60, 0x10, 0xba, 0xa1, 0x4f, 0x45, 0x55, 0x79, 0x54, 0x99, 0x22, 0x8e, 0xbb, 0x3b, 0x39, 0xe6,
	0xe8, 0x8a, 0xe2, 0x6a, 0xe4, 0xeb, 0x77, 0xf7, 0x72, 0xc4, 0xbc, 0x80, 0xd8, 0xcd, 0x26, 0x36,
	0x8b, 0x3f, 0x28, 0x26, 0x25, 0x28, 0xfd, 0xbb, 0x7c, 0x6a, 0x15, 0x33, 0x1c, 0x85, 0xcf, 0xdf,
	0xfd, 0xb4, 0x4a, 0xee, 0x8c, 0x8c, 0x91, 0xc7, 0x28, 0x04, 0x24, 0x55, 0xe2, 0x39, 0x13, 0x2c,
	0x07, 0x75, 0x0e, 0xc6, 0x70, 0x91, 0xeb, 0xe0, 0x3e, 0xd9, 0xd2, 0x97, 0x5c, 0xd2, 0x59, 0x96,
	0xd1, 0x4a, 0x2d, 0xd2, 0xb0, 0xd1, 0x6d, 0xf4, 0xd6, 0xc6, 0xad, 0x2a, 0xfc, 0x26, 0xcb, 0x8e,
	0x5c, 0x30, 0x68, 0x93, 0xa5, 0x2b, 0xce, 0xc2, 0xbf, 0xba, 0x8d, 0xde, 0xfa, 0xb8, 0x1a, 0x06,
	0x31, 0xd9, 0xa9, 0x20, 0x61, 0x4b, 0x6a, 0x94, 0xd5, 0x06, 0x52, 0x3a, 0x45, 0xa9, 0xc3, 0xa5,
	0x6e, 0xa3, 0xd7, 0x1a, 0x6f, 0xcf, 0xb2, 0xec, 0xcc, 0x96, 0x17, 0x3e, 0x33, 0x42, 0xa9, 0x83,
	0x11, 0x09, 0xac, 0x06, 0xaa, 0xa0, 0x44, 0x03, 0x94, 0xa5, 0xa9, 0x02, 0xad, 0xc3, 0xe5, 0x6e,
	0xa3, 0xb7, 0x71, 0xb0, 0x1b, 0xf9, 0x93, 0x44, 0xf5, 0x49, 0xa2, 0x21, 0x62, 0xf1, 0x9a, 0x15,
	0x16, 0xc6, 0x6d, 0xab, 0x61, 0xec, 0xa0, 0x23, 0xcf, 0x04, 0x4f, 0x49, 0x27, 0x07, 0x01, 0x8a,
	0x99, 0x4a, 0xf7, 0xde, 0x82, 0x36, 0x94, 0xa7, 0xe1, 0xdf, 0x7f, 0x54, 0x6d, 0xd7, 0xd8, 0xd8,
	0x53, 0xa7, 0x69, 0xf0, 0x80, 0x04, 0x52, 0xe1, 0xec, 0x9a, 0x0e, 0xfa, 0x7d, 0x9a, 0xa0, 0x30,
	0x5c, 0x58, 0x08, 0x57, 0x5c, 0x0f, 0xda, 0x2e, 0x33, 0xe8, 0xf7, 0x8f, 0xe7, 0xf1, 0xe0, 0x05,
	0xe9, 0x68, 0xa3, 0x80, 0x95, 0x94, 0xa7, 0x05, 0x50, 0xc3, 0x4b, 0x40, 0x6b, 0xc2, 0x55, 0xb7,
	0xf2, 0x7f, 0xb7, 0x56, 0x3e, 0x99, 0xff, 0x1d, 0xc3, 0xe5, 0x8f, 0x5f, 0xf6, 0x1b, 0xe3, 0x6d,
	0xcf, 0x9e, 0xa6, 0x05, 0x5c, 0x78, 0x32, 0x18, 0x92, 0xe6, 0x4f, 0xa6, 0xb5, 0xc5, 0x4c, 0x1b,
	0xfc, 0x07, 0xc7, 0x4b, 0xf2, 0x6f, 0xc9, 0x66, 0x37, 0x9d, 0x98, 0x02, 0x4b, 0x41, 0x69, 0x7a,
	0x39, 0x09, 0xd7, 0x9d, 0xed, 0xff, 0x5b, 0xb6, 0x57, 0xa7, 0xc2, 0x1c, 0x1e, 0xf8, 0x9e, 0x74,
	0x4a, 0x36, 0x9b, 0xb7, 0x63, 0xe4, 0xc9, 0x67, 0x93, 0x60, 0x44, 0xb6, 0x6a, 0x5d, 0xbd, 0x33,
	0xb2, 0xd8, 0xce, 0x36, 0xe7, 0x5c, 0xbd, 0xb9, 0x13, 0xd2, 0x4a, 0x15, 0xe3, 0xe2, 0xc6, 0xd3,
	0x5c, 0xcc, 0xd3, 0x74, 0x54, 0x6d, 0x39, 0x27, 0xff, 0xa4, 0x50, 0xb0, 0x6b, 0x48, 0x69, 0x52,
	0xa0, 0xfe, 0xde, 0xaf, 0xd6, 0x62, 0xb6, 0xce, 0x9c, 0x3e, 0xae, 0xe0, 0x5a, 0xba, 0x4f, 0x36,
	0x34, 0xa8, 0x2b, 0x50, 0x54, 0xb0, 0x12, 0xc2, 0x4d, 0x77, 0xb7, 0x89, 0x0f, 0x9d, 0xb1, 0x12,
	0x82, 0x7b, 0x64, 0x93, 0x25, 0x09, 0x48, 0x43, 0xa7, 0xc6, 0x48, 0x3a, 0xe8, 0x87, 0x5b, 0xee,
	0x5e, 0x34, 0x7d, 0xb4, 0xfa, 0xb2, 0x06, 0xfd, 0xe0, 0x11, 0x09, 0x53, 0xc8, 0x98, 0x2d, 0x0c,
	0x9d, 0xa2, 0x36, 0x34, 0x43, 0x75, 0x53, 0xdf, 0x76, 0xce, 0x9d, 0x79, 0x7e, 0x84, 0xda, 0x3c,
	0x41, 0xe5, 0xb9, 0xe1, 0xf0, 0xf3, 0xd7, 0xbd, 0xc6, 0xdb, 0xc7, 0x8b, 0x3d, 0x2b, 0xf2, 0x32,
	0xff, 0xc5, 0xd3, 0x32, 0x59, 0x71, 0x07, 0x3e, 0xfc, 0x16, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x89,
	0x9b, 0x0f, 0x9d, 0x04, 0x00, 0x00,
}

func (this *HttpConnectionManagerSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HttpConnectionManagerSettings)
	if !ok {
		that2, ok := that.(HttpConnectionManagerSettings)
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
	if this.SkipXffAppend != that1.SkipXffAppend {
		return false
	}
	if this.Via != that1.Via {
		return false
	}
	if this.XffNumTrustedHops != that1.XffNumTrustedHops {
		return false
	}
	if !this.UseRemoteAddress.Equal(that1.UseRemoteAddress) {
		return false
	}
	if !this.GenerateRequestId.Equal(that1.GenerateRequestId) {
		return false
	}
	if this.Proxy_100Continue != that1.Proxy_100Continue {
		return false
	}
	if this.StreamIdleTimeout != nil && that1.StreamIdleTimeout != nil {
		if *this.StreamIdleTimeout != *that1.StreamIdleTimeout {
			return false
		}
	} else if this.StreamIdleTimeout != nil {
		return false
	} else if that1.StreamIdleTimeout != nil {
		return false
	}
	if this.IdleTimeout != nil && that1.IdleTimeout != nil {
		if *this.IdleTimeout != *that1.IdleTimeout {
			return false
		}
	} else if this.IdleTimeout != nil {
		return false
	} else if that1.IdleTimeout != nil {
		return false
	}
	if !this.MaxRequestHeadersKb.Equal(that1.MaxRequestHeadersKb) {
		return false
	}
	if this.RequestTimeout != nil && that1.RequestTimeout != nil {
		if *this.RequestTimeout != *that1.RequestTimeout {
			return false
		}
	} else if this.RequestTimeout != nil {
		return false
	} else if that1.RequestTimeout != nil {
		return false
	}
	if this.DrainTimeout != nil && that1.DrainTimeout != nil {
		if *this.DrainTimeout != *that1.DrainTimeout {
			return false
		}
	} else if this.DrainTimeout != nil {
		return false
	} else if that1.DrainTimeout != nil {
		return false
	}
	if this.DelayedCloseTimeout != nil && that1.DelayedCloseTimeout != nil {
		if *this.DelayedCloseTimeout != *that1.DelayedCloseTimeout {
			return false
		}
	} else if this.DelayedCloseTimeout != nil {
		return false
	} else if that1.DelayedCloseTimeout != nil {
		return false
	}
	if this.ServerName != that1.ServerName {
		return false
	}
	if this.AcceptHttp_10 != that1.AcceptHttp_10 {
		return false
	}
	if this.DefaultHostForHttp_10 != that1.DefaultHostForHttp_10 {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
