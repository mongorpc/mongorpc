package client

import (
	"context"

	"github.com/mongorpc/mongorpc/lib/decoder"
	"github.com/mongorpc/mongorpc/lib/mongorpc"
)

func (c *Collection) Listen(ctx context.Context) (chan interface{}, error) {
	database := c.parent
	collection := c

	// create channel
	ch := make(chan interface{})

	// crate mongorpc get document request
	listner, err := c.client.mongorpc.Listen(ctx, &mongorpc.ListenRequest{
		Database:   database.name,
		Collection: collection.name,
		Options: &mongorpc.ChangeStreamOptions{
			FullDocument: true,
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
