package tests

import (
	"testing"
	"time"

	cfg_builder "github.com/quack-pot/cfg_builder/pkg"
)

const test_json_name string = "test.json"

type t_GolfJson struct {
	Hotel float64                    `json:"hotel"`
	India cfg_builder.DurationString `json:"india"`
}

func (j *t_GolfJson) CheckValues(t *testing.T) {
	AssertEqual(t, j.Hotel, 3.14)
	AssertEqual(t, j.India, cfg_builder.DurationString(time.Duration(61)*time.Second))
}

type t_FoxtrotJson struct {
	FoxtrotAlpha string  `json:"foxtrot_alpha"`
	FoxtrotBravo float64 `json:"foxtrot_bravo"`
}

func (j *t_FoxtrotJson) CheckValues(t *testing.T, idx int) {
	switch idx {
	case 0:
		AssertEqual(t, j.FoxtrotAlpha, "Item A")
		AssertEqual(t, j.FoxtrotBravo, -100.50)
	case 1:
		AssertEqual(t, j.FoxtrotAlpha, "Item B")
		AssertEqual(t, j.FoxtrotBravo, 1.00)
	case 2:
		AssertEqual(t, j.FoxtrotAlpha, "Item C")
		AssertEqual(t, j.FoxtrotBravo, -99.99)
	case 3:
		AssertEqual(t, j.FoxtrotAlpha, "Item D")
		AssertEqual(t, j.FoxtrotBravo, 1234.56)
	default:
		t.Fatalf("Unknown t_FoxtrotJson Value Of Index %d.", idx)
	}
}

type t_CharlieJson struct {
	Delta   uint            `json:"delta"`
	Echo    bool            `json:"echo"`
	Foxtrot []t_FoxtrotJson `json:"foxtrot"`
	Golf    t_GolfJson      `json:"golf"`
}

func (j *t_CharlieJson) CheckValues(t *testing.T) {
	AssertEqual(t, j.Delta, 2)
	AssertEqual(t, j.Echo, false)

	AssertEqual(t, len(j.Foxtrot), 4)
	for idx := range len(j.Foxtrot) {
		j.Foxtrot[idx].CheckValues(t, idx)
	}

	j.Golf.CheckValues(t)
}

type t_TestJson struct {
	Alpha   float64       `json:"alpha"`
	Bravo   string        `json:"bravo"`
	Charlie t_CharlieJson `json:"charlie"`
	Juliet  int           `json:"juliet"`
	Kilo    bool          `json:"kilo"`
	Lima    []int         `json:"lima"`

	Unused   int `json:"unused"`
	Untagged int
}

func (j *t_TestJson) CheckValues(t *testing.T) {
	AssertEqual(t, j.Alpha, 5.0)
	AssertEqual(t, j.Bravo, "Hello, World!")

	j.Charlie.CheckValues(t)

	AssertEqual(t, j.Juliet, -7)
	AssertEqual(t, j.Kilo, true)

	AssertEqual(t, len(j.Lima), 5)
	for idx := range len(j.Lima) {
		AssertEqual(t, j.Lima[idx], idx)
	}
}
