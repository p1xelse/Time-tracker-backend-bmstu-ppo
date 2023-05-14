package usecase

import (
	"encoding/json"
	"fmt"
	"strconv"
	projectRep "timetracker/internal/Project/repository"
	"timetracker/internal/cache"
	"timetracker/models"

	"github.com/pkg/errors"
)

type UsecaseI interface {
	CreateProject(e *models.Project) error
	UpdateProject(e *models.Project) error
	GetProject(id uint64) (*models.Project, error)
	DeleteProject(id uint64, userID uint64) error
	GetUserProjects(userID uint64) ([]*models.Project, error)
}

type usecase struct {
	projectRepository projectRep.RepositoryI
	redisStorage      *cache.StorageRedis
}

func New(pRep projectRep.RepositoryI) UsecaseI {
	return &usecase{
		projectRepository: pRep,
	}
}

func (u *usecase) CreateProject(e *models.Project) error {
	err := u.projectRepository.CreateProject(e)

	if err != nil {
		return errors.Wrap(err, "Error in func project.Usecase.CreateProject")
	}

	return nil
}

func (u *usecase) UpdateProject(p *models.Project) error {
	_, err := u.projectRepository.GetProject(p.ID)

	if err != nil {
		return errors.Wrap(err, "Error in func project.Usecase.Update.GetProject")
	}

	err = u.projectRepository.UpdateProject(p)

	if err != nil {
		return errors.Wrap(err, "Error in func project.Usecase.Update")
	}

	return nil
}

func (u *usecase) GetProject(id uint64) (*models.Project, error) {
	resProject, err := u.projectRepository.GetProject(id)

	if err != nil {
		return nil, errors.Wrap(err, "project.usecase.GetProject error while get project info")
	}

	return resProject, nil
}

func (u *usecase) DeleteProject(id uint64, userID uint64) error {
	existedProject, err := u.projectRepository.GetProject(id)
	if err != nil {
		return err
	}

	if existedProject == nil {
		return errors.New("Project not found") //TODO models error
	}

	if *existedProject.UserID != userID {
		return errors.New("Permission denied")
	}

	err = u.projectRepository.DeleteProject(id)

	if err != nil {
		return errors.Wrap(err, "Error in func project.Usecase.DeleteProject repository")
	}

	return nil
}

func (u *usecase) GetUserProjects(userID uint64) ([]*models.Project, error) {
	strUserID := strconv.FormatUint(userID, 10)
	dataFromCache, err := u.redisStorage.Get(strUserID)
	if err != nil {
		return nil, err
	}

	if dataFromCache != nil {
		res := make([]*models.Project, 0, 10)
		err = json.Unmarshal(dataFromCache, &res)
		if err != nil {
			return nil, fmt.Errorf("fail to unmarshal similar ids: %w", err)
		}
		fmt.Println("from cache")
		return res, nil
	}

	projects, err := u.projectRepository.GetUserProjects(userID)

	if err != nil {
		return nil, errors.Wrap(err, "Error in func project.Usecase.GetUserPosts")
	}

	err = u.redisStorage.Set(strUserID, projects)
	if err != nil {
		return nil, fmt.Errorf("can not set data to redis cache")
	}

	return projects, nil
}
