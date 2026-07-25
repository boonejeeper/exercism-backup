package elon

import "fmt"

func (c *Car) Drive() {
    if c.batteryDrain > c.battery {
        return
    }
    c.distance += c.speed
    c.battery -= c.batteryDrain
}

func (c *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c *Car) CanFinish(trackDistance int) bool {
    numberOfCycles := trackDistance / c.speed
    if trackDistance % c.speed != 0 {
        numberOfCycles++
    }
    expectedBatteryDrain := numberOfCycles * c.batteryDrain
    return c.battery >= expectedBatteryDrain
    
}
