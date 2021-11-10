package mongorpc

import (
	"time"

	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Encode input type to mongorpc proto value
func Encode(in interface{}) *proto.Value {
	result := &proto.Value{}
	switch value := in.(type) {

	// native types
	case int:
		result = &proto.Value{
			Type: &proto.Value_IntegerValue{
				IntegerValue: int64(value),
			},
		}

	case int32:
		result = &proto.Value{
			Type: &proto.Value_IntegerValue{
				IntegerValue: int64(value),
			},
		}

	case int64:
		result = &proto.Value{
			Type: &proto.Value_IntegerValue{
				IntegerValue: value,
			},
		}

	case float32:
		result = &proto.Value{
			Type: &proto.Value_DoubleValue{
				DoubleValue: float64(value),
			},
		}

	case float64:
		result = &proto.Value{
			Type: &proto.Value_DoubleValue{
				DoubleValue: value,
			},
		}

	case string:
		result = &proto.Value{
			Type: &proto.Value_StringValue{
				StringValue: value,
			},
		}
	case bool:
		result = &proto.Value{
			Type: &proto.Value_BoolValue{
				BoolValue: value,
			},
		}

	case []interface{}:
		arr := proto.ArrayValue{}
		for _, v := range value {
			arr.Values = append(arr.Values, Encode(v))
		}

		result = &proto.Value{
			Type: &proto.Value_ArrayValue{
				ArrayValue: &arr,
			},
		}

	case map[string]interface{}:
		fields := map[string]*proto.Value{}
		for k, v := range value {
			fields[k] = Encode(v)
		}

		result = &proto.Value{
			Type: &proto.Value_MapValue{
				MapValue: &proto.MapValue{
					Fields: fields,
				},
			},
		}

	case nil:
		result = &proto.Value{
			Type: &proto.Value_NullValue{
				NullValue: proto.NullValue_NULL_VALUE,
			},
		}

	// other native types
	case time.Time:
		result = &proto.Value{
			Type: &proto.Value_DateValue{
				DateValue: &proto.Timestamp{
					Seconds: value.Unix(),
					Nanos:   int32(value.Nanosecond()),
				},
			},
		}

	// MongoDB Types
	case primitive.ObjectID:
		result = &proto.Value{
			Type: &proto.Value_ObjectIdValue{
				ObjectIdValue: &proto.ObjectID{
					Id: value.Hex(),
				},
			},
		}

	case primitive.D:
		arr := proto.ArrayValue{}
		for _, v := range value {
			arr.Values = append(arr.Values, Encode(v))
		}

		result = &proto.Value{
			Type: &proto.Value_ArrayValue{
				ArrayValue: &arr,
			},
		}

	case primitive.E:
		fields := map[string]*proto.Value{}
		fields[value.Key] = Encode(value.Value)

		result = &proto.Value{
			Type: &proto.Value_MapValue{
				MapValue: &proto.MapValue{
					Fields: fields,
				},
			},
		}

	case primitive.A:
		arr := proto.ArrayValue{}
		for _, v := range value {
			arr.Values = append(arr.Values, Encode(v))
		}

		result = &proto.Value{
			Type: &proto.Value_ArrayValue{
				ArrayValue: &arr,
			},
		}

	case primitive.M:
		fields := map[string]*proto.Value{}
		for k, v := range value {
			fields[k] = Encode(v)
		}

		result = &proto.Value{
			Type: &proto.Value_MapValue{
				MapValue: &proto.MapValue{
					Fields: fields,
				},
			},
		}

	case primitive.Binary:
		result = &proto.Value{
			Type: &proto.Value_BytesValue{
				BytesValue: value.Data,
			},
		}

	case primitive.Timestamp:
		result = &proto.Value{
			Type: &proto.Value_DateValue{
				DateValue: &proto.Timestamp{
					Seconds: int64(value.T),
					Nanos:   int32(value.I),
				},
			},
		}

	case primitive.DateTime:
		result = &proto.Value{
			Type: &proto.Value_DateValue{
				DateValue: &proto.Timestamp{
					Seconds: int64(value),
					Nanos:   0,
				},
			},
		}

	case primitive.Null:
		result = &proto.Value{
			Type: &proto.Value_NullValue{
				NullValue: proto.NullValue_NULL_VALUE,
			},
		}

	default:
		logrus.Errorf("encoder.encode unsupported type: %T", value)
	}

	return result
}
