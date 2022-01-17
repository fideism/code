package nginx

import (
	"errors"
)

// Servers ...
type Servers struct {
	server int
	rst    []string
}

// Add ...
func (r *Servers) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}
	addr := params[0]
	r.rst = append(r.rst, addr)

	return nil
}

// Next ...
func (r *Servers) Next() string {
	if len(r.rst) == 0 {
		return ""
	}
	lens := len(r.rst)
	if r.server >= lens {
		r.server = 0
	}
	curAdd := r.rst[r.server]
	r.server = (r.server + 1) % lens
	return curAdd
}
