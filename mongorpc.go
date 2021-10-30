package mongorpc

import (
	"reflect"

	"github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoRPC Server
type MongoRPCServer struct {
	proto.UnimplementedMongoRPCServer
	DB *mongo.Client
}

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

// func EncodeValue(v interface{}) *proto.Value {

// 	switch v.(type) {
// 	case nil:
// 		return &proto.Value{Type: &proto.Value_NullValue{}}
// 	case int:
// 		return &proto.Value{Type: &proto.Value_IntegerValue{IntegerValue: v.(int)}}
// 	case string:
// 		return &proto.Value{Type: &proto.Value_StringValue{StringValue: v.(string)}}
// 	case bool:
// 		return &proto.Value{Type: &proto.Value_BoolValue{BoolValue: v.(bool)}}
// 	case float64:
// 		return &proto.Value{Type: &proto.Value_DoubleValue{DoubleValue: v.(float64)}}
// 	case []interface{}:
// 		return &proto.Value{Type: &proto.Value_ArrayValue{ArrayValue: EncodeArray(v.([]interface{})).(*proto.Array)}}
// 	case map[string]interface{}:
// 		return &proto.Value{Type: &proto.Value_MapValue{MapValue: EncodeMap(v.(map[string]interface{})).(*proto.Map)}}
// 	default:
// 		logrus.Error("Unsupported type: ", reflect.TypeOf(v))
// 	}

// 	return nil
// }

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
