package coozy

import (
	"time"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
)

// StringsValidEqual returns true if both a and b are valid and contain equal values.
func StringsValidEqual(a *nulls.String, b *nulls.String) bool {
	return (a.Valid && b.Valid && a.String == b.String)
}

// UUIDValidEqual returns true if both a and b are non-zero and have equal values.
func UUIDValidEqual(a uuid.UUID, b uuid.UUID) bool {
	return (a != uuid.Nil && b != uuid.Nil && uuid.Equal(a, b))
}

// TimesValidEqual returns true if both a and b are valid and contain equal values.
func TimesValidEqual(a *nulls.Time, b *nulls.Time) bool {
	return (a.Valid && b.Valid && a.Time.Equal(b.Time))
}

// TimesVary returns true if either a or b are valid and contain unequal values.
func TimesVary(a *nulls.Time, b *nulls.Time) (result bool) {

	if a.Valid != true && b.Valid != true {
		result = false
		return
	}

	if a.Valid && b.Valid {
		result = !a.Time.Equal(b.Time)
	} else {
		result = true
	}

	return
}

// TimesVaryMoreThan returns true if either a or b are valid and contain values that differ greater than the specified skew.
func TimesVaryMoreThan(a *nulls.Time, b *nulls.Time, skew time.Duration) (result bool) {

	if a.Valid != true && b.Valid != true {
		result = false
		return
	}

	if a.Valid && b.Valid {
		diff := a.Time.Sub(b.Time)
		if diff < 0 {
			diff = diff * -1
		}
		result = diff > skew
	} else {
		result = true
	}

	return
}
