package product

type ProductUseCase struct {
	repo *ProductRepository
}

func NewProductUseCase(repo *ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (u *ProductUseCase) GetProduct(id int) (*Product, error) {
	return u.repo.GetProduct(id)
}
