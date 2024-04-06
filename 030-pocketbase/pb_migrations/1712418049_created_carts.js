/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "t6x1jklzqf4kt43",
    "created": "2024-04-06 15:40:49.368Z",
    "updated": "2024-04-06 15:40:49.368Z",
    "name": "carts",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "iet0qhgb",
        "name": "user",
        "type": "relation",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "collectionId": "_pb_users_auth_",
          "cascadeDelete": false,
          "minSelect": null,
          "maxSelect": 1,
          "displayFields": null
        }
      },
      {
        "system": false,
        "id": "cjae9tgy",
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
      },
      {
        "system": false,
        "id": "cqvkwfwh",
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
      },
      {
        "system": false,
        "id": "qjelogtg",
        "name": "Ordered",
        "type": "bool",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {}
      }
    ],
    "indexes": [],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("t6x1jklzqf4kt43");

  return dao.deleteCollection(collection);
})
