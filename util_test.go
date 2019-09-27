package coozy

import (
	"testing"
	"time"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gofrs/uuid"
)

func TestStringsValidEqual(t *testing.T) {
	nv := nulls.String{}
	va := nulls.NewString("apple")
	vA := nulls.NewString("apple")
	vb := nulls.NewString("banana")

	if StringsValidEqual(&nv, &vb) || StringsValidEqual(&va, &nv) {
		t.Error("invalid parameters should return false")
	}

	if StringsValidEqual(&va, &vb) {
		t.Error("mismatch parameters should return false")
	}

	if !StringsValidEqual(&va, &vA) {
		t.Error("matching parameters should return true")
	}
}

func TestUUIDValidEqual(t *testing.T) {

	nv := uuid.Nil

	va, err := uuid.NewV1()
	if err != nil {
		t.Fatal(err)
	}

	vA, err := uuid.FromString(va.String())
	if err != nil {
		t.Fatal(err)
	}

	vb, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}

	if UUIDValidEqual(va, nv) || UUIDValidEqual(nv, vb) {
		t.Error("invalid parameters should return false")
	}

	if UUIDValidEqual(va, vb) {
		t.Error("mismatch parameters should return false")
	}

	if !UUIDValidEqual(va, vA) {
		t.Error("matching parameters should return true")
	}
}

func TestTimesValidEqual(t *testing.T) {
	nv := nulls.Time{}
	tm := time.Now()
	va := nulls.NewTime(tm)
	vA := nulls.NewTime(tm)
	tm = tm.Add(-36 * time.Hour)
	vb := nulls.NewTime(tm)

	if TimesValidEqual(&nv, &vb) || TimesValidEqual(&va, &nv) {
		t.Error("invalid parameters should return false")
	}

	if TimesValidEqual(&va, &vb) {
		t.Error("mismatch parameters should return false")
	}

	if !TimesValidEqual(&va, &vA) {
		t.Error("matching parameters should return true")
	}
}

func TestTimesVary(t *testing.T) {
	nv := nulls.Time{}
	tm := time.Now()
	va := nulls.NewTime(tm)
	vA := nulls.NewTime(tm)
	tm = tm.Add(-36 * time.Hour)
	vb := nulls.NewTime(tm)

	if (!TimesVary(&nv, &vb)) || (!TimesVary(&va, &nv)) {
		t.Error("invalid parameters should return true")
	}

	if !TimesVary(&va, &vb) {
		t.Error("mismatch parameters should return true")
	}

	if TimesVary(&va, &vA) {
		t.Error("matching parameters should return false")
	}
}

func TestTimesVaryMoreThan(t *testing.T) {
	nv := nulls.Time{}
	tm := time.Now()
	va := nulls.NewTime(tm)
	vA := nulls.NewTime(tm.Add(time.Hour * 12))
	vb := nulls.NewTime(tm.Add(time.Hour * 36))

	day := time.Hour * 24

	if (!TimesVaryMoreThan(&nv, &vb, day)) || (!TimesVaryMoreThan(&va, &nv, day)) {
		t.Error("invalid parameters should return true")
	}

	if (!TimesVaryMoreThan(&va, &vb, day)) || (!TimesVaryMoreThan(&vb, &va, day)) {
		t.Error("mismatch parameters should return true")
	}

	if TimesVaryMoreThan(&va, &vA, day) || TimesVaryMoreThan(&vA, &va, day) {
		t.Error("matching parameters should return false")
	}
}
