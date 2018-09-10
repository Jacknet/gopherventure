# GopherVenture
__A single player text adventure game for all ages
made in Go from the ground up.__

You are a gopher who ended up in a foreign universe.
Along the way, you try to find the exit to get back to
your normal world. But beware, as there's monsters along
your journey that you must fight.

### TO-DOs:
* A more developed plot is yet to be developed.

In search of storywriters.

~~* Implement RNG fight mechanics for monster battles.~~

RNG mechanic successfully added for first battle! Game
generates a number seed at the start.

* Player name would still contain line breaks.

Maybe start using `fmt.Scanln()` instead of a reader for
names? But it can only grab one word... This might need
some experimentation. At the moment, this issue us under
low priority as it does not affect normal gameplay.