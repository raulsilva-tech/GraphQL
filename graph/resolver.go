package graph

import "github.com/raulsilva-tech/GraphQL/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDAO *database.Category
	CourseDAO   *database.Course
}
