// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: doctor.proto

package doctor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	DoctorService_CreatePrescription_FullMethodName    = "/doctor.DoctorService/CreatePrescription"
	DoctorService_ReviewPrescription_FullMethodName    = "/doctor.DoctorService/ReviewPrescription"
	DoctorService_GetPrescriptionList_FullMethodName   = "/doctor.DoctorService/GetPrescriptionList"
	DoctorService_GetPrescriptionDetail_FullMethodName = "/doctor.DoctorService/GetPrescriptionDetail"
	DoctorService_SelectMedicines_FullMethodName       = "/doctor.DoctorService/SelectMedicines"
	DoctorService_GetMedicinesList_FullMethodName      = "/doctor.DoctorService/GetMedicinesList"
	DoctorService_GetMedicinesDetail_FullMethodName    = "/doctor.DoctorService/GetMedicinesDetail"
	DoctorService_CreateCareHistory_FullMethodName     = "/doctor.DoctorService/CreateCareHistory"
	DoctorService_GetCareHistoryList_FullMethodName    = "/doctor.DoctorService/GetCareHistoryList"
	DoctorService_GetCareHistoryDetail_FullMethodName  = "/doctor.DoctorService/GetCareHistoryDetail"
	DoctorService_GetWeeklyRanking_FullMethodName      = "/doctor.DoctorService/GetWeeklyRanking"
	DoctorService_GenerateWeeklyRanking_FullMethodName = "/doctor.DoctorService/GenerateWeeklyRanking"
	DoctorService_GetDoctorPerformance_FullMethodName  = "/doctor.DoctorService/GetDoctorPerformance"
)

// DoctorServiceClient is the client API for DoctorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 医生服务
type DoctorServiceClient interface {
	// 开处方
	CreatePrescription(ctx context.Context, in *CreatePrescriptionReq, opts ...grpc.CallOption) (*CreatePrescriptionResp, error)
	// 审核处方
	ReviewPrescription(ctx context.Context, in *ReviewPrescriptionReq, opts ...grpc.CallOption) (*ReviewPrescriptionResp, error)
	// 获取处方列表
	GetPrescriptionList(ctx context.Context, in *GetPrescriptionListReq, opts ...grpc.CallOption) (*GetPrescriptionListResp, error)
	// 获取处方详情
	GetPrescriptionDetail(ctx context.Context, in *GetPrescriptionDetailReq, opts ...grpc.CallOption) (*GetPrescriptionDetailResp, error)
	// 选择药品
	SelectMedicines(ctx context.Context, in *SelectMedicinesReq, opts ...grpc.CallOption) (*SelectMedicinesResp, error)
	// 获取药品列表
	GetMedicinesList(ctx context.Context, in *GetMedicinesListReq, opts ...grpc.CallOption) (*GetMedicinesListResp, error)
	// 获取药品详情
	GetMedicinesDetail(ctx context.Context, in *GetMedicinesDetailReq, opts ...grpc.CallOption) (*GetMedicinesDetailResp, error)
	// 创建病例
	CreateCareHistory(ctx context.Context, in *CreateCareHistoryReq, opts ...grpc.CallOption) (*CreateCareHistoryResp, error)
	// 获取病例列表
	GetCareHistoryList(ctx context.Context, in *GetCareHistoryListReq, opts ...grpc.CallOption) (*GetCareHistoryListResp, error)
	// 获取病例详情
	GetCareHistoryDetail(ctx context.Context, in *GetCareHistoryDetailReq, opts ...grpc.CallOption) (*GetCareHistoryDetailResp, error)
	// 获取周排行榜
	GetWeeklyRanking(ctx context.Context, in *GetWeeklyRankingReq, opts ...grpc.CallOption) (*GetWeeklyRankingResp, error)
	// 生成周排行榜
	GenerateWeeklyRanking(ctx context.Context, in *GenerateWeeklyRankingReq, opts ...grpc.CallOption) (*GenerateWeeklyRankingResp, error)
	// 获取医生业绩
	GetDoctorPerformance(ctx context.Context, in *GetDoctorPerformanceReq, opts ...grpc.CallOption) (*GetDoctorPerformanceResp, error)
}

type doctorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorServiceClient(cc grpc.ClientConnInterface) DoctorServiceClient {
	return &doctorServiceClient{cc}
}

func (c *doctorServiceClient) CreatePrescription(ctx context.Context, in *CreatePrescriptionReq, opts ...grpc.CallOption) (*CreatePrescriptionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePrescriptionResp)
	err := c.cc.Invoke(ctx, DoctorService_CreatePrescription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) ReviewPrescription(ctx context.Context, in *ReviewPrescriptionReq, opts ...grpc.CallOption) (*ReviewPrescriptionResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReviewPrescriptionResp)
	err := c.cc.Invoke(ctx, DoctorService_ReviewPrescription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetPrescriptionList(ctx context.Context, in *GetPrescriptionListReq, opts ...grpc.CallOption) (*GetPrescriptionListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPrescriptionListResp)
	err := c.cc.Invoke(ctx, DoctorService_GetPrescriptionList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetPrescriptionDetail(ctx context.Context, in *GetPrescriptionDetailReq, opts ...grpc.CallOption) (*GetPrescriptionDetailResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPrescriptionDetailResp)
	err := c.cc.Invoke(ctx, DoctorService_GetPrescriptionDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) SelectMedicines(ctx context.Context, in *SelectMedicinesReq, opts ...grpc.CallOption) (*SelectMedicinesResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SelectMedicinesResp)
	err := c.cc.Invoke(ctx, DoctorService_SelectMedicines_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetMedicinesList(ctx context.Context, in *GetMedicinesListReq, opts ...grpc.CallOption) (*GetMedicinesListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMedicinesListResp)
	err := c.cc.Invoke(ctx, DoctorService_GetMedicinesList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetMedicinesDetail(ctx context.Context, in *GetMedicinesDetailReq, opts ...grpc.CallOption) (*GetMedicinesDetailResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMedicinesDetailResp)
	err := c.cc.Invoke(ctx, DoctorService_GetMedicinesDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) CreateCareHistory(ctx context.Context, in *CreateCareHistoryReq, opts ...grpc.CallOption) (*CreateCareHistoryResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCareHistoryResp)
	err := c.cc.Invoke(ctx, DoctorService_CreateCareHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetCareHistoryList(ctx context.Context, in *GetCareHistoryListReq, opts ...grpc.CallOption) (*GetCareHistoryListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCareHistoryListResp)
	err := c.cc.Invoke(ctx, DoctorService_GetCareHistoryList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetCareHistoryDetail(ctx context.Context, in *GetCareHistoryDetailReq, opts ...grpc.CallOption) (*GetCareHistoryDetailResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCareHistoryDetailResp)
	err := c.cc.Invoke(ctx, DoctorService_GetCareHistoryDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetWeeklyRanking(ctx context.Context, in *GetWeeklyRankingReq, opts ...grpc.CallOption) (*GetWeeklyRankingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWeeklyRankingResp)
	err := c.cc.Invoke(ctx, DoctorService_GetWeeklyRanking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GenerateWeeklyRanking(ctx context.Context, in *GenerateWeeklyRankingReq, opts ...grpc.CallOption) (*GenerateWeeklyRankingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateWeeklyRankingResp)
	err := c.cc.Invoke(ctx, DoctorService_GenerateWeeklyRanking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetDoctorPerformance(ctx context.Context, in *GetDoctorPerformanceReq, opts ...grpc.CallOption) (*GetDoctorPerformanceResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDoctorPerformanceResp)
	err := c.cc.Invoke(ctx, DoctorService_GetDoctorPerformance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorServiceServer is the server API for DoctorService service.
// All implementations must embed UnimplementedDoctorServiceServer
// for forward compatibility
//
// 医生服务
type DoctorServiceServer interface {
	// 开处方
	CreatePrescription(context.Context, *CreatePrescriptionReq) (*CreatePrescriptionResp, error)
	// 审核处方
	ReviewPrescription(context.Context, *ReviewPrescriptionReq) (*ReviewPrescriptionResp, error)
	// 获取处方列表
	GetPrescriptionList(context.Context, *GetPrescriptionListReq) (*GetPrescriptionListResp, error)
	// 获取处方详情
	GetPrescriptionDetail(context.Context, *GetPrescriptionDetailReq) (*GetPrescriptionDetailResp, error)
	// 选择药品
	SelectMedicines(context.Context, *SelectMedicinesReq) (*SelectMedicinesResp, error)
	// 获取药品列表
	GetMedicinesList(context.Context, *GetMedicinesListReq) (*GetMedicinesListResp, error)
	// 获取药品详情
	GetMedicinesDetail(context.Context, *GetMedicinesDetailReq) (*GetMedicinesDetailResp, error)
	// 创建病例
	CreateCareHistory(context.Context, *CreateCareHistoryReq) (*CreateCareHistoryResp, error)
	// 获取病例列表
	GetCareHistoryList(context.Context, *GetCareHistoryListReq) (*GetCareHistoryListResp, error)
	// 获取病例详情
	GetCareHistoryDetail(context.Context, *GetCareHistoryDetailReq) (*GetCareHistoryDetailResp, error)
	// 获取周排行榜
	GetWeeklyRanking(context.Context, *GetWeeklyRankingReq) (*GetWeeklyRankingResp, error)
	// 生成周排行榜
	GenerateWeeklyRanking(context.Context, *GenerateWeeklyRankingReq) (*GenerateWeeklyRankingResp, error)
	// 获取医生业绩
	GetDoctorPerformance(context.Context, *GetDoctorPerformanceReq) (*GetDoctorPerformanceResp, error)
	mustEmbedUnimplementedDoctorServiceServer()
}

// UnimplementedDoctorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDoctorServiceServer struct {
}

func (UnimplementedDoctorServiceServer) CreatePrescription(context.Context, *CreatePrescriptionReq) (*CreatePrescriptionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrescription not implemented")
}
func (UnimplementedDoctorServiceServer) ReviewPrescription(context.Context, *ReviewPrescriptionReq) (*ReviewPrescriptionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewPrescription not implemented")
}
func (UnimplementedDoctorServiceServer) GetPrescriptionList(context.Context, *GetPrescriptionListReq) (*GetPrescriptionListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrescriptionList not implemented")
}
func (UnimplementedDoctorServiceServer) GetPrescriptionDetail(context.Context, *GetPrescriptionDetailReq) (*GetPrescriptionDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrescriptionDetail not implemented")
}
func (UnimplementedDoctorServiceServer) SelectMedicines(context.Context, *SelectMedicinesReq) (*SelectMedicinesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectMedicines not implemented")
}
func (UnimplementedDoctorServiceServer) GetMedicinesList(context.Context, *GetMedicinesListReq) (*GetMedicinesListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedicinesList not implemented")
}
func (UnimplementedDoctorServiceServer) GetMedicinesDetail(context.Context, *GetMedicinesDetailReq) (*GetMedicinesDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedicinesDetail not implemented")
}
func (UnimplementedDoctorServiceServer) CreateCareHistory(context.Context, *CreateCareHistoryReq) (*CreateCareHistoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCareHistory not implemented")
}
func (UnimplementedDoctorServiceServer) GetCareHistoryList(context.Context, *GetCareHistoryListReq) (*GetCareHistoryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCareHistoryList not implemented")
}
func (UnimplementedDoctorServiceServer) GetCareHistoryDetail(context.Context, *GetCareHistoryDetailReq) (*GetCareHistoryDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCareHistoryDetail not implemented")
}
func (UnimplementedDoctorServiceServer) GetWeeklyRanking(context.Context, *GetWeeklyRankingReq) (*GetWeeklyRankingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeeklyRanking not implemented")
}
func (UnimplementedDoctorServiceServer) GenerateWeeklyRanking(context.Context, *GenerateWeeklyRankingReq) (*GenerateWeeklyRankingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateWeeklyRanking not implemented")
}
func (UnimplementedDoctorServiceServer) GetDoctorPerformance(context.Context, *GetDoctorPerformanceReq) (*GetDoctorPerformanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctorPerformance not implemented")
}
func (UnimplementedDoctorServiceServer) mustEmbedUnimplementedDoctorServiceServer() {}

// UnsafeDoctorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorServiceServer will
// result in compilation errors.
type UnsafeDoctorServiceServer interface {
	mustEmbedUnimplementedDoctorServiceServer()
}

func RegisterDoctorServiceServer(s grpc.ServiceRegistrar, srv DoctorServiceServer) {
	s.RegisterService(&DoctorService_ServiceDesc, srv)
}

func _DoctorService_CreatePrescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrescriptionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).CreatePrescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_CreatePrescription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).CreatePrescription(ctx, req.(*CreatePrescriptionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_ReviewPrescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewPrescriptionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).ReviewPrescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_ReviewPrescription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).ReviewPrescription(ctx, req.(*ReviewPrescriptionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetPrescriptionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrescriptionListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetPrescriptionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetPrescriptionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetPrescriptionList(ctx, req.(*GetPrescriptionListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetPrescriptionDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrescriptionDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetPrescriptionDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetPrescriptionDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetPrescriptionDetail(ctx, req.(*GetPrescriptionDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_SelectMedicines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectMedicinesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).SelectMedicines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_SelectMedicines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).SelectMedicines(ctx, req.(*SelectMedicinesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetMedicinesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedicinesListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetMedicinesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetMedicinesList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetMedicinesList(ctx, req.(*GetMedicinesListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetMedicinesDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedicinesDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetMedicinesDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetMedicinesDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetMedicinesDetail(ctx, req.(*GetMedicinesDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_CreateCareHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCareHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).CreateCareHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_CreateCareHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).CreateCareHistory(ctx, req.(*CreateCareHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetCareHistoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCareHistoryListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetCareHistoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetCareHistoryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetCareHistoryList(ctx, req.(*GetCareHistoryListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetCareHistoryDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCareHistoryDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetCareHistoryDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetCareHistoryDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetCareHistoryDetail(ctx, req.(*GetCareHistoryDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetWeeklyRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeeklyRankingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetWeeklyRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetWeeklyRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetWeeklyRanking(ctx, req.(*GetWeeklyRankingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GenerateWeeklyRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateWeeklyRankingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GenerateWeeklyRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GenerateWeeklyRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GenerateWeeklyRanking(ctx, req.(*GenerateWeeklyRankingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetDoctorPerformance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDoctorPerformanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetDoctorPerformance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetDoctorPerformance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetDoctorPerformance(ctx, req.(*GetDoctorPerformanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DoctorService_ServiceDesc is the grpc.ServiceDesc for DoctorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoctorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doctor.DoctorService",
	HandlerType: (*DoctorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePrescription",
			Handler:    _DoctorService_CreatePrescription_Handler,
		},
		{
			MethodName: "ReviewPrescription",
			Handler:    _DoctorService_ReviewPrescription_Handler,
		},
		{
			MethodName: "GetPrescriptionList",
			Handler:    _DoctorService_GetPrescriptionList_Handler,
		},
		{
			MethodName: "GetPrescriptionDetail",
			Handler:    _DoctorService_GetPrescriptionDetail_Handler,
		},
		{
			MethodName: "SelectMedicines",
			Handler:    _DoctorService_SelectMedicines_Handler,
		},
		{
			MethodName: "GetMedicinesList",
			Handler:    _DoctorService_GetMedicinesList_Handler,
		},
		{
			MethodName: "GetMedicinesDetail",
			Handler:    _DoctorService_GetMedicinesDetail_Handler,
		},
		{
			MethodName: "CreateCareHistory",
			Handler:    _DoctorService_CreateCareHistory_Handler,
		},
		{
			MethodName: "GetCareHistoryList",
			Handler:    _DoctorService_GetCareHistoryList_Handler,
		},
		{
			MethodName: "GetCareHistoryDetail",
			Handler:    _DoctorService_GetCareHistoryDetail_Handler,
		},
		{
			MethodName: "GetWeeklyRanking",
			Handler:    _DoctorService_GetWeeklyRanking_Handler,
		},
		{
			MethodName: "GenerateWeeklyRanking",
			Handler:    _DoctorService_GenerateWeeklyRanking_Handler,
		},
		{
			MethodName: "GetDoctorPerformance",
			Handler:    _DoctorService_GetDoctorPerformance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doctor.proto",
}
