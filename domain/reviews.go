package domain

//go:generate mockgen -source=reviews.go -destination=mocks/reviews.go

type ReviewsStorage interface {
	AddReview(Review) (*Review, error)
	EditReview(Review) error
	GetReview(int) (Review, error)
	DeleteReview(int) (bool, error)
	GetReviewsProduct(int) ([]Review, error)
	GetReviews() ([]Review, error)
}

type ReviewsService interface {
	AddReview(Review) (Review, error)
	EditReview(Review) (Review, error)
	GetReview(int) (Review, error)
	DeleteReview(int) error
	GetAllReviewsProduct(int) ([]Review, error)
	GetReviews() ([]Review, error)
}

type Review struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	ProductID   int    `json:"product_id"`
	Description string `json:"description"`
	Grade       int    `json:"grade"`
}
