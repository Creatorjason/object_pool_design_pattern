package main

import (
	"fmt"
	"sync"
)


type pool struct{
	idle []iPoolObject
	active []iPoolObject
	capacity int
	mulock *sync.Mutex

}

// InitPool initializes the pool

func InitPool(poolObjects []iPoolObject)(*pool, error){
	poolObjLen := len(poolObjects)
	if poolObjLen == 0{
		return nil, fmt.Errorf("cannot return a pool of 0 length ")
	}
	active := make([]iPoolObject, 0)
	pool := &pool{
		idle: poolObjects,
		active: active,
		capacity: poolObjLen,
		mulock: new(sync.Mutex),
	}
	return pool, nil
}

//  Removes from idle pool and adds to active pool

func (p *pool) loan() (iPoolObject, error){
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0{
		return nil, fmt.Errorf("no pool object free, please request after some time")
	}
	// Lets assume that the client is grabbing the first object from the idle object pool 
	obj := p.idle[0]
	// update the pool, "letting it know that an object has been gotten from it"
	p.idle = p.idle[1:]
	//  add the borrowed object into the active pool
	p.active = append(p.active, obj)
	fmt.Printf("Loan object with id: %s\n", obj.getID())
	return obj, nil
}

// create a method that receives loan pool object from active pool back to the idle pool
func(p *pool) receive(target iPoolObject) error{
	p.mulock.Lock()
	defer p.mulock.Unlock()
	err := p.remove(target)
	if err != nil{
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("Return pool object with id: %s\n", target.getID())
	return nil

}


//  removes a pool object from active pool
func (p *pool) remove(target iPoolObject) error{
	currentActiveLength := len(p.active)
	for i, obj := range p.active{
		if obj.getID() == target.getID(){
			// swap the last pool object with this current pool object
			p.active[currentActiveLength - 1], p.active[i] = p.active[i], p.active[currentActiveLength - 1]
			p.active = p.active[:currentActiveLength - 1]
			return nil
		}
	}
	return fmt.Errorf("target pool object doesn't belong to pool")
}