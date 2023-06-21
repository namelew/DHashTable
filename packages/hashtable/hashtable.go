package hashtable

type Table[I any, K any] struct {
	size int
}

type Key[K any] interface {
	Index() int
}

func (t *Table[I, K]) hash(id Key[I]) int {
	return id.Index() % t.size
}

func (t *Table[I, K]) Insert(id Key[I], data K) error {
	return nil
}

func (t *Table[I, K]) Remove(id Key[I]) error {
	return nil
}

func (t *Table[I, K]) Get(id Key[I]) (K, error) {
	var result K
	return result, nil
}

func (t *Table[I, K]) Update(id Key[I], new K) error {
	return nil
}
