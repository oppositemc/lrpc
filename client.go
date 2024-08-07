package lrpc

import (
	"context"
	feed "github.com/emalak/lrpc/rpc/feed"
	storage "github.com/emalak/lrpc/rpc/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Settings struct {
	FeedOpts    *FeedOptions
	StorageOpts *StorageOptions
}

type FeedOptions struct {
	Address string
}

type StorageOptions struct {
	Address string
}

type Client struct {
	Feed    *Feed
	Storage *Storage
}

type Feed struct {
	conn   *grpc.ClientConn
	Client feed.LandmarkFeedClient
}

func newFeed(ctx context.Context, s Settings) (*Feed, error) {
	conn, err := grpc.DialContext(
		ctx,
		s.FeedOpts.Address,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := feed.NewLandmarkFeedClient(conn)
	f := Feed{
		conn:   conn,
		Client: client,
	}
	return &f, nil
}

func (f *Feed) Close() error {
	return f.conn.Close()
}

type Storage struct {
	conn   *grpc.ClientConn
	Client storage.StorageServiceClient
}

func newStorage(ctx context.Context, s Settings) (*Storage, error) {
	conn, err := grpc.DialContext(
		ctx,
		s.StorageOpts.Address,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := storage.NewStorageServiceClient(conn)
	st := Storage{
		conn:   conn,
		Client: client,
	}
	return &st, nil
}

func (s *Storage) Close() error {
	return s.conn.Close()
}

func New(ctx context.Context, s Settings) (*Client, error) {
	client := Client{}
	if s.StorageOpts != nil {
		f, err := newStorage(ctx, s)
		if err != nil {
			return nil, err
		}
		client.Storage = f
	}
	if s.FeedOpts != nil {
		f, err := newFeed(ctx, s)
		if err != nil {
			return nil, err
		}
		client.Feed = f
	}
	return &client, nil
}

func (c *Client) closeFeedConn() error {
	if c.Feed != nil {
		return c.Feed.Close()
	}
	return nil
}

func (c *Client) closeStorageConn() error {
	if c.Storage != nil {
		return c.Storage.Close()
	}
	return nil
}

func (c *Client) Close() error {
	if err := c.closeStorageConn(); err != nil {
		return err
	}
	if err := c.closeFeedConn(); err != nil {
		return err
	}
	return nil
}
