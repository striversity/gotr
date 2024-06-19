/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("95gkt4f4xdfyfvw")

  collection.indexes = [
    "CREATE INDEX `idx_ei501K7` ON `posts` (\n  `title`,\n  `body`\n)",
    "CREATE INDEX `idx_bQZ9R0n` ON `posts` (`body`)",
    "CREATE INDEX `idx_a1skTBA` ON `posts` (`title`)",
    "CREATE UNIQUE INDEX `idx_Tki8J2w` ON `posts` (`images`)"
  ]

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("95gkt4f4xdfyfvw")

  collection.indexes = []

  return dao.saveCollection(collection)
})
