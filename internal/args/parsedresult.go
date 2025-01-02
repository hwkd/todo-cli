package args

// ParsedAddActionValues is a struct that holds the parsed values of the add action.
type ParsedAddActionValues struct {
	Title       string
	Description string
}

// ParsedUpdateActionValues is a struct that holds the parsed values of the update action.
type ParsedUpdateActionValues struct {
	ID          string
	Title       string
	Description string
}

// ParsedIdValues is a struct that holds the parsed ids in the command line arguments.
type ParsedIdValues struct {
	IDs []string
}

type ParsedValues map[string]interface{}

// ParsedResult is a struct that holds the parsed action and values of the command line arguments.
type ParsedResult struct {
	Action string
	Values ParsedValues
}

// ParseAddActionValues wraps the parsed values in a typed struct for ease of use and safety.
func (r *ParsedResult) ParseAddActionValues() ParsedAddActionValues {
	values := ParsedAddActionValues{
		Title: r.Values["title"].(string),
	}
	if desc, ok := r.Values["description"]; ok {
		values.Description = desc.(string)
	}
	return values
}

// ParseUpdateActionValues wraps the parsed values in a typed struct for ease of use and safety.
func (r *ParsedResult) ParseUpdateActionValues() ParsedUpdateActionValues {
	values := ParsedUpdateActionValues{
		ID: r.Values["id"].(string),
	}
	if title, ok := r.Values["title"]; ok {
		values.Title = title.(string)
	}
	if desc, ok := r.Values["description"]; ok {
		values.Description = desc.(string)
	}
	return values
}

// ParseDeleteActionValues wraps the parsed values in a typed struct for ease of use and safety.
func (r *ParsedResult) ParseDeleteActionValues() ParsedIdValues {
	return ParsedIdValues{
		IDs: r.Values["ids"].([]string),
	}
}

// ParseMarkCompleteActionValues wraps the parsed values in a typed struct for ease of use and safety.
func (r *ParsedResult) ParseMarkCompleteActionValues() ParsedIdValues {
	return ParsedIdValues{
		IDs: r.Values["ids"].([]string),
	}
}

// ParseMarkIncompleteActionValues wraps the parsed values in a typed struct for ease of use and safety.
func (r *ParsedResult) ParseMarkIncompleteActionValues() ParsedIdValues {
	return ParsedIdValues{
		IDs: r.Values["ids"].([]string),
	}
}
