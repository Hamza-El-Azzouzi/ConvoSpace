package services

import (
	"fmt"
	"forum/internal/models"
	"forum/internal/repositories"

	"github.com/gofrs/uuid/v5"
)

type PostService struct {
	PostRepo *repositories.PostRepository
	CategotyRepo *repositories.CategoryRepository
}

func (p *PostService) PostSave(userId uuid.UUID, title, content string, category []string) error {
	postId := uuid.Must(uuid.NewV4())
	post := &models.Post{
		ID:      postId,
		UserID:  userId,
		Title:   title,
		Content: content,
	}
	for _,id :=  range category{
		postCategory := &models.PostCategory{
			PostID: postId,
			CategoryID: id,
		}
		err := p.PostRepo.PostCatgorie(postCategory)
		if err != nil{
			return fmt.Errorf("error F categorie : %v ",err)
		}
	}

	return p.PostRepo.Create(post) 
}
func (p *PostService)  AllPosts ()([]models.PostWithUser ,error){
	posts , err := p.PostRepo.AllPosts()
	if err != nil {
		return nil,fmt.Errorf("error Kayn f All Post service")
	}
	return posts,nil
}