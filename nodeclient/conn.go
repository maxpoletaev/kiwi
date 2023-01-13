package nodeclient

//go:generate mockgen -source=conn.go -destination=conn_mock.go -package=clust

import (
	"context"
	"sync/atomic"

	"google.golang.org/protobuf/types/known/emptypb"

	faildetectorpb "github.com/maxpoletaev/kiwi/faildetector/proto"
	"github.com/maxpoletaev/kiwi/internal/multierror"
	membershippb "github.com/maxpoletaev/kiwi/membership/proto"
	storagepb "github.com/maxpoletaev/kiwi/storage/proto"
)

// Conn is an interface to interact with a remote cluster member.
type Conn interface {
	Join(ctx context.Context, req *membershippb.JoinRequest) (*membershippb.JoinResponse, error)
	Members(ctx context.Context) (*membershippb.MembersResponse, error)
	Get(ctx context.Context, req *storagepb.GetRequest) (*storagepb.GetResponse, error)
	Put(ctx context.Context, req *storagepb.PutRequest) (*storagepb.PutResponse, error)
	PingDirect(ctx context.Context) (*faildetectorpb.PingResponse, error)
	PingIndirect(ctx context.Context, req *faildetectorpb.PingRequest) (*faildetectorpb.PingResponse, error)
	IsClosed() bool
	Close() error
}

// GrpcConn encapsulates several GRPC clients to interact with cluster members.
type GrpcConn struct {
	faildetectorClient faildetectorpb.FailDetectorServiceClient
	membershipClient   membershippb.MembershipServiceClient
	storageClient      storagepb.StorageServiceClient
	onClose            []func() error
	closed             uint32
}

func (c *GrpcConn) addOnCloseHook(f func() error) {
	c.onClose = append(c.onClose, f)
}

// Close closes the underlying GRPC connection. Please note that the connection may
// be used by other goroutines and closing it may cause some operations to fail.
func (c *GrpcConn) Close() error {
	if !atomic.CompareAndSwapUint32(&c.closed, 0, 1) {
		return nil // already closed
	}

	errs := multierror.New[int]()
	for idx, f := range c.onClose {
		if err := f(); err != nil {
			errs.Add(idx, err)
		}
	}

	return errs.Ret()
}

// IsClosed returns true if the connection is closed.
func (c *GrpcConn) IsClosed() bool {
	return atomic.LoadUint32(&c.closed) == 1
}

// Get returns the value for the given key. If the key does not exist, it returns an empty array.
func (c *GrpcConn) Get(ctx context.Context, req *storagepb.GetRequest) (*storagepb.GetResponse, error) {
	return c.storageClient.Get(ctx, req)
}

// Put sets the value for the given key. If the key already exists, it is overwritten.
func (c *GrpcConn) Put(ctx context.Context, req *storagepb.PutRequest) (*storagepb.PutResponse, error) {
	return c.storageClient.Put(ctx, req)
}

// Join attempts to join the cluster. It returns the list of current cluster members before the join.
func (c *GrpcConn) Join(ctx context.Context, req *membershippb.JoinRequest) (*membershippb.JoinResponse, error) {
	return c.membershipClient.Join(ctx, req)
}

// Info returns the info about all cluster members.
func (c *GrpcConn) Members(ctx context.Context) (*membershippb.MembersResponse, error) {
	return c.membershipClient.Members(ctx, &emptypb.Empty{})
}

// PingDirect sends a ping to the given node directly.
func (c *GrpcConn) PingDirect(ctx context.Context) (*faildetectorpb.PingResponse, error) {
	return c.faildetectorClient.Ping(ctx, &faildetectorpb.PingRequest{})
}

// PingIndirect sends a ping to the given node through an intermediate node.
func (c *GrpcConn) PingIndirect(ctx context.Context, req *faildetectorpb.PingRequest) (*faildetectorpb.PingResponse, error) {
	return c.faildetectorClient.Ping(ctx, req)
}
