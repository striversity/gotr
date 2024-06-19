/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("7nqr3k9mzmc0jlr")

  collection.options = {
    "query": "SELECT id, user, discount, paymentMethod, ordered, created, updated FROM carts;"
  }

  // remove
  collection.schema.removeField("eqww1liy")

  // remove
  collection.schema.removeField("dd2gsijo")

  // remove
  collection.schema.removeField("1phndavw")

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

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("7nqr3k9mzmc0jlr")

  collection.options = {
    "query": "SELECT id, user, discount, paymentMethod, created, updated FROM carts;"
  }

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "eqww1liy",
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
    "id": "dd2gsijo",
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
    "id": "1phndavw",
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

  // remove
  collection.schema.removeField("bxxdm82s")

  // remove
  collection.schema.removeField("vvziu1pj")

  // remove
  collection.schema.removeField("52yecctn")

  // remove
  collection.schema.removeField("i3nng661")

  return dao.saveCollection(collection)
})
