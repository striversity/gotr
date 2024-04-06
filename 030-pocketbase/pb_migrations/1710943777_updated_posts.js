/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("x1qvi72e55d13uo")

  collection.listRule = "@request.auth.verified = true"
  collection.viewRule = "@request.auth.verified = true"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("x1qvi72e55d13uo")

  collection.listRule = "id = @request.auth.id"
  collection.viewRule = "id = @request.auth.id"

  return dao.saveCollection(collection)
})
