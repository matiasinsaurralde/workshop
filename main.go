package main

import "errors"

var (
	errNotEnoughBalance      = errors.New("not enough balance")
	errEquipmentNotAvailable = errors.New("equipment not available")
	errNotExists             = errors.New("not exists")
	errAlreadyExists         = errors.New("already exists")
)

// Member is the member's data structure
type Member struct {
	Name    string
	Balance int
}

// Rent is called by member when renting equipment
func (m *Member) Rent(inventory *Inventory, equipmentsToRent []string, startDate string, endDate string, numberOfDays int) error {
	// Iterate through all equipments in the given inventory
	// if (i) equipment with the given name exists and
	// (ii) equipment is marked as available, then
	// add it to availableEquipment slice which will be
	// used for the rest of the flow:
	availableEquipment := make([]*Equipment, 0)
	for _, equipment := range inventory.Equipments {
		for _, equipmentToRent := range equipmentsToRent {
			if equipment.Name == equipmentToRent && equipment.Availability {
				availableEquipment = append(availableEquipment, equipment)
			}
		}
	}

	// Not all equipment is available if this condition fails:
	if len(availableEquipment) != len(equipmentsToRent) {
		return errEquipmentNotAvailable
	}

	// Now calculate total rental cost:
	var totalRentalCost int
	for _, equipment := range availableEquipment {
		totalRentalCost += equipment.RentalCost * numberOfDays
	}
	if totalRentalCost > m.Balance {
		return errNotEnoughBalance
	}
	m.Balance -= totalRentalCost
	return nil
}

// Equipment holds equipment information
type Equipment struct {
	Name         string
	RentalCost   int
	Availability bool
}

// Inventory holds an array of equipments
type Inventory struct {
	Equipments []*Equipment
}

// AddEquipment adds a new equipment to the inventory:
func (i *Inventory) AddEquipment(e *Equipment) error {
	// First check if item already exists or not:
	for _, existingEquipment := range i.Equipments {
		if e.Name == existingEquipment.Name {
			return errAlreadyExists
		}
	}
	// Set availability to true for newly added items:
	e.Availability = true
	i.Equipments = append(i.Equipments, e)
	return nil
}

// RemoveEquipment removes an equipment from the inventory:
func (i *Inventory) RemoveEquipment(name string) error {
	newEquipments := make([]*Equipment, 0)
	var match bool
	for _, e := range i.Equipments {
		if e.Name == name {
			match = true

			continue
		}
		newEquipments = append(newEquipments, e)
	}
	i.Equipments = newEquipments
	if !match {
		return errNotExists
	}
	return nil
}

// Size is an alias for len(i.Equipments)
func (i *Inventory) Size() int {
	return len(i.Equipments)
}

// UpdateRentalCost is used to update the rental cost of a given equipment:
func (i *Inventory) UpdateRentalCost(name string, newCost int) error {
	var match bool
	for _, e := range i.Equipments {
		if e.Name != name {
			continue
		}
		match = true
		e.RentalCost = newCost
	}
	if !match {
		return errNotExists
	}
	return nil
}

func main() {}
