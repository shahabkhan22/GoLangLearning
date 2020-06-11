/*
	Composition goes beyond the mechanics of type embedding.
	It’s a paradigm we can leverage to design better APIs and
	to build larger programs from smaller parts. It all starts
	from the declaration and implementation of types that have
	a single purpose. Programs that are architected with composition
	in mind have a better chance to grow and adapt to changing needs.
	They are also much easier to read and reason about.package Functions

	With composition, we can design for both flexibility and readability.
	Every identifier that is exported from a package makes up the package’s API.
	This includes all the constants, variables, types, methods and functions that
	are exported. Comments are a frequently-overlooked aspect of every package’s API,
	so be very clear and concise when communicating information to the user of the package.

	The idea behind this program is we have a contractor that we hired to renovate our house.
	In particular, there are some boards in the house that have rotted and need to be yanked out,
	as well as new boards that need to be nailed in. The contractor will be given a supply of nails,
	boards to work with and tools to perform the work.
*/

package main

import "fmt"

//Board represents a surface we can work on
type Board struct {
	NailsNeeded int
	NailsDriven int
}

//NailDriver represents behavior to drive nails into a board
type NailDriver interface {
	DriveNail(nailSupply *int, b *Board)
}

//NailPuller represents behavior to remove nails into a board
type NailPuller interface {
	PullNail(nailSupply *int, b *Board)
}

//NailDrivePuller represents behavior to drive/remove nails into a board
type NailDrivePuller interface { //composing two existing interfaces into a new one
	NailDriver
	NailPuller
}

//Mallet is a tool that pounds in nails
type Mallet struct{} //empty structure as no state needs to be maintained only behavior need to be implemented

//DriveNail pounds a nail into the specified board
func (Mallet) DriveNail(nailSupply *int, b *Board) {
	*nailSupply-- //Take a nail out of the supply

	b.NailsDriven++ //Pound a nail into the board

	fmt.Println("Mallet : pounded nail into the board")
}

//Crowbar is a tool that removes nails
type Crowbar struct{}

//Pullnail removes a nail out of the specified board
func (Crowbar) PullNail(nailSupply *int, b *Board) {
	b.NailDriven-- //Remove a nail out of the board

	*nailSupply++ //Put that nail back into the supply

	fmt.Println("Crowbar: yanked nail out of the board")
}

//Contractor carries out the task of securing boards
type Contractor struct{}

//Fasten will drive the nails into the board
//Using NailDriver interface and not NailDrivPuller as only the behavior that is needed by Fasten methodis the ability to drive nails
func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board) {
	for b.NailsDriven < b.NailsNeeded {
		d.DriveNail(nailSupply, b)
	}
}

//Unfasten will remove nails from a board
//Same only on behavior is needed
func (Contractor) Unfaster(p NailPuller, nailSupply *int, b *Board) {
	for b.NailsDriven > b.NailsNeeded {
		p.PullNail(nailSupply, b)
	}
}

//
