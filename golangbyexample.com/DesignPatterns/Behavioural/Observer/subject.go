package main

type subject interface {
	register(obs observer)
	deregister(obs observer)
	notifyAll()
}
