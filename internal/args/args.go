package args

import (
	"errors"
	"fmt"
)

/*
Usage:
  List todolist:
    todo

  Add todo:
    todo -a <title> [description]

  Update field:
    todo -u [-t title] [-d description]

  Delete:
    todo -d <id>

  Mark todo as complete:
    todo -c <id>

  Mark todo as incomplete:
    todo -r <id>
*/

const (
	ActionUndefined      = "undefined"
	ActionAdd            = "add"
	ActionUpdate         = "update"
	ActionDelete         = "delete"
	ActionMarkComplete   = "mark_complete"
	ActionMarkIncomplete = "mark_incomplete"
)

var (
	ErrUnsupportedAction = errors.New("Action not supported")
	ErrWrongFlag         = errors.New("Parsing wrong flag")
	ErrMissingArg        = errors.New("Missing argument")
)

// Parse parses the arguments, checks for syntax, and returns error if any, or returns ParseResult otherwise
func Parse(args []string) (*ParsedResult, error) {
	p := newParser(args)
	return p.parse()
}

type ParsedResult struct {
	Action string
	Values ParsedValues
}

type ParsedValues map[string]interface{}

type parser struct {
	idx     int
	readIdx int
	arg     *string
	args    []string
}

func newParser(args []string) parser {
	p := parser{args: args}
	p.read()
	return p
}

func (p *parser) parse() (*ParsedResult, error) {
	if len(p.args) == 0 {
		return &ParsedResult{
			Action: ActionUndefined,
			Values: nil,
		}, nil
	}

	return p.parseAction()
}

func (p *parser) parseAction() (*ParsedResult, error) {
	switch *p.arg {
	case "-a":
		return p.parseAddAction()
	case "-u":
		return p.parseUpdateAction()
	case "-d":
		return p.parseDeleteAction()
	case "-c":
		return p.parseMarkCompleteAction()
	case "-r":
		return p.parseMarkIncompleteAction()
	default:
		return nil, ErrUnsupportedAction
	}
}

// Parse `todo -a <title> [description]`
func (p *parser) parseAddAction() (*ParsedResult, error) {
	err := p.checkFlag("-a")
	if err != nil {
		return nil, err
	}

	p.read()
	if p.arg == nil {
		return nil, fmt.Errorf("%w: title", ErrMissingArg)
	}

	result := &ParsedResult{
		Action: ActionAdd,
		Values: ParsedValues{
			"title": *p.arg,
		},
	}

	p.read()
	if p.arg != nil {
		result.Values["description"] = *p.arg
	}

	return result, nil
}

// Parse `todo -u [-t title] [-d description]`
func (p *parser) parseUpdateAction() (*ParsedResult, error) {
	err := p.checkFlag("-u")
	if err != nil {
		return nil, err
	}

	p.read()
	if p.arg == nil {
		return nil, fmt.Errorf("%w: Expected -t or -d flag", ErrMissingArg)
	}

	var title, description *string

	for p.arg != nil {
		if *p.arg == "-t" {
			p.read()
			if p.arg == nil {
				return nil, fmt.Errorf("%w: title", ErrMissingArg)
			}
			title = p.arg
			p.read()
		} else if *p.arg == "-d" {
			p.read()
			if p.arg == nil {
				return nil, fmt.Errorf("%w: description", ErrMissingArg)
			}
			description = p.arg
			p.read()
		}
	}

	result := &ParsedResult{
		Action: ActionUpdate,
		Values: ParsedValues{},
	}
	if title != nil {
		result.Values["title"] = *title
	}
	if description != nil {
		result.Values["description"] = *description
	}

	return result, nil
}

// Parse `todo -d <id>`
func (p *parser) parseDeleteAction() (*ParsedResult, error) {
	err := p.checkFlag("-d")
	if err != nil {
		return nil, err
	}

	p.read()
	if p.arg == nil {
		return nil, fmt.Errorf("%w: id", ErrMissingArg)
	}

	result := &ParsedResult{
		Action: ActionDelete,
		Values: ParsedValues{
			"id": *p.arg,
		},
	}

	return result, nil
}

// Parse `todo -c <id>`
func (p *parser) parseMarkCompleteAction() (*ParsedResult, error) {
	err := p.checkFlag("-c")
	if err != nil {
		return nil, err
	}

	p.read()
	if p.arg == nil {
		return nil, fmt.Errorf("%w: id", ErrMissingArg)
	}

	result := &ParsedResult{
		Action: ActionMarkComplete,
		Values: ParsedValues{
			"id": *p.arg,
		},
	}

	return result, nil
}

// Parse `todo -r <id>`
func (p *parser) parseMarkIncompleteAction() (*ParsedResult, error) {
	err := p.checkFlag("-r")
	if err != nil {
		return nil, err
	}

	p.read()
	if p.arg == nil {
		return nil, fmt.Errorf("%w: id", ErrMissingArg)
	}

	result := &ParsedResult{
		Action: ActionMarkIncomplete,
		Values: ParsedValues{
			"id": *p.arg,
		},
	}

	return result, nil
}

func (p *parser) read() {
	if p.readIdx >= len(p.args) {
		p.arg = nil
	} else {
		p.arg = &p.args[p.readIdx]
	}
	p.idx = p.readIdx
	p.readIdx++
}

func (p *parser) checkFlag(flag string) error {
	if p.arg == nil {
		return fmt.Errorf("%w: Expected %s, got nil", ErrWrongFlag, flag)
	}
	if *p.arg != flag {
		return fmt.Errorf("%w: Expected %s, got %s", ErrWrongFlag, flag, *p.arg)
	}
	return nil
}
