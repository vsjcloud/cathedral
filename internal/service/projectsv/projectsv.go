package projectsv

import (
	"github.com/erathorus/quickstore"

	"cathedral/api/proto/model"
	"cathedral/internal/key"
)

type ProjectService struct {
	store *quickstore.Store
}

func (s *ProjectService) AddProject(project *model.Project) (quickstore.Key, error) {
	projectKey := key.GenerateProjectKey()
	err := s.store.Insert(projectKey, project)
	if err != nil {
		return quickstore.Key{}, err
	}
	return projectKey, nil
}
