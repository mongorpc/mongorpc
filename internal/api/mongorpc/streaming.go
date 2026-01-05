// Package mongorpc provides streaming RPC implementations.
package mongorpc

import (
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	mongorpcv1 "github.com/mongorpc/mongorpc/gen/mongorpc/v1"
	"github.com/mongorpc/mongorpc/internal/rules"
)

// Aggregate runs an aggregation pipeline with streaming results.
func (s *Server) Aggregate(req *mongorpcv1.AggregateRequest, stream mongorpcv1.MongoRPC_AggregateServer) error {
	if req.Pipeline == nil {
		return status.Error(codes.InvalidArgument, "pipeline is required")
	}
	if req.Pipeline.Database == "" || req.Pipeline.Collection == "" {
		return status.Error(codes.InvalidArgument, "database and collection are required")
	}

	ctx := stream.Context()
	slog.Debug("Aggregate", "database", req.Pipeline.Database, "collection", req.Pipeline.Collection)

	// Get collection
	coll := s.db.Client().Database(req.Pipeline.Database).Collection(req.Pipeline.Collection)

	// Convert pipeline stages to BSON
	pipeline, err := aggregationPipelineToBSON(req.Pipeline.Stages)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid pipeline: %v", err)
	}

	// Build options
	opts := options.Aggregate()
	if req.Pipeline.Options != nil {
		if req.Pipeline.Options.AllowDiskUse {
			opts.SetAllowDiskUse(true)
		}
		if req.Pipeline.Options.BatchSize > 0 {
			opts.SetBatchSize(req.Pipeline.Options.BatchSize)
		}
	}

	// Execute aggregation
	cursor, err := coll.Aggregate(ctx, pipeline, opts)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to execute aggregation: %v", err)
	}
	defer cursor.Close(ctx)

	// Stream results
	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return status.Errorf(codes.Internal, "failed to decode result: %v", err)
		}

		doc, err := bsonToDocument(result)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to convert result: %v", err)
		}

		if err := stream.Send(&mongorpcv1.AggregateResponse{Document: doc}); err != nil {
			return status.Errorf(codes.Internal, "failed to send result: %v", err)
		}
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, "cursor error: %v", err)
	}

	return nil
}

// RunQuery runs a query with streaming results.
func (s *Server) RunQuery(req *mongorpcv1.RunQueryRequest, stream mongorpcv1.MongoRPC_RunQueryServer) error {
	if req.Query == nil {
		return status.Error(codes.InvalidArgument, "query is required")
	}
	if req.Query.Database == "" || req.Query.Collection == "" {
		return status.Error(codes.InvalidArgument, "database and collection are required")
	}

	ctx := stream.Context()
	slog.Debug("RunQuery", "database", req.Query.Database, "collection", req.Query.Collection)

	// Get collection
	coll := s.db.Client().Database(req.Query.Database).Collection(req.Query.Collection)

	// Build filter
	filter := bson.D{}
	if req.Query.Filter != nil {
		var err error
		filter, err = filterToBSON(req.Query.Filter)
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "invalid filter: %v", err)
		}
	}

	// Build options
	findOpts := options.Find()
	if req.Query.Limit != nil {
		findOpts.SetLimit(req.Query.Limit.Value)
	}
	if req.Query.Skip > 0 {
		findOpts.SetSkip(req.Query.Skip)
	}
	if req.BatchSize > 0 {
		findOpts.SetBatchSize(req.BatchSize)
	}

	// Handle sort
	if len(req.Query.Sort) > 0 {
		sortDoc := bson.D{}
		for _, sort := range req.Query.Sort {
			direction := 1
			if sort.Direction == mongorpcv1.SortDirection_DESCENDING {
				direction = -1
			}
			sortDoc = append(sortDoc, bson.E{Key: sort.Field, Value: direction})
		}
		findOpts.SetSort(sortDoc)
	}

	// Execute query
	cursor, err := coll.Find(ctx, filter, findOpts)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to execute query: %v", err)
	}
	defer cursor.Close(ctx)

	// Stream results
	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return status.Errorf(codes.Internal, "failed to decode document: %v", err)
		}

		doc, err := bsonToDocument(result)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to convert document: %v", err)
		}

		if err := stream.Send(&mongorpcv1.RunQueryResponse{Document: doc}); err != nil {
			return status.Errorf(codes.Internal, "failed to send document: %v", err)
		}
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, "cursor error: %v", err)
	}

	return nil
}

// Watch watches a collection for changes (change streams).
func (s *Server) Watch(req *mongorpcv1.WatchRequest, stream mongorpcv1.MongoRPC_WatchServer) error {
	if req.Database == "" || req.Collection == "" {
		return status.Error(codes.InvalidArgument, "database and collection are required")
	}

	ctx := stream.Context()
	slog.Debug("Watch", "database", req.Database, "collection", req.Collection)

	// Get collection
	coll := s.db.Client().Database(req.Database).Collection(req.Collection)

	// Build pipeline for change stream
	pipeline := bson.A{}
	if len(req.Pipeline) > 0 {
		for _, stage := range req.Pipeline {
			stageDoc, err := pipelineStageToBSON(stage)
			if err != nil {
				return status.Errorf(codes.InvalidArgument, "invalid pipeline stage: %v", err)
			}
			pipeline = append(pipeline, stageDoc)
		}
	}

	// Build options
	opts := options.ChangeStream()
	// FORCE UpdateLookup to ensure we have data for rules
	opts.SetFullDocument(options.UpdateLookup)

	if req.Options != nil {
		if req.Options.BatchSize > 0 {
			opts.SetBatchSize(req.Options.BatchSize)
		}
		if len(req.Options.ResumeAfter) > 0 {
			var resumeToken bson.Raw
			if err := bson.Unmarshal(req.Options.ResumeAfter, &resumeToken); err == nil {
				opts.SetResumeAfter(resumeToken)
			}
		}
	}

	// FORCE UpdateLookup to ensure we have data for rules
	opts.SetFullDocument(options.UpdateLookup)

	// Start change stream
	changeStream, err := coll.Watch(ctx, pipeline, opts)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to start change stream: %v", err)
	}
	defer changeStream.Close(ctx)

	// Stream changes
	for changeStream.Next(ctx) {
		var change bson.M
		if err := changeStream.Decode(&change); err != nil {
			return status.Errorf(codes.Internal, "failed to decode change: %v", err)
		}

		// Security Check: Evaluate rules
		var op rules.Operation = rules.OpRead
		var resource map[string]interface{}

		// Use fullDocument as resource data
		if fullDoc, ok := change["fullDocument"].(bson.M); ok {
			resource = fullDoc
		}

		// Handle Deletes: Deletes might not have fullDocument.
		// For MVP: If no fullDocument (delete), we allow if rule matches ID or doesn't check 'resource.data'
		if opType, ok := change["operationType"].(string); ok && opType == "delete" {
			// Extract document key for delete check if needed
			if docKey, ok := change["documentKey"].(bson.M); ok {
				if id, ok := docKey["_id"].(bson.ObjectID); ok {
					if resource == nil {
						resource = make(map[string]interface{})
					}
					resource["_id"] = id.Hex()
				}
			}
		}

		ruleReq := &rules.Request{
			Operation:    op,
			Database:     req.Database,
			Collection:   req.Collection,
			ExistingData: resource,
		}

		if s.rules != nil {
			allowed, err := s.rules.Evaluate(ctx, ruleReq)
			if err != nil {
				slog.Error("Rule evaluation error", "error", err)
				// Fail shut: if evaluation errors, deny access
				continue
			}
			if !allowed {
				// Silently drop unauthorized event
				continue
			}
		}

		// Build change event
		event := &mongorpcv1.ChangeEvent{
			Database:   req.Database,
			Collection: req.Collection,
		}

		// Extract operation type
		if opType, ok := change["operationType"].(string); ok {
			switch opType {
			case "insert":
				event.OperationType = mongorpcv1.ChangeEventType_INSERT
			case "update":
				event.OperationType = mongorpcv1.ChangeEventType_UPDATE
			case "replace":
				event.OperationType = mongorpcv1.ChangeEventType_REPLACE
			case "delete":
				event.OperationType = mongorpcv1.ChangeEventType_DELETE
			case "drop":
				event.OperationType = mongorpcv1.ChangeEventType_DROP
			case "invalidate":
				event.OperationType = mongorpcv1.ChangeEventType_INVALIDATE
			}
		}

		// Extract full document if present
		if fullDoc, ok := change["fullDocument"].(bson.M); ok {
			doc, _ := bsonToDocument(fullDoc)
			event.FullDocument = doc
		}

		// Extract document key
		if docKey, ok := change["documentKey"].(bson.M); ok {
			if id, ok := docKey["_id"].(bson.ObjectID); ok {
				event.DocumentKey = &mongorpcv1.ObjectId{Hex: id.Hex()}
			}
		}

		// Build response
		resp := &mongorpcv1.WatchResponse{
			Event: event,
		}

		// Extract resume token
		if token := changeStream.ResumeToken(); token != nil {
			tokenBytes, _ := bson.Marshal(token)
			resp.ResumeToken = tokenBytes
		}

		if err := stream.Send(resp); err != nil {
			return status.Errorf(codes.Internal, "failed to send change: %v", err)
		}
	}

	if err := changeStream.Err(); err != nil {
		return status.Errorf(codes.Internal, "change stream error: %v", err)
	}

	return nil
}

// BatchGetDocuments gets multiple documents by ID with streaming.
func (s *Server) BatchGetDocuments(req *mongorpcv1.BatchGetDocumentsRequest, stream mongorpcv1.MongoRPC_BatchGetDocumentsServer) error {
	if req.Database == "" || req.Collection == "" {
		return status.Error(codes.InvalidArgument, "database and collection are required")
	}
	if len(req.Ids) == 0 {
		return status.Error(codes.InvalidArgument, "at least one id is required")
	}

	ctx := stream.Context()
	slog.Debug("BatchGetDocuments", "database", req.Database, "collection", req.Collection, "count", len(req.Ids))

	// Get collection
	coll := s.db.Client().Database(req.Database).Collection(req.Collection)

	// Process each ID
	for _, id := range req.Ids {
		if id == nil || id.Hex == "" {
			continue
		}

		objectID, err := bson.ObjectIDFromHex(id.Hex)
		if err != nil {
			// Send as missing
			if err := stream.Send(&mongorpcv1.BatchGetDocumentsResponse{
				Result: &mongorpcv1.BatchGetDocumentsResponse_Missing{Missing: id},
			}); err != nil {
				return status.Errorf(codes.Internal, "failed to send response: %v", err)
			}
			continue
		}

		filter := bson.D{{Key: "_id", Value: objectID}}
		var result bson.M
		err = coll.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			// Send as missing
			if err := stream.Send(&mongorpcv1.BatchGetDocumentsResponse{
				Result: &mongorpcv1.BatchGetDocumentsResponse_Missing{Missing: id},
			}); err != nil {
				return status.Errorf(codes.Internal, "failed to send response: %v", err)
			}
			continue
		}

		doc, err := bsonToDocument(result)
		if err != nil {
			continue
		}

		if err := stream.Send(&mongorpcv1.BatchGetDocumentsResponse{
			Result: &mongorpcv1.BatchGetDocumentsResponse_Found{Found: doc},
		}); err != nil {
			return status.Errorf(codes.Internal, "failed to send response: %v", err)
		}
	}

	return nil
}

// aggregationPipelineToBSON converts pipeline stages to BSON array.
func aggregationPipelineToBSON(stages []*mongorpcv1.PipelineStage) (bson.A, error) {
	pipeline := make(bson.A, 0, len(stages))

	for _, stage := range stages {
		stageDoc, err := pipelineStageToBSON(stage)
		if err != nil {
			return nil, err
		}
		pipeline = append(pipeline, stageDoc)
	}

	return pipeline, nil
}

// pipelineStageToBSON converts a single pipeline stage to BSON.
func pipelineStageToBSON(stage *mongorpcv1.PipelineStage) (bson.D, error) {
	if stage == nil {
		return bson.D{}, nil
	}

	switch s := stage.StageType.(type) {
	case *mongorpcv1.PipelineStage_Match:
		filter, err := filterToBSON(s.Match.Filter)
		if err != nil {
			return nil, err
		}
		return bson.D{{Key: "$match", Value: filter}}, nil

	case *mongorpcv1.PipelineStage_Limit:
		return bson.D{{Key: "$limit", Value: s.Limit.Limit}}, nil

	case *mongorpcv1.PipelineStage_Skip:
		return bson.D{{Key: "$skip", Value: s.Skip.Skip}}, nil

	case *mongorpcv1.PipelineStage_Sort:
		sortDoc := bson.D{}
		for _, order := range s.Sort.Sort {
			direction := 1
			if order.Direction == mongorpcv1.SortDirection_DESCENDING {
				direction = -1
			}
			sortDoc = append(sortDoc, bson.E{Key: order.Field, Value: direction})
		}
		return bson.D{{Key: "$sort", Value: sortDoc}}, nil

	case *mongorpcv1.PipelineStage_Count:
		return bson.D{{Key: "$count", Value: s.Count.Field}}, nil

	case *mongorpcv1.PipelineStage_Sample:
		return bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: s.Sample.Size}}}}, nil

	case *mongorpcv1.PipelineStage_Unwind:
		return bson.D{{Key: "$unwind", Value: s.Unwind.Path}}, nil

	case *mongorpcv1.PipelineStage_Group:
		groupDoc := bson.D{}
		if s.Group.Id != nil {
			idVal, err := aggregationExpressionToBSON(s.Group.Id)
			if err != nil {
				return nil, err
			}
			groupDoc = append(groupDoc, bson.E{Key: "_id", Value: idVal})
		} else {
			groupDoc = append(groupDoc, bson.E{Key: "_id", Value: nil})
		}
		for field, acc := range s.Group.Accumulators {
			accVal, err := accumulatorToBSON(acc)
			if err != nil {
				return nil, err
			}
			groupDoc = append(groupDoc, bson.E{Key: field, Value: accVal})
		}
		return bson.D{{Key: "$group", Value: groupDoc}}, nil

	case *mongorpcv1.PipelineStage_Project:
		projectDoc := bson.D{}
		for field, spec := range s.Project.Fields {
			if spec.GetInclude() {
				projectDoc = append(projectDoc, bson.E{Key: field, Value: 1})
			} else if spec.GetExpression() != nil {
				exprVal, err := aggregationExpressionToBSON(spec.GetExpression())
				if err != nil {
					return nil, err
				}
				projectDoc = append(projectDoc, bson.E{Key: field, Value: exprVal})
			} else if spec.GetLiteral() != nil {
				litVal, err := valueToBSON(spec.GetLiteral())
				if err != nil {
					return nil, err
				}
				projectDoc = append(projectDoc, bson.E{Key: field, Value: litVal})
			}
		}
		return bson.D{{Key: "$project", Value: projectDoc}}, nil

	case *mongorpcv1.PipelineStage_Lookup:
		lookupDoc := bson.D{
			{Key: "from", Value: s.Lookup.From},
			{Key: "localField", Value: s.Lookup.LocalField},
			{Key: "foreignField", Value: s.Lookup.ForeignField},
			{Key: "as", Value: s.Lookup.As},
		}
		return bson.D{{Key: "$lookup", Value: lookupDoc}}, nil

	case *mongorpcv1.PipelineStage_AddFields:
		addFieldsDoc := bson.D{}
		for field, expr := range s.AddFields.Fields {
			exprVal, err := aggregationExpressionToBSON(expr)
			if err != nil {
				return nil, err
			}
			addFieldsDoc = append(addFieldsDoc, bson.E{Key: field, Value: exprVal})
		}
		return bson.D{{Key: "$addFields", Value: addFieldsDoc}}, nil

	case *mongorpcv1.PipelineStage_ReplaceRoot:
		newRootVal, err := aggregationExpressionToBSON(s.ReplaceRoot.NewRoot)
		if err != nil {
			return nil, err
		}
		return bson.D{{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: newRootVal}}}}, nil

	case *mongorpcv1.PipelineStage_Bucket:
		bucketDoc := bson.D{}
		if s.Bucket.GroupBy != nil {
			groupByVal, err := aggregationExpressionToBSON(s.Bucket.GroupBy)
			if err != nil {
				return nil, err
			}
			bucketDoc = append(bucketDoc, bson.E{Key: "groupBy", Value: groupByVal})
		}
		if len(s.Bucket.Boundaries) > 0 {
			boundaries := make(bson.A, len(s.Bucket.Boundaries))
			for i, b := range s.Bucket.Boundaries {
				bVal, err := valueToBSON(b)
				if err != nil {
					return nil, err
				}
				boundaries[i] = bVal
			}
			bucketDoc = append(bucketDoc, bson.E{Key: "boundaries", Value: boundaries})
		}
		if s.Bucket.Default != nil {
			defVal, err := valueToBSON(s.Bucket.Default)
			if err != nil {
				return nil, err
			}
			bucketDoc = append(bucketDoc, bson.E{Key: "default", Value: defVal})
		}
		return bson.D{{Key: "$bucket", Value: bucketDoc}}, nil

	case *mongorpcv1.PipelineStage_Out:
		return bson.D{{Key: "$out", Value: s.Out.Collection}}, nil

	case *mongorpcv1.PipelineStage_Merge:
		mergeDoc := bson.D{
			{Key: "into", Value: s.Merge.IntoCollection},
		}
		if len(s.Merge.On) > 0 {
			mergeDoc = append(mergeDoc, bson.E{Key: "on", Value: s.Merge.On})
		}
		return bson.D{{Key: "$merge", Value: mergeDoc}}, nil

	case *mongorpcv1.PipelineStage_Raw:
		// Raw BSON stage
		rawDoc := bson.D{}
		for k, v := range s.Raw.Fields {
			bsonVal, err := valueToBSON(v)
			if err != nil {
				return nil, err
			}
			rawDoc = append(rawDoc, bson.E{Key: k, Value: bsonVal})
		}
		return rawDoc, nil

	default:
		return bson.D{}, nil
	}
}

// aggregationExpressionToBSON converts an AggregationExpression to BSON.
func aggregationExpressionToBSON(expr *mongorpcv1.AggregationExpression) (interface{}, error) {
	if expr == nil {
		return nil, nil
	}

	// Check which type of expression is set
	if fieldRef := expr.GetFieldRef(); fieldRef != "" {
		return "$" + fieldRef, nil
	}
	if literal := expr.GetLiteral(); literal != nil {
		return valueToBSON(literal)
	}
	if variable := expr.GetVariable(); variable != "" {
		return "$$" + variable, nil
	}
	if op := expr.GetOperator(); op != nil {
		return operatorExpressionToBSON(op)
	}

	return nil, nil
}

// operatorExpressionToBSON converts an OperatorExpression to BSON.
func operatorExpressionToBSON(op *mongorpcv1.OperatorExpression) (bson.D, error) {
	if op == nil {
		return bson.D{}, nil
	}

	args := make(bson.A, len(op.Args))
	for i, arg := range op.Args {
		argVal, err := aggregationExpressionToBSON(arg)
		if err != nil {
			return nil, err
		}
		args[i] = argVal
	}

	return bson.D{{Key: "$" + op.Op, Value: args}}, nil
}

// accumulatorToBSON converts an Accumulator to BSON.
func accumulatorToBSON(acc *mongorpcv1.Accumulator) (bson.D, error) {
	if acc == nil {
		return bson.D{}, nil
	}

	var operator string
	switch acc.Type {
	case mongorpcv1.Accumulator_SUM:
		operator = "$sum"
	case mongorpcv1.Accumulator_AVG:
		operator = "$avg"
	case mongorpcv1.Accumulator_MIN:
		operator = "$min"
	case mongorpcv1.Accumulator_MAX:
		operator = "$max"
	case mongorpcv1.Accumulator_FIRST:
		operator = "$first"
	case mongorpcv1.Accumulator_LAST:
		operator = "$last"
	case mongorpcv1.Accumulator_PUSH:
		operator = "$push"
	case mongorpcv1.Accumulator_ADD_TO_SET:
		operator = "$addToSet"
	case mongorpcv1.Accumulator_COUNT:
		return bson.D{{Key: "$count", Value: bson.D{}}}, nil
	default:
		operator = "$sum"
	}

	exprVal, err := aggregationExpressionToBSON(acc.Expression)
	if err != nil {
		return nil, err
	}

	return bson.D{{Key: operator, Value: exprVal}}, nil
}
