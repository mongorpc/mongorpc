package client

type Collection struct {
	client *MongoRPCClient
	name   string
	parent *Database
}

// Initilize a new collection
func (db *Database) Collection(name string) *Collection {

	// Create a new collection
	return &Collection{
		name:   name,
		client: db.client,
		parent: db,
	}
}
