package blogs

import (
	"context"
	_blogDomain "otoklix-blog/business/blogs"

	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(database *gorm.DB) _blogDomain.Repository {
	return BlogRepository{
		db: database,
	}
}

func (repo *BlogRepository) GetData(ctx context.Context) ([]_blogDomain.Domain, error) {
	blog := []Blog{}
	result := repo.db.Find(&blog)
	if result.Error != nil {
		return []_blogDomain.Domain{}, result.Error
	}
	return ToListDomain(blog), nil
}
func (repo *BlogRepository) GetById(ctx context.Context, id uint) (_blogDomain.Domain, error) {
	blog := Blog{}
	result := repo.db.Find(&blog, id)
	if result.Error != nil {
		return _blogDomain.Domain{}, result.Error
	}
	return ToDomain(blog), nil
}
func (repo *BlogRepository) Create(ctx context.Context, data _blogDomain.Domain) (_blogDomain.Domain, error) {
	blog := Blog{
		Title:       data.Title,
		Content:     data.Content,
		PublishedAt: data.Published_at,
	}
	result := repo.db.Create(blog)
	if result.Error != nil {
		return _blogDomain.Domain{}, result.Error
	}
	return ToDomain(blog), nil
}
func (repo *BlogRepository) Update(ctx context.Context, data _blogDomain.Domain) (_blogDomain.Domain, error) {
	blog := Blog{}

	// result := repo.db.Model(&blog).Update(Blog{Title: data.Title, Content: data.Content, PublishedAt: data.Published_at})
	// if result.Error != nil {
	// 	return _blogDomain.Domain{}, result.Error
	// }
	repo.db.Find(&blog, data.Id)
	return ToDomain(blog), nil
}
func (repo *BlogRepository) Delete(ctx context.Context, id uint) (_blogDomain.Domain, error) {
	var blog _blogDomain.Domain
	repo.db.Find(&blog, id)
	result := repo.db.Delete(&_blogDomain.Domain{}, id)
	if result.Error != nil {
		return _blogDomain.Domain{}, result.Error
	}
	return blog, nil
}