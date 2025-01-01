package repository

import (
	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) GetModulesByUser(userID uuid.UUID) ([]model.CourseModule, error) {
	var modules []model.CourseModule
	if err := r.db.Where("user_id = ?", userID).Find(&modules).Error; err != nil {
		return nil, err
	}
	return modules, nil
}

func (r *CourseRepository) GetAllModules() ([]model.CourseModule, error) {
	var modules []model.CourseModule
	if err := r.db.Find(&modules).Error; err != nil {
		return nil, err
	}
	return modules, nil
}

func (r *CourseRepository) GetCoursesByUser(userID uuid.UUID) ([]model.Course, error) {
	var courses []model.Course
	if err := r.db.Preload("Modules").Preload("Category").Where("user_id = ?", userID).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *CourseRepository) GetAllCourses() ([]model.Course, error) {
	var courses []model.Course
	if err := r.db.Preload("Modules").Preload("Category").Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *CourseRepository) GetCourseByIDAndUser(id, userID uuid.UUID) (*model.Course, error) {
	var course model.Course
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepository) CreateCourse(course *model.Course) error {
	return r.db.Create(course).Error
}

func (r *CourseRepository) UpdateCourse(course *model.Course) error {
	return r.db.Save(course).Error
}

func (r *CourseRepository) DeleteCourse(id uuid.UUID) error {
	return r.db.Delete(&model.Course{}, "id = ?", id).Error
}

func (r *CourseRepository) GetCategoriesByUser(userID uuid.UUID) ([]model.CourseCategory, error) {
	var categories []model.CourseCategory
	if err := r.db.Preload("Courses").Preload("Courses.Modules").Preload("Courses.Category").
		Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CourseRepository) GetAllCategories() ([]model.CourseCategory, error) {
	var categories []model.CourseCategory
	if err := r.db.Preload("Courses").Preload("Courses.Modules").Preload("Courses.Category").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CourseRepository) GetCategoryByIDAndUser(id, userID uuid.UUID) (*model.CourseCategory, error) {
	var category model.CourseCategory
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CourseRepository) CreateCategory(category *model.CourseCategory) error {
	return r.db.Create(category).Error
}

func (r *CourseRepository) UpdateCategory(category *model.CourseCategory) error {
	return r.db.Save(category).Error
}

func (r *CourseRepository) DeleteCategory(id uuid.UUID) error {
	return r.db.Delete(&model.CourseCategory{}, "id = ?", id).Error
}

func (r *CourseRepository) GetModuleByIDAndUser(id, userID uuid.UUID) (*model.CourseModule, error) {
	var module model.CourseModule
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&module).Error; err != nil {
		return nil, err
	}
	return &module, nil
}

func (r *CourseRepository) CreateModule(module *model.CourseModule) error {
	return r.db.Create(module).Error
}

func (r *CourseRepository) UpdateModule(module *model.CourseModule) error {
	return r.db.Save(module).Error
}

func (r *CourseRepository) DeleteModule(id uuid.UUID) error {
	return r.db.Delete(&model.CourseModule{}, "id = ?", id).Error
}

func (r *CourseRepository) CountModulesByIDsAndUser(ids []uuid.UUID, userID uuid.UUID) (int64, error) {
	var count int64
	if err := r.db.Model(&model.CourseModule{}).
		Where("id IN ? AND user_id = ?", ids, userID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CourseRepository) CountCoursesByIDsAndUser(ids []uuid.UUID, userID uuid.UUID) (int64, error) {
	var count int64
	if err := r.db.Model(&model.Course{}).
		Where("id IN ? AND user_id = ?", ids, userID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CourseRepository) CountItemsByIDsAndUser(ids []uuid.UUID, userID uuid.UUID) (int64, error) {
	var count int64
	if err := r.db.Model(&model.CourseItem{}).
		Where("id IN ? AND user_id = ?", ids, userID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CourseRepository) ReplaceModuleItems(moduleID uuid.UUID, itemIDs []uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if len(itemIDs) == 0 {
			return tx.Where("course_module_id = ?", moduleID).
				Delete(&model.CourseModuleItem{}).Error
		}
		if err := tx.Where("course_module_id = ? AND course_item_id NOT IN ?", moduleID, itemIDs).
			Delete(&model.CourseModuleItem{}).Error; err != nil {
			return err
		}
		rows := make([]model.CourseModuleItem, 0, len(itemIDs))
		for _, itemID := range itemIDs {
			rows = append(rows, model.CourseModuleItem{
				CourseModuleID: moduleID,
				CourseItemID:   itemID,
			})
		}
	return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&rows).Error
	})
}

func (r *CourseRepository) ReplaceCourseModules(courseID uuid.UUID, moduleIDs []uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if len(moduleIDs) == 0 {
			return tx.Where("course_id = ?", courseID).
				Delete(&model.CourseCourseModule{}).Error
		}
		if err := tx.Where("course_id = ? AND course_module_id NOT IN ?", courseID, moduleIDs).
			Delete(&model.CourseCourseModule{}).Error; err != nil {
			return err
		}
		rows := make([]model.CourseCourseModule, 0, len(moduleIDs))
		for _, moduleID := range moduleIDs {
			rows = append(rows, model.CourseCourseModule{
				CourseID:       courseID,
				CourseModuleID: moduleID,
			})
		}
		return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&rows).Error
	})
}

func (r *CourseRepository) AddCourseModules(courseID uuid.UUID, moduleIDs []uuid.UUID) error {
	if len(moduleIDs) == 0 {
		return nil
	}
	rows := make([]model.CourseCourseModule, 0, len(moduleIDs))
	for _, moduleID := range moduleIDs {
		rows = append(rows, model.CourseCourseModule{
			CourseID:       courseID,
			CourseModuleID: moduleID,
		})
	}
	return r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&rows).Error
}

func (r *CourseRepository) ReplaceItemModules(itemID uuid.UUID, moduleIDs []uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if len(moduleIDs) == 0 {
			return tx.Where("course_item_id = ?", itemID).
				Delete(&model.CourseModuleItem{}).Error
		}
		if err := tx.Where("course_item_id = ? AND course_module_id NOT IN ?", itemID, moduleIDs).
			Delete(&model.CourseModuleItem{}).Error; err != nil {
			return err
		}
		rows := make([]model.CourseModuleItem, 0, len(moduleIDs))
		for _, moduleID := range moduleIDs {
			rows = append(rows, model.CourseModuleItem{
				CourseModuleID: moduleID,
				CourseItemID:   itemID,
			})
		}
		return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&rows).Error
	})
}

func (r *CourseRepository) SetCoursesCategoryID(userID, categoryID uuid.UUID, ids []uuid.UUID) error {
	return r.db.Model(&model.Course{}).
		Where("id IN ? AND user_id = ?", ids, userID).
		Update("category_id", categoryID).Error
}

func (r *CourseRepository) ClearCategoryFromOtherCourses(userID, categoryID uuid.UUID, keepIDs []uuid.UUID) error {
	query := r.db.Model(&model.Course{}).
		Where("category_id = ? AND user_id = ?", categoryID, userID)
	if len(keepIDs) > 0 {
		query = query.Where("id NOT IN ?", keepIDs)
	}
	return query.Update("category_id", nil).Error
}

func (r *CourseRepository) GetItemsByUser(userID uuid.UUID) ([]model.CourseItem, error) {
	var items []model.CourseItem
	if err := r.db.Preload("Modules").Where("user_id = ?", userID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *CourseRepository) GetAllItems() ([]model.CourseItem, error) {
	var items []model.CourseItem
	if err := r.db.Preload("Modules").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *CourseRepository) GetItemByIDAndUser(id, userID uuid.UUID) (*model.CourseItem, error) {
	var item model.CourseItem
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *CourseRepository) CreateItem(item *model.CourseItem) error {
	return r.db.Create(item).Error
}

func (r *CourseRepository) UpdateItem(item *model.CourseItem) error {
	return r.db.Save(item).Error
}

func (r *CourseRepository) DeleteItem(id uuid.UUID) error {
	return r.db.Delete(&model.CourseItem{}, "id = ?", id).Error
}
