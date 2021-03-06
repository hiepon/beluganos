// -*- coding: utf-8 -*-

// Copyright (C) 2018 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"
	"strings"
	"sync"

	lib "fabricflow/fibs/fibslib"
)

//
// OidMapEntry is OidMapTable entry
//
type OidMapEntry struct {
	Name      string
	Global    string
	Local     string
	GlobalOid []uint
	LocalOid  []uint
}

//
// NewOidMapEntry creates new entry.
//
func NewOidMapEntry(name, gOid, lOid string) *OidMapEntry {
	return &OidMapEntry{
		Name:      name,
		Global:    gOid,
		Local:     lOid,
		GlobalOid: lib.ParseOID(gOid),
		LocalOid:  lib.ParseOID(lOid),
	}
}

func oidHasPrefix(oid string, prefix string) bool {
	if ok := strings.HasSuffix(oid, "."); !ok {
		oid = fmt.Sprintf("%s.", oid)
	}

	if ok := strings.HasSuffix(prefix, "."); !ok {
		prefix = fmt.Sprintf("%s.", prefix)
	}

	return strings.HasPrefix(oid, prefix)
}

//
// String returns reable contenf.
//
func (e *OidMapEntry) String() string {
	return fmt.Sprintf("%s global:'%s', local:'%s'", e.Name, e.Global, e.Local)
}

//
// MatchGlobal compares with global oid.
//
func (e *OidMapEntry) MatchGlobal(oid string) bool {
	return oidHasPrefix(oid, e.Global)
}

//
// MatchLocal compares with local oid.
//
func (e *OidMapEntry) MatchLocal(oid string) bool {
	return oidHasPrefix(oid, e.Local)
}

//
// OidMapTable is table for Global/Local exchange.
//
type OidMapTable struct {
	entries []*OidMapEntry
	mutex   sync.Mutex
}

//
// NewOidMapTable creates new table.
//
func NewOidMapTable() *OidMapTable {
	return &OidMapTable{
		entries: []*OidMapEntry{},
	}
}

func (t *OidMapTable) WriteTo(w io.Writer) (sum int64, err error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for _, e := range t.entries {
		var n int
		n, err = fmt.Fprintf(w, "OidMap %s\n", e)
		sum += int64(n)
		if err != nil {
			return
		}
	}
	return
}

//
// matchByGlobal finds entry(global oid).
//
func (t *OidMapTable) matchByGlobal(oid string) (*OidMapEntry, bool) {
	for _, e := range t.entries {
		if e.MatchGlobal(oid) {
			return e, true
		}
	}
	return nil, false
}

//
// MatchByGlobal finds entry(global oid).
//
func (t *OidMapTable) MatchByGlobal(oid string) (*OidMapEntry, bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.matchByGlobal(oid)
}

//
// matchByGlobal finds entry(local oid)
//
func (t *OidMapTable) matchByLocal(oid string) (*OidMapEntry, bool) {
	for _, e := range t.entries {
		if e.MatchLocal(oid) {
			return e, true
		}
	}
	return nil, false
}

//
// MatchByGlobal finds entry(local oid)
//
func (t *OidMapTable) MatchByLocal(oid string) (*OidMapEntry, bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.matchByLocal(oid)
}

//
// add appends entry.
//
func (t *OidMapTable) add(entry *OidMapEntry) bool {
	for _, e := range t.entries {
		if e.Global == entry.Global {
			return false
		}
		if e.Local == entry.Local {
			return false
		}
	}

	t.entries = append(t.entries, entry)
	return true
}

//
// Add appends entry.
//
func (t *OidMapTable) Add(entry *OidMapEntry) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.add(entry)
}
