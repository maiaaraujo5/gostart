package health

import "context"

type health struct {
	Name        string
	Description string
	Checker     Checker
	Enabled     bool
}

type result struct {
	Health *health
	Err    error
}

var healths []*health

func NewHealth(name, description string, checker Checker, enabled bool) *health {
	return &health{
		Name:        name,
		Description: description,
		Checker:     checker,
		Enabled:     enabled,
	}
}

func Add(health *health) {
	healths = append(healths, health)
}

func Check(ctx context.Context) []*result {
	var results []*result

	for _, health := range healths {
		if !health.Enabled {
			continue
		}

		err := health.Checker.Check(ctx)

		results = append(results, NewResult(health, err))
	}

	return results
}

func NewResult(health *health, err error) *result {
	return &result{
		Health: health,
		Err:    err,
	}
}
