// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networkServer.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NetworkServer struct {
	// Network-server ID.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Network-server name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Network-server server.
	// Format: hostname:ip (e.g. localhost:8000).
	Server string `protobuf:"bytes,3,opt,name=server,proto3" json:"server,omitempty"`
	// CA certificate (optional).
	CaCert string `protobuf:"bytes,4,opt,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	// TLS (client) certificate for connecting to the network-server (optional).
	TlsCert string `protobuf:"bytes,5,opt,name=tls_cert,json=tlsCert,proto3" json:"tls_cert,omitempty"`
	// TLS (client) key for connecting to the network-server (optional).
	TlsKey string `protobuf:"bytes,6,opt,name=tls_key,json=tlsKey,proto3" json:"tls_key,omitempty"`
	// Routing-profile ca certificate (used by the network-server to connect
	// back to the application-server) (optional).
	RoutingProfileCaCert string `protobuf:"bytes,7,opt,name=routing_profile_ca_cert,json=routingProfileCACert,proto3" json:"routing_profile_ca_cert,omitempty"`
	// Routing-profile TLS certificate (used by the network-server to connect
	// back to the application-server) (optional).
	RoutingProfileTlsCert string `protobuf:"bytes,8,opt,name=routing_profile_tls_cert,json=routingProfileTLSCert,proto3" json:"routing_profile_tls_cert,omitempty"`
	// Routing-profile TLS key (used by the network-server to connect
	// back to the application-server) (optional).
	RoutingProfileTlsKey string `protobuf:"bytes,9,opt,name=routing_profile_tls_key,json=routingProfileTLSKey,proto3" json:"routing_profile_tls_key,omitempty"`
	// Enable gateway discovery for this network-server.
	GatewayDiscoveryEnabled bool `protobuf:"varint,10,opt,name=gateway_discovery_enabled,json=gatewayDiscoveryEnabled,proto3" json:"gateway_discovery_enabled,omitempty"`
	// The number of times per day the gateway discovery 'ping' must be
	// broadcasted per gateway.
	GatewayDiscoveryInterval uint32 `protobuf:"varint,11,opt,name=gateway_discovery_interval,json=gatewayDiscoveryInterval,proto3" json:"gateway_discovery_interval,omitempty"`
	// The frequency (Hz) of the gateway discovery 'ping'.
	GatewayDiscoveryTxFrequency uint32 `protobuf:"varint,12,opt,name=gateway_discovery_tx_frequency,json=gatewayDiscoveryTXFrequency,proto3" json:"gateway_discovery_tx_frequency,omitempty"`
	// The data-rate of the gateway discovery 'ping'.
	GatewayDiscoveryDr   uint32   `protobuf:"varint,13,opt,name=gateway_discovery_dr,json=gatewayDiscoveryDR,proto3" json:"gateway_discovery_dr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkServer) Reset()         { *m = NetworkServer{} }
func (m *NetworkServer) String() string { return proto.CompactTextString(m) }
func (*NetworkServer) ProtoMessage()    {}
func (*NetworkServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_networkServer_e5aa2d4ba970efdf, []int{0}
}
func (m *NetworkServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServer.Unmarshal(m, b)
}
func (m *NetworkServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServer.Marshal(b, m, deterministic)
}
func (dst *NetworkServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServer.Merge(dst, src)
}
func (m *NetworkServer) XXX_Size() int {
	return xxx_messageInfo_NetworkServer.Size(m)
}
func (m *NetworkServer) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServer.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServer proto.InternalMessageInfo

func (m *NetworkServer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NetworkServer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkServer) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *NetworkServer) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func (m *NetworkServer) GetTlsCert() string {
	if m != nil {
		return m.TlsCert
	}
	return ""
}

func (m *NetworkServer) GetTlsKey() string {
	if m != nil {
		return m.TlsKey
	}
	return ""
}

func (m *NetworkServer) GetRoutingProfileCaCert() string {
	if m != nil {
		return m.RoutingProfileCaCert
	}
	return ""
}

func (m *NetworkServer) GetRoutingProfileTlsCert() string {
	if m != nil {
		return m.RoutingProfileTlsCert
	}
	return ""
}

func (m *NetworkServer) GetRoutingProfileTlsKey() string {
	if m != nil {
		return m.RoutingProfileTlsKey
	}
	return ""
}

func (m *NetworkServer) GetGatewayDiscoveryEnabled() bool {
	if m != nil {
		return m.GatewayDiscoveryEnabled
	}
	return false
}

func (m *NetworkServer) GetGatewayDiscoveryInterval() uint32 {
	if m != nil {
		return m.GatewayDiscoveryInterval
	}
	return 0
}

func (m *NetworkServer) GetGatewayDiscoveryTxFrequency() uint32 {
	if m != nil {
		return m.GatewayDiscoveryTxFrequency
	}
	return 0
}

func (m *NetworkServer) GetGatewayDiscoveryDr() uint32 {
	if m != nil {
		return m.GatewayDiscoveryDr
	}
	return 0
}

type NetworkServerListItem struct {
	// Network-server ID.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Network-server name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Network-server server.
	// Format: hostname:ip (e.g. localhost:8000).
	Server string `protobuf:"bytes,3,opt,name=server,proto3" json:"server,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NetworkServerListItem) Reset()         { *m = NetworkServerListItem{} }
func (m *NetworkServerListItem) String() string { return proto.CompactTextString(m) }
func (*NetworkServerListItem) ProtoMessage()    {}
func (*NetworkServerListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_networkServer_e5aa2d4ba970efdf, []int{1}
}
func (m *NetworkServerListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServerListItem.Unmarshal(m, b)
}
func (m *NetworkServerListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServerListItem.Marshal(b, m, deterministic)
}
func (dst *NetworkServerListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServerListItem.Merge(dst, src)
}
func (m *NetworkServerListItem) XXX_Size() int {
	return xxx_messageInfo_NetworkServerListItem.Size(m)
}
func (m *NetworkServerListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServerListItem.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServerListItem proto.InternalMessageInfo

func (m *NetworkServerListItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NetworkServerListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkServerListItem) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *NetworkServerListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *NetworkServerListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateNetworkServerRequest struct {
	// Network-server object to create.
	NetworkServer        *NetworkServer `protobuf:"bytes,1,opt,name=network_server,json=networkServer,proto3" json:"network_server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateNetworkServerRequest) Reset()         { *m = CreateNetworkServerRequest{} }
func (m *CreateNetworkServerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNetworkServerRequest) ProtoMessage()    {}
func (*CreateNetworkServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_networkServer_e5aa2d4ba970efdf, []int{2}
}
func (m *CreateNetworkServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNetworkServerRequest.Unmarshal(m, b)
}
func (m *CreateNetworkServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNetworkServerRequest.Marshal(b, m, deterministic)
}
func (dst *CreateNetworkServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNetworkServerRequest.Merge(dst, src)
}
func (m *CreateNetworkServerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNetworkServerRequest.Size(m)
}
func (m *CreateNetworkServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNetworkServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNetworkServerRequest proto.InternalMessageInfo

func (m *CreateNetworkServerRequest) GetNetworkServer() *NetworkServer {
	if m != nil {
		return m.NetworkServer
	}
	return nil
}

type CreateNetworkServerResponse struct {
	// Network-server ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNetworkServerResponse) Reset()         { *m = CreateNetworkServerResponse{} }
func (m *CreateNetworkServerResponse) String() string { return proto.CompactTextString(m) }
func (*CreateNetworkServerResponse) ProtoMessage()    {}
func (*CreateNetworkServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_networkServer_e5aa2d4ba970efdf, []int{3}
}
func (m *CreateNetworkServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNetworkServerResponse.Unmarshal(m, b)
}
func (m *CreateNetworkServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNetworkServerResponse.Marshal(b, m, deterministic)
}
func (dst *CreateNetworkServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNetworkServerResponse.Merge(dst, src)
}
func (m *CreateNetworkServerResponse) XXX_Size() int {
	return xxx_messageInfo_CreateNetworkServerResponse.Size(m)
}
func (m *CreateNetworkServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNetworkServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNetworkServerResponse proto.InternalMessageInfo

func (m *CreateNetworkServerResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetNetworkServerRequest struct {
	// Network-server ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNetworkServerRequest) Reset()         { *m = GetNetworkServerRequest{} }
func (m *GetNetworkServerRequest) String() string { return proto.CompactTextString(m) }
func (*GetNetworkServerRequest) ProtoMessage()    {}
func (*GetNetworkServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_networkServer_e5aa2d4ba970efdf, []int{4}
}
func (m *GetNetworkServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNetworkServerRequest.Unmarshal(m, b)
}
func (m *GetNetworkServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNetworkServerRequest.Marshal(b, m, deterministic)
}
func (dst *GetNetworkServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNetworkServerRequest.Merge(dst, src)
}
func (m *GetNetworkServerRequest) XXX_Size() int {
	return xxx_messageInfo_GetNetworkServerRequest.Size(m)
}
func (m *GetNetworkServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNetworkServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNetworkServerRequest proto.InternalMessageInfo

func (m *GetNetworkServerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetNetworkServerResponse struct {
	// Network-server object.
	NetworkServer *NetworkServer `protobuf:"bytes,1,opt,name=network_server,json=networkServer,proto3" json:"network_server,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// The LoRa Server version.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// The LoRa Server region configured.
	Region               string   `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNetworkServerResponse) Reset()         { *m = GetNetworkServerResponse{} }
func (m *GetNetworkServerResponse) String() string { return proto.CompactTextString(m) }
func (*GetNetworkServerResponse) ProtoMessage()    {}
func (*GetNetworkServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_networkServer_e5aa2d4ba970efdf, []int{5}
}
func (m *GetNetworkServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNetworkServerResponse.Unmarshal(m, b)
}
func (m *GetNetworkServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNetworkServerResponse.Marshal(b, m, deterministic)
}
func (dst *GetNetworkServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNetworkServerResponse.Merge(dst, src)
}
func (m *GetNetworkServerResponse) XXX_Size() int {
	return xxx_messageInfo_GetNetworkServerResponse.Size(m)
}
func (m *GetNetworkServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNetworkServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNetworkServerResponse proto.InternalMessageInfo

func (m *GetNetworkServerResponse) GetNetworkServer() *NetworkServer {
	if m != nil {
		return m.NetworkServer
	}
	return nil
}

func (m *GetNetworkServerResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetNetworkServerResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *GetNetworkServerResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetNetworkServerResponse) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

type UpdateNetworkServerRequest struct {
	// Network-server object to update.
	NetworkServer        *NetworkServer `protobuf:"bytes,1,opt,name=network_server,json=networkServer,proto3" json:"network_server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateNetworkServerRequest) Reset()         { *m = UpdateNetworkServerRequest{} }
func (m *UpdateNetworkServerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateNetworkServerRequest) ProtoMessage()    {}
func (*UpdateNetworkServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_networkServer_e5aa2d4ba970efdf, []int{6}
}
func (m *UpdateNetworkServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNetworkServerRequest.Unmarshal(m, b)
}
func (m *UpdateNetworkServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNetworkServerRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateNetworkServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNetworkServerRequest.Merge(dst, src)
}
func (m *UpdateNetworkServerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateNetworkServerRequest.Size(m)
}
func (m *UpdateNetworkServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNetworkServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNetworkServerRequest proto.InternalMessageInfo

func (m *UpdateNetworkServerRequest) GetNetworkServer() *NetworkServer {
	if m != nil {
		return m.NetworkServer
	}
	return nil
}

type DeleteNetworkServerRequest struct {
	// Network-server ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNetworkServerRequest) Reset()         { *m = DeleteNetworkServerRequest{} }
func (m *DeleteNetworkServerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteNetworkServerRequest) ProtoMessage()    {}
func (*DeleteNetworkServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_networkServer_e5aa2d4ba970efdf, []int{7}
}
func (m *DeleteNetworkServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNetworkServerRequest.Unmarshal(m, b)
}
func (m *DeleteNetworkServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNetworkServerRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteNetworkServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNetworkServerRequest.Merge(dst, src)
}
func (m *DeleteNetworkServerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteNetworkServerRequest.Size(m)
}
func (m *DeleteNetworkServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNetworkServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNetworkServerRequest proto.InternalMessageInfo

func (m *DeleteNetworkServerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListNetworkServerRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Organization id to filter on.
	OrganizationId       int64    `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNetworkServerRequest) Reset()         { *m = ListNetworkServerRequest{} }
func (m *ListNetworkServerRequest) String() string { return proto.CompactTextString(m) }
func (*ListNetworkServerRequest) ProtoMessage()    {}
func (*ListNetworkServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_networkServer_e5aa2d4ba970efdf, []int{8}
}
func (m *ListNetworkServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNetworkServerRequest.Unmarshal(m, b)
}
func (m *ListNetworkServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNetworkServerRequest.Marshal(b, m, deterministic)
}
func (dst *ListNetworkServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNetworkServerRequest.Merge(dst, src)
}
func (m *ListNetworkServerRequest) XXX_Size() int {
	return xxx_messageInfo_ListNetworkServerRequest.Size(m)
}
func (m *ListNetworkServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNetworkServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNetworkServerRequest proto.InternalMessageInfo

func (m *ListNetworkServerRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListNetworkServerRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListNetworkServerRequest) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

type ListNetworkServerResponse struct {
	// Total number of network-servers.
	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Network-servers within the result-set.
	Result               []*NetworkServerListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ListNetworkServerResponse) Reset()         { *m = ListNetworkServerResponse{} }
func (m *ListNetworkServerResponse) String() string { return proto.CompactTextString(m) }
func (*ListNetworkServerResponse) ProtoMessage()    {}
func (*ListNetworkServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_networkServer_e5aa2d4ba970efdf, []int{9}
}
func (m *ListNetworkServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNetworkServerResponse.Unmarshal(m, b)
}
func (m *ListNetworkServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNetworkServerResponse.Marshal(b, m, deterministic)
}
func (dst *ListNetworkServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNetworkServerResponse.Merge(dst, src)
}
func (m *ListNetworkServerResponse) XXX_Size() int {
	return xxx_messageInfo_ListNetworkServerResponse.Size(m)
}
func (m *ListNetworkServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNetworkServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNetworkServerResponse proto.InternalMessageInfo

func (m *ListNetworkServerResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListNetworkServerResponse) GetResult() []*NetworkServerListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkServer)(nil), "api.NetworkServer")
	proto.RegisterType((*NetworkServerListItem)(nil), "api.NetworkServerListItem")
	proto.RegisterType((*CreateNetworkServerRequest)(nil), "api.CreateNetworkServerRequest")
	proto.RegisterType((*CreateNetworkServerResponse)(nil), "api.CreateNetworkServerResponse")
	proto.RegisterType((*GetNetworkServerRequest)(nil), "api.GetNetworkServerRequest")
	proto.RegisterType((*GetNetworkServerResponse)(nil), "api.GetNetworkServerResponse")
	proto.RegisterType((*UpdateNetworkServerRequest)(nil), "api.UpdateNetworkServerRequest")
	proto.RegisterType((*DeleteNetworkServerRequest)(nil), "api.DeleteNetworkServerRequest")
	proto.RegisterType((*ListNetworkServerRequest)(nil), "api.ListNetworkServerRequest")
	proto.RegisterType((*ListNetworkServerResponse)(nil), "api.ListNetworkServerResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServerServiceClient is the client API for NetworkServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServerServiceClient interface {
	// Create creates the given network-server.
	Create(ctx context.Context, in *CreateNetworkServerRequest, opts ...grpc.CallOption) (*CreateNetworkServerResponse, error)
	// Get returns the network-server matching the given id.
	Get(ctx context.Context, in *GetNetworkServerRequest, opts ...grpc.CallOption) (*GetNetworkServerResponse, error)
	// Update updates the given network-server.
	Update(ctx context.Context, in *UpdateNetworkServerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete deletes the network-server matching the given id.
	Delete(ctx context.Context, in *DeleteNetworkServerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List lists the available network-servers.
	List(ctx context.Context, in *ListNetworkServerRequest, opts ...grpc.CallOption) (*ListNetworkServerResponse, error)
}

type networkServerServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServerServiceClient(cc *grpc.ClientConn) NetworkServerServiceClient {
	return &networkServerServiceClient{cc}
}

func (c *networkServerServiceClient) Create(ctx context.Context, in *CreateNetworkServerRequest, opts ...grpc.CallOption) (*CreateNetworkServerResponse, error) {
	out := new(CreateNetworkServerResponse)
	err := c.cc.Invoke(ctx, "/api.NetworkServerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerServiceClient) Get(ctx context.Context, in *GetNetworkServerRequest, opts ...grpc.CallOption) (*GetNetworkServerResponse, error) {
	out := new(GetNetworkServerResponse)
	err := c.cc.Invoke(ctx, "/api.NetworkServerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerServiceClient) Update(ctx context.Context, in *UpdateNetworkServerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.NetworkServerService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerServiceClient) Delete(ctx context.Context, in *DeleteNetworkServerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.NetworkServerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerServiceClient) List(ctx context.Context, in *ListNetworkServerRequest, opts ...grpc.CallOption) (*ListNetworkServerResponse, error) {
	out := new(ListNetworkServerResponse)
	err := c.cc.Invoke(ctx, "/api.NetworkServerService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServerServiceServer is the server API for NetworkServerService service.
type NetworkServerServiceServer interface {
	// Create creates the given network-server.
	Create(context.Context, *CreateNetworkServerRequest) (*CreateNetworkServerResponse, error)
	// Get returns the network-server matching the given id.
	Get(context.Context, *GetNetworkServerRequest) (*GetNetworkServerResponse, error)
	// Update updates the given network-server.
	Update(context.Context, *UpdateNetworkServerRequest) (*empty.Empty, error)
	// Delete deletes the network-server matching the given id.
	Delete(context.Context, *DeleteNetworkServerRequest) (*empty.Empty, error)
	// List lists the available network-servers.
	List(context.Context, *ListNetworkServerRequest) (*ListNetworkServerResponse, error)
}

func RegisterNetworkServerServiceServer(s *grpc.Server, srv NetworkServerServiceServer) {
	s.RegisterService(&_NetworkServerService_serviceDesc, srv)
}

func _NetworkServerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServiceServer).Create(ctx, req.(*CreateNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServiceServer).Get(ctx, req.(*GetNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServerService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServiceServer).Update(ctx, req.(*UpdateNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServiceServer).Delete(ctx, req.(*DeleteNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworkServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NetworkServerService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServiceServer).List(ctx, req.(*ListNetworkServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.NetworkServerService",
	HandlerType: (*NetworkServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NetworkServerService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _NetworkServerService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NetworkServerService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NetworkServerService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NetworkServerService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networkServer.proto",
}

func init() { proto.RegisterFile("networkServer.proto", fileDescriptor_networkServer_e5aa2d4ba970efdf) }

var fileDescriptor_networkServer_e5aa2d4ba970efdf = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x96, 0xe3, 0x34, 0x6d, 0x4f, 0x69, 0x91, 0x86, 0x6c, 0xe3, 0xb8, 0x65, 0x1b, 0x7c, 0x43,
	0x58, 0xd1, 0x04, 0x75, 0x85, 0x10, 0x2b, 0x6e, 0xaa, 0x64, 0x59, 0x55, 0x54, 0x08, 0xb9, 0x41,
	0x70, 0x67, 0x4d, 0xed, 0x93, 0x68, 0xb4, 0x8e, 0xc7, 0x3b, 0x9e, 0x64, 0x09, 0x68, 0x6f, 0x78,
	0x05, 0xde, 0x03, 0xf1, 0x1a, 0x5c, 0xf3, 0x0a, 0xbc, 0x06, 0x12, 0x9a, 0x1f, 0xa3, 0x8d, 0x63,
	0x83, 0x58, 0xb8, 0x89, 0x32, 0xfe, 0xbe, 0xef, 0x7c, 0x33, 0xe7, 0x7c, 0x63, 0xc3, 0x3b, 0x19,
	0xca, 0x97, 0x5c, 0x3c, 0xbf, 0x43, 0xb1, 0x46, 0x31, 0xca, 0x05, 0x97, 0x9c, 0xb8, 0x34, 0x67,
	0xfe, 0xf9, 0x82, 0xf3, 0x45, 0x8a, 0x63, 0x9a, 0xb3, 0x31, 0xcd, 0x32, 0x2e, 0xa9, 0x64, 0x3c,
	0x2b, 0x0c, 0xc5, 0xbf, 0xb0, 0xa8, 0x5e, 0xdd, 0xaf, 0xe6, 0x63, 0xc9, 0x96, 0x58, 0x48, 0xba,
	0xcc, 0x2d, 0xe1, 0xac, 0x4a, 0xc0, 0x65, 0x2e, 0x37, 0x06, 0x0c, 0x7e, 0x69, 0xc3, 0xf1, 0x97,
	0xaf, 0x1b, 0x93, 0x13, 0x68, 0xb1, 0xc4, 0x73, 0x06, 0xce, 0xd0, 0x0d, 0x5b, 0x2c, 0x21, 0x04,
	0xda, 0x19, 0x5d, 0xa2, 0xd7, 0x1a, 0x38, 0xc3, 0xc3, 0x50, 0xff, 0x27, 0xa7, 0xd0, 0x29, 0x34,
	0xdb, 0x73, 0xf5, 0x53, 0xbb, 0x22, 0x3d, 0xd8, 0x8f, 0x69, 0x14, 0xa3, 0x90, 0x5e, 0xdb, 0x00,
	0x31, 0x9d, 0xa0, 0x90, 0xa4, 0x0f, 0x07, 0x32, 0x2d, 0x0c, 0xb2, 0xa7, 0x91, 0x7d, 0x99, 0x16,
	0x1a, 0xea, 0x81, 0xfa, 0x1b, 0x3d, 0xc7, 0x8d, 0xd7, 0x31, 0x1a, 0x99, 0x16, 0x5f, 0xe0, 0x86,
	0x7c, 0x0c, 0x3d, 0xc1, 0x57, 0x92, 0x65, 0x8b, 0x28, 0x17, 0x7c, 0xce, 0x52, 0x8c, 0xca, 0xe2,
	0xfb, 0x9a, 0xd8, 0xb5, 0xf0, 0x57, 0x06, 0x9d, 0x5c, 0xeb, 0x7a, 0x9f, 0x80, 0x57, 0x95, 0xfd,
	0x65, 0x7d, 0xa0, 0x75, 0x0f, 0xb6, 0x75, 0xb3, 0xdb, 0x3b, 0x2d, 0xac, 0xf1, 0x2b, 0x37, 0x76,
	0x58, 0xe7, 0x37, 0xbb, 0xbd, 0x53, 0xdb, 0x7c, 0x02, 0xfd, 0x05, 0x95, 0xf8, 0x92, 0x6e, 0xa2,
	0x84, 0x15, 0x31, 0x5f, 0xa3, 0xd8, 0x44, 0x98, 0xd1, 0xfb, 0x14, 0x13, 0x0f, 0x06, 0xce, 0xf0,
	0x20, 0xec, 0x59, 0xc2, 0xb4, 0xc4, 0x9f, 0x1a, 0x98, 0x7c, 0x06, 0xfe, 0xae, 0x96, 0x65, 0x12,
	0xc5, 0x9a, 0xa6, 0xde, 0xd1, 0xc0, 0x19, 0x1e, 0x87, 0x5e, 0x55, 0x7c, 0x63, 0x71, 0x32, 0x81,
	0x87, 0xbb, 0x6a, 0xf9, 0x5d, 0x34, 0x17, 0xf8, 0x62, 0x85, 0x59, 0xbc, 0xf1, 0xde, 0xd2, 0x15,
	0xce, 0xaa, 0x15, 0x66, 0xdf, 0x7e, 0x5e, 0x52, 0xc8, 0x47, 0xd0, 0xdd, 0x2d, 0x92, 0x08, 0xef,
	0x58, 0x4b, 0x49, 0x55, 0x3a, 0x0d, 0x83, 0x5f, 0x1d, 0x78, 0xb0, 0x15, 0x99, 0x5b, 0x56, 0xc8,
	0x1b, 0x89, 0xcb, 0xff, 0x14, 0x9d, 0x4f, 0x01, 0x62, 0x81, 0x54, 0x62, 0x12, 0x51, 0x93, 0x9e,
	0xa3, 0x2b, 0x7f, 0x64, 0xa2, 0x3b, 0x2a, 0xa3, 0x3b, 0x9a, 0x95, 0xd9, 0x0e, 0x0f, 0x2d, 0xfb,
	0x5a, 0x2a, 0xe9, 0x2a, 0x4f, 0x4a, 0xe9, 0xde, 0x3f, 0x4b, 0x2d, 0xfb, 0x5a, 0x06, 0xdf, 0x80,
	0x3f, 0xd1, 0x75, 0xb6, 0x0e, 0x14, 0xaa, 0xe6, 0x14, 0xaa, 0xf0, 0x89, 0xbd, 0x94, 0x91, 0xdd,
	0xb3, 0xa3, 0x8b, 0x93, 0x11, 0xcd, 0xd9, 0x68, 0x5b, 0x72, 0xbc, 0x75, 0x7d, 0x83, 0x4b, 0x38,
	0xab, 0x2d, 0x5c, 0xe4, 0x3c, 0x2b, 0xb0, 0xda, 0xa9, 0xe0, 0x03, 0xe8, 0x3d, 0x43, 0x59, 0xbb,
	0x89, 0x2a, 0xf5, 0x0f, 0x07, 0xbc, 0x5d, 0xae, 0xad, 0xfb, 0xe6, 0x3b, 0xae, 0x0c, 0xa0, 0xf5,
	0xe6, 0x03, 0x70, 0xff, 0xc5, 0x00, 0x88, 0x07, 0xfb, 0x6b, 0x14, 0x05, 0xe3, 0x99, 0x7d, 0x63,
	0x94, 0x4b, 0x15, 0x14, 0x81, 0x0b, 0x05, 0x98, 0x17, 0x86, 0x5d, 0xa9, 0x91, 0x7d, 0xad, 0xe5,
	0xff, 0xf7, 0xc8, 0x3e, 0x04, 0x7f, 0x8a, 0x29, 0x36, 0x14, 0xae, 0x8e, 0xe1, 0x05, 0x78, 0x2a,
	0xf7, 0xb5, 0xdc, 0x2e, 0xec, 0xa5, 0x6c, 0xc9, 0xa4, 0xa5, 0x9b, 0x85, 0x3a, 0x10, 0x9f, 0xcf,
	0x0b, 0x34, 0xcd, 0x75, 0x43, 0xbb, 0x22, 0xef, 0xc3, 0xdb, 0x5c, 0x2c, 0x68, 0xc6, 0xbe, 0xd7,
	0xef, 0xf5, 0x88, 0x25, 0xba, 0x85, 0x6e, 0x78, 0xf2, 0xfa, 0xe3, 0x9b, 0x69, 0x90, 0x43, 0xbf,
	0xc6, 0xd2, 0x4e, 0xfe, 0x02, 0x8e, 0x24, 0x97, 0x34, 0x8d, 0x62, 0xbe, 0xca, 0x4a, 0x67, 0xd0,
	0x8f, 0x26, 0xea, 0x09, 0xb9, 0x52, 0xfd, 0x2c, 0x56, 0xa9, 0xb2, 0x77, 0xf5, 0x80, 0x76, 0x3a,
	0x52, 0x5e, 0xe4, 0xd0, 0x32, 0xaf, 0x7e, 0x6e, 0x43, 0x77, 0x8b, 0xa1, 0x7e, 0x59, 0x8c, 0x24,
	0x85, 0x8e, 0x89, 0x37, 0xb9, 0xd0, 0x65, 0x9a, 0x2f, 0x91, 0x3f, 0x68, 0x26, 0x98, 0xad, 0x07,
	0x17, 0x3f, 0xfe, 0xf6, 0xfb, 0x4f, 0xad, 0x7e, 0xd0, 0xd5, 0x5f, 0x38, 0x3b, 0x94, 0x4b, 0x33,
	0xbe, 0xe2, 0x89, 0xf3, 0x88, 0x20, 0xb8, 0xcf, 0x50, 0x92, 0x73, 0x5d, 0xa9, 0xe1, 0x9e, 0xf8,
	0xef, 0x36, 0xa0, 0xd6, 0xe4, 0x3d, 0x6d, 0x72, 0x46, 0xfa, 0x75, 0x26, 0xe3, 0x1f, 0x58, 0xf2,
	0x8a, 0xac, 0xa1, 0x63, 0x92, 0x65, 0x0f, 0xd5, 0x1c, 0x33, 0xff, 0x74, 0x27, 0xdd, 0x4f, 0xd5,
	0x47, 0x35, 0x78, 0xac, 0x5d, 0x2e, 0xfd, 0x61, 0xbd, 0xcb, 0x76, 0x34, 0x47, 0x2c, 0x79, 0xa5,
	0x8e, 0x97, 0x40, 0xc7, 0x04, 0xcf, 0xfa, 0x36, 0xa7, 0xb0, 0xd1, 0xd7, 0x9e, 0xee, 0xd1, 0xdf,
	0x9c, 0x2e, 0x86, 0xb6, 0x9a, 0x2f, 0x31, 0x7d, 0x6a, 0xca, 0xae, 0xff, 0xb0, 0x09, 0xb6, 0x7d,
	0x3c, 0xd7, 0x4e, 0xa7, 0xa4, 0x76, 0x58, 0xf7, 0x1d, 0xbd, 0xaf, 0xc7, 0x7f, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x2f, 0x2e, 0xd5, 0xe2, 0xcd, 0x08, 0x00, 0x00,
}
