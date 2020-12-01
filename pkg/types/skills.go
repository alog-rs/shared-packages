package types

import (
	"sort"

	"github.com/alog-rs/shared-packages/pkg/constants"
)

// Skill represents one of the Runescape skills
type Skill int

// Runescape skills
const (
	Attack Skill = iota
	Defence
	Strength
	Constitution
	Ranged
	Prayer
	Magic
	Cooking
	Woodcutting
	Fletching
	Fishing
	Firemaking
	Crafting
	Smithing
	Mining
	Herblore
	Agility
	Thieving
	Slayer
	Farming
	Runecrafting
	Hunter
	Construction
	Summoning
	Dungeoneering
	Divination
	Invention
	Archaeology
)

// GetVirtualLevel gets the virtual level for the given XP
func (s Skill) GetVirtualLevel(XP int64) int {
	table := constants.StandardXPTable

	if s.IsEliteSkill() {
		table = constants.EliteXPTable
	}

	return sort.Search(len(table), func(i int) bool {
		return table[i] > XP
	})
}

// IsStandardSkill determines if the provided skill is a standard skill
func (s Skill) IsStandardSkill() bool {
	if s == Invention {
		return false
	}

	return true
}

// IsEliteSkill determines if the provided skill is a elite skill
func (s Skill) IsEliteSkill() bool {
	return !s.IsStandardSkill()
}
