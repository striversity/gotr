/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("t6x1jklzqf4kt43")

  collection.listRule = "(user = @request.auth.id) && (ordered = false)"
  collection.viewRule = "(user = @request.auth.id) && (ordered = false)"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("t6x1jklzqf4kt43")

  collection.listRule = "user = @request.auth.id"
  collection.viewRule = "user = @request.auth.id"

  return dao.saveCollection(collection)
})
