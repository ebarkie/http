// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package request

// Generator is an interface for adding key generators which are
// invoked when using the Set* methods with the f (formatted) suffix.
type Generator interface {
	Generate(Data) (key string)
}

// Data represents a group of key/value pairs.
type Data struct {
	vals map[string]string
}

// Clear clears all data that was previously set.
func (d *Data) Clear() {
	d.vals = make(map[string]string)
}

// Exists indicates if a given key has been set.
func (d Data) Exists(key string) bool {
	_, exists := d.vals[key]
	return exists
}

// SetFloat adds a float value.
func (d *Data) SetFloat(key string, f float64) {
	d.SetString(key, Float(f))
}

// SetFloatf adds a float value using a generated key.
func (d *Data) SetFloatf(g Generator, f float64) {
	d.SetStringf(g, Float(f))
}

// SetInt adds an integer value.
func (d *Data) SetInt(key string, i int) {
	d.SetString(key, Int(i))
}

// SetIntf adds an integer value using a generated key.
func (d *Data) SetIntf(g Generator, i int) {
	d.SetStringf(g, Int(i))
}

// SetString adds a string value.
func (d *Data) SetString(key string, s string) {
	if d.vals == nil {
		d.Clear()
	}
	d.vals[key] = s
}

// SetStringf adds a string value using a generated key.
func (d *Data) SetStringf(g Generator, s string) {
	d.SetString(g.Generate(*d), s)
}

// Values returns all data key/value pairs that were previously set.
func (d Data) Values() map[string]string {
	return d.vals
}
