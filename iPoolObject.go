package main

type iPoolObject interface{
	getID() string// any id that help diferentiate objects from each other in the object pool
}

