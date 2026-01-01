package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/repository"
)

type CourseService struct {
	repo *repository.CourseRepository
}

func NewCourseService(repo *repository.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) GetMyModules(userID uuid.UUID) ([]model.CourseModule, error) {
	return s.repo.GetModulesByUser(userID)
}

func (s *CourseService) GetAllModules() ([]model.CourseModule, error) {
	return s.repo.GetAllModules()
}

func (s *CourseService) GetMyCourses(userID uuid.UUID) ([]model.Course, error) {
	return s.repo.GetCoursesByUser(userID)
}

func (s *CourseService) GetAllCourses() ([]model.Course, error) {
	return s.repo.GetAllCourses()
}

func (s *CourseService) GetMyCategories(userID uuid.UUID) ([]model.CourseCategory, error) {
	return s.repo.GetCategoriesByUser(userID)
}

func (s *CourseService) GetAllCategories() ([]model.CourseCategory, error) {
	return s.repo.GetAllCategories()
}

func (s *CourseService) CreateCourse(userID uuid.UUID, req *model.CreateCourseRequest) (*model.Course, error) {
	if req.CategoriaID != nil {
		if _, err := s.repo.GetCategoryByIDAndUser(*req.CategoriaID, userID); err != nil {
			return nil, errors.New("category not found")
		}
	}
	course := &model.Course{
		UserID:   userID,
		Name:     req.Nome,
		ImageURL: req.Imagem,
		CategoryID: req.CategoriaID,
	}
	if err := s.repo.CreateCourse(course); err != nil {
		return nil, err
	}
	if len(req.ModulosIDs) > 0 {
		if err := s.attachModulesToCourse(userID, course.ID, req.ModulosIDs); err != nil {
			return nil, err
		}
	}
	return course, nil
}

func (s *CourseService) UpdateCourse(userID, courseID uuid.UUID, req *model.UpdateCourseRequest) (*model.Course, error) {
	course, err := s.repo.GetCourseByIDAndUser(courseID, userID)
	if err != nil {
		return nil, err
	}
	if req.Nome != "" {
		course.Name = req.Nome
	}
	if req.Imagem != "" {
		course.ImageURL = req.Imagem
	}
	if req.CategoriaID != nil {
		if _, err := s.repo.GetCategoryByIDAndUser(*req.CategoriaID, userID); err != nil {
			return nil, errors.New("category not found")
		}
		course.CategoryID = req.CategoriaID
	}
	if err := s.repo.UpdateCourse(course); err != nil {
		return nil, err
	}
	if req.ModulosIDs != nil {
		if err := s.attachModulesToCourse(userID, course.ID, *req.ModulosIDs); err != nil {
			return nil, err
		}
	}
	return course, nil
}

func (s *CourseService) DeleteCourse(userID, courseID uuid.UUID) error {
	if _, err := s.repo.GetCourseByIDAndUser(courseID, userID); err != nil {
		return err
	}
	return s.repo.DeleteCourse(courseID)
}

func (s *CourseService) CreateCategory(userID uuid.UUID, req *model.CreateCourseCategoryRequest) (*model.CourseCategory, error) {
	category := &model.CourseCategory{
		UserID:   userID,
		Name:     req.Nome,
		ImageURL: req.Imagem,
	}
	if err := s.repo.CreateCategory(category); err != nil {
		return nil, err
	}
	if len(req.CursosIDs) > 0 {
		if err := s.attachCoursesToCategory(userID, category.ID, req.CursosIDs); err != nil {
			return nil, err
		}
	}
	return category, nil
}

func (s *CourseService) UpdateCategory(userID, categoryID uuid.UUID, req *model.UpdateCourseCategoryRequest) (*model.CourseCategory, error) {
	category, err := s.repo.GetCategoryByIDAndUser(categoryID, userID)
	if err != nil {
		return nil, err
	}
	if req.Nome != "" {
		category.Name = req.Nome
	}
	if req.Imagem != "" {
		category.ImageURL = req.Imagem
	}
	if err := s.repo.UpdateCategory(category); err != nil {
		return nil, err
	}
	if req.CursosIDs != nil {
		if err := s.attachCoursesToCategory(userID, category.ID, *req.CursosIDs); err != nil {
			return nil, err
		}
	}
	return category, nil
}

func (s *CourseService) DeleteCategory(userID, categoryID uuid.UUID) error {
	if _, err := s.repo.GetCategoryByIDAndUser(categoryID, userID); err != nil {
		return err
	}
	if err := s.repo.ClearCategoryFromOtherCourses(userID, categoryID, nil); err != nil {
		return err
	}
	return s.repo.DeleteCategory(categoryID)
}

func (s *CourseService) CreateModule(userID uuid.UUID, req *model.CreateCourseModuleRequest) (*model.CourseModule, error) {
	if req.CursoID != nil {
		if _, err := s.repo.GetCourseByIDAndUser(*req.CursoID, userID); err != nil {
			return nil, errors.New("course not found")
		}
	}
	module := &model.CourseModule{
		UserID:    userID,
		Title:     req.Modulo,
	}
	if err := s.repo.CreateModule(module); err != nil {
		return nil, err
	}
	if req.CursoID != nil {
		if err := s.repo.AddCourseModules(*req.CursoID, []uuid.UUID{module.ID}); err != nil {
			return nil, err
		}
	}
	if len(req.ItensIDs) > 0 {
		if err := s.attachItemsToModule(userID, module.ID, req.ItensIDs); err != nil {
			return nil, err
		}
	}
	return module, nil
}

func (s *CourseService) UpdateModule(userID, moduleID uuid.UUID, req *model.UpdateCourseModuleRequest) (*model.CourseModule, error) {
	module, err := s.repo.GetModuleByIDAndUser(moduleID, userID)
	if err != nil {
		return nil, err
	}
	if req.Modulo != "" {
		module.Title = req.Modulo
	}
	if req.CursoID != nil {
		if _, err := s.repo.GetCourseByIDAndUser(*req.CursoID, userID); err != nil {
			return nil, errors.New("course not found")
		}
	}
	if err := s.repo.UpdateModule(module); err != nil {
		return nil, err
	}
	if req.CursoID != nil {
		if err := s.repo.AddCourseModules(*req.CursoID, []uuid.UUID{module.ID}); err != nil {
			return nil, err
		}
	}
	if req.ItensIDs != nil {
		if err := s.attachItemsToModule(userID, module.ID, *req.ItensIDs); err != nil {
			return nil, err
		}
	}
	return module, nil
}

func (s *CourseService) DeleteModule(userID, moduleID uuid.UUID) error {
	if _, err := s.repo.GetModuleByIDAndUser(moduleID, userID); err != nil {
		return err
	}
	return s.repo.DeleteModule(moduleID)
}

func (s *CourseService) GetMyItems(userID uuid.UUID) ([]model.CourseItem, error) {
	return s.repo.GetItemsByUser(userID)
}

func (s *CourseService) GetAllItems() ([]model.CourseItem, error) {
	return s.repo.GetAllItems()
}

func (s *CourseService) CreateItem(userID uuid.UUID, req *model.CreateCourseItemRequest) (*model.CourseItem, error) {
	moduleIDs := req.ModulosIDs
	if len(moduleIDs) == 0 && req.ModuloID != nil {
		moduleIDs = []uuid.UUID{*req.ModuloID}
	}
	item := &model.CourseItem{
		UserID:  userID,
		Title:   req.Titulo,
		Type:    req.Tipo,
		Content: req.Conteudo,
	}
	if req.ModuloID != nil && len(req.ModulosIDs) == 0 {
		item.ModuleID = req.ModuloID
	}
	if err := s.repo.CreateItem(item); err != nil {
		return nil, err
	}
	if len(moduleIDs) > 0 {
		if err := s.attachModulesToItem(userID, item.ID, moduleIDs); err != nil {
			return nil, err
		}
	}
	return item, nil
}

func (s *CourseService) UpdateItem(userID, itemID uuid.UUID, req *model.UpdateCourseItemRequest) (*model.CourseItem, error) {
	item, err := s.repo.GetItemByIDAndUser(itemID, userID)
	if err != nil {
		return nil, err
	}
	moduleIDsProvided := false
	var moduleIDs []uuid.UUID
	if req.ModulosIDs != nil {
		moduleIDsProvided = true
		moduleIDs = *req.ModulosIDs
		item.ModuleID = nil
	} else if req.ModuloID != nil {
		moduleIDsProvided = true
		moduleIDs = []uuid.UUID{*req.ModuloID}
		item.ModuleID = req.ModuloID
	}
	if req.Titulo != "" {
		item.Title = req.Titulo
	}
	if req.Tipo != "" {
		item.Type = req.Tipo
	}
	if req.Conteudo != "" {
		item.Content = req.Conteudo
	}
	if err := s.repo.UpdateItem(item); err != nil {
		return nil, err
	}
	if moduleIDsProvided {
		if err := s.attachModulesToItem(userID, item.ID, moduleIDs); err != nil {
			return nil, err
		}
	}
	return item, nil
}

func (s *CourseService) DeleteItem(userID, itemID uuid.UUID) error {
	if _, err := s.repo.GetItemByIDAndUser(itemID, userID); err != nil {
		return err
	}
	return s.repo.DeleteItem(itemID)
}

func (s *CourseService) attachModulesToCourse(userID, courseID uuid.UUID, moduleIDs []uuid.UUID) error {
	if len(moduleIDs) > 0 {
		count, err := s.repo.CountModulesByIDsAndUser(moduleIDs, userID)
		if err != nil {
			return err
		}
		if count != int64(len(moduleIDs)) {
			return errors.New("module not found")
		}
	}
	return s.repo.ReplaceCourseModules(courseID, moduleIDs)
}

func (s *CourseService) attachCoursesToCategory(userID, categoryID uuid.UUID, courseIDs []uuid.UUID) error {
	if len(courseIDs) > 0 {
		count, err := s.repo.CountCoursesByIDsAndUser(courseIDs, userID)
		if err != nil {
			return err
		}
		if count != int64(len(courseIDs)) {
			return errors.New("course not found")
		}
		if err := s.repo.SetCoursesCategoryID(userID, categoryID, courseIDs); err != nil {
			return err
		}
	}
	return s.repo.ClearCategoryFromOtherCourses(userID, categoryID, courseIDs)
}

func (s *CourseService) attachItemsToModule(userID, moduleID uuid.UUID, itemIDs []uuid.UUID) error {
	if len(itemIDs) > 0 {
		count, err := s.repo.CountItemsByIDsAndUser(itemIDs, userID)
		if err != nil {
			return err
		}
		if count != int64(len(itemIDs)) {
			return errors.New("item not found")
		}
	}
	return s.repo.ReplaceModuleItems(moduleID, itemIDs)
}

func (s *CourseService) attachModulesToItem(userID, itemID uuid.UUID, moduleIDs []uuid.UUID) error {
	if len(moduleIDs) > 0 {
		count, err := s.repo.CountModulesByIDsAndUser(moduleIDs, userID)
		if err != nil {
			return err
		}
		if count != int64(len(moduleIDs)) {
			return errors.New("module not found")
		}
	}
	return s.repo.ReplaceItemModules(itemID, moduleIDs)
}
