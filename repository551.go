package repository551

type Repository struct{}

var repositoryInstance *Repository

func Load() *Repository {
	if repositoryInstance != nil {
		return repositoryInstance
	}

	repositoryInstance = &Repository{}

	return repositoryInstance
}
