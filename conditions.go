package main

type Conditions struct {
	Branch string
	Owner  bool
	Admin  bool
	Master bool
}

func (c Conditions) AreMet(p PushEvent) bool {
	if c.Branch != "" && p.Branch() != c.Branch {
		return false
	}

	if c.Owner && !p.ByOwner() {
		return false
	}

	if c.Admin && !p.ByAdmin() {
		return false
	}

	if c.Master && !p.ToMaster() {
		return false
	}

	return true
}
