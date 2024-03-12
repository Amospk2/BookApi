package category

func (category Category) Valid() bool {

	if len(category.Name) == 0 || category.Name == "" {
		return false
	}

	return true
}
