// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/endpoint.proto

package envoy_api_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ClusterLoadAssignment struct {
	ClusterName          string                          `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Endpoints            []*endpoint.LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	NamedEndpoints       map[string]*endpoint.Endpoint   `protobuf:"bytes,5,rep,name=named_endpoints,json=namedEndpoints,proto3" json:"named_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Policy               *ClusterLoadAssignment_Policy   `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ClusterLoadAssignment) Reset()         { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()    {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dd3a9d301e6758b, []int{0}
}

func (m *ClusterLoadAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment.Merge(m, src)
}
func (m *ClusterLoadAssignment) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment.Size(m)
}
func (m *ClusterLoadAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment proto.InternalMessageInfo

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*endpoint.LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetNamedEndpoints() map[string]*endpoint.Endpoint {
	if m != nil {
		return m.NamedEndpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ClusterLoadAssignment_Policy struct {
	DropOverloads           []*ClusterLoadAssignment_Policy_DropOverload `protobuf:"bytes,2,rep,name=drop_overloads,json=dropOverloads,proto3" json:"drop_overloads,omitempty"`
	OverprovisioningFactor  *wrappers.UInt32Value                        `protobuf:"bytes,3,opt,name=overprovisioning_factor,json=overprovisioningFactor,proto3" json:"overprovisioning_factor,omitempty"`
	EndpointStaleAfter      *duration.Duration                           `protobuf:"bytes,4,opt,name=endpoint_stale_after,json=endpointStaleAfter,proto3" json:"endpoint_stale_after,omitempty"`
	DisableOverprovisioning bool                                         `protobuf:"varint,5,opt,name=disable_overprovisioning,json=disableOverprovisioning,proto3" json:"disable_overprovisioning,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral    struct{}                                     `json:"-"`
	XXX_unrecognized        []byte                                       `json:"-"`
	XXX_sizecache           int32                                        `json:"-"`
}

func (m *ClusterLoadAssignment_Policy) Reset()         { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dd3a9d301e6758b, []int{0, 0}
}

func (m *ClusterLoadAssignment_Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Size(m)
}
func (m *ClusterLoadAssignment_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy) GetDropOverloads() []*ClusterLoadAssignment_Policy_DropOverload {
	if m != nil {
		return m.DropOverloads
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetOverprovisioningFactor() *wrappers.UInt32Value {
	if m != nil {
		return m.OverprovisioningFactor
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetEndpointStaleAfter() *duration.Duration {
	if m != nil {
		return m.EndpointStaleAfter
	}
	return nil
}

// Deprecated: Do not use.
func (m *ClusterLoadAssignment_Policy) GetDisableOverprovisioning() bool {
	if m != nil {
		return m.DisableOverprovisioning
	}
	return false
}

type ClusterLoadAssignment_Policy_DropOverload struct {
	Category             string                   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	DropPercentage       *_type.FractionalPercent `protobuf:"bytes,2,opt,name=drop_percentage,json=dropPercentage,proto3" json:"drop_percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ClusterLoadAssignment_Policy_DropOverload) Reset() {
	*m = ClusterLoadAssignment_Policy_DropOverload{}
}
func (m *ClusterLoadAssignment_Policy_DropOverload) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy_DropOverload) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy_DropOverload) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dd3a9d301e6758b, []int{0, 0, 0}
}

func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Size(m)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy_DropOverload) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ClusterLoadAssignment_Policy_DropOverload) GetDropPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.DropPercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.api.v2.ClusterLoadAssignment")
	proto.RegisterMapType((map[string]*endpoint.Endpoint)(nil), "envoy.api.v2.ClusterLoadAssignment.NamedEndpointsEntry")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy")
	proto.RegisterType((*ClusterLoadAssignment_Policy_DropOverload)(nil), "envoy.api.v2.ClusterLoadAssignment.Policy.DropOverload")
}

func init() { proto.RegisterFile("envoy/api/v2/endpoint.proto", fileDescriptor_8dd3a9d301e6758b) }

var fileDescriptor_8dd3a9d301e6758b = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0xc7, 0x6b, 0xa7, 0x09, 0xe9, 0xf4, 0x53, 0xc3, 0x47, 0x8d, 0x69, 0x21, 0x82, 0x4d, 0x95,
	0x85, 0x2d, 0xa5, 0x42, 0x45, 0x48, 0x2c, 0x6a, 0xda, 0x0a, 0x50, 0x45, 0x23, 0xa3, 0xb2, 0xac,
	0x99, 0xd8, 0x13, 0x6b, 0x84, 0x33, 0x33, 0x1a, 0x8f, 0x0d, 0x16, 0x1b, 0x6e, 0xc0, 0x96, 0x33,
	0x70, 0x04, 0x4e, 0xc0, 0x96, 0x23, 0x70, 0x05, 0x96, 0x45, 0x42, 0xc8, 0xe3, 0xb1, 0x9b, 0xa6,
	0xa9, 0xc4, 0x6e, 0xec, 0xff, 0xfb, 0xff, 0xde, 0x9b, 0xf7, 0x9e, 0x0d, 0xee, 0x61, 0x9a, 0xb3,
	0xc2, 0x45, 0x9c, 0xb8, 0xf9, 0xc0, 0xc5, 0x34, 0xe2, 0x8c, 0x50, 0xe9, 0x70, 0xc1, 0x24, 0x83,
	0x2b, 0x4a, 0x74, 0x10, 0x27, 0x4e, 0x3e, 0xb0, 0xdd, 0xb9, 0xa1, 0xcd, 0x21, 0x08, 0xd9, 0x84,
	0x33, 0x8a, 0xa9, 0x4c, 0x2b, 0xbb, 0x6d, 0x55, 0x06, 0x59, 0x70, 0xec, 0x72, 0x2c, 0x42, 0x5c,
	0x83, 0xed, 0xad, 0x98, 0xb1, 0x38, 0xc1, 0x8a, 0x85, 0x28, 0x65, 0x12, 0x49, 0xc2, 0x68, 0xed,
	0xbb, 0xaf, 0x55, 0xf5, 0x34, 0xca, 0xc6, 0x6e, 0x94, 0x09, 0x15, 0x70, 0x9d, 0xfe, 0x41, 0x20,
	0xce, 0xb1, 0x68, 0xfc, 0x59, 0xc4, 0xd1, 0x34, 0xd7, 0x9d, 0x90, 0x58, 0x20, 0x89, 0xb5, 0xbe,
	0x7d, 0x45, 0x4f, 0x25, 0x92, 0x59, 0x6d, 0xdf, 0xcc, 0x51, 0x42, 0x22, 0x24, 0xb1, 0x5b, 0x1f,
	0x2a, 0xe1, 0xe1, 0x9f, 0x0e, 0xb8, 0xfd, 0x3c, 0xc9, 0x52, 0x89, 0xc5, 0x31, 0x43, 0xd1, 0x7e,
	0x9a, 0x92, 0x98, 0x4e, 0x30, 0x95, 0xb0, 0x0f, 0x56, 0xc2, 0x4a, 0x08, 0x28, 0x9a, 0x60, 0xcb,
	0xe8, 0x19, 0x3b, 0x4b, 0xde, 0x8d, 0x73, 0x6f, 0x51, 0x98, 0x3d, 0xc3, 0x5f, 0xd6, 0xe2, 0x6b,
	0x34, 0xc1, 0xf0, 0x05, 0x58, 0xaa, 0x5b, 0x96, 0x5a, 0x66, 0xaf, 0xb5, 0xb3, 0x3c, 0xe8, 0x3b,
	0xd3, 0x8d, 0x76, 0x9a, 0x29, 0x1c, 0xb3, 0x10, 0x25, 0x44, 0x16, 0xc7, 0xa3, 0xc3, 0xda, 0xe1,
	0x5f, 0x98, 0xe1, 0x3b, 0xb0, 0x5e, 0x66, 0x8b, 0x82, 0x0b, 0x5e, 0x5b, 0xf1, 0xf6, 0x2e, 0xf3,
	0xe6, 0xd6, 0xec, 0x94, 0xc5, 0x44, 0x0d, 0xf7, 0x90, 0x4a, 0x51, 0xf8, 0x6b, 0xf4, 0xd2, 0x4b,
	0xe8, 0x81, 0x0e, 0x67, 0x09, 0x09, 0x0b, 0x6b, 0xb1, 0x67, 0x5c, 0x2d, 0x74, 0x3e, 0x78, 0xa8,
	0x1c, 0xbe, 0x76, 0xda, 0xbf, 0x5a, 0xa0, 0x53, 0xbd, 0x82, 0x67, 0x60, 0x2d, 0x12, 0x8c, 0x07,
	0x2c, 0xc7, 0x22, 0x61, 0x28, 0xaa, 0xef, 0xbf, 0xf7, 0xff, 0x58, 0xe7, 0x40, 0x30, 0x7e, 0xa2,
	0xfd, 0xfe, 0x6a, 0x34, 0xf5, 0x94, 0xc2, 0x33, 0xb0, 0x59, 0xa2, 0xb9, 0x60, 0x39, 0x49, 0x09,
	0xa3, 0x84, 0xc6, 0xc1, 0x18, 0x85, 0x92, 0x09, 0xab, 0xa5, 0xea, 0xdf, 0x72, 0xaa, 0xd5, 0x71,
	0xea, 0xd5, 0x71, 0x4e, 0x5f, 0x52, 0xb9, 0x3b, 0x78, 0x8b, 0x92, 0x0c, 0xab, 0x79, 0xf5, 0xcd,
	0xde, 0x82, 0x7f, 0x67, 0x96, 0x72, 0xa4, 0x20, 0xf0, 0x14, 0xdc, 0x6a, 0xb6, 0x3d, 0x95, 0x28,
	0xc1, 0x01, 0x1a, 0x4b, 0x2c, 0x74, 0x73, 0xee, 0x5e, 0x81, 0x1f, 0xe8, 0xbd, 0xf5, 0xba, 0xe7,
	0x5e, 0xfb, 0x9b, 0x61, 0xf6, 0x17, 0x7c, 0x58, 0x03, 0xde, 0x94, 0xfe, 0xfd, 0xd2, 0x0e, 0x9f,
	0x01, 0x2b, 0x22, 0x29, 0x1a, 0x25, 0x38, 0x98, 0x4d, 0x6c, 0xb5, 0x7b, 0xc6, 0x4e, 0xd7, 0x33,
	0x2d, 0xc3, 0xdf, 0xd4, 0x31, 0x27, 0x33, 0x21, 0xf6, 0x27, 0xb0, 0x32, 0xdd, 0x14, 0xf8, 0x08,
	0x74, 0x43, 0x24, 0x71, 0xcc, 0x44, 0x31, 0xbb, 0x88, 0x8d, 0x00, 0x8f, 0xc0, 0xba, 0x1a, 0x85,
	0xfe, 0x2e, 0x51, 0x8c, 0x2d, 0x53, 0xdd, 0x62, 0x5b, 0xcf, 0xa2, 0xfc, 0x6a, 0x9d, 0x23, 0x81,
	0xc2, 0xf2, 0x02, 0x28, 0x19, 0x56, 0x71, 0xbe, 0x1a, 0xe0, 0xb0, 0x31, 0xbd, 0x5a, 0xec, 0x1a,
	0x1b, 0xa6, 0x3d, 0x02, 0x37, 0xe7, 0xac, 0x13, 0xdc, 0x00, 0xad, 0xf7, 0x58, 0x17, 0xe1, 0x97,
	0x47, 0xf8, 0x18, 0xb4, 0xf3, 0xb2, 0xd7, 0x3a, 0xd9, 0x83, 0x6b, 0x16, 0xbf, 0xe6, 0xf8, 0x55,
	0xf4, 0x53, 0xf3, 0x89, 0xe1, 0x9d, 0xfe, 0xfe, 0xfa, 0xf7, 0x4b, 0xdb, 0x86, 0xd5, 0x5f, 0xc5,
	0x09, 0x19, 0x1d, 0x93, 0xf8, 0xc2, 0x92, 0xef, 0x7e, 0xff, 0xfc, 0xe3, 0x67, 0xc7, 0xdc, 0x30,
	0x80, 0x4d, 0x58, 0xc5, 0xe5, 0x82, 0x7d, 0x2c, 0x2e, 0xa5, 0xf0, 0x56, 0x6b, 0xf4, 0xb0, 0x1c,
	0xd2, 0xd0, 0x18, 0x75, 0xd4, 0xb4, 0x76, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xef, 0xcb, 0xbf,
	0xb6, 0x09, 0x05, 0x00, 0x00,
}