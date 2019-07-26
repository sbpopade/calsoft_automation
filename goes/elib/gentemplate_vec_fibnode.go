// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=elib -id fibNode -d VecType=fibNodeVec -d Type=fibNode vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elib

type fibNodeVec []fibNode

func (p *fibNodeVec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := NextResizeCap(new_len)
		q := make([]fibNode, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *fibNodeVec) validate(new_len uint, zero fibNode) *fibNode {
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

func (p *fibNodeVec) validateSlowPath(zero fibNode, old_cap, new_len, old_len uint) *fibNode {
	if new_len > old_cap {
		new_cap := NextResizeCap(new_len)
		q := make([]fibNode, new_cap, new_cap)
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

func (p *fibNodeVec) Validate(i uint) *fibNode {
	var zero fibNode
	return p.validate(i+1, zero)
}

func (p *fibNodeVec) ValidateInit(i uint, zero fibNode) *fibNode {
	return p.validate(i+1, zero)
}

func (p *fibNodeVec) ValidateLen(l uint) (v *fibNode) {
	if l > 0 {
		var zero fibNode
		v = p.validate(l, zero)
	}
	return
}

func (p *fibNodeVec) ValidateLenInit(l uint, zero fibNode) (v *fibNode) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *fibNodeVec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p fibNodeVec) Len() uint { return uint(len(p)) }
