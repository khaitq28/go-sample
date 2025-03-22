package calculator

// Calculator provides basic arithmetic operations
type Calculator struct {
	lastResult float64
}

// Add adds two numbers
func (c *Calculator) Add(a, b float64) float64 {
	c.lastResult = a + b
	return c.lastResult
}

// Subtract subtracts two numbers
func (c *Calculator) Subtract(a, b float64) float64 {
	c.lastResult = a - b
	return c.lastResult
}

// GetLastResult returns the last calculation result
func (c *Calculator) GetLastResult() float64 {
	return c.lastResult
}
