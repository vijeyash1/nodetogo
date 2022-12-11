// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColUInt128 represents UInt128 column.
type ColUInt128 []UInt128

// Compile-time assertions for ColUInt128.
var (
	_ ColInput  = ColUInt128{}
	_ ColResult = (*ColUInt128)(nil)
	_ Column    = (*ColUInt128)(nil)
)

// Rows returns count of rows in column.
func (c ColUInt128) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColUInt128) Reset() {
	*c = (*c)[:0]
}

// Type returns ColumnType of UInt128.
func (ColUInt128) Type() ColumnType {
	return ColumnTypeUInt128
}

// Row returns i-th row of column.
func (c ColUInt128) Row(i int) UInt128 {
	return c[i]
}

// Append UInt128 to column.
func (c *ColUInt128) Append(v UInt128) {
	*c = append(*c, v)
}

// LowCardinality returns LowCardinality for UInt128 .
func (c *ColUInt128) LowCardinality() *ColLowCardinality[UInt128] {
	return &ColLowCardinality[UInt128]{
		index: c,
	}
}

// Array is helper that creates Array of UInt128.
func (c *ColUInt128) Array() *ColArr[UInt128] {
	return &ColArr[UInt128]{
		Data: c,
	}
}

// Nullable is helper that creates Nullable(UInt128).
func (c *ColUInt128) Nullable() *ColNullable[UInt128] {
	return &ColNullable[UInt128]{
		Values: c,
	}
}

// NewArrUInt128 returns new Array(UInt128).
func NewArrUInt128() *ColArr[UInt128] {
	return &ColArr[UInt128]{
		Data: new(ColUInt128),
	}
}