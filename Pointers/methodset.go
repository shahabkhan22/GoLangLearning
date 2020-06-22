package main

/*
A type may have a method set associated with it. The method set of an iterface type is ist interface.
The Method Set of any other type T consists of all methods declared with receiver type T.
The method set of correspoding pointer type is *T is the set of all methdds declared with receiver *T or T i.e. it also contains
the method set of T

In a method set, each method must have a unique non-blank method name.

In a nustshell,
METHOD SETS determine WHAT METHODS attach to a TYPE.

Receivers		|       Values
---------------------------------
(t T)			|	T and *T
(t *T)			|	*T


Basically 3 scenarios come into shape -
1. Value to either a non-pointer receiver OR pointer receiver is TRUE
2. Pointer Value to pointer receiver is TRUE
3. Non-pointer Value to Pointer Receiver is FALSE
*/
