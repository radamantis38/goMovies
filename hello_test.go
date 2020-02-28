package main

import "testing"

func TestSum(t *testing.T) {
	main()
	_, err:= Sum(5, 5)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAmenorB(t *testing.T) {
	err:= AmenorB(4, 5)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = AmenorB(6, 5)
   if err == nil {
	   t.Errorf("Fallo")
   }

}