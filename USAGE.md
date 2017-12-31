# request
```
import "github.com/ebarkie/request"
```

Package request helps build a grouping of key/value pairs for a data set. It
simplifies indexing keys and untyping data.

## Usage

#### func  Float

```go
func Float(f float64) string
```
Float converts a float to a string.

#### func  Int

```go
func Int(i int) string
```
Int converts a integer to a string.

#### type Data

```go
type Data struct {
}
```

Data represents a group of key/value pairs.

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
	Generate(Data) (key string)
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

    Weather Underground
    tempf  First outdoor temperature
    temp2f Second outdoor temperature

    Indexed{Format: "temp#f", Begin: 1, Zero: 2}

    WeatherCloud
    temp01 First outdoor temperature
    temp02 Second outdoor temperature

    Indexed{Format: "temp#", Begin: 1, Width: 2}

#### func (Indexed) Generate

```go
func (i Indexed) Generate(d Data) (key string)
```
Generate generates a key using the formatting criteria and finding the lowest a
vailable index.
