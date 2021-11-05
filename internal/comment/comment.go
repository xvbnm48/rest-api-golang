package comment

import (
	"github.com/jinzhu/gorm"
)

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

// PostCOment - adds a new comment to the database
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

// update comment - updates comment by idd with new comment info
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}

	if result := s.DB.Model(&comment).Update(newComment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// deletecomment - delete comment from database by id
func (s *Service) DeleteComment(ID uint) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}

	return nil
}

// getALlComments - retrieves all coments from database
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}

	return comments, nil
}
