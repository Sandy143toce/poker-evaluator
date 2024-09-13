schema "public" {
}
table "game_results" {
  schema = schema.public
  column "id" {
    null = false
    type = integer
    identity {
      generated = ALWAYS
    }
  }

  column "player_best_hand" {
    null = false
    type = jsonb
  }

  column "other_best_hands" {
    null = false
    type = jsonb
  }

  column "created_at" {
    type    = timestamp
    default = sql("CURRENT_TIMESTAMP")
    null    = false
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_game_results_created_at" {
    columns = [column.created_at]
    type    = "btree"
    order   = "DESC"
  }
}
