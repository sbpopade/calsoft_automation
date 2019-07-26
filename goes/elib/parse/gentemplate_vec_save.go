// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=parse -id save -d VecType=saveVec -d Type=save github.com/platinasystems/go/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse

import (
	"github.com/platinasystems/go/elib"
)

type saveVec []save

func (p *saveVec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]save, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *saveVec) validate(new_len uint, zero save) *save {
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

func (p *saveVec) validateSlowPath(zero save, old_cap, new_len, old_len uint) *save {
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]save, new_cap, new_cap)
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

func (p *saveVec) Validate(i uint) *save {
	var zero save
	return p.validate(i+1, zero)
}

func (p *saveVec) ValidateInit(i uint, zero save) *save {
	return p.validate(i+1, zero)
}

func (p *saveVec) ValidateLen(l uint) (v *save) {
	if l > 0 {
		var zero save
		v = p.validate(l, zero)
	}
	return
}

func (p *saveVec) ValidateLenInit(l uint, zero save) (v *save) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *saveVec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p saveVec) Len() uint { return uint(len(p)) }