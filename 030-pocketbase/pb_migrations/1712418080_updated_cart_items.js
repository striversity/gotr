/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("979nwt5avzbw18j")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "nziconos",
    "name": "cart",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "t6x1jklzqf4kt43",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("979nwt5avzbw18j")

  // remove
  collection.schema.removeField("nziconos")

  return dao.saveCollection(collection)
})
