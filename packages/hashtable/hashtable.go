package hashtable

import (
	"errors"

	"github.com/tidwall/btree"
	"golang.org/x/exp/constraints"
)

type Table[I constraints.Ordered, K any] struct {
	size  int
	slots []btree.Map[I, K]
}

type Key[K constraints.Ordered] interface {
	Value() K
	Index() int
}

func (t *Table[I, K]) hash(id Key[I]) int {
	return id.Index() % t.size
}

func (t *Table[I, K]) Write(id Key[I], data K) {
	t.slots[t.hash(id)].Set(id.Value(), data)
}

func (t *Table[I, K]) Delete(id Key[I]) error {
	_, notFound := t.slots[t.hash(id)].Delete(id.Value())

	if notFound {
		return errors.New("delete action failed: unable to find register")
	}

	return nil
}

func (t *Table[I, K]) Get(id Key[I]) (K, error) {
	result, notFound := t.slots[t.hash(id)].Get(id.Value())

	if notFound {
		return result, errors.New("query action failed: unable to find register")
	}

	return result, nil
}
