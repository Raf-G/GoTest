package domain

type ReviewsStorage interface {
	GetReview(int) (Review, error)
	GetReviewsProduct(int) ([]Review, error)
}

type ReviewsService interface {
	GetOneReview(int) (Review, error)
	GetAllReviewsProduct(int) ([]Review, error)
}

type Review struct {
	ID          int
	UserID      int
	ProductID   int
	Description string
	Grade       int
}
