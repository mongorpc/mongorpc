# mongorpc

MongoDB + gRPC = mongorpc

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
    - [ ] Return Whole Collection when some change has been done
- [x] Create Index
- [x] Get Indexes
- [x] Delete Index
- [ ] Reindex
- [x] Ping
- [ ] Collection Stats
- [ ] Create Collection
- [ ] Rename Collection
- [ ] Delete Collection
- [ ] Bulk Insert
- [ ] Bulk Update
- [ ] Bulk Delete
- [ ] Bulk Replace
- [ ] Authendication (~jwt rpc middleware)
- [ ] Permission Based Collection Read & Write (~permission middleware) (~firestore roles)
    - [ ] Roles Compiler
    - [ ] Roles Validator and Enforcer

**mongorpc client's**

- [ ] Web (mongorpc-js)
- [ ] iOS (mongorpc-swift)
- [ ] Flutter (mongorpc-dart)
- [ ] Android (mongorpc-kotlin)
- [ ] Node.js (mongorpc-node)
- [ ] Go (mongorpc-go)

**mongorpc client's offline capability concept**
- [ ] Use Key/Value storage to store DocumentID as Key and Document as Value.
- [ ] Store Collection Name as Key and Document ID's Array as Value.
- [ ] Store Last Updated Cursor, When Last Cursor changed, do some sync operations.
