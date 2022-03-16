package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquipmentRent(t *testing.T) {
	member := Member{
		Name:    "Wojciech",
		Balance: 50,
	}
	inventory := Inventory{
		Equipments: []*Equipment{
			{
				Name:         "W03 Pyranha Burn",
				RentalCost:   1,
				Availability: true,
			},
			{
				Name:         "P05 TNP Rapa",
				RentalCost:   1,
				Availability: true,
			},
		},
	}

	err := member.Rent(&inventory, []string{
		"W03 Pyranha Burn",
		"P05 TNP Rapa",
	}, "2022-03-15", "2022-03-16", 2)

	assert.NoError(t, err)

	assert.Equal(t, member.Balance, 46)
}

func TestEquipmentRentNotEnoughBalance(t *testing.T) {
	member := Member{
		Name:    "Wojciech",
		Balance: 3,
	}
	inventory := Inventory{
		Equipments: []*Equipment{
			{
				Name:         "W03 Pyranha Burn",
				RentalCost:   1,
				Availability: true,
			},
			{
				Name:         "P05 TNP Rapa",
				RentalCost:   1,
				Availability: true,
			},
		},
	}

	err := member.Rent(&inventory, []string{
		"W03 Pyranha Burn",
		"P05 TNP Rapa",
	}, "2022-03-15", "2022-03-16", 2)
	assert.EqualError(t, err, "not enough balance")

	assert.Equal(t, member.Balance, 3)
}

func TestEquipmentRentNotAvailable(t *testing.T) {
	member := Member{
		Name:    "Wojciech",
		Balance: 50,
	}
	inventory := Inventory{
		Equipments: []*Equipment{
			{
				Name:         "W03 Pyranha Burn",
				RentalCost:   1,
				Availability: false,
			},
			{
				Name:         "P05 TNP Rapa",
				RentalCost:   1,
				Availability: true,
			},
		},
	}

	err := member.Rent(&inventory, []string{
		"W03 Pyranha Burn",
		"P05 TNP Rapa",
	}, "2022-03-15", "2022-03-16", 2)
	assert.EqualError(t, err, "equipment not available")
}

func TestInventoryAddingNewItem(t *testing.T) {
	inventory := Inventory{
		Equipments: []*Equipment{
			{
				Name:         "W03 Pyranha Burn",
				RentalCost:   1,
				Availability: true,
			},
			{
				Name:         "P05 TNP Rapa",
				RentalCost:   1,
				Availability: true,
			},
		},
	}

	newKayak := Equipment{
		Name:       "W04 Wavesport Diesel",
		RentalCost: 2,
	}
	err := inventory.AddEquipment(&newKayak)
	assert.NoError(t, err)
	assert.Equal(t, inventory.Size(), 3)
	assert.Equal(t, inventory.Equipments[2], &newKayak)
	// Check if availability is true for newly added item:
	assert.True(t, inventory.Equipments[2].Availability)

	// Try to re-add the existing item:
	err = inventory.AddEquipment(&newKayak)
	assert.EqualError(t, err, "already exists")
}
func TestInventoryRemovingAnItem(t *testing.T) {
	inventory := Inventory{
		Equipments: []*Equipment{
			{
				Name:         "W03 Pyranha Burn",
				RentalCost:   1,
				Availability: true,
			},
			{
				Name:         "P05 TNP Rapa",
				RentalCost:   1,
				Availability: true,
			},
		},
	}

	err := inventory.RemoveEquipment("P05 TNP Rapa")
	assert.NoError(t, err)
	assert.Equal(t, inventory.Size(), 1)

	// Try to remove nonexistent item:
	err = inventory.RemoveEquipment("W02 Pyranha Burn")
	assert.EqualError(t, err, "not exists")
}

func TestInventorUpdatingAnItem(t *testing.T) {
	inventory := Inventory{
		Equipments: []*Equipment{
			{
				Name:         "W03 Pyranha Burn",
				RentalCost:   1,
				Availability: true,
			},
			{
				Name:         "P05 TNP Rapa",
				RentalCost:   1,
				Availability: true,
			},
		},
	}
	err := inventory.UpdateRentalCost("W03 Pyranha Burn", 3)
	assert.NoError(t, err)
	assert.Equal(t, inventory.Equipments[0].RentalCost, 3)

	// Try to update nonexistent item:
	err = inventory.UpdateRentalCost("W02 Pyranha Burn", 3)
	// assert.EqualError(t, err, "not exists")
	assert.NoError(t, err)
}
