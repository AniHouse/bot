package clubs

import "time"

type cfg struct {
	Price               uint          `yaml:"price,omitempty"`
	MinimumMembers      uint          `yaml:"minimum_members,omitempty"`
	NotVerifiedLifetime time.Duration `yaml:"not_verified_lifetime,omitempty"`
}
