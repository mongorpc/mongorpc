# mongorpc

MongoDB + gRPC = mongorpc

mongorpc is a server that can be used to communicate with **MongoDB** using **gRPC**. mongorpc can be used with **microservices**, or from any **mobile** or **web** applications. you can extend mongorpc with custom interceptors to add **authentication**, **logging**, **casbin**, **tracing**, etc.

mongorpc has many client libraries for different langauges and platforms. that provides a **ORM** **like** **interface** to interact with MongoDB.

**mongorpc is in really early development, and things may change before we release version 1.0.**

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

- **Dart** https://github.com/mongorpc/mongorpc-dart

```dart
import 'package:mongorpc/mongorpc.dart';

Future<void> main(List<String> args) async {
  var options = ChannelOptions(credentials: ChannelCredentials.insecure());
  var client = await mongorpc("localhost", port: 1203, options: options);

  var database = client.database("sample_mflix");
  var collection = database.collection("movies");

  var documents = await collection
      .documents()
      .limit(10)
      .sort(by: "title")
      .where("title", isEqualTo: "Blacksmith Scene")
      .get();
  print(documents);
}
```


---
**Deploy**
[Run on Google Cloud](https://deploy.cloud.run)
