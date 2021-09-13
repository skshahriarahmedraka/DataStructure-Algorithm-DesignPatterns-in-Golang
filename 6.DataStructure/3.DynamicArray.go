/*
 * Name : sk shahriar ahmed raka
 * Email : skshahriarahmedraka@gmail.com
 * Telegram : https://www.t.me/shahriarraka
 * Github : https://github.com/skshahriarahmedraka
 */

// A dynamic array is quite similar to a regular array, but its Size is modifiable during program runtime,
// very similar to how a slice in Go works. The implementation is for educational purposes and explains
// how one might go about implementing their own version of slices.
package main


import (
"errors"
	"fmt"
)

var defaultCapacity = 10

type DynamicArray struct {
	Size        int
	Capacity    int
	ElementData []interface{}
}

// Put function is change/update the value in array with the index and new value
func (da *DynamicArray) Put(index int, element interface{}) error {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return err
	}

	da.ElementData[index] = element

	return nil
}

// Add function is add new element to our array
func (da *DynamicArray) Add(element interface{}) {
	if da.Size == da.Capacity {
		da.NewCapacity()
	}

	da.ElementData[da.Size] = element
	da.Size++
}

// Remove function is remove an element with the index
func (da *DynamicArray) Remove(index int) error {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return err
	}

	copy(da.ElementData[index:], da.ElementData[index+1:da.Size])
	da.ElementData[da.Size-1] = nil

	da.Size--

	return nil
}

// Get function is return one element with the index of array
func (da *DynamicArray) Get(index int) (interface{}, error) {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return nil, err
	}

	return da.ElementData[index], nil
}

// IsEmpty function is check that the array has value or not
func (da *DynamicArray) IsEmpty() bool {
	return da.Size == 0
}

// GetData function return all value of array
func (da *DynamicArray) GetData() []interface{} {
	return da.ElementData[:da.Size]
}

// CheckRangeFromIndex function it will check the range from the index
func (da *DynamicArray) CheckRangeFromIndex(index int) error {
	if index >= da.Size || index < 0 {
		return errors.New("index out of range")
	}
	return nil
}

// NewCapacity function increase the Capacity
func (da *DynamicArray) NewCapacity() {
	if da.Capacity == 0 {
		da.Capacity = defaultCapacity
	} else {
		da.Capacity = da.Capacity << 1
	}

	newDataElement := make([]interface{}, da.Capacity)

	copy(newDataElement, da.ElementData)

	da.ElementData = newDataElement
}


func main (){
	d := &DynamicArray{}
	d.Add(3)
	d.Add(9)
	d.Add(5)

	li:=d.GetData()
	fmt.Println(li)

}