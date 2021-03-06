// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package terminal

import (
	"fmt"
	"strings"
)

type TableInfo struct {
	maxColNum  int
	maxColWide []int
	infos      []string
	heads      []string
	records    [][]string
}

func (t *TableInfo) Infos(infos ...string) {
	t.infos = append(t.infos, infos...)
}

func (t *TableInfo) Add(items ...string) {
	for i, item := range items {
		if i < t.maxColNum && t.maxColWide[i] < len(item) {
			t.maxColWide[i] = len(item)
		}
	}
	t.records = append(t.records, items)
}

func (t *TableInfo) String() string {
	sb := &strings.Builder{}

	// information print
	for _, info := range t.infos {
		sb.WriteString(info)
		sb.WriteByte('\n')
	}
	if len(t.infos) > 0 && len(t.records) > 0 {
		sb.WriteByte('\n')
	}

	// table header info
	if len(t.records) == 0 {
		return sb.String()
	}
	lastColIdx := len(t.heads) - 1
	for i, head := range t.heads {
		t.writeItem(sb, i, head)
		if i == lastColIdx {
			sb.WriteByte('\n')
		}
	}

	// record info print
	lastRecordIdx := len(t.records) - 1
	for i, recItems := range t.records {
		var idx int
		for ri, item := range recItems {
			t.writeItem(sb, ri, item)
			idx = ri
		}
		if idx <= lastColIdx && i != lastRecordIdx {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func (t *TableInfo) writeItem(sb *strings.Builder, idxWide int, item string) {
	sb.WriteString(fmt.Sprintf("%-*s", t.maxColWide[idxWide]+2, item))
}

func NewTableInfo(heads ...string) *TableInfo {
	ti := &TableInfo{
		maxColNum:  len(heads),
		maxColWide: make([]int, len(heads)),
		heads:      heads,
	}
	for i, head := range heads {
		ti.maxColWide[i] = len(head)
	}
	return ti
}
