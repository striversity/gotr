/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("979nwt5avzbw18j")

  collection.name = "cartitems"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("979nwt5avzbw18j")

  collection.name = "cart_items"

  return dao.saveCollection(collection)
})
