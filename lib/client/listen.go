package client

import (
	"context"
	"fmt"

	"github.com/mongorpc/mongorpc/lib/decoder"
	"github.com/mongorpc/mongorpc/lib/encoder"
	"github.com/mongorpc/mongorpc/lib/mongorpc"
	"github.com/sirupsen/logrus"
)

// mongodb operationType enum
type OperationType string

const (
	Insert OperationType = "insert"
	Update OperationType = "update"
	Delete OperationType = "delete"
)

func (q *QueryBuilder) Listen(ctx context.Context, operation *OperationType) (chan interface{}, error) {
	database := q.collection.parent
	collection := q.collection

	// create channel
	ch := make(chan interface{})

	filter := make(map[string]interface{})
	if operation != nil {
		filter["operationType"] = string(*operation)
	}
	if q.filter != nil {
		for k, v := range q.filter {
			filter[fmt.Sprintf("fullDocument.%s", k)] = v
		}
	}

	match := make(map[string]interface{})
	match["$match"] = filter

	logrus.Println([]*mongorpc.Value{
		encoder.Encode(match),
	})

	// crate mongorpc get document request
	listner, err := q.client.mongorpc.Listen(ctx, &mongorpc.ListenRequest{
		Database:   database.name,
		Collection: collection.name,
		Options: &mongorpc.ChangeStreamOptions{
			FullDocument: true,
		},
		Pipeline: []*mongorpc.Value{
			encoder.Encode(match),
		},
	})

	if err != nil {
		return nil, err
	}

	// start listening
	go func() {
		for {
			// get document
			doc, err := listner.Recv()
			if err != nil {
				close(ch)
				return
			}

			// decode document
			data := decoder.Decode(doc.Changes)

			// send document
			ch <- data
		}
	}()

	return ch, nil
}
