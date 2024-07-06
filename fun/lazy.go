package fun

/*
Performs an arbitrary process and blocks as long as it takes.
Returns arbitrary data.
*/
type Lazy[S any] func() S

/*
Performs an arbitrary process and blocks as long as it takes.
Returns arbitrary data and errors.
*/
type LazyWithError[S any] func() (S, error)
