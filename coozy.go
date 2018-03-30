package coozy

import (
	"github.com/gobuffalo/pop"
)

// Version contains the package version.
const Version = "0.0.1"

// WhereLiker defines the WhereLike() method used by FindPop.
type WhereLiker interface {
	WhereLike(q *pop.Query) *pop.Query
}

// SavePop saves a record struct into the named environment.
func SavePop(popEnvironmentName string, rec interface{}) (err error) {

	var tx *pop.Connection
	tx, err = pop.Connect(popEnvironmentName)
	if err != nil {
		return
	}

	err = tx.Save(rec)
	if err != nil {
		return
	}

	return
}

// FindPop queries the named environment for records based on the supplied criteria.
func FindPop(popEnvironmentName string, criteria WhereLiker, recs interface{}, eagerFetchFields ...string) (err error) {

	var tx *pop.Connection
	tx, err = pop.Connect(popEnvironmentName)
	if err != nil {
		return
	}

	var q *pop.Query
	if len(eagerFetchFields) > 0 {
		q = tx.Eager(eagerFetchFields...)
	} else {
		q = tx.Q()
	}

	err = criteria.WhereLike(q).All(recs)
	return
}
