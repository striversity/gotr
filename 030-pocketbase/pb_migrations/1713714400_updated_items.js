/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("rhbw5bb5l6zm3cl")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "mhwswvms",
    "name": "images",
    "type": "file",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "mimeTypes": [
        "image/png",
        "image/jpeg",
        "image/gif"
      ],
      "thumbs": [],
      "maxSelect": 99,
      "maxSize": 5242880,
      "protected": false
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("rhbw5bb5l6zm3cl")

  // remove
  collection.schema.removeField("mhwswvms")

  return dao.saveCollection(collection)
})
