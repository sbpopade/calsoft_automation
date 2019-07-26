// autogenerated: do not edit!
// generated from gentemplate [gentemplate -d Package=event -id timedEvent -d PoolType=timedEventPool -d Type=TimedActor -d Data=events github.com/platinasystems/go/elib/pool.tmpl]

// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package event

import (
	"github.com/platinasystems/go/elib"
)

type timedEventPool struct {
	elib.Pool
	events []TimedActor
}

func (p *timedEventPool) GetIndex() (i uint) {
	l := uint(len(p.events))
	i = p.Pool.GetIndex(l)
	if i >= l {
		p.Validate(i)
	}
	return i
}

func (p *timedEventPool) PutIndex(i uint) (ok bool) {
	return p.Pool.PutIndex(i)
}

func (p *timedEventPool) IsFree(i uint) (v bool) {
	v = i >= uint(len(p.events))
	if !v {
		v = p.Pool.IsFree(i)
	}
	return
}

func (p *timedEventPool) Resize(n uint) {
	c := uint(cap(p.events))
	l := uint(len(p.events) + int(n))
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]TimedActor, l, c)
		copy(q, p.events)
		p.events = q
	}
	p.events = p.events[:l]
}

func (p *timedEventPool) Validate(i uint) {
	c := uint(cap(p.events))
	l := uint(i) + 1
	if l > c {
		c = elib.NextResizeCap(l)
		q := make([]TimedActor, l, c)
		copy(q, p.events)
		p.events = q
	}
	if l > uint(len(p.events)) {
		p.events = p.events[:l]
	}
}

func (p *timedEventPool) Elts() uint {
	return uint(len(p.events)) - p.FreeLen()
}

func (p *timedEventPool) Len() uint {
	return uint(len(p.events))
}

func (p *timedEventPool) Foreach(f func(x TimedActor)) {
	for i := range p.events {
		if !p.Pool.IsFree(uint(i)) {
			f(p.events[i])
		}
	}
}

func (p *timedEventPool) ForeachIndex(f func(i uint)) {
	for i := range p.events {
		if !p.Pool.IsFree(uint(i)) {
			f(uint(i))
		}
	}
}

func (p *timedEventPool) Reset() {
	p.Pool.Reset()
	if len(p.events) > 0 {
		p.events = p.events[:0]
	}
}