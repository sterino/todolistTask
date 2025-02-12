package task

import "time"

type Entity struct {
	ID          string    `db:"id" bson:"_id"`
	Title       *string   `db:"title" bson:"title"`
	Description *string   `db:"description" bson:"description"`
	Status      *string   `db:"status" bson:"status"`
	Priority    *string   `db:"priority" bson:"priority"`
	StartDate   time.Time `db:"start_date" bson:"start_date"`
	EndDate     time.Time `db:"end_date" bson:"end_date"`
}
