package types_test

import (
	"strconv"
	"testing"

	"github.com/alog-rs/shared-packages/pkg/types"
)

func TestIsStandardSkill(t *testing.T) {
	cases := []struct {
		name     string
		skill    types.Skill
		expected bool
	}{
		{"elite skill", types.Invention, false},
		{"standard skill", types.Attack, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.skill.IsStandardSkill()

			if result != tc.expected {
				t.Fatalf("got %s expected %s", strconv.FormatBool(result), strconv.FormatBool(tc.expected))
			}
		})
	}
}

func TestIsEliteSkill(t *testing.T) {
	cases := []struct {
		name     string
		skill    types.Skill
		expected bool
	}{
		{"elite skill", types.Invention, true},
		{"standard skill", types.Strength, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.skill.IsEliteSkill()

			if result != tc.expected {
				t.Fatalf("got %s expected %s", strconv.FormatBool(result), strconv.FormatBool(tc.expected))
			}
		})
	}
}

func TestGetVirtualLevel(t *testing.T) {
	cases := []struct {
		name     string
		skill    types.Skill
		XP       int64
		expected int
	}{
		{"standard skill 0 xp", types.Attack, 0, 1},
		{"elite skill 0 xp", types.Invention, 0, 1},
		{"standard skill 200m xp", types.Attack, 200000000, 126},
		{"elite skill 200m xp", types.Invention, 200000000, 150},
		{"standard skill 99", types.Prayer, 13034431, 99},
		{"elite skill 99", types.Invention, 36073511, 99},
		{"standard skill 120", types.Constitution, 104273167, 120},
		{"elite skill 120", types.Invention, 80618654, 120},
		{"standard skill 75", types.Magic, 1336400, 75},
		{"elite skill 75", types.Invention, 11882200, 75},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.skill.GetVirtualLevel(tc.XP)

			if result != tc.expected {
				t.Fatalf("got %d expected %d", result, tc.expected)
			}
		})
	}
}
