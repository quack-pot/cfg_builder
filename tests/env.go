package tests

import (
	"testing"
	"time"

	cfg_builder "github.com/quack-pot/cfg_builder/pkg"
)

const test_env_name string = "test.env"

type t_GolfEnv struct {
	Hotel float64                    `json:"hotel"`
	India cfg_builder.DurationString `json:"india"`
}

func (e *t_GolfEnv) CheckValues(t *testing.T) {
	AssertEqual(t, e.Hotel, 4.25)
	AssertEqual(
		t,
		e.India,
		cfg_builder.DurationString(
			(time.Duration(3)*time.Hour)+
				(time.Duration(2)*time.Minute)+
				(time.Duration(1)*time.Second),
		),
	)
}

type t_CharlieEnv struct {
	Delta uint      `json:"delta"`
	Echo  bool      `json:"echo"`
	Golf  t_GolfEnv `json:"golf"`
}

func (e *t_CharlieEnv) CheckValues(t *testing.T) {
	AssertEqual(t, e.Delta, 3)
	AssertEqual(t, e.Echo, true)
	e.Golf.CheckValues(t)
}

type t_TestEnv struct {
	Alpha   float64      `json:"alpha"`
	Bravo   string       `json:"bravo"`
	Charlie t_CharlieEnv `json:"charlie"`
	Juliet  int          `json:"juliet"`
	Kilo    bool         `json:"kilo"`
	Lima    []int        `json:"lima"`

	Unused   int `json:"unused"`
	Untagged int
}

func (e *t_TestEnv) CheckValues(t *testing.T) {
	AssertEqual(t, e.Alpha, 6.0)
	AssertEqual(t, e.Bravo, "Goodbye, Universe!")

	e.Charlie.CheckValues(t)

	AssertEqual(t, e.Juliet, -8)
	AssertEqual(t, e.Kilo, false)

	AssertEqual(t, len(e.Lima), 5)
	for idx := range len(e.Lima) {
		AssertEqual(t, e.Lima[idx], idx+1)
	}
}
