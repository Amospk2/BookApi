package author

func (author Author) Valid() bool {

	if len(author.Name) == 0 || author.Name == "" {
		return false
	}

	return true
}
