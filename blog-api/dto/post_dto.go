package dto

type CreatePostRequest struct {
	Title          string `json:"title"`
	Content        string `json:"content"`
	IsPublished    bool   `json:"is_published"`
	AuthorID       uint   `json:"author_id"`
	PostCategoryID uint   `json:"post_category_id"`
}

type UpdatePostRequest struct {
	Title          string `json:"title"`
	Content        string `json:"content"`
	IsPublished    bool   `json:"is_published"`
	PostCategoryID uint   `json:"post_category_id"`
}

type PostResponse struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	IsPublished  bool   `json:"is_published"`
	Author       string `json:"author"`
	PostCategory string `json:"post_category"`
}
