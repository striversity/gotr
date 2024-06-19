/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "7nqr3k9mzmc0jlr",
    "created": "2024-05-21 13:54:58.592Z",
    "updated": "2024-05-21 13:54:58.592Z",
    "name": "carts_view",
    "type": "view",
    "system": false,
    "schema": [
      {
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
      },
      {
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
      },
      {
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
      }
    ],
    "indexes": [],
    "listRule": "user = @request.auth.id",
    "viewRule": "user = @request.auth.id",
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {
      "query": "SELECT id, user, discount, paymentMethod, created, updated FROM carts;"
    }
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("7nqr3k9mzmc0jlr");

  return dao.deleteCollection(collection);
})
