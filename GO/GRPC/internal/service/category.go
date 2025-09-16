package service

import (
	context "context"
	"github.com/Alencar26/Estudo/GO/GRPC/internal/database"
	"github.com/Alencar26/Estudo/GO/GRPC/internal/pb"
)

type CategoyService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoyService {
	return &CategoyService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoyService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil

}
