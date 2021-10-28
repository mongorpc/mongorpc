package mongorpc

import (
	"github.com/mongorpc/mongorpc/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRPCServer struct {
	proto.UnimplementedMongoRPCServer
	DB *mongo.Client
}

// func MapToValue(json map[string]interface{}) map[string]*proto.Value {
// 	res := map[string]*proto.Value{}
// 	for k, v := range json {
// 		switch value := v.(type) {
// 		case nil:
// 			res[k] = &proto.Value{
// 				ValueType: &proto.Value_NullValue{
// 					NullValue: proto.NullValue_NULL_VALUE,
// 				},
// 			}
// 		case bool:
// 			res[k] = &proto.Value{
// 				ValueType: &proto.Value_BooleanValue{
// 					BooleanValue: value,
// 				},
// 			}
// 		case int64:
// 			res[k] = &proto.Value{
// 				ValueType: &proto.Value_IntegerValue{
// 					IntegerValue: value,
// 				},
// 			}
// 		case float64:
// 			res[k] = &proto.Value{
// 				ValueType: &proto.Value_DoubleValue{
// 					DoubleValue: value,
// 				},
// 			}
// 		case proto.Timestamp:
// 			res[k] = &proto.Value{
// 				ValueType: &proto.Value_TimestampValue{
// 					TimestampValue: &value,
// 				},
// 			}
// 		case string:
// 			res[k] = &proto.Value{
// 				ValueType: &proto.Value_StringValue{
// 					StringValue: value,
// 				},
// 			}
// 		case []byte:
// 			res[k] = &proto.Value{
// 				ValueType: &proto.Value_BytesValue{
// 					BytesValue: value,
// 				},
// 			}
// 		case proto.LatLng:
// 			res[k] = &proto.Value{
// 				ValueType: &proto.Value_GeoPointValue{
// 					GeoPointValue: &value,
// 				},
// 			}
// 		case map[string]interface{}:
// 			res[k] = &proto.Value{
// 				ValueType: &proto.Value_MapValue{
// 					MapValue: &proto.MapValue{
// 						Fields: MapToValue(value),
// 					},
// 				},
// 			}
// 		case []interface{}:
// 			res[k] = &proto.Value{
// 				ValueType: &proto.Value_ArrayValue{
// 					ArrayValue: &proto.ArrayValue{
// 						Values: ArrayToValue(value),
// 					},
// 				},
// 			}
// 		default:
// 			logrus.Println("UNKNOWN TYPE: ", value)
// 		}
// 	}

// 	return res
// }

// func ArrayToValue(array []interface{}) []*proto.Value {
// 	res := []*proto.Value{}
// 	for _, v := range array {
// 		switch value := v.(type) {
// 		case nil:
// 			res = append(res, &proto.Value{
// 				ValueType: &proto.Value_NullValue{
// 					NullValue: proto.NullValue_NULL_VALUE,
// 				},
// 			})
// 		case bool:
// 			res = append(res, &proto.Value{
// 				ValueType: &proto.Value_BooleanValue{
// 					BooleanValue: value,
// 				},
// 			})
// 		case int64:
// 			res = append(res, &proto.Value{
// 				ValueType: &proto.Value_IntegerValue{
// 					IntegerValue: value,
// 				},
// 			})
// 		case float64:
// 			res = append(res, &proto.Value{
// 				ValueType: &proto.Value_DoubleValue{
// 					DoubleValue: value,
// 				},
// 			})
// 		case proto.Timestamp:
// 			res = append(res, &proto.Value{
// 				ValueType: &proto.Value_TimestampValue{
// 					TimestampValue: &value,
// 				},
// 			})
// 		case string:
// 			res = append(res, &proto.Value{
// 				ValueType: &proto.Value_StringValue{
// 					StringValue: value,
// 				},
// 			})
// 		case []byte:
// 			res = append(res, &proto.Value{
// 				ValueType: &proto.Value_BytesValue{
// 					BytesValue: value,
// 				},
// 			})
// 		case proto.LatLng:
// 			res = append(res, &proto.Value{
// 				ValueType: &proto.Value_GeoPointValue{
// 					GeoPointValue: &value,
// 				},
// 			})
// 		case map[string]interface{}:
// 			res = append(res, &proto.Value{
// 				ValueType: &proto.Value_MapValue{
// 					MapValue: &proto.MapValue{
// 						Fields: MapToValue(value),
// 					},
// 				},
// 			})
// 		case []interface{}:
// 			res = append(res, &proto.Value{
// 				ValueType: &proto.Value_ArrayValue{
// 					ArrayValue: &proto.ArrayValue{
// 						Values: ArrayToValue(value),
// 					},
// 				},
// 			})
// 		default:
// 			logrus.Println("UNKNOWN TYPE: ", value)
// 		}
// 	}
// 	return res
// }

// func ValueToMap(value *proto.MapValue) map[string]interface{} {
// 	result := map[string]interface{}{}
// 	for k, v := range value.Fields {
// 		logrus.Println("key: ", k, v)
// 		switch v.ValueType.(type) {
// 		case *proto.Value_NullValue:
// 			result[k] = nil
// 		case *proto.Value_BooleanValue:
// 			result[k] = v.GetBooleanValue()
// 		case *proto.Value_IntegerValue:
// 			result[k] = v.GetIntegerValue()
// 		case *proto.Value_DoubleValue:
// 			result[k] = v.GetDoubleValue()
// 		case *proto.Value_TimestampValue:
// 			result[k] = v.GetTimestampValue()
// 		case *proto.Value_StringValue:
// 			result[k] = v.GetStringValue()
// 		case *proto.Value_BytesValue:
// 			result[k] = v.GetBytesValue()
// 		case *proto.Value_GeoPointValue:
// 			result[k] = v.GetGeoPointValue()
// 		// case *proto.Value_MapValue:
// 		// 	result[k] = ValueToMap(v.GetMapValue())
// 		// case *proto.Value_ArrayValue:
// 		// 	result[k] = ValueToArray(v.GetArrayValue())
// 		default:
// 			logrus.Println("UNKNOWN TYPE: ", v)
// 		}
// 	}
// 	return result
// }

// func ValueToArray(value *proto.ArrayValue) []interface{} {
// 	result := []interface{}{}
// 	for _, v := range value.Values {
// 		switch v.ValueType.(type) {
// 		case *proto.Value_NullValue:
// 			result = append(result, nil)
// 		case *proto.Value_BooleanValue:
// 			result = append(result, v.GetBooleanValue())
// 		case *proto.Value_IntegerValue:
// 			result = append(result, v.GetIntegerValue())
// 		case *proto.Value_DoubleValue:
// 			result = append(result, v.GetDoubleValue())
// 		case *proto.Value_TimestampValue:
// 			result = append(result, v.GetTimestampValue())
// 		case *proto.Value_StringValue:
// 			result = append(result, v.GetStringValue())
// 		case *proto.Value_BytesValue:
// 			result = append(result, v.GetBytesValue())
// 		case *proto.Value_GeoPointValue:
// 			result = append(result, v.GetGeoPointValue())
// 		case *proto.Value_MapValue:
// 			result = append(result, ValueToMap(v.GetMapValue()))
// 		case *proto.Value_ArrayValue:
// 			result = append(result, ValueToArray(v.GetArrayValue()))
// 		default:
// 			logrus.Println("UNKNOWN TYPE: ", v)
// 		}
// 	}
// 	return result
// }
