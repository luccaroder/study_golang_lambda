package main

type CommandError struct {
	Err error
}

func (r *CommandError) Error() string {
	return r.Err.Error()
}
