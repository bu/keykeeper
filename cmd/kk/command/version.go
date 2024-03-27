package command

type Version struct {
}

func (v *Version) Run(*Env) error {
	println("kk version 0.1.0")
	return nil
}
