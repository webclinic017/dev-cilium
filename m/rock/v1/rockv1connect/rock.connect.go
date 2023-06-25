// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: rock/v1/rock.proto

package rockv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/defn/dev/m/defn/dev/rock/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// RockStoreServiceName is the fully-qualified name of the RockStoreService service.
	RockStoreServiceName = "rock.v1.RockStoreService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RockStoreServiceGetRockProcedure is the fully-qualified name of the RockStoreService's GetRock
	// RPC.
	RockStoreServiceGetRockProcedure = "/rock.v1.RockStoreService/GetRock"
	// RockStoreServicePutRockProcedure is the fully-qualified name of the RockStoreService's PutRock
	// RPC.
	RockStoreServicePutRockProcedure = "/rock.v1.RockStoreService/PutRock"
	// RockStoreServiceDeleteRockProcedure is the fully-qualified name of the RockStoreService's
	// DeleteRock RPC.
	RockStoreServiceDeleteRockProcedure = "/rock.v1.RockStoreService/DeleteRock"
)

// RockStoreServiceClient is a client for the rock.v1.RockStoreService service.
type RockStoreServiceClient interface {
	GetRock(context.Context, *connect_go.Request[v1.GetRockRequest]) (*connect_go.Response[v1.GetRockResponse], error)
	PutRock(context.Context, *connect_go.Request[v1.PutRockRequest]) (*connect_go.Response[v1.PutRockResponse], error)
	DeleteRock(context.Context, *connect_go.Request[v1.DeleteRockRequest]) (*connect_go.Response[v1.DeleteRockResponse], error)
}

// NewRockStoreServiceClient constructs a client for the rock.v1.RockStoreService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRockStoreServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) RockStoreServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &rockStoreServiceClient{
		getRock: connect_go.NewClient[v1.GetRockRequest, v1.GetRockResponse](
			httpClient,
			baseURL+RockStoreServiceGetRockProcedure,
			opts...,
		),
		putRock: connect_go.NewClient[v1.PutRockRequest, v1.PutRockResponse](
			httpClient,
			baseURL+RockStoreServicePutRockProcedure,
			opts...,
		),
		deleteRock: connect_go.NewClient[v1.DeleteRockRequest, v1.DeleteRockResponse](
			httpClient,
			baseURL+RockStoreServiceDeleteRockProcedure,
			opts...,
		),
	}
}

// rockStoreServiceClient implements RockStoreServiceClient.
type rockStoreServiceClient struct {
	getRock    *connect_go.Client[v1.GetRockRequest, v1.GetRockResponse]
	putRock    *connect_go.Client[v1.PutRockRequest, v1.PutRockResponse]
	deleteRock *connect_go.Client[v1.DeleteRockRequest, v1.DeleteRockResponse]
}

// GetRock calls rock.v1.RockStoreService.GetRock.
func (c *rockStoreServiceClient) GetRock(ctx context.Context, req *connect_go.Request[v1.GetRockRequest]) (*connect_go.Response[v1.GetRockResponse], error) {
	return c.getRock.CallUnary(ctx, req)
}

// PutRock calls rock.v1.RockStoreService.PutRock.
func (c *rockStoreServiceClient) PutRock(ctx context.Context, req *connect_go.Request[v1.PutRockRequest]) (*connect_go.Response[v1.PutRockResponse], error) {
	return c.putRock.CallUnary(ctx, req)
}

// DeleteRock calls rock.v1.RockStoreService.DeleteRock.
func (c *rockStoreServiceClient) DeleteRock(ctx context.Context, req *connect_go.Request[v1.DeleteRockRequest]) (*connect_go.Response[v1.DeleteRockResponse], error) {
	return c.deleteRock.CallUnary(ctx, req)
}

// RockStoreServiceHandler is an implementation of the rock.v1.RockStoreService service.
type RockStoreServiceHandler interface {
	GetRock(context.Context, *connect_go.Request[v1.GetRockRequest]) (*connect_go.Response[v1.GetRockResponse], error)
	PutRock(context.Context, *connect_go.Request[v1.PutRockRequest]) (*connect_go.Response[v1.PutRockResponse], error)
	DeleteRock(context.Context, *connect_go.Request[v1.DeleteRockRequest]) (*connect_go.Response[v1.DeleteRockResponse], error)
}

// NewRockStoreServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRockStoreServiceHandler(svc RockStoreServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(RockStoreServiceGetRockProcedure, connect_go.NewUnaryHandler(
		RockStoreServiceGetRockProcedure,
		svc.GetRock,
		opts...,
	))
	mux.Handle(RockStoreServicePutRockProcedure, connect_go.NewUnaryHandler(
		RockStoreServicePutRockProcedure,
		svc.PutRock,
		opts...,
	))
	mux.Handle(RockStoreServiceDeleteRockProcedure, connect_go.NewUnaryHandler(
		RockStoreServiceDeleteRockProcedure,
		svc.DeleteRock,
		opts...,
	))
	return "/rock.v1.RockStoreService/", mux
}

// UnimplementedRockStoreServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRockStoreServiceHandler struct{}

func (UnimplementedRockStoreServiceHandler) GetRock(context.Context, *connect_go.Request[v1.GetRockRequest]) (*connect_go.Response[v1.GetRockResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("rock.v1.RockStoreService.GetRock is not implemented"))
}

func (UnimplementedRockStoreServiceHandler) PutRock(context.Context, *connect_go.Request[v1.PutRockRequest]) (*connect_go.Response[v1.PutRockResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("rock.v1.RockStoreService.PutRock is not implemented"))
}

func (UnimplementedRockStoreServiceHandler) DeleteRock(context.Context, *connect_go.Request[v1.DeleteRockRequest]) (*connect_go.Response[v1.DeleteRockResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("rock.v1.RockStoreService.DeleteRock is not implemented"))
}
