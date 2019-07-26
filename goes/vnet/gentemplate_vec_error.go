// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=vnet -id error -d VecType=errVec -d Type=err github.com/platinasystems/go/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vnet

import (
	"github.com/platinasystems/go/elib"
)

type errVec []err

func (p *errVec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]err, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *errVec) validate(new_len uint, zero err) *err {
	old_cap := uint(cap(*p))
	old_len := uint(len(*p))
	if new_len <= old_cap {
		// Need to reslice to larger length?
		if new_len > old_len {
			*p = (*p)[:new_len]
			for i := old_len; i < new_len; i++ {
				(*p)[i] = zero
			}
		}
		return &(*p)[new_len-1]
	}
	return p.validateSlowPath(zero, old_cap, new_len, old_len)
}

func (p *errVec) validateSlowPath(zero err, old_cap, new_len, old_len uint) *err {
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]err, new_cap, new_cap)
		copy(q, *p)
		for i := old_len; i < new_cap; i++ {
			q[i] = zero
		}
		*p = q[:new_len]
	}
	if new_len > old_len {
		*p = (*p)[:new_len]
	}
	return &(*p)[new_len-1]
}

func (p *errVec) Validate(i uint) *err {
	var zero err
	return p.validate(i+1, zero)
}

func (p *errVec) ValidateInit(i uint, zero err) *err {
	return p.validate(i+1, zero)
}

func (p *errVec) ValidateLen(l uint) (v *err) {
	if l > 0 {
		var zero err
		v = p.validate(l, zero)
	}
	return
}

func (p *errVec) ValidateLenInit(l uint, zero err) (v *err) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *errVec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p errVec) Len() uint { return uint(len(p)) }