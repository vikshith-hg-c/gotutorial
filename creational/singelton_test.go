package creational

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	Instance1 := GetInstance()
	if Instance1 == nil {
		t.Error("excpted one instance to Singleton After calling GetInstance(), but nil")
	}

	Instance2 := GetInstance()
	if Instance2 != Instance1 {
		t.Error("excpted same instance to Singleton After calling GetInstance(), not nil")
	}

	currentCount := Instance1.AddOne()
	if currentCount != 1 {
		t.Errorf("firtime call count =1 but %d \n", currentCount)
	}

	currentCount = Instance2.AddOne()
	if currentCount != 2 {
		t.Errorf("firtime call count =2 but %d\n", currentCount)
	}

}
