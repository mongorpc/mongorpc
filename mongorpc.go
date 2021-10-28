package mongorpc

import (
	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoRPC Server
type MongoRPCServer struct {
	proto.UnimplementedMongoRPCServer
	DB *mongo.Client
}

// Encoder encodes a values to a mongorpc proto types
type MongoRPCEncoder interface {
	EncodeArray([]interface{}) proto.Array
}

// Decoder decodes a mongorpc proto types to a values
type MongoRPCDecoder interface {
	DecodeArray(proto.Array) []interface{}
}

// EncodeMap encodes a map to a mongorpc proto types
func EncodeMap(m map[string]interface{}) *proto.Map {
	result := map[string]*proto.Value{}

	// iterate over the map
	for k, v := range m {

		// check if the value is nil then return proto null
		if v == nil {
			result[k] = &proto.Value{
				Type: &proto.Value_NullValue{
					NullValue: proto.Null_NULL_VALUE,
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
			}
		}
	}

	// return the proto map
	return &proto.Map{
		Fields: result,
	}
}

// DecodeMap decodes a mongorpc proto types to a map
func EncodeArray(arr []interface{}) *proto.Array {
	result := []*proto.Value{}

	// iterate over the array
	for _, v := range arr {

		// check if the value is nil then return proto null
		if v == nil {
			result = append(result, &proto.Value{
				Type: &proto.Value_NullValue{
					NullValue: proto.Null_NULL_VALUE,
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
			}
		}
	}

	// return the proto array
	return &proto.Array{
		Values: result,
	}
}
