package mongorpc

import (
	"fmt"

	pb "github.com/mongorpc/mongorpc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRPCServer struct {
	pb.UnimplementedMongoRPCServer
	DB *mongo.Client
}

func parseJSON(json map[string]interface{}) map[string]*pb.Value {
	res := map[string]*pb.Value{}
	for k, v := range json {

		switch value := v.(type) {
		case nil:
			res[k] = &pb.Value{
				ValueType: &pb.Value_NullValue{
					// TODO: Assign something to null value
				},
			}
		case bool:
			res[k] = &pb.Value{
				ValueType: &pb.Value_BooleanValue{
					BooleanValue: value,
				},
			}
		case int64:
			res[k] = &pb.Value{
				ValueType: &pb.Value_IntegerValue{
					IntegerValue: value,
				},
			}
		case float64:
			res[k] = &pb.Value{
				ValueType: &pb.Value_DoubleValue{
					DoubleValue: value,
				},
			}
		case pb.Timestamp:
			res[k] = &pb.Value{
				ValueType: &pb.Value_TimestampValue{
					TimestampValue: &value,
				},
			}
		case string:
			res[k] = &pb.Value{
				ValueType: &pb.Value_StringValue{
					StringValue: value,
				},
			}
		case []byte:
			res[k] = &pb.Value{
				ValueType: &pb.Value_BytesValue{
					BytesValue: value,
				},
			}
		case pb.LatLng:
			res[k] = &pb.Value{
				ValueType: &pb.Value_GeoPointValue{
					GeoPointValue: &value,
				},
			}
		case map[string]interface{}:
			res[k] = &pb.Value{
				ValueType: &pb.Value_MapValue{
					MapValue: &pb.MapValue{
						Fields: parseJSON(value),
					},
				},
			}
		case []interface{}:
			res[k] = &pb.Value{
				ValueType: &pb.Value_ArrayValue{
					ArrayValue: &pb.ArrayValue{
						Values: parseArray(value),
					},
				},
			}
		default:
			fmt.Println(value)
			//check for timestamp
			logrus.Println("UNKNOWN TYPE: ", value)
		}
	}

	return res
}

func parseArray(array []interface{}) []*pb.Value {
	res := []*pb.Value{}
	for _, v := range array {
		// xType := fmt.Sprintf("%T", v)
		// fmt.Println(xType)
		switch value := v.(type) {

		case nil:
			res = append(res, &pb.Value{
				ValueType: &pb.Value_NullValue{
					// TODO: Assign something to null value
				},
			})
		case bool:
			res = append(res, &pb.Value{
				ValueType: &pb.Value_BooleanValue{
					BooleanValue: value,
				},
			})
		case int64:
			res = append(res, &pb.Value{
				ValueType: &pb.Value_IntegerValue{
					IntegerValue: value,
				},
			})
		case float64:
			res = append(res, &pb.Value{
				ValueType: &pb.Value_DoubleValue{
					DoubleValue: value,
				},
			})
		case pb.Timestamp:
			res = append(res, &pb.Value{
				ValueType: &pb.Value_TimestampValue{
					TimestampValue: &value,
				},
			})
		case string:
			res = append(res, &pb.Value{
				ValueType: &pb.Value_StringValue{
					StringValue: value,
				},
			})
		case []byte:
			res = append(res, &pb.Value{
				ValueType: &pb.Value_BytesValue{
					BytesValue: value,
				},
			})
		case pb.LatLng:
			res = append(res, &pb.Value{
				ValueType: &pb.Value_GeoPointValue{
					GeoPointValue: &value,
				},
			})
		case map[string]interface{}:
			res = append(res, &pb.Value{
				ValueType: &pb.Value_MapValue{
					MapValue: &pb.MapValue{
						Fields: parseJSON(value),
					},
				},
			})
		case []interface{}:
			res = append(res, &pb.Value{
				ValueType: &pb.Value_ArrayValue{
					ArrayValue: &pb.ArrayValue{
						Values: parseArray(value),
					},
				},
			})
		default:
			//check for timestamp
			logrus.Println("UNKNOWN TYPE: ", value)
		}
	}
	return res
}
