package models

//go:generate reform

//reform:news
type News struct {
	ID      int64  `reform:"id,pk"`
	Title   string `reform:"title"`
	Content string `reform:"content"`
}
