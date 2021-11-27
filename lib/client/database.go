package client

type Database struct {
	client *MongoRPCClient
	name   string
}

// Initiliaze a new database
func (client *MongoRPCClient) Database(name string) *Database {
	// Return a new database
	return &Database{
		name:   name,
		client: client,
	}
}
