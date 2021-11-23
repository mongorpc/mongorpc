package decoder

import (
	"github.com/mongorpc/mongorpc/lib/mongorpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Decode mongorpc value to value
func Decode(value *mongorpc.Value) interface{} {
	switch value.Type.(type) {
	case *mongorpc.Value_NullValue:
		return nil

	case *mongorpc.Value_Integer32Value:
		return value.GetInteger32Value()

	case *mongorpc.Value_Integer64Value:
		return value.GetInteger64Value()

	case *mongorpc.Value_DoubleValue:
		return value.GetDoubleValue()

	case *mongorpc.Value_StringValue:
		return value.GetStringValue()

	case *mongorpc.Value_BooleanValue:
		return value.GetBooleanValue()

	case *mongorpc.Value_ArrayValue:
		arr := []interface{}{}
		for _, v := range value.GetArrayValue().Values {
			arr = append(arr, Decode(v))
		}
		return arr

	case *mongorpc.Value_MapValue:
		m := map[string]interface{}{}
		for k, v := range value.GetMapValue().Fields {
			m[k] = Decode(v)
		}
		return m

	case *mongorpc.Value_ObjectIdValue:
		id, err := primitive.ObjectIDFromHex(value.GetObjectIdValue().Id)
		if err != nil {
			return nil
		}
		return id

	case *mongorpc.Value_TimestampValue:
		return primitive.DateTime(value.GetTimestampValue().Seconds)
	}

	return nil
}
