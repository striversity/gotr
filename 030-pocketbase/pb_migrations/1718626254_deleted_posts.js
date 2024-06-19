/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("95gkt4f4xdfyfvw");

  return dao.deleteCollection(collection);
}, (db) => {
  const collection = new Collection({
    "id": "95gkt4f4xdfyfvw",
    "created": "2024-06-15 10:35:18.162Z",
    "updated": "2024-06-15 10:37:10.541Z",
    "name": "posts",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "zf4w951k",
        "name": "title",
        "type": "text",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "qcwiitz4",
        "name": "images",
        "type": "file",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "mimeTypes": [],
          "thumbs": [],
          "maxSelect": 99,
          "maxSize": 5242880,
          "protected": false
        }
      },
      {
        "system": false,
        "id": "xujrwnd0",
        "name": "body",
        "type": "editor",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "convertUrls": false
        }
      }
    ],
    "indexes": [
      "CREATE INDEX `idx_ei501K7` ON `posts` (\n  `title`,\n  `body`\n)",
      "CREATE INDEX `idx_bQZ9R0n` ON `posts` (`body`)",
      "CREATE INDEX `idx_a1skTBA` ON `posts` (`title`)",
      "CREATE UNIQUE INDEX `idx_Tki8J2w` ON `posts` (`images`)"
    ],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
})
