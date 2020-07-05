# query

```go
import "github.com/ebarkie/http/query"
```

Package query helps build HTTP query key/value pairs by indexing keys and
untyping values.

## Usage

#### type Data

```go
type Data struct {
}
```

Data holds the query data.

#### func (*Data) Clear

```go
func (d *Data) Clear()
```
Clear clears all data that was previously set.

#### func (Data) Exists

```go
func (d Data) Exists(key string) bool
```
Exists indicates if a given key has been set.

#### func (*Data) SetFloat

```go
func (d *Data) SetFloat(key string, f float64)
```
SetFloat adds a float value.

#### func (*Data) SetFloatf

```go
func (d *Data) SetFloatf(g Generator, f float64)
```
SetFloatf adds a float value using a generated key.

#### func (*Data) SetInt

```go
func (d *Data) SetInt(key string, i int)
```
SetInt adds an integer value.

#### func (*Data) SetIntf

```go
func (d *Data) SetIntf(g Generator, i int)
```
SetIntf adds an integer value using a generated key.

#### func (*Data) SetString

```go
func (d *Data) SetString(key string, s string)
```
SetString adds a string value.

#### func (*Data) SetStringf

```go
func (d *Data) SetStringf(g Generator, s string)
```
SetStringf adds a string value using a generated key.

#### func (Data) Values

```go
func (d Data) Values() map[string]string
```
Values returns all data key/value pairs that were previously set.

#### type Generator

```go
type Generator interface {
	Generate(Values) (key string)
}
```

Generator is an interface for adding key generators which are invoked when using
the Set* methods with the f (formatted) suffix.

#### type Indexed

```go
type Indexed struct {
	Format string // Template where the sharp symbol(#) will be substituted for the number
	Begin  int    // Begin index number
	Width  int    // Index number width
	Zero   int    // Any index number below the zero value not be printed (it's implied)
}
```

Indexed generates an indexed key based on formatting criteria.

Examples:

    WeatherCloud
    temp   First outdoor temperature
    temp02 Second outdoor temperature

    Indexed{Format: "temp#", Begin: 1, Zero: 1, Width: 2}

    Weather Underground
    tempf  First outdoor temperature
    temp2f Second outdoor temperature

    Indexed{Format: "temp#f", Begin: 1, Zero: 1}

#### func (Indexed) Generate

```go
func (i Indexed) Generate(v Values) (key string)
```
Generate generates a key using the formatting criteria and finding the lowest a
vailable index.

#### type Values

```go
type Values interface {
	Clear()
	Exists(key string) bool
	Values() map[string]string
}
```

Values is an high level interface for accessing the query values.
