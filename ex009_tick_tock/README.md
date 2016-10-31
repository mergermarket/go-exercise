# Tick Tock

At GoLunch we looked at [Tickers]. A `Ticker` is a simple struct around a `<-chan
Time` (a recieve-only channel that emits things of type `Time`), that emits the
current time at a regular interval.

There is an even simpler implementation of the same functionality: the function
[`Tick`] returns a plain `<-chan Time` which also emits the current Time at
a regular interval. The signature for `Tick` is:

```
func Tick(d Duration) <-chan Time
```

Write your own implementation of `Tick` which behaves in the same way. Try and
explain why Go created a separate `Ticker` type, and `NewTicker` function. What
benefits do they provide that `Tick` does not?


[Tickers]: https://golang.org/pkg/time/#Ticker
[NewTicker]: https://golang.org/pkg/time/#NewTicker
[`Tick`]: https://golang.org/pkg/time/#Tick
