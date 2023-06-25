package sync

import (
	"context"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"google.golang.org/protobuf/proto"
)

func (s *Service) validateBuilderBid(ctx context.Context, id peer.ID, message *pubsub.Message) (pubsub.ValidationResult, error) {
	// TODO: Add builder bid validation.
	return pubsub.ValidationAccept, nil
}

func (s *Service) builderBidSubscriber(ctx context.Context, message proto.Message) error {
	// TODO: Save builder bid to op pool.
	return nil
}
