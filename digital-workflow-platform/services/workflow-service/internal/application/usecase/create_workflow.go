package usecase

import (
	"workflow-service/internal/domain/entity"
	"workflow-service/internal/domain/repository"
)

type CreateWorkflowUsecase struct {
	repo repository.WorkflowRepository
}

func NewCreateWorkflowUsecase(r repository.WorkflowRepository) *CreateWorkflowUsecase {
	return &CreateWorkflowUsecase{repo: r}
}

func (uc *CreateWorkflowUsecase) Create(name string) error {
	workflow := &entity.Workflow{
		Name:   name,
		Status: "CREATED",
	}

	return uc.repo.Create(workflow)
}
