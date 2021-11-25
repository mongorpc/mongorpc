package client

// TODO: test all query builder methods

func (q *QueryBuilder) Exists(field string, value bool) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$exists"] = value
	return q
}

func (q *QueryBuilder) EqualTo(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$eq"] = value
	return q
}

func (q *QueryBuilder) NotEqualTo(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$ne"] = value
	return q
}

func (q *QueryBuilder) GreaterThan(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$gt"] = value
	return q
}

func (q *QueryBuilder) GreaterThanOrEqualTo(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$gte"] = value
	return q
}

func (q *QueryBuilder) LessThan(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$lt"] = value
	return q
}

func (q *QueryBuilder) LessThanOrEqualTo(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$lte"] = value
	return q
}

func (q *QueryBuilder) In(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$in"] = value
	return q
}

func (q *QueryBuilder) NotIn(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$nin"] = value
	return q
}

func (q *QueryBuilder) All(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$all"] = value
	return q
}

func (q *QueryBuilder) Size(field string, value int32) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$size"] = value
	return q
}

func (q *QueryBuilder) Mod(field string, value int32, mod int32) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$mod"] = []interface{}{value, mod}
	return q
}

func (q *QueryBuilder) Regex(field string, value string) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$regex"] = value
	return q
}

func (q *QueryBuilder) NotRegex(field string, value string) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$not"] = map[string]interface{}{"$regex": value}
	return q
}

func (q *QueryBuilder) RegexOptions(field string, options string) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$options"] = options
	return q
}

func (q *QueryBuilder) Not(field string) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$not"] = map[string]interface{}{}
	return q
}

func (q *QueryBuilder) Near(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$near"] = value
	return q
}

func (q *QueryBuilder) NearSphere(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$nearSphere"] = value
	return q
}

func (q *QueryBuilder) Within(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$within"] = value
	return q
}

func (q *QueryBuilder) WithinBox(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$box"] = value
	return q
}

func (q *QueryBuilder) WithinCenter(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$center"] = value
	return q
}

func (q *QueryBuilder) WithinCenterSphere(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$centerSphere"] = value
	return q
}

func (q *QueryBuilder) WithinPolygon(field string, value interface{}) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	if q.filter[field] == nil {
		q.filter[field] = make(map[string]interface{})
	}
	q.filter[field]["$polygon"] = value
	return q
}

// text search
func (q *QueryBuilder) Search(value string) *QueryBuilder {
	if q.filter == nil {
		q.filter = make(map[string]map[string]interface{})
	}
	q.filter["$text"] = map[string]interface{}{"$search": value}
	return q
}
