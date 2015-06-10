//go:generate sh -c "jsonenums -type=UintDF; jsonenums -type=UintPS"
package enums

// Constants for DF sorting criterion.
const (
	FS UintDF = iota
	MP
	TOTAL
	USED
	AVAIL
)

// Constants for PS sorting criterion.
const (
	PID UintPS = iota
	PRI
	NICE
	VIRT
	RES
	TIME
	NAME
	UID
	USER
)

// Uint-derived types:

// UintDF is a derived Uint for constants.
type UintDF Uint

// UintPS is a derived Uint for constants.
type UintPS Uint

// Uint is a positive or 0 number.
type Uint uint