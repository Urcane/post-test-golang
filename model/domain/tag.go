package domain

type Tag struct {
	ID    int
	Label string
	Posts []*Post //many to many
}
