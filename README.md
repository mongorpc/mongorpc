# mongorpc

MongoDB + gRPC = mongorpc

mongorpc is a server that can be used to communicate with **MongoDB** using **gRPC**. mongorpc can be used with **microservices**, or from any **mobile** or **web** applications. you can extend mongorpc with custom interceptors to add **authentication**, **logging**, **casbin**, **tracing**, etc.

mongorpc has many client libraries for different langauges and platforms. that provides a **ORM** **like** **interface** to interact with MongoDB.

mongorpc is in really early development, and things may change before we release version 1.0.

Clients
------



- **Web** https://github.com/mongorpc/mongorpc-js

```ts
const client = new MongoRPC("http://localhost:8080");

const document = await client
  .database("sample_mflix")
  .collection("movies")
  .document("573a13b0f29313caabd35231")
  .get();

console.log(document);

// Text Search Documents (MongoDB 4.0+) (https://docs.mongodb.com/manual/text-search/)
// Requires to create index on `title` field first
const documents = await client
  .database("sample_mflix")
  .collection("movies")
  .search("The Shawshank Redemption")
  .get();

console.log(documents);
```


- **Golang** https://github.com/mongorpc/mongorpc-go
```go
// Initilize database
db := client.Database("sample_mflix")

// Get Document By ID
doc, err := db.Collection("movies").Document("573a13b0f29313caabd35231").Get(context.TODO())
if err != nil {
  fmt.Println(err)
}
fmt.Println(doc)

// Search Documents
docs, err := db.Collection("movies").Documents().Search("Batman").Get(context.TODO())
if err != nil {
  fmt.Println(err)
}
fmt.Println(docs)

```


- **Swift** https://github.com/mongorpc/mongorpc-swift

```swift
import MongoRPC


let client = MongoRPC(host: "localhost", port: 27051)

client.database("sample_mflix").collection("movies").document(id: "573a13b0f29313caabd35231").get { result in

    switch result {
    case let .success(document):
        print(document)
    case let .failure(error):
        print(error.localizedDescription)
    }
}

```

**Deployment**
[Run on Google Cloud](https://deploy.cloud.run)
