package dbx

import "github.com/upper/db/v4"

type Q[T any, ID any] struct {
	db        db.Session
	TableName string
}

func NewQ[T any, ID any](db db.Session, table string) *Q[T, ID] {
	return &Q[T, ID]{db: db, TableName: table}
}

// Get get a record by id
func (q *Q[T, ID]) Get(id ID) (*T, error) {
	return q.GetBy("id", id)
}

func (q *Q[T, ID]) GetMulti(ids []ID) (ts []T, err error) {
	return
}

func (q *Q[T, ID]) GetBy(field string, value any) (*T, error) {
	var t T
	err := q.db.Collection(q.TableName).Find(db.Cond{field: value}).One(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (q *Q[T, ID]) LikeBy(prefix string) ([]T, error) {
	return nil, nil
}

// Save save a record
func (q *Q[T, ID]) Save(t db.Record) error {
	return q.db.Save(t)
}
