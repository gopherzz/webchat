package repository

func NewMemoryRepository() *Repository {
	return &Repository{
		Messages: NewMemoryMessages(),
	}
}
