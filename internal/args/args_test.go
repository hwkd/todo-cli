package args

import "testing"

// TODO: Add failure test cases

type StringMap map[string]string

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
			input: []string{"-u", "-t", "abc"},
			want: ParsedResult{
				Action: ActionUpdate,
				Values: ParsedValues{
					"title": "abc",
				},
			},
		},
		{
			name:  "Update only description",
			input: []string{"-u", "-d", "defgh"},
			want: ParsedResult{
				Action: ActionUpdate,
				Values: ParsedValues{
					"description": "defgh",
				},
			},
		},
		{
			name:  "Update title and description",
			input: []string{"-u", "-t", "abc", "-d", "defgh"},
			want: ParsedResult{
				Action: ActionUpdate,
				Values: ParsedValues{
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

func TestParsingDelete(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  ParsedResult
	}{
		{
			"Get ID",
			[]string{"-d", "123"},
			ParsedResult{
				Action: ActionDelete,
				Values: ParsedValues{
					"id": "123",
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

			if id, ok := tt.want.Values["id"]; ok {
				if result.Values["id"] != id {
					t.Errorf("Expected %s, got %s", id, result.Values["id"])
					return
				}
			} else if id, ok := result.Values["id"]; ok {
				t.Errorf("Expected nil, got %s", id)
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
					"id": "123",
				},
			},
		},
		{
			"Mark todo with id 123 complete",
			[]string{"-c", "123"},
			ParsedResult{
				Action: ActionMarkComplete,
				Values: ParsedValues{
					"id": "123",
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

			if id, ok := tt.want.Values["id"]; ok {
				if result.Values["id"] != id {
					t.Errorf("Expected %s, got %s", id, result.Values["id"])
					return
				}
			} else if id, ok := result.Values["id"]; ok {
				t.Errorf("Expected nil, got %s", id)
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
					"id": "123",
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

			if id, ok := tt.want.Values["id"]; ok {
				if result.Values["id"] != id {
					t.Errorf("Expected %s, got %s", id, result.Values["id"])
					return
				}
			} else if id, ok := result.Values["id"]; ok {
				t.Errorf("Expected nil, got %s", id)
				return
			}
		})
	}
}
