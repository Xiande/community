package viewmodel

type PagingListForTQ struct {
	PagingList
	IsMyFavorite bool
	TagID        int
	TagTitle     string
}
