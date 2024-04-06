/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("uz29na52e5arxg5")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "vdp1o70g",
    "name": "price",
    "type": "number",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "noDecimal": false
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("uz29na52e5arxg5")

  // remove
  collection.schema.removeField("vdp1o70g")

  return dao.saveCollection(collection)
})
