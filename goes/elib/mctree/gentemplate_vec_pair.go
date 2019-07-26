// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=mctree -id Pair -d VecType=pair_vec -d Type=Pair github.com/platinasystems/go/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mctree

import (
	"github.com/platinasystems/go/elib"
)

type pair_vec []Pair

func (p *pair_vec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]Pair, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *pair_vec) validate(new_len uint, zero Pair) *Pair {
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

func (p *pair_vec) validateSlowPath(zero Pair, old_cap, new_len, old_len uint) *Pair {
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]Pair, new_cap, new_cap)
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

func (p *pair_vec) Validate(i uint) *Pair {
	var zero Pair
	return p.validate(i+1, zero)
}

func (p *pair_vec) ValidateInit(i uint, zero Pair) *Pair {
	return p.validate(i+1, zero)
}

func (p *pair_vec) ValidateLen(l uint) (v *Pair) {
	if l > 0 {
		var zero Pair
		v = p.validate(l, zero)
	}
	return
}

func (p *pair_vec) ValidateLenInit(l uint, zero Pair) (v *Pair) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *pair_vec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p pair_vec) Len() uint { return uint(len(p)) }