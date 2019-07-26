// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=elib -id Word -d VecType=WordVec -d Type=Word vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elib

type WordVec []Word

func (p *WordVec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := NextResizeCap(new_len)
		q := make([]Word, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *WordVec) validate(new_len uint, zero Word) *Word {
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

func (p *WordVec) validateSlowPath(zero Word, old_cap, new_len, old_len uint) *Word {
	if new_len > old_cap {
		new_cap := NextResizeCap(new_len)
		q := make([]Word, new_cap, new_cap)
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

func (p *WordVec) Validate(i uint) *Word {
	var zero Word
	return p.validate(i+1, zero)
}

func (p *WordVec) ValidateInit(i uint, zero Word) *Word {
	return p.validate(i+1, zero)
}

func (p *WordVec) ValidateLen(l uint) (v *Word) {
	if l > 0 {
		var zero Word
		v = p.validate(l, zero)
	}
	return
}

func (p *WordVec) ValidateLenInit(l uint, zero Word) (v *Word) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *WordVec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p WordVec) Len() uint { return uint(len(p)) }
