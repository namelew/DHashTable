package hashtable

import (
	"errors"

	"github.com/tidwall/btree"
	"golang.org/x/exp/constraints"
)

type Key[K constraints.Ordered] interface {
	Value() K
	Index() int
}

type HashTable[I constraints.Ordered, K any] interface {
	hash(id Key[I]) int
	Insert(id Key[I], data K) error
	Delete(id Key[I]) error
	Search(id Key[I]) (K, error)
}

type Linked[I constraints.Ordered, K any] struct {
	size  int
	slots []btree.Map[I, K]
}

type Open[I constraints.Ordered, K any] struct {
	size  int
	slots []K
}

func (t *Linked[I, K]) init(size int) {
	if t.size == 0 {
		t.size = size
	}
}

func (t *Linked[I, K]) hash(id Key[I]) int {
	return id.Index() % t.size
}

func (t *Linked[I, K]) Insert(id Key[I], data K) error {
	t.slots[t.hash(id)].Set(id.Value(), data)
	return nil
}

func (t *Linked[I, K]) Delete(id Key[I]) error {
	_, notFound := t.slots[t.hash(id)].Delete(id.Value())

	if notFound {
		return errors.New("delete action failed: unable to find register")
	}

	return nil
}

func (t *Linked[I, K]) Search(id Key[I]) (K, error) {
	result, notFound := t.slots[t.hash(id)].Get(id.Value())

	if notFound {
		return result, errors.New("query action failed: unable to find register")
	}

	return result, nil
}

func (t *Open[I, K]) init(size int) {
	if t.size == 0 {
		t.size = size
	}
}

func (t *Open[I, K]) hash(id Key[I]) int {
	return id.Index() % t.size
}

func (t *Open[I, K]) Insert(id Key[I], data K) error {
	return nil
}

func (t *Open[I, K]) Delete(id Key[I]) error {
	return nil
}

func (t *Open[I, K]) Search(id Key[I]) (K, error) {
	var result K
	return result, nil
}
