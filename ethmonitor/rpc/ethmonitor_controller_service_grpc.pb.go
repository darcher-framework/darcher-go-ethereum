// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EthmonitorControllerServiceClient is the client API for EthmonitorControllerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EthmonitorControllerServiceClient interface {
	NotifyTxReceived(ctx context.Context, in *TxReceivedMsg, opts ...grpc.CallOption) (*empty.Empty, error)
	NotifyTxFinished(ctx context.Context, in *TxFinishedMsg, opts ...grpc.CallOption) (*empty.Empty, error)
	NotifyTxTraverseStart(ctx context.Context, in *TxTraverseStartMsg, opts ...grpc.CallOption) (*empty.Empty, error)
	NotifyTxStateChangeMsg(ctx context.Context, in *TxStateChangeMsg, opts ...grpc.CallOption) (*empty.Empty, error)
	AskForNextState(ctx context.Context, in *TxStateControlMsg, opts ...grpc.CallOption) (*TxStateControlMsg, error)
	SelectTx(ctx context.Context, in *SelectTxControlMsg, opts ...grpc.CallOption) (*SelectTxControlMsg, error)
	// notify tx error to darcher analyzer
	NotifyTxError(ctx context.Context, in *TxErrorMsg, opts ...grpc.CallOption) (*empty.Empty, error)
	// notify contract vulnerability to darcher analyzer
	NotifyContractVulnerability(ctx context.Context, in *ContractVulReport, opts ...grpc.CallOption) (*empty.Empty, error)
}

type ethmonitorControllerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEthmonitorControllerServiceClient(cc grpc.ClientConnInterface) EthmonitorControllerServiceClient {
	return &ethmonitorControllerServiceClient{cc}
}

func (c *ethmonitorControllerServiceClient) NotifyTxReceived(ctx context.Context, in *TxReceivedMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/darcher.EthmonitorControllerService/notifyTxReceived", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethmonitorControllerServiceClient) NotifyTxFinished(ctx context.Context, in *TxFinishedMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/darcher.EthmonitorControllerService/notifyTxFinished", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethmonitorControllerServiceClient) NotifyTxTraverseStart(ctx context.Context, in *TxTraverseStartMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/darcher.EthmonitorControllerService/notifyTxTraverseStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethmonitorControllerServiceClient) NotifyTxStateChangeMsg(ctx context.Context, in *TxStateChangeMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/darcher.EthmonitorControllerService/notifyTxStateChangeMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethmonitorControllerServiceClient) AskForNextState(ctx context.Context, in *TxStateControlMsg, opts ...grpc.CallOption) (*TxStateControlMsg, error) {
	out := new(TxStateControlMsg)
	err := c.cc.Invoke(ctx, "/darcher.EthmonitorControllerService/askForNextState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethmonitorControllerServiceClient) SelectTx(ctx context.Context, in *SelectTxControlMsg, opts ...grpc.CallOption) (*SelectTxControlMsg, error) {
	out := new(SelectTxControlMsg)
	err := c.cc.Invoke(ctx, "/darcher.EthmonitorControllerService/selectTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethmonitorControllerServiceClient) NotifyTxError(ctx context.Context, in *TxErrorMsg, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/darcher.EthmonitorControllerService/notifyTxError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ethmonitorControllerServiceClient) NotifyContractVulnerability(ctx context.Context, in *ContractVulReport, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/darcher.EthmonitorControllerService/notifyContractVulnerability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EthmonitorControllerServiceServer is the server API for EthmonitorControllerService service.
// All implementations must embed UnimplementedEthmonitorControllerServiceServer
// for forward compatibility
type EthmonitorControllerServiceServer interface {
	NotifyTxReceived(context.Context, *TxReceivedMsg) (*empty.Empty, error)
	NotifyTxFinished(context.Context, *TxFinishedMsg) (*empty.Empty, error)
	NotifyTxTraverseStart(context.Context, *TxTraverseStartMsg) (*empty.Empty, error)
	NotifyTxStateChangeMsg(context.Context, *TxStateChangeMsg) (*empty.Empty, error)
	AskForNextState(context.Context, *TxStateControlMsg) (*TxStateControlMsg, error)
	SelectTx(context.Context, *SelectTxControlMsg) (*SelectTxControlMsg, error)
	// notify tx error to darcher analyzer
	NotifyTxError(context.Context, *TxErrorMsg) (*empty.Empty, error)
	// notify contract vulnerability to darcher analyzer
	NotifyContractVulnerability(context.Context, *ContractVulReport) (*empty.Empty, error)
	mustEmbedUnimplementedEthmonitorControllerServiceServer()
}

// UnimplementedEthmonitorControllerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEthmonitorControllerServiceServer struct {
}

func (*UnimplementedEthmonitorControllerServiceServer) NotifyTxReceived(context.Context, *TxReceivedMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyTxReceived not implemented")
}
func (*UnimplementedEthmonitorControllerServiceServer) NotifyTxFinished(context.Context, *TxFinishedMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyTxFinished not implemented")
}
func (*UnimplementedEthmonitorControllerServiceServer) NotifyTxTraverseStart(context.Context, *TxTraverseStartMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyTxTraverseStart not implemented")
}
func (*UnimplementedEthmonitorControllerServiceServer) NotifyTxStateChangeMsg(context.Context, *TxStateChangeMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyTxStateChangeMsg not implemented")
}
func (*UnimplementedEthmonitorControllerServiceServer) AskForNextState(context.Context, *TxStateControlMsg) (*TxStateControlMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskForNextState not implemented")
}
func (*UnimplementedEthmonitorControllerServiceServer) SelectTx(context.Context, *SelectTxControlMsg) (*SelectTxControlMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectTx not implemented")
}
func (*UnimplementedEthmonitorControllerServiceServer) NotifyTxError(context.Context, *TxErrorMsg) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyTxError not implemented")
}
func (*UnimplementedEthmonitorControllerServiceServer) NotifyContractVulnerability(context.Context, *ContractVulReport) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyContractVulnerability not implemented")
}
func (*UnimplementedEthmonitorControllerServiceServer) mustEmbedUnimplementedEthmonitorControllerServiceServer() {
}

func RegisterEthmonitorControllerServiceServer(s *grpc.Server, srv EthmonitorControllerServiceServer) {
	s.RegisterService(&_EthmonitorControllerService_serviceDesc, srv)
}

func _EthmonitorControllerService_NotifyTxReceived_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxReceivedMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthmonitorControllerServiceServer).NotifyTxReceived(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/darcher.EthmonitorControllerService/NotifyTxReceived",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthmonitorControllerServiceServer).NotifyTxReceived(ctx, req.(*TxReceivedMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthmonitorControllerService_NotifyTxFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxFinishedMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthmonitorControllerServiceServer).NotifyTxFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/darcher.EthmonitorControllerService/NotifyTxFinished",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthmonitorControllerServiceServer).NotifyTxFinished(ctx, req.(*TxFinishedMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthmonitorControllerService_NotifyTxTraverseStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxTraverseStartMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthmonitorControllerServiceServer).NotifyTxTraverseStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/darcher.EthmonitorControllerService/NotifyTxTraverseStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthmonitorControllerServiceServer).NotifyTxTraverseStart(ctx, req.(*TxTraverseStartMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthmonitorControllerService_NotifyTxStateChangeMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxStateChangeMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthmonitorControllerServiceServer).NotifyTxStateChangeMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/darcher.EthmonitorControllerService/NotifyTxStateChangeMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthmonitorControllerServiceServer).NotifyTxStateChangeMsg(ctx, req.(*TxStateChangeMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthmonitorControllerService_AskForNextState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxStateControlMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthmonitorControllerServiceServer).AskForNextState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/darcher.EthmonitorControllerService/AskForNextState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthmonitorControllerServiceServer).AskForNextState(ctx, req.(*TxStateControlMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthmonitorControllerService_SelectTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectTxControlMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthmonitorControllerServiceServer).SelectTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/darcher.EthmonitorControllerService/SelectTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthmonitorControllerServiceServer).SelectTx(ctx, req.(*SelectTxControlMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthmonitorControllerService_NotifyTxError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxErrorMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthmonitorControllerServiceServer).NotifyTxError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/darcher.EthmonitorControllerService/NotifyTxError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthmonitorControllerServiceServer).NotifyTxError(ctx, req.(*TxErrorMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _EthmonitorControllerService_NotifyContractVulnerability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContractVulReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EthmonitorControllerServiceServer).NotifyContractVulnerability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/darcher.EthmonitorControllerService/NotifyContractVulnerability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EthmonitorControllerServiceServer).NotifyContractVulnerability(ctx, req.(*ContractVulReport))
	}
	return interceptor(ctx, in, info, handler)
}

var _EthmonitorControllerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "darcher.EthmonitorControllerService",
	HandlerType: (*EthmonitorControllerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "notifyTxReceived",
			Handler:    _EthmonitorControllerService_NotifyTxReceived_Handler,
		},
		{
			MethodName: "notifyTxFinished",
			Handler:    _EthmonitorControllerService_NotifyTxFinished_Handler,
		},
		{
			MethodName: "notifyTxTraverseStart",
			Handler:    _EthmonitorControllerService_NotifyTxTraverseStart_Handler,
		},
		{
			MethodName: "notifyTxStateChangeMsg",
			Handler:    _EthmonitorControllerService_NotifyTxStateChangeMsg_Handler,
		},
		{
			MethodName: "askForNextState",
			Handler:    _EthmonitorControllerService_AskForNextState_Handler,
		},
		{
			MethodName: "selectTx",
			Handler:    _EthmonitorControllerService_SelectTx_Handler,
		},
		{
			MethodName: "notifyTxError",
			Handler:    _EthmonitorControllerService_NotifyTxError_Handler,
		},
		{
			MethodName: "notifyContractVulnerability",
			Handler:    _EthmonitorControllerService_NotifyContractVulnerability_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ethmonitor_controller_service.proto",
}