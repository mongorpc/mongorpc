package mongorpc

import (
	"reflect"

	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Encoder encodes a values to a mongorpc proto types
type MongoRPCEncoder interface {
	EncodeArray([]interface{}) proto.ArrayValue
}

// Decoder decodes a mongorpc proto types to a values
type MongoRPCDecoder interface {
	DecodeArray(proto.ArrayValue) []interface{}
}

// Decode Value to a map
func DecodeValue(v *proto.Value) interface{} {

	switch v.Type.(type) {
	case *proto.Value_NullValue:
		return nil
	case *proto.Value_IntegerValue:
		return v.GetIntegerValue()
	case *proto.Value_StringValue:
		return v.GetStringValue()
	case *proto.Value_BoolValue:
		return v.GetBoolValue()
	case *proto.Value_DoubleValue:
		return v.GetDoubleValue()
	case *proto.Value_ArrayValue:
		return DecodeArray(v.GetArrayValue())
	case *proto.Value_MapValue:
		return DecodeMap(v.GetMapValue())
	default:
		logrus.Error("Unsupported type: ", reflect.TypeOf(v.Type))
	}

	return nil
}

// DecodeMap decodes a mongorpc proto types to a map
func DecodeMap(m *proto.MapValue) map[string]interface{} {
	result := map[string]interface{}{}

	// iterate over the map
	for k, v := range m.Fields {

		switch v.Type.(type) {
		case *proto.Value_NullValue:
			result[k] = nil
		case *proto.Value_IntegerValue:
			result[k] = v.GetIntegerValue()
		case *proto.Value_StringValue:
			result[k] = v.GetStringValue()
		case *proto.Value_BoolValue:
			result[k] = v.GetBoolValue()
		case *proto.Value_DoubleValue:
			result[k] = v.GetDoubleValue()
		case *proto.Value_ArrayValue:
			result[k] = DecodeArray(v.GetArrayValue())
		case *proto.Value_MapValue:
			result[k] = DecodeMap(v.GetMapValue())
		default:
			logrus.Error("Unsupported type: ", reflect.TypeOf(v.Type))
		}
	}

	// return the map
	return result
}

func DecodeArray(a *proto.ArrayValue) []interface{} {
	result := []interface{}{}

	// iterate over the array
	for _, v := range a.Values {
		switch v.Type.(type) {
		case *proto.Value_NullValue:
			result = append(result, nil)
		case *proto.Value_IntegerValue:
			result = append(result, v.GetIntegerValue())
		case *proto.Value_StringValue:
			result = append(result, v.GetStringValue())
		case *proto.Value_BoolValue:
			result = append(result, v.GetBoolValue())
		case *proto.Value_DoubleValue:
			result = append(result, v.GetDoubleValue())
		case *proto.Value_ArrayValue:
			result = append(result, DecodeArray(v.GetArrayValue()))
		case *proto.Value_MapValue:
			result = append(result, DecodeMap(v.GetMapValue()))
		default:
			logrus.Error("Unsupported type: ", reflect.TypeOf(v.Type))
		}
	}

	// return the array
	return result
}

// EncodeMap encodes a map to a mongorpc proto types
func EncodeMap(m map[string]interface{}) *proto.MapValue {
	result := map[string]*proto.Value{}

	// iterate over the map
	for k, v := range m {

		// check if the value is nil then return proto null
		if v == nil {
			result[k] = &proto.Value{
				Type: &proto.Value_NullValue{
					NullValue: proto.NullValue_NULL_VALUE,
				},
			}
		} else {

			// check value type and encode to mongorpc proto types
			switch value := v.(type) {
			case int:
				result[k] = &proto.Value{
					Type: &proto.Value_IntegerValue{
						IntegerValue: int64(value),
					},
				}
			case int32:
				result[k] = &proto.Value{
					Type: &proto.Value_IntegerValue{
						IntegerValue: int64(value),
					},
				}
			case int64:
				result[k] = &proto.Value{
					Type: &proto.Value_IntegerValue{
						IntegerValue: int64(value),
					},
				}
			case string:
				result[k] = &proto.Value{
					Type: &proto.Value_StringValue{
						StringValue: value,
					},
				}
			case bool:
				result[k] = &proto.Value{
					Type: &proto.Value_BoolValue{
						BoolValue: value,
					},
				}
			case float64:
				result[k] = &proto.Value{
					Type: &proto.Value_DoubleValue{
						DoubleValue: value,
					},
				}
			case []interface{}:
				result[k] = &proto.Value{
					Type: &proto.Value_ArrayValue{
						ArrayValue: EncodeArray(value),
					},
				}
			case map[string]interface{}:
				result[k] = &proto.Value{
					Type: &proto.Value_MapValue{
						MapValue: EncodeMap(value),
					},
				}

			// Mongo Types
			case primitive.ObjectID:
				result[k] = &proto.Value{
					Type: &proto.Value_ObjectIdValue{
						ObjectIdValue: &proto.ObjectID{
							Id: value.Hex(),
						},
					},
				}

			case primitive.DateTime:
				result[k] = &proto.Value{
					Type: &proto.Value_DateValue{
						DateValue: &proto.Timestamp{
							Seconds: int64(value) / 1000,
							Nanos:   int32(int64(value) % 1000 * 1000000),
						},
					},
				}
			case primitive.A:
				result[k] = &proto.Value{
					Type: &proto.Value_ArrayValue{
						ArrayValue: EncodeArray(value),
					},
				}
			default:
				logrus.Error("Unsupported type: ", reflect.TypeOf(value))
			}
		}
	}

	// return the proto map
	return &proto.MapValue{
		Fields: result,
	}
}

// DecodeMap decodes a mongorpc proto types to a map
func EncodeArray(arr []interface{}) *proto.ArrayValue {
	result := []*proto.Value{}

	// iterate over the array
	for _, v := range arr {

		// check if the value is nil then return proto null
		if v == nil {
			result = append(result, &proto.Value{
				Type: &proto.Value_NullValue{
					NullValue: proto.NullValue_NULL_VALUE,
				},
			})
		} else {

			// check value type and encode to mongorpc proto types
			switch value := v.(type) {
			case int:
				result = append(result, &proto.Value{
					Type: &proto.Value_IntegerValue{
						IntegerValue: int64(value),
					},
				})
			case int32:
				result = append(result, &proto.Value{
					Type: &proto.Value_IntegerValue{
						IntegerValue: int64(value),
					},
				})
			case int64:
				result = append(result, &proto.Value{
					Type: &proto.Value_IntegerValue{
						IntegerValue: int64(value),
					},
				})
			case string:
				result = append(result, &proto.Value{
					Type: &proto.Value_StringValue{
						StringValue: value,
					},
				})
			case bool:
				result = append(result, &proto.Value{
					Type: &proto.Value_BoolValue{
						BoolValue: value,
					},
				})
			case float64:
				result = append(result, &proto.Value{
					Type: &proto.Value_DoubleValue{
						DoubleValue: value,
					},
				})
			case []interface{}:
				result = append(result, &proto.Value{
					Type: &proto.Value_ArrayValue{
						ArrayValue: EncodeArray(value),
					},
				})
			case map[string]interface{}:
				result = append(result, &proto.Value{
					Type: &proto.Value_MapValue{
						MapValue: EncodeMap(value),
					},
				})

			// Mongo Types
			case primitive.ObjectID:
				result = append(result, &proto.Value{
					Type: &proto.Value_ObjectIdValue{
						ObjectIdValue: &proto.ObjectID{
							Id: value.Hex(),
						},
					},
				})

			case primitive.DateTime:
				result = append(result, &proto.Value{
					Type: &proto.Value_DateValue{
						DateValue: &proto.Timestamp{
							Seconds: int64(value) / 1000,
							Nanos:   int32(int64(value) % 1000 * 1000000),
						},
					},
				})
			case primitive.A:
				result = append(result, &proto.Value{
					Type: &proto.Value_ArrayValue{
						ArrayValue: EncodeArray(value),
					},
				})
			default:
				logrus.Error("Unsupported type: ", reflect.TypeOf(value))
			}
		}
	}

	// return the proto array
	return &proto.ArrayValue{
		Values: result,
	}
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
