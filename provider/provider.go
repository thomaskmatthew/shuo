package provider

import "time"

type Repo struct {
	Provider string
	ID       string
	FullName string
}

type Branch struct {
	Repo       Repo
	Name       string
	Author     string
	LastCommit time.Time
	Merged     bool // note: not all providers can populate this
	Protected  bool
	Default    bool
}
