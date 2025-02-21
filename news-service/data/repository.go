package data

type Repository interface {
	GetNewsList(limit, offset int) ([]*NewsWithCategories, int, error)
	EditNews(id int, title, content string, categories []int) error
}
