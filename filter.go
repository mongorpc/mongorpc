package mongorpc

import (
	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"github.com/slavabobik/greenleaf"
	"go.mongodb.org/mongo-driver/bson"
)

// proto ascending bool to mongodb sort order
func DecodeSortOrder(ascending bool) int {
	// ascending order for a field, set the field to 1 in the sort document.
	if ascending {
		return 1
	} else {
		// descending order for a field, set the field and -1 in the sort documents.
		return -1
	}
}

// Decode sort to bson
func DecodeSort(fields []*proto.Sort) bson.D {
	sort := bson.D{}
	for _, s := range fields {
		sort = append(sort, bson.E{
			Key:   s.Field,
			Value: DecodeSortOrder(s.Ascending),
		})
	}
	return sort
}

func DecodeFilter(filters []*proto.Filter) greenleaf.Document {
	filter := greenleaf.
		Filter()
	for _, f := range filters {
		switch op := f.Operator.(type) {
		case *proto.Filter_Equal:

			switch op.Equal.Value.Type.(type) {
			case *proto.Value_StringValue:
				filter = filter.EqString(op.Equal.Field, op.Equal.Value.GetStringValue())

				// TODO: add other types

			default:
				logrus.Errorln("unsupported type")
			}

		// TODO: add other operators

		default:
			logrus.Errorln("unsupported type")
		}

		// End of array
	}
	return filter.Build()
}
