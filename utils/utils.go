package utils

import (
	"strings"
	"unicode"
)

func GenerateSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove non-alphanumeric characters except hyphens
	slug = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' {
			return r
		}
		return -1
	}, slug)

	return slug
}

// func GenerateUniqueSlug(title string, db *gorm.DB) string {
// 	baseSlug := GenerateSlug(title)
// 	slug := baseSlug
// 	counter := 1

// 	for {
// 		var count int64
// 		db.Model(&models.Blog{}).Where("slug = ?", slug).Count(&count)
// 		if count == 0 {
// 			break
// 		}
// 		slug = fmt.Sprintf("%s-%d", baseSlug, counter)
// 		counter++
// 	}

// 	return slug
// }
