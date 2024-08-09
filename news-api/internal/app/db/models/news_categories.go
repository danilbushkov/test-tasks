package models

//go:generate reform

// reform:news_categories
type NewsCategories struct {
	NewID      int64 `reform:"news_id"`
	CategoryId int64 `reform:"category_id"`
}
