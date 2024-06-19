/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("b3f36dkenb4oaup");

  return dao.deleteCollection(collection);
}, (db) => {
  const collection = new Collection({
    "id": "b3f36dkenb4oaup",
    "created": "2024-05-19 15:36:41.424Z",
    "updated": "2024-05-19 17:07:39.044Z",
    "name": "carts_view",
    "type": "view",
    "system": false,
    "schema": [
      {
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
      },
      {
        "system": false,
        "id": "a8czdk4r",
        "name": "ordered",
        "type": "bool",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {}
      },
      {
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
      },
      {
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
      }
    ],
    "indexes": [],
    "listRule": "user = @request.auth.id",
    "viewRule": "user = @request.auth.id",
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {
      "query": "SELECT id, user, ordered, paymentMethod, discount, created, updated FROM carts\n  WHERE ordered = true;"
    }
  });

  return Dao(db).saveCollection(collection);
})
