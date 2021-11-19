# mongorpc

MongoDB + gRPC = mongorpc

mongorpc is a server that can be used to communicate with **MongoDB** using **gRPC**. mongorpc can be used with **microservices**, or from any **mobile** or **web** applications. you can extend mongorpc with custom interceptors to add **authentication**, **logging**, **casbin**, **tracing**, etc.

mongorpc has many client libraries for different langauges and platforms. that provides a **ORM** **like** **interface** to interact with MongoDB.

in future there will be support for `Text Search`, `Geospatial Queries` and `Bulk Operations`

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
```


- **Golang** https://github.com/mongorpc/mongorpc-go
```go
// Initilize database
db := client.Database("sample_mflix")

// List Collections
collections, err := db.ListCollectionNames(context.TODO())
if err != nil {
  fmt.Println(err)
}
fmt.Println(collections)

// Get Document By ID
doc, err := db.Collection("movies").Document("573a13b0f29313caabd35231").Get(context.TODO())
if err != nil {
  fmt.Println(err)
}
fmt.Println(doc)

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


## 🚧 **Roadmap** 🚧


**mongorpc**

- [x] List Collections
- [x] Get Document
- [x] List Documents
- [x] Create Document
- [x] Update Document
- [x] Delete Document
- [x] Count Documents
- [x] Listen Collection Changes
    - [x] Return Changed Documents
    - [ ] Add Filter, Sort and Limit in Listen Requests
    - [ ] Return Whole Collection when some change has been done
- [x] Create Index
- [x] Get Indexes
- [x] Delete Index
- [ ] Reindex
- [x] Ping
- [x] Health Check
- [x] Collection Stats
- [x] Create Collection
- [x] Rename Collection
- [x] Delete Collection
- [ ] Bulk Insert
- [ ] Bulk Update
- [ ] Bulk Delete
- [ ] Bulk Replace
- [ ] **Text Search**
- [ ] **Geospatial Queries**
- [x] Middlewares (All Interceptor moved to it's seprate libraries)
    - [x] JWT/Auth Interceptor
    - [x] Oso Interceptor
    
**mongorpc client's**

- [x] Web (mongorpc-js)
- [ ] iOS (mongorpc-swift)
- [ ] Flutter (mongorpc-dart)
- [ ] Android (mongorpc-kotlin)
- [ ] Node.js (mongorpc-node)
- [ ] Go (mongorpc-go)

**mongorpc client's offline capability concept**
- [ ] Use Key/Value storage to store DocumentID as Key and Document as Value.
- [ ] Store Collection Name as Key and Document ID's Array as Value.
- [ ] Store Last Updated Cursor, When Last Cursor changed, do some sync operations.
