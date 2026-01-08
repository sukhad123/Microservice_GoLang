package products

import (
	"context"

	repo "github.com/learn_go/internal/adapters/postgresql.sqlc"
)

type Service interface {
	// Add your service methods here
	ListProducts(ctx context.Context) ([]repo.Product, error)
}

type svc struct {
	//repo
	repo *repo.Queries
}

func NewService(repo *repo.Queries) Service {
	return &svc{
		repo: repo,
	}
}
func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	// products, err := s.repo.ListProducts(ctx)
	// //Handle error
	// if err != nil {
	// 	fmt.Print("Error at somwehre in service go", err)
	// }
	//return products if done
	return s.repo.ListProducts(ctx)
}

func (s *svc) CreateProduct(ctx context.Context, data repo.CreateProductParams) (repo.Product, error) {
	//this function will take that service and return product
	return s.repo.CreateProduct(ctx, data)

}
