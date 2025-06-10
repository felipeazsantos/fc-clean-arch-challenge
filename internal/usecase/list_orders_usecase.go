package usecase

type ListOrdersInput struct {

}

type ListOrdersOutput struct {}

type ListOrdersUseCase struct {

}

func NewListOrdersUseCase() *ListOrdersUseCase {
	return &ListOrdersUseCase{}
}

func (u *ListOrdersUseCase) Execute(input ListOrdersInput) (ListOrdersOutput, error) {
	
	return ListOrdersOutput{}, nil
}