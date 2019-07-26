// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=vnet -id ifThread -d VecType=ifThreadVec -d Type=*InterfaceThread github.com/platinasystems/go/elib/vec.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vnet

import (
	"github.com/platinasystems/go/elib"
)

type ifThreadVec []*InterfaceThread

func (p *ifThreadVec) Resize(n uint) {
	old_cap := uint(cap(*p))
	new_len := uint(len(*p)) + n
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]*InterfaceThread, new_len, new_cap)
		copy(q, *p)
		*p = q
	}
	*p = (*p)[:new_len]
}

func (p *ifThreadVec) validate(new_len uint, zero *InterfaceThread) **InterfaceThread {
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

func (p *ifThreadVec) validateSlowPath(zero *InterfaceThread, old_cap, new_len, old_len uint) **InterfaceThread {
	if new_len > old_cap {
		new_cap := elib.NextResizeCap(new_len)
		q := make([]*InterfaceThread, new_cap, new_cap)
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

func (p *ifThreadVec) Validate(i uint) **InterfaceThread {
	var zero *InterfaceThread
	return p.validate(i+1, zero)
}

func (p *ifThreadVec) ValidateInit(i uint, zero *InterfaceThread) **InterfaceThread {
	return p.validate(i+1, zero)
}

func (p *ifThreadVec) ValidateLen(l uint) (v **InterfaceThread) {
	if l > 0 {
		var zero *InterfaceThread
		v = p.validate(l, zero)
	}
	return
}

func (p *ifThreadVec) ValidateLenInit(l uint, zero *InterfaceThread) (v **InterfaceThread) {
	if l > 0 {
		v = p.validate(l, zero)
	}
	return
}

func (p *ifThreadVec) ResetLen() {
	if *p != nil {
		*p = (*p)[:0]
	}
}

func (p ifThreadVec) Len() uint { return uint(len(p)) }