---
title: MongoDB
icon: fa-leaf
primary: "#47A248"
lang: javascript
locale: zhs
---

## fa-pen-to-square CRUD 操作

```javascript
db.users.insertOne({ name: "Alice", age: 30 });
db.users.insertMany([
  { name: "Bob", age: 25 },
  { name: "Charlie", age: 35 }
]);

db.users.findOne({ name: "Alice" });
db.users.find({ age: { $gte: 25 } }).limit(10).skip(5);

db.users.updateOne({ name: "Alice" }, { $set: { age: 31 } });
db.users.updateMany({}, { $inc: { age: 1 } });

db.users.deleteOne({ name: "Charlie" });
db.users.deleteMany({ age: { $lt: 20 } });
```

## fa-filter 查询操作符

```javascript
db.orders.find({ status: { $in: ["active", "pending"] } });
db.orders.find({ $or: [{ status: "active" }, { total: { $gt: 1000 } }] });
db.orders.find({ $and: [{ age: { $gte: 18 } }, { age: { $lte: 65 } }] });

db.products.find({ tags: { $all: ["electronics", "sale"] } });
db.products.find({ qty: { $exists: true } });
db.users.find({ name: { $regex: "^Al", $options: "i" } });

db.items.find({ field: { $type: "string" } });
db.items.find({ $expr: { $gt: ["$budget", "$spent"] } });
```

## fa-columns 投影

```javascript
db.users.find({}, { name: 1, email: 1, _id: 0 });
db.users.find({}, { password: 0 });

db.posts.find({}, { "author.name": 1 });

db.books.find({}, { title: 1, $slice: ["comments", 5] });
db.books.find({}, { title: 1, $elemMatch: { "comments.score": { $gt: 8 } } });
```

## fa-wrench 更新操作符

```javascript
db.users.updateOne({ _id: 1 }, {
  $set: { status: "active" },
  $unset: { tempField: "" },
  $inc: { loginCount: 1 },
  $push: { tags: "new" },
  $addToSet: { roles: "admin" },
  $pull: { tags: "deprecated" },
  $rename: { oldName: "newName" }
});

db.users.updateOne({ _id: 1 }, { $push: { scores: { $each: [90, 85, 95], $sort: -1, $slice: 5 } } });
db.users.updateOne({ _id: 1 }, { $pop: { tags: 1 } });
db.users.updateOne({ _id: 1 }, { $mul: { price: 1.1 } });
```

## fa-layer-group 聚合管道

```javascript
db.orders.aggregate([
  { $match: { status: "completed" } },
  { $group: { _id: "$category", total: { $sum: "$amount" }, count: { $sum: 1 } } },
  { $sort: { total: -1 } },
  { $limit: 10 }
]);

db.orders.aggregate([
  { $lookup: { from: "users", localField: "userId", foreignField: "_id", as: "user" } },
  { $unwind: "$user" },
  { $project: { orderTotal: "$amount", userName: "$user.name" } }
]);

db.logs.aggregate([
  { $bucket: { groupBy: "$timestamp", boundaries: [0, 100, 200, 300], default: "other" } }
]);

db.events.aggregate([
  { $facet: {
    byStatus: [{ $group: { _id: "$status", count: { $sum: 1 } } }],
    recent: [{ $sort: { createdAt: -1 } }, { $limit: 5 }]
  }}
]);
```

## fa-magnifying-glass 索引

```javascript
db.users.createIndex({ email: 1 }, { unique: true });
db.users.createIndex({ name: 1, age: -1 });
db.users.createIndex({ location: "2dsphere" });
db.users.createIndex({ description: "text" });
db.users.createIndex({ createdAt: 1 }, { expireAfterSeconds: 3600 });

db.users.getIndexes();
db.users.dropIndex("email_1");
db.users.dropIndexes();

db.orders.createIndex({ status: 1, createdAt: -1 }, { partialFilterExpression: { status: "active" } });
db.users.createIndex({ email: 1 }, { unique: true, sparse: true });
```

## fa-font 文本搜索

```javascript
db.articles.createIndex({ title: "text", body: "text" });
db.articles.find({ $text: { $search: "mongodb indexing" } });
db.articles.find({ $text: { $search: "\"exact phrase\"" } });
db.articles.find({ $text: { $search: "mongodb -nosql" } });

db.articles.find(
  { $text: { $search: "mongodb" } },
  { score: { $meta: "textScore" } }
).sort({ score: { $meta: "textScore" } });

db.articles.createIndex({ "$**": "text" }, { name: "allTextIndex" });
```

## fa-location-dot 地理空间

```javascript
db.places.createIndex({ location: "2dsphere" });

db.places.insertOne({
  name: "Central Park",
  location: { type: "Point", coordinates: [-73.97, 40.77] }
});

db.places.find({
  location: {
    $near: {
      $geometry: { type: "Point", coordinates: [-73.97, 40.77] },
      $maxDistance: 1000
    }
  }
});

db.places.find({
  location: {
    $geoWithin: {
      $polygon: [[-73.98, 40.76], [-73.96, 40.76], [-73.96, 40.78], [-73.98, 40.78]]
    }
  }
});

db.places.find({
  location: { $geoIntersects: { $geometry: { type: "LineString", coordinates: [[-73.98, 40.76], [-73.96, 40.78]] } } }
});
```

## fa-clone 复制集

```javascript
rs.initiate({ _id: "rs0", members: [
  { _id: 0, host: "mongo1:27017" },
  { _id: 1, host: "mongo2:27017" },
  { _id: 2, host: "mongo3:27017" }
]});

rs.status();
rs.conf();
rs.add("mongo4:27017");
rs.remove("mongo4:27017");
rs.stepDown(60);

db.getMongo().setReadPref("secondary");
db.getMongo().setReadPref("secondaryPreferred");
```

## fa-grip 分片

```javascript
sh.enableSharding("mydb");
sh.shardCollection("mydb.users", { email: 1 });
sh.shardCollection("mydb.logs", { timestamp: 1 });

sh.addShard("rs1/mongo1:27017,mongo2:27017");
sh.status();

db.adminCommand({ balancerStart: 1 });
db.adminCommand({ balancerStop: 1 });
db.adminCommand({ balancerStatus: 1 });

sh.moveChunk("mydb.users", { email: "m" }, "shard002");
db.adminCommand({ splitVector: "mydb.users", find: { email: "m" } });
```

## fa-arrows-rotate 事务

```javascript
const session = db.getMongo().startSession();
session.startTransaction();
try {
  db.accounts.updateOne({ _id: 1 }, { $inc: { balance: -100 } }, { session });
  db.accounts.updateOne({ _id: 2 }, { $inc: { balance: 100 } }, { session });
  session.commitTransaction();
} catch (e) {
  session.abortTransaction();
} finally {
  session.endSession();
}

session.startTransaction({ readConcern: "snapshot", writeConcern: { w: "majority" } });
```

## fa-tower-broadcast 变更流

```javascript
const changeStream = db.users.watch();
changeStream.on("change", (next) => { printjson(next); });

db.orders.watch([{ $match: { "fullDocument.status": "shipped" } }], { fullDocument: "updateLookup" });

db.users.watch([], { resumeAfter: <resumeToken> });
db.users.watch([], { startAtOperationTime: Timestamp(1625097600, 1) });
```

## fa-shield-halved 模式验证

```javascript
db.createCollection("users", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["name", "email"],
      properties: {
        name: { bsonType: "string", minLength: 2, maxLength: 100 },
        email: { bsonType: "string", pattern: "^.+@.+$" },
        age: { bsonType: "int", minimum: 0, maximum: 150 },
        tags: { bsonType: "array", items: { bsonType: "string" } }
      }
    }
  },
  validationLevel: "strict",
  validationAction: "error"
});

db.runCommand({ collMod: "users", validator: { $jsonSchema: { bsonType: "object", required: ["name"] } } });
db.getCollectionInfos({ name: "users" })[0].options.validator;
```

## fa-terminal Shell 命令

```javascript
show dbs;
use mydb;
show collections;

db.stats();
db.users.stats();
db.users.countDocuments({ status: "active" });
db.users.estimatedDocumentCount();

db.users.distinct("city");
db.users.find().explain("executionStats");

db.copyDatabase("src", "dst");
db.dropDatabase();
db.users.renameCollection("customers");

db.adminCommand({ listDatabases: 1 });
db.serverStatus();
```
