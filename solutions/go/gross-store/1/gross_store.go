package gross

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
    units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    fmt.Println(bill)
    u, exists := units[unit]
    if(!exists) {
        return false
    }

    qty, exists := bill[item]
    if (exists) {
        bill[item] = qty + u
    } else {
        bill[item] = u
    }
    
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	currentQty, itemExists := bill[item]
    unitQty, unitExists := units[unit]
    if(!itemExists || !unitExists || currentQty < unitQty) {
        return false
    }
    newQty := currentQty - unitQty
    if (0 == newQty) {
        delete(bill, item)
    } else {
        bill[item] = newQty
    }
    return true
    
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exists := bill[item]
    if (exists) {
        return qty, true
    } else {
        return 0, false
    }
}
