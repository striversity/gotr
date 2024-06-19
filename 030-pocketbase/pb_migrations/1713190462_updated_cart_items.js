/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("979nwt5avzbw18j")

  collection.listRule = "@request.auth.id != \"\" && @request.auth.id = cart.user.id"
  collection.viewRule = "@request.auth.id != \"\" && @request.auth.id = cart.user.id"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("979nwt5avzbw18j")

  collection.listRule = "@request.auth.id != \"\""
  collection.viewRule = "@request.auth.id != \"\""

  return dao.saveCollection(collection)
})
