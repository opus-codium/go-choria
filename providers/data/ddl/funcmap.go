// Copyright (c) 2021, R.I. Pienaar and the Choria Project contributors
//
// SPDX-License-Identifier: Apache-2.0

package ddl

type FuncMapEntry struct {
	F    interface{}
	Name string
	DDL  *DDL
}

type FuncMap map[string]FuncMapEntry
