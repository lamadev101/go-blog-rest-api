package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lamadev101/blog-rest-api/utils"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Blogs     []Blog    `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Blog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string    `gorm:"type:varchar(200);not null"`
	Slug      string    `gorm:"type:varchar(200);unique;not null"`
	Content   string    `gorm:"type:text;not null"`
	AuthorID  string    `gorm:"type:uuid;not null"` // Foreign key
	Author    User      `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Blog) BeforeCreate(tx *gorm.DB) (err error) {
	if b.Slug == "" { // Generate slug only if it is empty
		b.Slug = utils.GenerateSlug(b.Title)
	}
	return
}
