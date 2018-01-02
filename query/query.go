// Copyright (c) 2016-2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package query helps build HTTP query key/value pairs by
// indexing keys and untyping values.
package query

// Generator is an interface for adding key generators which are
// invoked when using the Set* methods with the f (formatted) suffix.
type Generator interface {
	Generate(Values) (key string)
}

// Values is an high level interface for accessing the query values.
type Values interface {
	Clear()
	Exists(key string) bool
	Values() map[string]string
}
