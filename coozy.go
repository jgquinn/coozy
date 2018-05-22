package coozy

import (
	"encoding/json"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
)

// Version contains the package version.
const Version = "0.0.4-pre"

// WhereLiker defines the WhereLike() method used by FindPop.
type WhereLiker interface {
	WhereLike(q *pop.Query) *pop.Query
}

// NullableJSONMap returns map poplated from not-null JSON string when provided.
func NullableJSONMap(ns nulls.String) (m map[string]string) {
	if ns.Valid {
		m = make(map[string]string)
		_ = json.Unmarshal([]byte(ns.String), &m)
	}
	return
}

// SavePop saves a record struct into the named environment.
func SavePop(popEnvironmentName string, rec interface{}) (err error) {

	var tx *pop.Connection
	tx, err = pop.Connect(popEnvironmentName)
	if err != nil {
		return
	}

	err = SaveConnectedPop(tx, rec)
	return
}

// SaveConnectedPop saves a record struct into the provided Pop connection.
func SaveConnectedPop(tx *pop.Connection, rec interface{}) (err error) {
	err = tx.Save(rec)
	return
}

// FindPop queries the named environment for records based on the supplied criteria.
func FindPop(popEnvironmentName string, criteria WhereLiker, recs interface{}, eagerFetchFields ...string) (err error) {

	var tx *pop.Connection
	tx, err = pop.Connect(popEnvironmentName)
	if err != nil {
		return
	}

	err = FindConnectedPop(tx, criteria, recs, eagerFetchFields...)
	return
}

// FindConnectedPop queries the named environment for records based on the supplied criteria.
func FindConnectedPop(tx *pop.Connection, criteria WhereLiker, recs interface{}, eagerFetchFields ...string) (err error) {
	var q *pop.Query
	if len(eagerFetchFields) > 0 {
		q = tx.Eager(eagerFetchFields...)
	} else {
		q = tx.Q()
	}

	err = criteria.WhereLike(q).All(recs)
	return
}
