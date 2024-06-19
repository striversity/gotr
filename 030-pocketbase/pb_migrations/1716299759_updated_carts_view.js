/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("7nqr3k9mzmc0jlr")

  collection.options = {
    "query": "SELECT id, user, discount, paymentMethod, ordered, created, updated FROM carts\n  WHERE ordered = true;"
  }

  // remove
  collection.schema.removeField("bxxdm82s")

  // remove
  collection.schema.removeField("vvziu1pj")

  // remove
  collection.schema.removeField("52yecctn")

  // remove
  collection.schema.removeField("i3nng661")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "1fxvgngx",
    "name": "user",
    "type": "relation",
    "required": true,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "_pb_users_auth_",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": null
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "xl4ezr2k",
    "name": "discount",
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

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "xy5o5e3k",
    "name": "paymentMethod",
    "type": "select",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "Credit Card",
        "Cash",
        "Venmo",
        "PayPal",
        "Zelle"
      ]
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "3k02epfs",
    "name": "ordered",
    "type": "bool",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {}
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("7nqr3k9mzmc0jlr")

  collection.options = {
    "query": "SELECT id, user, discount, paymentMethod, ordered, created, updated FROM carts;"
  }

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "bxxdm82s",
    "name": "user",
    "type": "relation",
    "required": true,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "_pb_users_auth_",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": 1,
      "displayFields": null
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "vvziu1pj",
    "name": "discount",
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

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "52yecctn",
    "name": "paymentMethod",
    "type": "select",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "Credit Card",
        "Cash",
        "Venmo",
        "PayPal",
        "Zelle"
      ]
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "i3nng661",
    "name": "ordered",
    "type": "bool",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {}
  }))

  // remove
  collection.schema.removeField("1fxvgngx")

  // remove
  collection.schema.removeField("xl4ezr2k")

  // remove
  collection.schema.removeField("xy5o5e3k")

  // remove
  collection.schema.removeField("3k02epfs")

  return dao.saveCollection(collection)
})
