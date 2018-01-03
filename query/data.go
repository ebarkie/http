// Copyright (c) 2016 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package query

import "strconv"

// Data holds the query data.
type Data struct {
	m map[string]string
}

// Clear clears all data that was previously set.
func (d *Data) Clear() {
	d.m = make(map[string]string)
}

// Exists indicates if a given key has been set.
func (d Data) Exists(key string) bool {
	_, exists := d.m[key]
	return exists
}

// SetFloat adds a float value.
func (d *Data) SetFloat(key string, f float64) {
	d.SetString(key, strconv.FormatFloat(f, 'f', -1, 64))
}

// SetFloatf adds a float value using a generated key.
func (d *Data) SetFloatf(g Generator, f float64) {
	d.SetFloat(g.Generate(d), f)
}

// SetInt adds an integer value.
func (d *Data) SetInt(key string, i int) {
	d.SetString(key, strconv.Itoa(i))
}

// SetIntf adds an integer value using a generated key.
func (d *Data) SetIntf(g Generator, i int) {
	d.SetInt(g.Generate(d), i)
}

// SetString adds a string value.
func (d *Data) SetString(key string, s string) {
	if d.m == nil {
		d.Clear()
	}
	d.m[key] = s
}

// SetStringf adds a string value using a generated key.
func (d *Data) SetStringf(g Generator, s string) {
	d.SetString(g.Generate(d), s)
}

// Values returns all data key/value pairs that were previously set.
func (d Data) Values() map[string]string {
	return d.m
}
