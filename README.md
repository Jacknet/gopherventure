# GopherVenture
__A single player text adventure game for all ages
made in Go from the ground up.__

You are a gopher who ended up in a foreign universe.
Along the way, you try to find the exit to get back to
your normal world. But beware, as there's monsters along
your journey that you must fight.

### TO-DOs:
* A more developed plot is yet to be developed.
* At the moment, the project is split into two sources
due to the differences in line breaks on Windows and
Linux/*nix. Attempt to combine the two versions. Maybe
include code that checks what environment you're using?
* ~~Implement RNG fight mechanics for monster battles.~~
The `math/rand` package in Go is actually pseudo-random.
Maybe combining this with the `time` package might
give some more expected random values? To be experimented.

[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)