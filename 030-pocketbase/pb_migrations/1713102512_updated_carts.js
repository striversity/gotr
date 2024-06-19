/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("t6x1jklzqf4kt43")

  collection.listRule = "@request.auth.verified = true"
  collection.viewRule = "@request.auth.verified = true"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("t6x1jklzqf4kt43")

  collection.listRule = ""
  collection.viewRule = ""

  return dao.saveCollection(collection)
})
