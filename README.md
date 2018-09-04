# GopherVenture
__A single player text adventure game for all ages
made in Go from the ground up.__

You are a gopher who ended up in a foreign universe.
Along the way, you try to find the exit to get back to
your normal world. But beware, as there's monsters along
your journey that you must fight.

### TO-DOs:
* A more developed plot is yet to be developed.
* Implement RNG fight mechanics for monster battles.
Currently experimenting with mixing `time` and
`math/rand` to make the pseudo-system more varied.
Might change to a RNG that loops over and over until it
is called and needed, spews a number, and loops again.
Some systems might not take advantage of clock-based RNG,
so a looping RNG function might be implemented in place.
* Player name would still contain line breaks.
Maybe start using `fmt.Scanln()` instead of a reader for
names? But it can only grab one word... This might need
some experimentation.