package mongorpc

import (
	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRPCServer struct {
	proto.UnimplementedMongoRPCServer
	DB *mongo.Client
}

type MongoRPCEncoder interface {
	EncodeArray([]interface{}) proto.Array
}

type MongoRPCDecoder interface {
	DecodeArray(proto.Array) []interface{}
}

func EncodeMap(m map[string]interface{}) *proto.Map {
	result := map[string]*proto.Value{}
	for k, v := range m {
		if v == nil {
			result[k] = &proto.Value{
				Type: &proto.Value_NullValue{
					NullValue: proto.Null_NULL_VALUE,
				},
			}
		} else {
			switch v.(type) {
			case int:
				result[k] = &proto.Value{
					Type: &proto.Value_IntegerValue{
						IntegerValue: int64(v.(int)),
					},
				}
			case string:
				result[k] = &proto.Value{
					Type: &proto.Value_StringValue{
						StringValue: v.(string),
					},
				}
			case bool:
				result[k] = &proto.Value{
					Type: &proto.Value_BoolValue{
						BoolValue: v.(bool),
					},
				}
			case float64:
				result[k] = &proto.Value{
					Type: &proto.Value_DoubleValue{
						DoubleValue: v.(float64),
					},
				}
			case []interface{}:
				result[k] = &proto.Value{
					Type: &proto.Value_ArrayValue{
						ArrayValue: EncodeArray(v.([]interface{})),
					},
				}
			case map[string]interface{}:
				result[k] = &proto.Value{
					Type: &proto.Value_MapValue{
						MapValue: EncodeMap(v.(map[string]interface{})),
					},
				}
			}
		}
	}
	return &proto.Map{
		Fields: result,
	}
}

func EncodeArray(arr []interface{}) *proto.Array {
	result := []*proto.Value{}
	for _, v := range arr {
		if v == nil {
			result = append(result, &proto.Value{
				Type: &proto.Value_NullValue{
					NullValue: proto.Null_NULL_VALUE,
				},
			})
		} else {
			switch v.(type) {
			case int:
				result = append(result, &proto.Value{
					Type: &proto.Value_IntegerValue{
						IntegerValue: int64(v.(int)),
					},
				})
			case string:
				result = append(result, &proto.Value{
					Type: &proto.Value_StringValue{
						StringValue: v.(string),
					},
				})
			case bool:
				result = append(result, &proto.Value{
					Type: &proto.Value_BoolValue{
						BoolValue: v.(bool),
					},
				})
			case float64:
				result = append(result, &proto.Value{
					Type: &proto.Value_DoubleValue{
						DoubleValue: v.(float64),
					},
				})
			case []interface{}:
				result = append(result, &proto.Value{
					Type: &proto.Value_ArrayValue{
						ArrayValue: EncodeArray(v.([]interface{})),
					},
				})
			case map[string]interface{}:
				result = append(result, &proto.Value{
					Type: &proto.Value_MapValue{
						MapValue: EncodeMap(v.(map[string]interface{})),
					},
				})
			}
		}
	}
	return &proto.Array{
		Values: result,
	}
}
