package tests

import (
	"path/filepath"
	"testing"

	cfg_builder "github.com/quack-pot/cfg_builder/pkg"
)

func TestJson(t *testing.T) {
	builder := cfg_builder.NewConfigBuilder()

	builder.AddJSON(filepath.Join("./data/", test_json_name))

	config, err := builder.Build()

	if err != nil {
		t.Fatalf("Error building config: %v", err)
	}

	untagged_value, has_untagged_value := config.Get("untagged")
	AssertEqual(t, untagged_value.(float64), 777)
	AssertEqual(t, has_untagged_value, true)

	_, has_unused_value := config.Get("unused")
	AssertEqual(t, has_unused_value, false)

	var test_json t_TestJson = t_TestJson{Unused: 999, Untagged: 888}
	if err := config.Bind(&test_json, ""); err != nil {
		t.Fatalf("Error parsing test json data: %v", err)
	}

	AssertEqual(t, test_json.Unused, 999)
	AssertEqual(t, test_json.Untagged, 777)

	test_json.CheckValues(t)
}

func TestEnv(t *testing.T) {
	builder := cfg_builder.NewConfigBuilder()

	builder.AddENV(filepath.Join("./data/", test_env_name))

	config, err := builder.Build()

	if err != nil {
		t.Fatalf("Error building config: %v", err)
	}

	untagged_value, has_untagged_value := config.Get("untagged")
	AssertEqual(t, untagged_value.(float64), 777)
	AssertEqual(t, has_untagged_value, true)

	_, has_unused_value := config.Get("unused")
	AssertEqual(t, has_unused_value, false)

	var test_env t_TestEnv = t_TestEnv{Unused: 999, Untagged: 888}
	if err := config.Bind(&test_env, ""); err != nil {
		t.Fatalf("Error parsing test env data: %v", err)
	}

	AssertEqual(t, test_env.Unused, 999)
	AssertEqual(t, test_env.Untagged, 777)

	test_env.CheckValues(t)
}
