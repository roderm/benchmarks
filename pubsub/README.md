# Pubsub
In go we have channels for multithreaded communication between go-routines. In my first benchmark-test I'm wondering how fast they are compared to callbacks.

# plain
without any go-routines

# routines
start for every pub-message a new routine

# workers
add a bund of workers to not explode goroutines