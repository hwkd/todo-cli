package args

import "testing"

// TODO: Add failure test cases

func TestParsingHelp(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  ParsedResult
	}{
		{
			"Help",
			[]string{"-h"},
			ParsedResult{
				Action: ActionHelp,
				Values: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Errorf("Expected nil, got `%s`", err)
				return
			}

			if result.Action != tt.want.Action {
				t.Errorf("Expected %s, got %s", tt.want.Action, result.Action)
				return
			}

			if tt.want.Values != nil {
				t.Errorf("Expected nil, got %v", tt.want.Values)
				return
			}
		})
	}
}

func TestParsingEmpty(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  ParsedResult
	}{
		{
			"Empty input should default to list",
			[]string{},
			ParsedResult{
				Action: ActionList,
				Values: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Errorf("Expected nil, got `%s`", err)
				return
			}

			if result.Action != tt.want.Action {
				t.Errorf("Expected %s, got %s", tt.want.Action, result.Action)
				return
			}

			if tt.want.Values != nil {
				t.Errorf("Expected nil, got %v", tt.want.Values)
				return
			}
		})
	}
}

func TestParsingList(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  ParsedResult
	}{
		{
			"List all todos",
			[]string{"-l"},
			ParsedResult{
				Action: ActionList,
				Values: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Errorf("Expected nil, got `%s`", err)
				return
			}

			if result.Action != tt.want.Action {
				t.Errorf("Expected %s, got %s", tt.want.Action, result.Action)
				return
			}

			if tt.want.Values != nil {
				t.Errorf("Expected nil, got %v", tt.want.Values)
				return
			}
		})
	}
}

func TestParsingAdd(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  ParsedResult
	}{
		{
			"Add only title",
			[]string{"-a", "some title"},
			ParsedResult{
				Action: ActionAdd,
				Values: ParsedValues{
					"title": "some title",
				},
			},
		},
		{
			"Add title and description",
			[]string{"-a", "hello", "world and goodbye"},
			ParsedResult{
				Action: ActionAdd,
				Values: ParsedValues{
					"title":       "hello",
					"description": "world and goodbye",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Errorf("Expected nil, got `%s`", err)
				return
			}

			if result.Action != tt.want.Action {
				t.Errorf("Expected %s, got %s", tt.want.Action, result.Action)
				return
			}

			for _, field := range []string{"title", "description"} {
				if value, ok := tt.want.Values[field]; ok {
					if result.Values[field] != value {
						t.Errorf("Expected %s, got %s", value, result.Values[field])
						return
					}
				} else if value, ok := result.Values[field]; ok {
					t.Errorf("Expected nil, got %s", value)
					return
				}
			}
		})
	}
}

func TestParsingUpdate(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  ParsedResult
	}{
		{
			name:  "Update only title",
			input: []string{"-u", "1", "-t", "abc"},
			want: ParsedResult{
				Action: ActionUpdate,
				Values: ParsedValues{
					"id":    "1",
					"title": "abc",
				},
			},
		},
		{
			name:  "Update only description",
			input: []string{"-u", "2", "-d", "defgh"},
			want: ParsedResult{
				Action: ActionUpdate,
				Values: ParsedValues{
					"id":          "2",
					"description": "defgh",
				},
			},
		},
		{
			name:  "Update title and description",
			input: []string{"-u", "3", "-t", "abc", "-d", "defgh"},
			want: ParsedResult{
				Action: ActionUpdate,
				Values: ParsedValues{
					"id":          "3",
					"title":       "abc",
					"description": "defgh",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Errorf("Expected nil, got `%s`", err)
				return
			}

			if result.Action != tt.want.Action {
				t.Errorf("Expected %s, got %s", tt.want.Action, result.Action)
				return
			}

			for _, field := range []string{"id", "title", "description"} {
				if value, ok := tt.want.Values[field]; ok {
					if result.Values[field] != value {
						t.Errorf("Expected %s, got %s", value, result.Values[field])
						return
					}
				} else if value, ok := result.Values[field]; ok {
					t.Errorf("Expected nil, got %s", value)
					return
				}
			}
		})
	}
}

func TestParsingDelete(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  ParsedResult
	}{
		{
			"Get ID",
			[]string{"-d", "123", "456"},
			ParsedResult{
				Action: ActionDelete,
				Values: ParsedValues{
					"ids": []string{"123", "456"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Errorf("Expected nil, got `%s`", err)
				return
			}

			if result.Action != tt.want.Action {
				t.Errorf("Expected %s, got %s", tt.want.Action, result.Action)
				return
			}

			if ids, ok := tt.want.Values["ids"].([]string); ok {
				resultIds := result.Values["ids"].([]string)
				for i, id := range ids {
					if resultIds[i] != id {
						t.Errorf("Expected %s, got %s", id, resultIds[i])
						return
					}
				}
			} else if ids, ok := result.Values["ids"]; ok {
				t.Errorf("Expected nil, got %v", ids)
				return
			}
		})
	}
}

func TestParsingMarkComplete(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  ParsedResult
	}{
		{
			"Mark todo with id 123 complete",
			[]string{"-c", "123"},
			ParsedResult{
				Action: ActionMarkComplete,
				Values: ParsedValues{
					"ids": []string{"123"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Errorf("Expected nil, got %s", err)
				return
			}

			if result.Action != tt.want.Action {
				t.Errorf("Expected %s, got %s", tt.want.Action, result.Action)
				return
			}

			if ids, ok := tt.want.Values["ids"].([]string); ok {
				resultIds := result.Values["ids"].([]string)
				for i, id := range ids {
					if resultIds[i] != id {
						t.Errorf("Expected %s, got %s", id, resultIds[i])
						return
					}
				}
			} else if ids, ok := result.Values["ids"]; ok {
				t.Errorf("Expected nil, got %v", ids)
				return
			}
		})
	}
}

func TestParsingMarkIncomplete(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  ParsedResult
	}{
		{
			"Mark todo with id 123 as incomplete",
			[]string{"-r", "123"},
			ParsedResult{
				Action: ActionMarkIncomplete,
				Values: ParsedValues{
					"ids": []string{"123"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Errorf("Expected nil, got %s", err)
				return
			}

			if result.Action != tt.want.Action {
				t.Errorf("Expected %s, got %s", tt.want.Action, result.Action)
				return
			}

			if ids, ok := tt.want.Values["ids"].([]string); ok {
				resultIds := result.Values["ids"].([]string)
				for i, id := range ids {
					if resultIds[i] != id {
						t.Errorf("Expected %s, got %s", id, resultIds[i])
						return
					}
				}
			} else if ids, ok := result.Values["ids"]; ok {
				t.Errorf("Expected nil, got %v", ids)
				return
			}
		})
	}
}
