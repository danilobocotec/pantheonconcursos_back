package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseModule struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	Title       string         `gorm:"not null" json:"modulo"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (c *CourseModule) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

type CourseItem struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	ModuleID  *uuid.UUID     `gorm:"type:uuid;index" json:"modulo_id,omitempty"`
	Title     string         `gorm:"not null" json:"titulo"`
	Type      string         `gorm:"column:tipo" json:"tipo"`
	Content   string         `gorm:"type:text" json:"conteudo"`
	Modules   []CourseModule `gorm:"many2many:course_module_items;" json:"modulos,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (c *CourseItem) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

type CourseModuleItem struct {
	CourseModuleID uuid.UUID `gorm:"type:uuid;primaryKey;column:course_module_id" json:"course_module_id"`
	CourseItemID   uuid.UUID `gorm:"type:uuid;primaryKey;column:course_item_id" json:"course_item_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type CourseCourseModule struct {
	CourseID       uuid.UUID `gorm:"type:uuid;primaryKey;column:course_id" json:"course_id"`
	CourseModuleID uuid.UUID `gorm:"type:uuid;primaryKey;column:course_module_id" json:"course_module_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreateCourseModuleRequest struct {
	Modulo   string      `json:"modulo" binding:"required,min=2"`
	CursoID  *uuid.UUID  `json:"curso_id"`
	ItensIDs []uuid.UUID `json:"itens_ids"`
}

type UpdateCourseModuleRequest struct {
	Modulo   string       `json:"modulo" binding:"omitempty,min=2"`
	CursoID  *uuid.UUID   `json:"curso_id"`
	ItensIDs *[]uuid.UUID `json:"itens_ids"`
}

type Course struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	CategoryID *uuid.UUID    `gorm:"type:uuid;index" json:"categoria_id,omitempty"`
	Name      string         `gorm:"not null" json:"nome"`
	ImageURL  string         `gorm:"column:image" json:"imagem"`
	Category  *CourseCategory `gorm:"foreignKey:CategoryID" json:"categoria,omitempty"`
	Modules   []CourseModule `gorm:"many2many:course_course_modules;" json:"modulos"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (c *Course) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

type CreateCourseRequest struct {
	Nome       string      `json:"nome" binding:"required,min=2"`
	CategoriaID *uuid.UUID `json:"categoria_id"`
	Imagem     string      `json:"imagem"`
	ModulosIDs []uuid.UUID `json:"modulos_ids"`
}

type UpdateCourseRequest struct {
	Nome       string        `json:"nome" binding:"omitempty,min=2"`
	CategoriaID *uuid.UUID   `json:"categoria_id"`
	Imagem     string        `json:"imagem"`
	ModulosIDs *[]uuid.UUID  `json:"modulos_ids"`
}

type CourseCategory struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	Name      string         `gorm:"not null" json:"nome"`
	ImageURL  string         `gorm:"column:image" json:"imagem"`
	Courses   []Course       `gorm:"foreignKey:CategoryID" json:"cursos"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (c *CourseCategory) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

type CreateCourseCategoryRequest struct {
	Nome      string      `json:"nome" binding:"required,min=2"`
	Imagem    string      `json:"imagem"`
	CursosIDs []uuid.UUID `json:"cursos_ids"`
}

type UpdateCourseCategoryRequest struct {
	Nome      string        `json:"nome" binding:"omitempty,min=2"`
	Imagem    string        `json:"imagem"`
	CursosIDs *[]uuid.UUID  `json:"cursos_ids"`
}

type CreateCourseItemRequest struct {
	ModuloID   *uuid.UUID `json:"modulo_id"`
	ModulosIDs []uuid.UUID `json:"modulos_ids"`
	Titulo     string     `json:"titulo" binding:"required,min=2"`
	Tipo       string     `json:"tipo" binding:"required,min=2"`
	Conteudo   string     `json:"conteudo" binding:"required,min=2"`
}

type UpdateCourseItemRequest struct {
	ModuloID   *uuid.UUID  `json:"modulo_id"`
	ModulosIDs *[]uuid.UUID `json:"modulos_ids"`
	Titulo     string      `json:"titulo" binding:"omitempty,min=2"`
	Tipo       string      `json:"tipo" binding:"omitempty,min=2"`
	Conteudo   string      `json:"conteudo" binding:"omitempty,min=2"`
}
