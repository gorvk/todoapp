package types

type Todo struct {
	Id          string `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	IsCompleted bool   `db:"is_completed" json:"isCompleted"`
	UserId      string `db:"user_id" json:"-"`
}
