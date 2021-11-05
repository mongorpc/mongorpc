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
- [x] Authendication & Permission
    - [x] JWT Verification Interceptor
    - [x] Casbin Enforcer
    - [ ] Add Casbin Policy RPC Routes
    - [ ] Custom JWT Generation
    - [ ] Admin Authendication Using api keys
    
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
