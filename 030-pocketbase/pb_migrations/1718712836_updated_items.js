/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("rhbw5bb5l6zm3cl")

  collection.indexes = [
    "CREATE INDEX `idx_QyXKJPq` ON `items` (`name`)",
    "CREATE INDEX `idx_ecQufUv` ON `items` (`description`)"
  ]

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("rhbw5bb5l6zm3cl")

  collection.indexes = []

  return dao.saveCollection(collection)
})
