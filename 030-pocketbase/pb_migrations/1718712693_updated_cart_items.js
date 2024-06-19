/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("979nwt5avzbw18j")

  collection.indexes = [
    "CREATE UNIQUE INDEX `idx_h96BOBV` ON `cart_items` (\n  `item`,\n  `cart`\n)"
  ]

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("979nwt5avzbw18j")

  collection.indexes = []

  return dao.saveCollection(collection)
})
