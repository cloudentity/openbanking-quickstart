package main 

type ConsentLister interface {
	List() []ClientConsents
}