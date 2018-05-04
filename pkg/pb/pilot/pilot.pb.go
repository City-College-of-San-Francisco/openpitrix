// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openpitrix/pilot/pilot.proto

package pbpilot // import "openpitrix.io/openpitrix/pkg/pb/pilot"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import types "openpitrix.io/openpitrix/pkg/pb/types"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PilotService service

type PilotServiceClient interface {
	GetPilotConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.PilotConfig, error)
	GetFrontdateList(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.FrontgateIdList, error)
	GetFrontgateConfig(ctx context.Context, in *types.FrontgateId, opts ...grpc.CallOption) (*types.FrontgateConfig, error)
	SetFrontgateConfig(ctx context.Context, in *types.FrontgateConfig, opts ...grpc.CallOption) (*types.Empty, error)
	GetDroneConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.DroneConfig, error)
	SetDroneConfig(ctx context.Context, in *types.SetDroneConfigRequest, opts ...grpc.CallOption) (*types.Empty, error)
	IsConfdRunning(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Bool, error)
	StartConfd(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Empty, error)
	StopConfd(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Empty, error)
	RegisterMetadata(ctx context.Context, in *types.SubTask_RegisterMetadata, opts ...grpc.CallOption) (*types.Empty, error)
	DeregisterMetadata(ctx context.Context, in *types.SubTask_DeregisterMetadata, opts ...grpc.CallOption) (*types.Empty, error)
	RegisterCmd(ctx context.Context, in *types.SubTask_RegisterCmd, opts ...grpc.CallOption) (*types.Empty, error)
	DeregisterCmd(ctx context.Context, in *types.SubTask_DeregisterCmd, opts ...grpc.CallOption) (*types.Empty, error)
	ReportSubTaskStatus(ctx context.Context, in *types.SubTaskStatus, opts ...grpc.CallOption) (*types.Empty, error)
	GetSubtaskStatus(ctx context.Context, in *types.SubTaskId, opts ...grpc.CallOption) (*types.SubTaskStatus, error)
	HandleSubtask(ctx context.Context, in *types.SubTaskMessage, opts ...grpc.CallOption) (*types.Empty, error)
	PingPilot(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
	PingFrontgate(ctx context.Context, in *types.FrontgateId, opts ...grpc.CallOption) (*types.Empty, error)
	PingDrone(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Empty, error)
	FrontgateChannel(ctx context.Context, opts ...grpc.CallOption) (PilotService_FrontgateChannelClient, error)
}

type pilotServiceClient struct {
	cc *grpc.ClientConn
}

func NewPilotServiceClient(cc *grpc.ClientConn) PilotServiceClient {
	return &pilotServiceClient{cc}
}

func (c *pilotServiceClient) GetPilotConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.PilotConfig, error) {
	out := new(types.PilotConfig)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/GetPilotConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) GetFrontdateList(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.FrontgateIdList, error) {
	out := new(types.FrontgateIdList)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/GetFrontdateList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) GetFrontgateConfig(ctx context.Context, in *types.FrontgateId, opts ...grpc.CallOption) (*types.FrontgateConfig, error) {
	out := new(types.FrontgateConfig)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/GetFrontgateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) SetFrontgateConfig(ctx context.Context, in *types.FrontgateConfig, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/SetFrontgateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) GetDroneConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.DroneConfig, error) {
	out := new(types.DroneConfig)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/GetDroneConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) SetDroneConfig(ctx context.Context, in *types.SetDroneConfigRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/SetDroneConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) IsConfdRunning(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Bool, error) {
	out := new(types.Bool)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/IsConfdRunning", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) StartConfd(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/StartConfd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) StopConfd(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/StopConfd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) RegisterMetadata(ctx context.Context, in *types.SubTask_RegisterMetadata, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/RegisterMetadata", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) DeregisterMetadata(ctx context.Context, in *types.SubTask_DeregisterMetadata, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/DeregisterMetadata", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) RegisterCmd(ctx context.Context, in *types.SubTask_RegisterCmd, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/RegisterCmd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) DeregisterCmd(ctx context.Context, in *types.SubTask_DeregisterCmd, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/DeregisterCmd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) ReportSubTaskStatus(ctx context.Context, in *types.SubTaskStatus, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/ReportSubTaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) GetSubtaskStatus(ctx context.Context, in *types.SubTaskId, opts ...grpc.CallOption) (*types.SubTaskStatus, error) {
	out := new(types.SubTaskStatus)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/GetSubtaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) HandleSubtask(ctx context.Context, in *types.SubTaskMessage, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/HandleSubtask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) PingPilot(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/PingPilot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) PingFrontgate(ctx context.Context, in *types.FrontgateId, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/PingFrontgate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) PingDrone(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/PingDrone", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) FrontgateChannel(ctx context.Context, opts ...grpc.CallOption) (PilotService_FrontgateChannelClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PilotService_serviceDesc.Streams[0], c.cc, "/openpitrix.pilot.PilotService/FrontgateChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &pilotServiceFrontgateChannelClient{stream}
	return x, nil
}

type PilotService_FrontgateChannelClient interface {
	Send(*types.Bytes) error
	Recv() (*types.Bytes, error)
	grpc.ClientStream
}

type pilotServiceFrontgateChannelClient struct {
	grpc.ClientStream
}

func (x *pilotServiceFrontgateChannelClient) Send(m *types.Bytes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pilotServiceFrontgateChannelClient) Recv() (*types.Bytes, error) {
	m := new(types.Bytes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PilotService service

type PilotServiceServer interface {
	GetPilotConfig(context.Context, *types.Empty) (*types.PilotConfig, error)
	GetFrontdateList(context.Context, *types.Empty) (*types.FrontgateIdList, error)
	GetFrontgateConfig(context.Context, *types.FrontgateId) (*types.FrontgateConfig, error)
	SetFrontgateConfig(context.Context, *types.FrontgateConfig) (*types.Empty, error)
	GetDroneConfig(context.Context, *types.Empty) (*types.DroneConfig, error)
	SetDroneConfig(context.Context, *types.SetDroneConfigRequest) (*types.Empty, error)
	IsConfdRunning(context.Context, *types.DroneEndpoint) (*types.Bool, error)
	StartConfd(context.Context, *types.DroneEndpoint) (*types.Empty, error)
	StopConfd(context.Context, *types.DroneEndpoint) (*types.Empty, error)
	RegisterMetadata(context.Context, *types.SubTask_RegisterMetadata) (*types.Empty, error)
	DeregisterMetadata(context.Context, *types.SubTask_DeregisterMetadata) (*types.Empty, error)
	RegisterCmd(context.Context, *types.SubTask_RegisterCmd) (*types.Empty, error)
	DeregisterCmd(context.Context, *types.SubTask_DeregisterCmd) (*types.Empty, error)
	ReportSubTaskStatus(context.Context, *types.SubTaskStatus) (*types.Empty, error)
	GetSubtaskStatus(context.Context, *types.SubTaskId) (*types.SubTaskStatus, error)
	HandleSubtask(context.Context, *types.SubTaskMessage) (*types.Empty, error)
	PingPilot(context.Context, *types.Empty) (*types.Empty, error)
	PingFrontgate(context.Context, *types.FrontgateId) (*types.Empty, error)
	PingDrone(context.Context, *types.DroneEndpoint) (*types.Empty, error)
	FrontgateChannel(PilotService_FrontgateChannelServer) error
}

func RegisterPilotServiceServer(s *grpc.Server, srv PilotServiceServer) {
	s.RegisterService(&_PilotService_serviceDesc, srv)
}

func _PilotService_GetPilotConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).GetPilotConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/GetPilotConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).GetPilotConfig(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_GetFrontdateList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).GetFrontdateList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/GetFrontdateList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).GetFrontdateList(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_GetFrontgateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).GetFrontgateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/GetFrontgateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).GetFrontgateConfig(ctx, req.(*types.FrontgateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_SetFrontgateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).SetFrontgateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/SetFrontgateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).SetFrontgateConfig(ctx, req.(*types.FrontgateConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_GetDroneConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).GetDroneConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/GetDroneConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).GetDroneConfig(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_SetDroneConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SetDroneConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).SetDroneConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/SetDroneConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).SetDroneConfig(ctx, req.(*types.SetDroneConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_IsConfdRunning_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).IsConfdRunning(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/IsConfdRunning",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).IsConfdRunning(ctx, req.(*types.DroneEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_StartConfd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).StartConfd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/StartConfd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).StartConfd(ctx, req.(*types.DroneEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_StopConfd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).StopConfd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/StopConfd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).StopConfd(ctx, req.(*types.DroneEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_RegisterMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTask_RegisterMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).RegisterMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/RegisterMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).RegisterMetadata(ctx, req.(*types.SubTask_RegisterMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_DeregisterMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTask_DeregisterMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).DeregisterMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/DeregisterMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).DeregisterMetadata(ctx, req.(*types.SubTask_DeregisterMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_RegisterCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTask_RegisterCmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).RegisterCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/RegisterCmd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).RegisterCmd(ctx, req.(*types.SubTask_RegisterCmd))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_DeregisterCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTask_DeregisterCmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).DeregisterCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/DeregisterCmd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).DeregisterCmd(ctx, req.(*types.SubTask_DeregisterCmd))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_ReportSubTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTaskStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).ReportSubTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/ReportSubTaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).ReportSubTaskStatus(ctx, req.(*types.SubTaskStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_GetSubtaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTaskId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).GetSubtaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/GetSubtaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).GetSubtaskStatus(ctx, req.(*types.SubTaskId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_HandleSubtask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTaskMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).HandleSubtask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/HandleSubtask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).HandleSubtask(ctx, req.(*types.SubTaskMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_PingPilot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).PingPilot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/PingPilot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).PingPilot(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_PingFrontgate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).PingFrontgate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/PingFrontgate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).PingFrontgate(ctx, req.(*types.FrontgateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_PingDrone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).PingDrone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/PingDrone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).PingDrone(ctx, req.(*types.DroneEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_FrontgateChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PilotServiceServer).FrontgateChannel(&pilotServiceFrontgateChannelServer{stream})
}

type PilotService_FrontgateChannelServer interface {
	Send(*types.Bytes) error
	Recv() (*types.Bytes, error)
	grpc.ServerStream
}

type pilotServiceFrontgateChannelServer struct {
	grpc.ServerStream
}

func (x *pilotServiceFrontgateChannelServer) Send(m *types.Bytes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pilotServiceFrontgateChannelServer) Recv() (*types.Bytes, error) {
	m := new(types.Bytes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PilotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.pilot.PilotService",
	HandlerType: (*PilotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPilotConfig",
			Handler:    _PilotService_GetPilotConfig_Handler,
		},
		{
			MethodName: "GetFrontdateList",
			Handler:    _PilotService_GetFrontdateList_Handler,
		},
		{
			MethodName: "GetFrontgateConfig",
			Handler:    _PilotService_GetFrontgateConfig_Handler,
		},
		{
			MethodName: "SetFrontgateConfig",
			Handler:    _PilotService_SetFrontgateConfig_Handler,
		},
		{
			MethodName: "GetDroneConfig",
			Handler:    _PilotService_GetDroneConfig_Handler,
		},
		{
			MethodName: "SetDroneConfig",
			Handler:    _PilotService_SetDroneConfig_Handler,
		},
		{
			MethodName: "IsConfdRunning",
			Handler:    _PilotService_IsConfdRunning_Handler,
		},
		{
			MethodName: "StartConfd",
			Handler:    _PilotService_StartConfd_Handler,
		},
		{
			MethodName: "StopConfd",
			Handler:    _PilotService_StopConfd_Handler,
		},
		{
			MethodName: "RegisterMetadata",
			Handler:    _PilotService_RegisterMetadata_Handler,
		},
		{
			MethodName: "DeregisterMetadata",
			Handler:    _PilotService_DeregisterMetadata_Handler,
		},
		{
			MethodName: "RegisterCmd",
			Handler:    _PilotService_RegisterCmd_Handler,
		},
		{
			MethodName: "DeregisterCmd",
			Handler:    _PilotService_DeregisterCmd_Handler,
		},
		{
			MethodName: "ReportSubTaskStatus",
			Handler:    _PilotService_ReportSubTaskStatus_Handler,
		},
		{
			MethodName: "GetSubtaskStatus",
			Handler:    _PilotService_GetSubtaskStatus_Handler,
		},
		{
			MethodName: "HandleSubtask",
			Handler:    _PilotService_HandleSubtask_Handler,
		},
		{
			MethodName: "PingPilot",
			Handler:    _PilotService_PingPilot_Handler,
		},
		{
			MethodName: "PingFrontgate",
			Handler:    _PilotService_PingFrontgate_Handler,
		},
		{
			MethodName: "PingDrone",
			Handler:    _PilotService_PingDrone_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FrontgateChannel",
			Handler:       _PilotService_FrontgateChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "openpitrix/pilot/pilot.proto",
}

func init() { proto.RegisterFile("openpitrix/pilot/pilot.proto", fileDescriptor_pilot_49201b8ce8bb4a41) }

var fileDescriptor_pilot_49201b8ce8bb4a41 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xd5, 0x0b, 0x52, 0x4d, 0x13, 0x45, 0x46, 0x02, 0x69, 0xdb, 0x8a, 0x72, 0x40, 0x45,
	0xa8, 0xc9, 0x22, 0x38, 0x22, 0x2e, 0x49, 0xf3, 0x0f, 0x11, 0x08, 0x59, 0x04, 0x12, 0x1c, 0x90,
	0x53, 0x4f, 0x8d, 0x95, 0xc4, 0x36, 0xf6, 0x04, 0xc8, 0x2b, 0xf1, 0x94, 0x68, 0xbd, 0x1b, 0xb2,
	0xe9, 0xc6, 0x59, 0x44, 0x2f, 0x6b, 0xc9, 0xf3, 0x7d, 0xbf, 0x19, 0xdb, 0xeb, 0x31, 0x39, 0xd1,
	0x06, 0x94, 0x91, 0x68, 0xe5, 0xaf, 0xd8, 0xc8, 0xb9, 0xc6, 0xec, 0xdb, 0x32, 0x56, 0xa3, 0xa6,
	0x8d, 0x4d, 0xb4, 0xe5, 0xe7, 0xa3, 0x13, 0xa1, 0xb5, 0x98, 0x43, 0xcc, 0x8c, 0x8c, 0x99, 0x52,
	0x1a, 0x19, 0x4a, 0xad, 0x5c, 0xa6, 0x8f, 0x2e, 0xfc, 0x70, 0xd5, 0x14, 0xa0, 0x9a, 0xee, 0x27,
	0x13, 0x02, 0x6c, 0xac, 0x8d, 0x57, 0xec, 0x50, 0x17, 0x73, 0xe3, 0xca, 0x80, 0xcb, 0xbe, 0xc1,
	0x28, 0xb7, 0x5a, 0x41, 0x1e, 0x3d, 0x2b, 0x45, 0xaf, 0xad, 0x56, 0x28, 0x18, 0x42, 0xd0, 0x5f,
	0x58, 0x59, 0x74, 0x5c, 0xce, 0xcd, 0xdc, 0x2c, 0x0b, 0x3e, 0xff, 0x7d, 0x44, 0x8e, 0xc6, 0xa9,
	0x38, 0x01, 0xfb, 0x43, 0x5e, 0x01, 0x1d, 0x90, 0x7a, 0x1f, 0xd0, 0x4f, 0x75, 0xb4, 0xba, 0x96,
	0x82, 0x3e, 0x68, 0x15, 0xb6, 0x26, 0x2b, 0xbb, 0xbb, 0x30, 0xb8, 0x8a, 0x4e, 0xcb, 0x81, 0xa2,
	0xef, 0x2d, 0x69, 0xf4, 0x01, 0x7b, 0x69, 0xad, 0x9c, 0x21, 0xbc, 0x91, 0x0e, 0xc3, 0xac, 0x47,
	0xe5, 0x40, 0x6f, 0xbd, 0xca, 0x21, 0xf7, 0xde, 0x8f, 0x84, 0xae, 0x79, 0xe9, 0x6c, 0x9e, 0xe5,
	0x74, 0xaf, 0x71, 0x2f, 0x37, 0x27, 0x8c, 0x09, 0x4d, 0xca, 0xdc, 0x6a, 0x63, 0x14, 0x5a, 0x4c,
	0xbe, 0x87, 0x97, 0xe9, 0x19, 0xfe, 0xc7, 0x1e, 0x16, 0x7d, 0x13, 0x52, 0x4f, 0xb6, 0x49, 0xe7,
	0x65, 0xc3, 0xb6, 0x62, 0x02, 0xdf, 0x97, 0xe0, 0x30, 0x5c, 0xdd, 0x90, 0xd4, 0x87, 0x2e, 0xd5,
	0xf2, 0xc9, 0x52, 0x29, 0xa9, 0x04, 0x7d, 0x18, 0x28, 0xa2, 0xab, 0xb8, 0xd1, 0x52, 0x61, 0x74,
	0xbf, 0x2c, 0x68, 0x6b, 0x3d, 0xa7, 0x3d, 0x42, 0x12, 0x64, 0xd6, 0x9f, 0x38, 0xaf, 0xc6, 0x04,
	0x4b, 0xea, 0x92, 0xc3, 0x04, 0xb5, 0xb9, 0x2d, 0xe6, 0x13, 0x69, 0x4c, 0x40, 0x48, 0x87, 0x60,
	0x47, 0x80, 0x8c, 0x33, 0x64, 0xf4, 0xe9, 0x8e, 0xfd, 0x5a, 0x4e, 0x3f, 0x30, 0x37, 0xfb, 0x7a,
	0x53, 0x1b, 0x06, 0x7f, 0x21, 0xf4, 0x12, 0xec, 0x4d, 0xf4, 0x45, 0x18, 0x5d, 0x56, 0x87, 0xe1,
	0x23, 0x72, 0x77, 0x5d, 0x49, 0x67, 0xc1, 0xe9, 0xe3, 0xea, 0x82, 0x3b, 0x0b, 0x1e, 0xc6, 0xbd,
	0x27, 0xb5, 0x4d, 0xf6, 0x14, 0x78, 0xfe, 0x2f, 0x65, 0xee, 0x45, 0xbe, 0x23, 0xf7, 0x26, 0x60,
	0xb4, 0xc5, 0xdc, 0x97, 0x20, 0xc3, 0xa5, 0xdb, 0x75, 0x50, 0x5b, 0x82, 0x30, 0x70, 0xec, 0x5b,
	0x43, 0xb2, 0x9c, 0xe2, 0x86, 0x76, 0x1c, 0xa4, 0x0d, 0x79, 0x54, 0x95, 0x8a, 0xbe, 0x26, 0xb5,
	0x01, 0x53, 0x7c, 0x0e, 0x39, 0x94, 0x9e, 0x05, 0x1d, 0x23, 0x70, 0x8e, 0x09, 0x08, 0x57, 0xf7,
	0x8a, 0x1c, 0x8e, 0xa5, 0x12, 0xbe, 0x97, 0x85, 0x6f, 0x6e, 0xd0, 0xde, 0x27, 0xb5, 0xd4, 0xfe,
	0xb7, 0x5b, 0x54, 0xb5, 0xa8, 0x7d, 0xb7, 0x22, 0x05, 0xf9, 0x9f, 0xff, 0x16, 0xb7, 0x62, 0x40,
	0x1a, 0x9b, 0xce, 0xf5, 0x8d, 0x29, 0x05, 0xf3, 0x5d, 0xab, 0x6a, 0xaf, 0x10, 0x76, 0x1e, 0x99,
	0x0f, 0x3c, 0x39, 0x78, 0x76, 0xd0, 0x8e, 0x3f, 0x37, 0x0b, 0x51, 0xa9, 0xe3, 0xe2, 0x8b, 0x3a,
	0x13, 0xb1, 0x99, 0x66, 0x0f, 0xcf, 0x4b, 0x33, 0xf5, 0xe3, 0xf4, 0x8e, 0x7f, 0x64, 0x5e, 0xfc,
	0x09, 0x00, 0x00, 0xff, 0xff, 0xba, 0xe7, 0x49, 0x2e, 0x7b, 0x07, 0x00, 0x00,
}
