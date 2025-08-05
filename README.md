# URL Shortener (golang)
The second project in the [Gophercises](https://gophercises.com/) course for
practicing and mastering Go.

I need to be honest - this one gave me some headaches while I wrapped my head
around how Go handles web routing (differences between `http.Handle`, `http.Handler`, `http.HandleFunc`, `http.HandlerFunc`,
and what even is a `mux`), and worst of all: YAML parsing.

Even worse than YAML is the fact that the code looks dead simple to me now -
an interesting contrast to the feeling I had until my synapses connected in
such a way to allow me to write the damn thing. There were several _Eureka!_
moments indeed.

## Bonus
- [x] Update the main/main.go source file to accept a YAML file as a flag and then
load the YAML from a file rather than from a string.
- [x] Build a JSONHandler that serves the same purpose, but reads from JSON data.
- [ ] Build a Handler that doesn't read from a map but instead reads from a database.
Whether you use BoltDB, SQL, or something else is entirely up to you.