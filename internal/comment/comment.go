package comment

import "github.com/jinzhu/gorm"

// service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}

//comment - defines our comment structure
type Comment struct {
	gorm.Model
	SLug   string
	Body   string
	Author string
}

// comment service - the interface for our comment
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// new service - return a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetComment - retrieves comments by their ID from the database
func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

//GetCommentBySlug - retrieves all comments by slug ( path - /article/name)
func (s *Service) GetCommentBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}

	return comments, nil
}

func (s *Service)