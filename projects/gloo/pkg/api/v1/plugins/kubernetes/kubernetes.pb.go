// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/kubernetes/kubernetes.proto

package kubernetes

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	plugins "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins"
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

// Upstream Spec for Kubernetes Upstreams
// Kubernetes Upstreams represent a set of one or more addressable pods for a Kubernetes Service
// the Gloo Kubernetes Upstream maps to a single service port. Because Kubernetes Services support multiple ports,
// Gloo requires that a different upstream be created for each port
// Kubernetes Upstreams are typically generated automatically by Gloo from the Kubernetes API
type UpstreamSpec struct {
	// The name of the Kubernetes Service
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The namespace where the Service lives
	ServiceNamespace string `protobuf:"bytes,2,opt,name=service_namespace,json=serviceNamespace,proto3" json:"service_namespace,omitempty"`
	// The access port port of the kubernetes service is listening. This port is used by Gloo to look up the corresponding
	// port on the pod for routing.
	ServicePort uint32 `protobuf:"varint,3,opt,name=service_port,json=servicePort,proto3" json:"service_port,omitempty"`
	// Allows finer-grained filtering of pods for the Upstream. Gloo will select pods based on their labels if
	// any are provided here.
	// (see [Kubernetes labels and selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
	Selector map[string]string `protobuf:"bytes,4,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec *plugins.ServiceSpec `protobuf:"bytes,5,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	// Subset configuration. For discovery sources that has labels (like kubernetes). this
	// configuration allows you to partition the upstream to a set of subsets.
	// for each unique set of keys and values, a subset will be created.
	SubsetSpec           *plugins.SubsetSpec `protobuf:"bytes,6,opt,name=subset_spec,json=subsetSpec,proto3" json:"subset_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_419a7b10c074c4e5, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *UpstreamSpec) GetServiceNamespace() string {
	if m != nil {
		return m.ServiceNamespace
	}
	return ""
}

func (m *UpstreamSpec) GetServicePort() uint32 {
	if m != nil {
		return m.ServicePort
	}
	return 0
}

func (m *UpstreamSpec) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *UpstreamSpec) GetServiceSpec() *plugins.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

func (m *UpstreamSpec) GetSubsetSpec() *plugins.SubsetSpec {
	if m != nil {
		return m.SubsetSpec
	}
	return nil
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "kubernetes.plugins.gloo.solo.io.UpstreamSpec")
	proto.RegisterMapType((map[string]string)(nil), "kubernetes.plugins.gloo.solo.io.UpstreamSpec.SelectorEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/kubernetes/kubernetes.proto", fileDescriptor_419a7b10c074c4e5)
}

var fileDescriptor_419a7b10c074c4e5 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0xf3, 0xb5, 0x7c, 0x4e, 0x5a, 0xa8, 0xa1, 0x8b, 0xd0, 0x85, 0xa6, 0xae, 0x02,
	0xe2, 0x04, 0xeb, 0x46, 0xec, 0xca, 0xff, 0xb8, 0x91, 0xd2, 0x22, 0x82, 0x1b, 0x49, 0xc2, 0x25,
	0xc6, 0x26, 0xb9, 0xc3, 0xcc, 0xa4, 0xd0, 0xb7, 0x71, 0xe9, 0x73, 0xf9, 0x24, 0x92, 0x49, 0x1a,
	0x47, 0xb0, 0x28, 0xdd, 0x9d, 0xb9, 0x9c, 0xf9, 0xcd, 0xc9, 0xb9, 0x21, 0xd3, 0x38, 0x91, 0x2f,
	0x45, 0x48, 0x23, 0xcc, 0x7c, 0x81, 0x29, 0x1e, 0x25, 0xe8, 0xc7, 0x29, 0xa2, 0xcf, 0x38, 0xbe,
	0x42, 0x24, 0x45, 0x75, 0x0a, 0x58, 0xe2, 0x2f, 0x8f, 0x7d, 0x96, 0x16, 0x71, 0x92, 0x0b, 0x7f,
	0x51, 0x84, 0xc0, 0x73, 0x90, 0xa0, 0x4b, 0xca, 0x38, 0x4a, 0xb4, 0xf7, 0xf5, 0x49, 0xe5, 0xa7,
	0x25, 0x83, 0x96, 0x78, 0x9a, 0xe0, 0x70, 0x10, 0x63, 0x8c, 0xca, 0xeb, 0x97, 0xaa, 0xba, 0x36,
	0xbc, 0xdd, 0x2a, 0x88, 0x00, 0xbe, 0x4c, 0x22, 0x78, 0x16, 0x0c, 0xa2, 0x1a, 0x74, 0xb3, 0x1d,
	0xa8, 0x08, 0x05, 0x48, 0x8d, 0x73, 0xf0, 0x66, 0x92, 0xee, 0x03, 0x13, 0x92, 0x43, 0x90, 0xcd,
	0x19, 0x44, 0xf6, 0x88, 0x74, 0xd7, 0xcf, 0xe5, 0x41, 0x06, 0x8e, 0xe1, 0x1a, 0xde, 0xce, 0xcc,
	0xaa, 0x67, 0xf7, 0x41, 0x06, 0xf6, 0x21, 0xd9, 0xd5, 0x2d, 0x82, 0x05, 0x11, 0x38, 0x2d, 0xe5,
	0xeb, 0x6b, 0x3e, 0x35, 0xd7, 0x79, 0x0c, 0xb9, 0x74, 0x4c, 0xd7, 0xf0, 0x7a, 0x0d, 0x6f, 0x8a,
	0x5c, 0xda, 0x8f, 0xe4, 0xbf, 0x80, 0x14, 0x22, 0x89, 0xdc, 0xf9, 0xe7, 0x9a, 0x9e, 0x35, 0x9e,
	0xd0, 0x5f, 0xea, 0xa5, 0x7a, 0x66, 0x3a, 0xaf, 0x6f, 0x5f, 0xe7, 0x92, 0xaf, 0x66, 0x0d, 0xcc,
	0xbe, 0xfa, 0x7a, 0xbb, 0xfc, 0x64, 0xa7, 0xed, 0x1a, 0x9e, 0x35, 0x1e, 0xfd, 0x4c, 0x9c, 0x57,
	0xce, 0x12, 0xd8, 0xc4, 0x53, 0x8d, 0x9c, 0x13, 0x4b, 0xeb, 0xcd, 0xe9, 0x28, 0x88, 0xbb, 0x01,
	0xa2, 0x8c, 0x8a, 0x41, 0x44, 0xa3, 0x87, 0x13, 0xd2, 0xfb, 0x96, 0xd1, 0xee, 0x13, 0x73, 0x01,
	0xab, 0xba, 0xdc, 0x52, 0xda, 0x03, 0xd2, 0x5e, 0x06, 0x69, 0xb1, 0x2e, 0xb2, 0x3a, 0x9c, 0xb5,
	0x4e, 0x8d, 0x8b, 0xbb, 0xf7, 0x8f, 0x3d, 0xe3, 0xe9, 0xf2, 0x6f, 0x0b, 0x67, 0x8b, 0x78, 0xf3,
	0x6f, 0x1c, 0x76, 0xd4, 0xd2, 0x4f, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x2f, 0x9f, 0xb4,
	0x10, 0x03, 0x00, 0x00,
}

func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if this.ServiceNamespace != that1.ServiceNamespace {
		return false
	}
	if this.ServicePort != that1.ServicePort {
		return false
	}
	if len(this.Selector) != len(that1.Selector) {
		return false
	}
	for i := range this.Selector {
		if this.Selector[i] != that1.Selector[i] {
			return false
		}
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if !this.SubsetSpec.Equal(that1.SubsetSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
