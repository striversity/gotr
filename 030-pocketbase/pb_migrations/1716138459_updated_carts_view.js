/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("b3f36dkenb4oaup")

  collection.options = {
    "query": "SELECT id, user, ordered, paymentMethod, discount, created, updated FROM carts\n  WHERE ordered = true;"
  }

  // remove
  collection.schema.removeField("3akaioic")

  // remove
  collection.schema.removeField("qambhezv")

  // remove
  collection.schema.removeField("8bcu3lpj")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "cqmhbr8i",
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
    "id": "a8czdk4r",
    "name": "ordered",
    "type": "bool",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {}
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "tloiotj0",
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
    "id": "ktkpnbl2",
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

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("b3f36dkenb4oaup")

  collection.options = {
    "query": "SELECT id, user, ordered, paymentMethod, created, updated FROM carts;"
  }

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "3akaioic",
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
    "id": "qambhezv",
    "name": "ordered",
    "type": "bool",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {}
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "8bcu3lpj",
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
  collection.schema.removeField("cqmhbr8i")

  // remove
  collection.schema.removeField("a8czdk4r")

  // remove
  collection.schema.removeField("tloiotj0")

  // remove
  collection.schema.removeField("ktkpnbl2")

  return dao.saveCollection(collection)
})
