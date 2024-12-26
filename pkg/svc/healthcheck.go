package svc

import (
	"context"

	"math/rand/v2"
)

func (s *Svc) SetUnhealthy(ctx context.Context, err error) {
	s.health = err
}

func (s *Svc) IsUnhealthy(ctx context.Context) bool {
	ctx, span := s.tracer.Start(ctx, "[SVC]_health_check")
	defer span.End()
	s.isResponsive(ctx)
	return s.health != nil
}

func (s *Svc) isResponsive(ctx context.Context) error {
	operations := []string{"+", "/", "*"}
	randomIndex := rand.IntN(len(operations))
	pick := operations[randomIndex]
	num1 := minRangeNumber1 + rand.IntN((maxRangeNumber1-minRangeNumber1)+minRangeNumber1)
	num2 := minRangeNumber2 + rand.IntN((maxRangeNumber2-minRangeNumber2)+minRangeNumber2)
	res := 0
	if pick == "+" {
		res = num1 + num2
	} else if pick == "/" {
		res = num1 / num2
	} else {
		res = num1 * num2
	}
	if res == 0 {
		s.SetUnhealthy(ctx, ErrUnexpected)
		return ErrUnexpected
	}
	return nil
}
