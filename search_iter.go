package neon

type SearchIter[T any] struct {
	cur             *T
	err             error
	searchContainer SearchContainer
	meta            *SearchMeta
	values          []T
}

func (it *SearchIter[T]) Current() *T {
	return it.cur
}

func (it *SearchIter[T]) SearchResult() SearchContainer {
	return it.searchContainer
}

func (it *SearchIter[T]) getPage() {
	it.meta = it.searchContainer.GetSearchMeta()
}

func (it *SearchIter[T]) Meta() *SearchMeta {
	return it.meta
}
