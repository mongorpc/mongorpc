package encoder

import (
	"time"

	"github.com/mongorpc/mongorpc/lib/mongorpc"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Encode input type to mongorpc mongorpc value
func Encode(in interface{}) *mongorpc.Value {
	result := &mongorpc.Value{}
	switch value := in.(type) {

	// native types
	case int:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_Integer64Value{
				Integer64Value: int64(value),
			},
		}

	case int32:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_Integer32Value{
				Integer32Value: value,
			},
		}

	case int64:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_Integer64Value{
				Integer64Value: value,
			},
		}

	case float32:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_DoubleValue{
				DoubleValue: float64(value),
			},
		}

	case float64:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_DoubleValue{
				DoubleValue: value,
			},
		}

	case string:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_StringValue{
				StringValue: value,
			},
		}
	case bool:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_BooleanValue{
				BooleanValue: value,
			},
		}

	case []interface{}:
		arr := mongorpc.ArrayValue{}
		for _, v := range value {
			arr.Values = append(arr.Values, Encode(v))
		}

		result = &mongorpc.Value{
			Type: &mongorpc.Value_ArrayValue{
				ArrayValue: &arr,
			},
		}

	case map[string]interface{}:
		fields := map[string]*mongorpc.Value{}
		for k, v := range value {
			fields[k] = Encode(v)
		}

		result = &mongorpc.Value{
			Type: &mongorpc.Value_MapValue{
				MapValue: &mongorpc.MapValue{
					Fields: fields,
				},
			},
		}

	case nil:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_NullValue{
				NullValue: mongorpc.NullValue_NULL_VALUE,
			},
		}

	// other native types
	case time.Time:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_TimestampValue{
				TimestampValue: &mongorpc.Timestamp{
					Seconds: value.Unix(),
					Nanos:   int32(value.Nanosecond()),
				},
			},
		}

	// MongoDB Types
	case primitive.ObjectID:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_ObjectIdValue{
				ObjectIdValue: &mongorpc.ObjectId{
					Id: value.Hex(),
				},
			},
		}

	case primitive.D:
		arr := mongorpc.ArrayValue{}
		for _, v := range value {
			arr.Values = append(arr.Values, Encode(v))
		}

		result = &mongorpc.Value{
			Type: &mongorpc.Value_ArrayValue{
				ArrayValue: &arr,
			},
		}

	case primitive.E:
		fields := map[string]*mongorpc.Value{}
		fields[value.Key] = Encode(value.Value)

		result = &mongorpc.Value{
			Type: &mongorpc.Value_MapValue{
				MapValue: &mongorpc.MapValue{
					Fields: fields,
				},
			},
		}

	case primitive.A:
		arr := mongorpc.ArrayValue{}
		for _, v := range value {
			arr.Values = append(arr.Values, Encode(v))
		}

		result = &mongorpc.Value{
			Type: &mongorpc.Value_ArrayValue{
				ArrayValue: &arr,
			},
		}

	case primitive.M:
		fields := map[string]*mongorpc.Value{}
		for k, v := range value {
			fields[k] = Encode(v)
		}

		result = &mongorpc.Value{
			Type: &mongorpc.Value_MapValue{
				MapValue: &mongorpc.MapValue{
					Fields: fields,
				},
			},
		}

	case primitive.Binary:
		logrus.Errorf("encoder.encode unsupported type: %T", value)

	case primitive.Timestamp:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_TimestampValue{
				TimestampValue: &mongorpc.Timestamp{
					Seconds: int64(value.T),
					Nanos:   int32(value.I),
				},
			},
		}

	case primitive.DateTime:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_TimestampValue{
				TimestampValue: &mongorpc.Timestamp{
					Seconds: int64(value),
					Nanos:   0,
				},
			},
		}

	case primitive.Null:
		result = &mongorpc.Value{
			Type: &mongorpc.Value_NullValue{
				NullValue: mongorpc.NullValue_NULL_VALUE,
			},
		}

	default:
		logrus.Errorf("encoder.encode unsupported type: %T", value)
	}

	return result
}
