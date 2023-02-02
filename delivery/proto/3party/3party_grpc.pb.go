// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: delivery/proto/3party/3party.proto

package openapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ThirdPartyServiceClient is the client API for ThirdPartyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThirdPartyServiceClient interface {
	Testing(ctx context.Context, in *TestingRequest, opts ...grpc.CallOption) (*TestingResponse, error)
	//SIKP
	VerificationSIKP(ctx context.Context, in *VerificationSIKPRequest, opts ...grpc.CallOption) (*VerificationSIKPReponse, error)
	CheckPlafondSIKP(ctx context.Context, in *CheckPlafondSIKPRequest, opts ...grpc.CallOption) (*CheckPlafondSIKPReponse, error)
	//====================== Region Jamkrindo ==============================//
	SubmitJamkrindoCalon(ctx context.Context, in *SubmitJamkrindoCalonRequest, opts ...grpc.CallOption) (*SubmitJamkrindoCalonResponse, error)
	JamkrindoKlaim(ctx context.Context, in *JamkrindoKlaimRequest, opts ...grpc.CallOption) (*JamkrindoKlaimResponse, error)
}

type thirdPartyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThirdPartyServiceClient(cc grpc.ClientConnInterface) ThirdPartyServiceClient {
	return &thirdPartyServiceClient{cc}
}

func (c *thirdPartyServiceClient) Testing(ctx context.Context, in *TestingRequest, opts ...grpc.CallOption) (*TestingResponse, error) {
	out := new(TestingResponse)
	err := c.cc.Invoke(ctx, "/proto.ThirdPartyService/Testing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyServiceClient) VerificationSIKP(ctx context.Context, in *VerificationSIKPRequest, opts ...grpc.CallOption) (*VerificationSIKPReponse, error) {
	out := new(VerificationSIKPReponse)
	err := c.cc.Invoke(ctx, "/proto.ThirdPartyService/VerificationSIKP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyServiceClient) CheckPlafondSIKP(ctx context.Context, in *CheckPlafondSIKPRequest, opts ...grpc.CallOption) (*CheckPlafondSIKPReponse, error) {
	out := new(CheckPlafondSIKPReponse)
	err := c.cc.Invoke(ctx, "/proto.ThirdPartyService/CheckPlafondSIKP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyServiceClient) SubmitJamkrindoCalon(ctx context.Context, in *SubmitJamkrindoCalonRequest, opts ...grpc.CallOption) (*SubmitJamkrindoCalonResponse, error) {
	out := new(SubmitJamkrindoCalonResponse)
	err := c.cc.Invoke(ctx, "/proto.ThirdPartyService/SubmitJamkrindoCalon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thirdPartyServiceClient) JamkrindoKlaim(ctx context.Context, in *JamkrindoKlaimRequest, opts ...grpc.CallOption) (*JamkrindoKlaimResponse, error) {
	out := new(JamkrindoKlaimResponse)
	err := c.cc.Invoke(ctx, "/proto.ThirdPartyService/JamkrindoKlaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThirdPartyServiceServer is the server API for ThirdPartyService service.
// All implementations must embed UnimplementedThirdPartyServiceServer
// for forward compatibility
type ThirdPartyServiceServer interface {
	Testing(context.Context, *TestingRequest) (*TestingResponse, error)
	//SIKP
	VerificationSIKP(context.Context, *VerificationSIKPRequest) (*VerificationSIKPReponse, error)
	CheckPlafondSIKP(context.Context, *CheckPlafondSIKPRequest) (*CheckPlafondSIKPReponse, error)
	//====================== Region Jamkrindo ==============================//
	SubmitJamkrindoCalon(context.Context, *SubmitJamkrindoCalonRequest) (*SubmitJamkrindoCalonResponse, error)
	JamkrindoKlaim(context.Context, *JamkrindoKlaimRequest) (*JamkrindoKlaimResponse, error)
	mustEmbedUnimplementedThirdPartyServiceServer()
}

// UnimplementedThirdPartyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedThirdPartyServiceServer struct {
}

func (UnimplementedThirdPartyServiceServer) Testing(context.Context, *TestingRequest) (*TestingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Testing not implemented")
}
func (UnimplementedThirdPartyServiceServer) VerificationSIKP(context.Context, *VerificationSIKPRequest) (*VerificationSIKPReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerificationSIKP not implemented")
}
func (UnimplementedThirdPartyServiceServer) CheckPlafondSIKP(context.Context, *CheckPlafondSIKPRequest) (*CheckPlafondSIKPReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPlafondSIKP not implemented")
}
func (UnimplementedThirdPartyServiceServer) SubmitJamkrindoCalon(context.Context, *SubmitJamkrindoCalonRequest) (*SubmitJamkrindoCalonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitJamkrindoCalon not implemented")
}
func (UnimplementedThirdPartyServiceServer) JamkrindoKlaim(context.Context, *JamkrindoKlaimRequest) (*JamkrindoKlaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JamkrindoKlaim not implemented")
}
func (UnimplementedThirdPartyServiceServer) mustEmbedUnimplementedThirdPartyServiceServer() {}

// UnsafeThirdPartyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThirdPartyServiceServer will
// result in compilation errors.
type UnsafeThirdPartyServiceServer interface {
	mustEmbedUnimplementedThirdPartyServiceServer()
}

func RegisterThirdPartyServiceServer(s grpc.ServiceRegistrar, srv ThirdPartyServiceServer) {
	s.RegisterService(&ThirdPartyService_ServiceDesc, srv)
}

func _ThirdPartyService_Testing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyServiceServer).Testing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThirdPartyService/Testing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyServiceServer).Testing(ctx, req.(*TestingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyService_VerificationSIKP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationSIKPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyServiceServer).VerificationSIKP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThirdPartyService/VerificationSIKP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyServiceServer).VerificationSIKP(ctx, req.(*VerificationSIKPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyService_CheckPlafondSIKP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPlafondSIKPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyServiceServer).CheckPlafondSIKP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThirdPartyService/CheckPlafondSIKP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyServiceServer).CheckPlafondSIKP(ctx, req.(*CheckPlafondSIKPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyService_SubmitJamkrindoCalon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitJamkrindoCalonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyServiceServer).SubmitJamkrindoCalon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThirdPartyService/SubmitJamkrindoCalon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyServiceServer).SubmitJamkrindoCalon(ctx, req.(*SubmitJamkrindoCalonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThirdPartyService_JamkrindoKlaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JamkrindoKlaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThirdPartyServiceServer).JamkrindoKlaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ThirdPartyService/JamkrindoKlaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThirdPartyServiceServer).JamkrindoKlaim(ctx, req.(*JamkrindoKlaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThirdPartyService_ServiceDesc is the grpc.ServiceDesc for ThirdPartyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThirdPartyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ThirdPartyService",
	HandlerType: (*ThirdPartyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Testing",
			Handler:    _ThirdPartyService_Testing_Handler,
		},
		{
			MethodName: "VerificationSIKP",
			Handler:    _ThirdPartyService_VerificationSIKP_Handler,
		},
		{
			MethodName: "CheckPlafondSIKP",
			Handler:    _ThirdPartyService_CheckPlafondSIKP_Handler,
		},
		{
			MethodName: "SubmitJamkrindoCalon",
			Handler:    _ThirdPartyService_SubmitJamkrindoCalon_Handler,
		},
		{
			MethodName: "JamkrindoKlaim",
			Handler:    _ThirdPartyService_JamkrindoKlaim_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "delivery/proto/3party/3party.proto",
}
