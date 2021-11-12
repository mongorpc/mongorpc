package mongorpc

import (
	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Decode proto value to value
func Decode(value *proto.Value) interface{} {
	switch value.Type.(type) {
	case *proto.Value_NullValue:
		return nil

	case *proto.Value_IntegerValue:
		return value.GetIntegerValue()

	case *proto.Value_DoubleValue:
		return value.GetDoubleValue()

	case *proto.Value_StringValue:
		return value.GetStringValue()

	case *proto.Value_BoolValue:
		return value.GetBoolValue()

	case *proto.Value_ArrayValue:
		arr := []interface{}{}
		for _, v := range value.GetArrayValue().Values {
			arr = append(arr, Decode(v))
		}
		return arr

	case *proto.Value_MapValue:
		m := map[string]interface{}{}
		for k, v := range value.GetMapValue().Fields {
			m[k] = Decode(v)
		}
		return m

	case *proto.Value_ObjectIdValue:
		return value.GetObjectIdValue().Id

	case *proto.Value_DateValue:
		return primitive.DateTime(value.GetDateValue().Seconds)

	}

	return nil
}

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

// Decode sort to bson sort
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

func DecodeSortToInterface(fields []*proto.Sort) []map[string]map[string]interface{} {
	sort := []map[string]map[string]interface{}{}
	for _, s := range fields {
		sort = append(sort, map[string]map[string]interface{}{
			s.Field: {
				"ascending": DecodeSortOrder(s.Ascending),
			},
		})
	}
	return sort
}

// Decode a mongorpc proto filter to a bson filter
func DecodeFilter(filters []*proto.Filter) map[string]map[string]interface{} {
	filter := Filter()
	for _, f := range filters {
		switch op := f.Operator.(type) {
		case *proto.Filter_Equal:

			switch op.Equal.Value.Type.(type) {
			case *proto.Value_StringValue:
				filter = filter.EqString(op.Equal.Field, op.Equal.Value.GetStringValue())

			case *proto.Value_IntegerValue:
				filter = filter.EqInt64(op.Equal.Field, op.Equal.Value.GetIntegerValue())

			case *proto.Value_DoubleValue:
				filter = filter.EqFloat64(op.Equal.Field, op.Equal.Value.GetDoubleValue())

			case *proto.Value_BoolValue:
				filter = filter.EqBool(op.Equal.Field, op.Equal.Value.GetBoolValue())

			// TODO: add support for other types

			default:
				logrus.Errorln("unsupported type")
			}

		case *proto.Filter_NotEqual:

			switch op.NotEqual.Value.Type.(type) {
			case *proto.Value_StringValue:
				filter = filter.NeString(op.NotEqual.Field, op.NotEqual.Value.GetStringValue())

			case *proto.Value_IntegerValue:
				filter = filter.NeInt64(op.NotEqual.Field, op.NotEqual.Value.GetIntegerValue())

			case *proto.Value_DoubleValue:
				filter = filter.NeFloat64(op.NotEqual.Field, op.NotEqual.Value.GetDoubleValue())

			case *proto.Value_BoolValue:
				filter = filter.NeBool(op.NotEqual.Field, op.NotEqual.Value.GetBoolValue())

			default:
				logrus.Errorln("unsupported type")
			}

		case *proto.Filter_Less:

			switch op.Less.Value.Type.(type) {
			case *proto.Value_StringValue:
				filter = filter.LtString(op.Less.Field, op.Less.Value.GetStringValue())

			case *proto.Value_IntegerValue:
				filter = filter.LtInt64(op.Less.Field, op.Less.Value.GetIntegerValue())

			case *proto.Value_DoubleValue:
				filter = filter.LtFloat64(op.Less.Field, op.Less.Value.GetDoubleValue())

			case *proto.Value_BoolValue:
				filter = filter.LtBool(op.Less.Field, op.Less.Value.GetBoolValue())

			default:
				logrus.Errorln("unsupported type")
			}

		case *proto.Filter_LessEqual:

			switch op.LessEqual.Value.Type.(type) {
			case *proto.Value_StringValue:
				filter = filter.LteString(op.LessEqual.Field, op.LessEqual.Value.GetStringValue())

			case *proto.Value_IntegerValue:
				filter = filter.LteInt64(op.LessEqual.Field, op.LessEqual.Value.GetIntegerValue())

			case *proto.Value_DoubleValue:
				filter = filter.LteFloat64(op.LessEqual.Field, op.LessEqual.Value.GetDoubleValue())

			case *proto.Value_BoolValue:
				filter = filter.LteBool(op.LessEqual.Field, op.LessEqual.Value.GetBoolValue())

			default:
				logrus.Errorln("unsupported type")
			}

		case *proto.Filter_Greater:

			switch op.Greater.Value.Type.(type) {
			case *proto.Value_StringValue:
				filter = filter.GtString(op.Greater.Field, op.Greater.Value.GetStringValue())

			case *proto.Value_IntegerValue:
				filter = filter.GtInt64(op.Greater.Field, op.Greater.Value.GetIntegerValue())

			case *proto.Value_DoubleValue:
				filter = filter.GtFloat64(op.Greater.Field, op.Greater.Value.GetDoubleValue())

			case *proto.Value_BoolValue:
				filter = filter.GtBool(op.Greater.Field, op.Greater.Value.GetBoolValue())

			default:
				logrus.Errorln("unsupported type")
			}

		case *proto.Filter_In:
			// TODO: support in

		case *proto.Filter_NotIn:
			// TODO: support not in

		case *proto.Filter_Exists:
			filter = filter.Exists(op.Exists.Field, true)

		case *proto.Filter_NotExists:
			filter = filter.Exists(op.NotExists.Field, false)

		default:
			logrus.Errorln("unsupported type")
		}

		// End of array
	}
	return filter.Build()
}
