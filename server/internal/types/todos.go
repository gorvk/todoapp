package types

type Todo struct {
	Id          string `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	IsCompleted bool   `db:"isCompleted" json:"isCompleted"`
}
