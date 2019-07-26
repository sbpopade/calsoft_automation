// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=elib -id Int32 -d VecType=Int32Vec -d Type=int32 vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elib

type Int32Vec []int32

func (p *Int32Vec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := NextResizeCap(new_len)
		q := make([]int32, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *Int32Vec) validate(new_len uint, zero int32) *int32 {
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

func (p *Int32Vec) validateSlowPath(zero int32, old_cap, new_len, old_len uint) *int32 {
	if new_len > old_cap {
		new_cap := NextResizeCap(new_len)
		q := make([]int32, new_cap, new_cap)
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

func (p *Int32Vec) Validate(i uint) *int32 {
	var zero int32
	return p.validate(i+1, zero)
}

func (p *Int32Vec) ValidateInit(i uint, zero int32) *int32 {
	return p.validate(i+1, zero)
}

func (p *Int32Vec) ValidateLen(l uint) (v *int32) {
	if l > 0 {
		var zero int32
		v = p.validate(l, zero)
	}
	return
}

func (p *Int32Vec) ValidateLenInit(l uint, zero int32) (v *int32) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *Int32Vec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p Int32Vec) Len() uint { return uint(len(p)) }