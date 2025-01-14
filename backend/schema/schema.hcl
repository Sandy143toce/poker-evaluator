schema "public" {
}
table "game_results" {
  schema = schema.public
  column "id" {
    type = serial
  }

  column "hand" {
    type = varchar(50)
    null = false
  }

  column "hand_rank" {
    type = integer
    null = false
  }

  column "cards" {
    type = sql("text[]")
    null = false
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
  }
}