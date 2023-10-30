// Package audioout implements the viam-labs:service:audioout API
package audioout

import (
	"context"

	"github.com/edaniels/golog"
	"go.viam.com/utils/protoutils"
	"go.viam.com/utils/rpc"

	pb "github.com/viam-labs/audioout-api/src/audioout_go/grpc"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/robot"
)

// API is the full API definition.
var API = resource.APINamespace("viam-labs").WithComponentType("audioout")

// Named is a helper for getting the named Audioout's typed resource name.
func Named(name string) resource.Name {
	return resource.NewName(API, name)
}

// FromRobot is a helper for getting the named Audioout from the given Robot.
func FromRobot(r robot.Robot, name string) (Audioout, error) {
	return robot.ResourceFromRobot[Audioout](r, Named(name))
}

func init() {
	resource.RegisterAPI(API, resource.APIRegistration[Audioout]{
		// Reconfigurable, and contents of reconfwrapper.go are only needed for standalone (non-module) uses.
		RPCServiceServerConstructor: NewRPCServiceServer,
		RPCServiceHandler:           pb.RegisterAudiooutServiceHandlerFromEndpoint,
		RPCServiceDesc:              &pb.DisplayService_ServiceDesc,
		RPCClient: func(
			ctx context.Context,
			conn rpc.ClientConn,
			remoteName string,
			name resource.Name,
			logger golog.Logger,
		) (Audioout, error) {
			return NewClientFromConn(conn, remoteName, name, logger), nil
		},
	})
}

// Audioout defines the Go interface for the component (should match the protobuf methods.)
type Audioout interface {
	resource.Resource
	Play(ctx context.Context, file_path string, loop_count, maxtime_ms, fadein_ms int) error
	Stop(ctx context.Context) error
}

// serviceServer implements the Audioout RPC service from audioout.proto.
type serviceServer struct {
	pb.UnimplementedDisplayServiceServer
	coll resource.APIResourceCollection[Audioout]
}

// NewRPCServiceServer returns a new RPC server for the Audioout API.
func NewRPCServiceServer(coll resource.APIResourceCollection[Audioout]) interface{} {
	return &serviceServer{coll: coll}
}

func (s *serviceServer) Play(ctx context.Context, req *pb.PlayRequest) (*pb.PlayResponse, error) {
	g, err := s.coll.Resource(req.Name)
	if err != nil {
		return nil, err
	}
	err = g.Play(ctx, req.file_path, int(req.loop_count), int(req.maxtime_ms), int(req.fadein_ms))
	if err != nil {
		return nil, err
	}
	return &pb.PlayResponse{}, nil
}

func (s *serviceServer) Stop(ctx context.Context, req *pb.StopRequest) (*pb.StopResponse, error) {
	g, err := s.coll.Resource(req.Name)
	if err != nil {
		return nil, err
	}
	err = g.Stop(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.StopResponse{}, nil
}

func (s *serviceServer) DoCommand(ctx context.Context, req *pb.DoCommandRequest) (*pb.DoCommandResponse, error) {
	g, err := s.coll.Resource(req.Name)
	if err != nil {
		return nil, err
	}
	resp, err := g.DoCommand(ctx, req.Command.AsMap())
	if err != nil {
		return nil, err
	}
	pbResp, err := protoutils.StructToStructPb(resp)
	if err != nil {
		return nil, err
	}
	return &pb.DoCommandResponse{Result: pbResp}, nil
}

// NewClientFromConn creates a new Audioout RPC client from an existing connection.
func NewClientFromConn(conn rpc.ClientConn, remoteName string, name resource.Name, logger golog.Logger) Audioout {
	sc := newSvcClientFromConn(conn, remoteName, name, logger)
	return clientFromSvcClient(sc, name.ShortName())
}

func newSvcClientFromConn(conn rpc.ClientConn, remoteName string, name resource.Name, logger golog.Logger) *serviceClient {
	client := pb.NewAudiooutServiceClient(conn)
	sc := &serviceClient{
		Named:  name.PrependRemote(remoteName).AsNamed(),
		client: client,
		logger: logger,
	}
	return sc
}

type serviceClient struct {
	resource.Named
	resource.AlwaysRebuild
	resource.TriviallyCloseable
	client pb.AudiooutServiceClient
	logger golog.Logger
}

// client is an gripper client.
type client struct {
	*serviceClient
	name string
}

func clientFromSvcClient(sc *serviceClient, name string) Audioout {
	return &client{sc, name}
}

func (c *client) Play(ctx context.Context, file_path string, loop_count, maxtime_ms, fadein_ms int) error {
	_, err := c.client.Play(ctx, &pb.PlayRequest{
		Name:       c.name,
		file_path:  file_path,
		loop_count: int32(loop_count),
		maxtime_ms: int32(maxtime_ms),
		fadein_ms:  int32(fadein_ms),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Stop(ctx context.Context) error {
	_, err := c.client.Reset(ctx, &pb.StopRequest{
		Name: c.name,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *client) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	command, err := protoutils.StructToStructPb(cmd)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.DoCommand(ctx, &pb.DoCommandRequest{
		Name:    c.name,
		Command: command,
	})
	if err != nil {
		return nil, err
	}
	return resp.Result.AsMap(), nil
}
