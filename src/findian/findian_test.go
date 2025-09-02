package main
/**
 * The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. 
 * The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”. 
 */

import (
	"testing"
)


func TestFindian(t *testing.T) {
	var found = [4]string{"ian", "Ian","iuiygaygn", "I d skd a efju N"}
	var missing = [3]string{"ihhhhhn", "ina", "xian"}

	for _, val := range found {
		result := Findian( val );
		if result == false {
			t.Errorf("Not found error")
		}
	}
	for _, val := range missing {
		result := Findian( val );
		if result == true {
			t.Errorf("Found error")
		}
	}
}

func TestStartsWithI(t *testing.T) {

	var found = [4]string{"ian", "Ian","iuiygaygn", "I d skd a efju N"}
	var missing = [3]string{"hhhhhn", "na", "xian"}

	for _, val := range found {
		result := StartsWithI( val );
		if result == false {
			t.Errorf("Not found error")
		}
	}
	for _, val := range missing {
		result := StartsWithI( val );
		if result == true {
			t.Errorf("Found error")
		}
	}

}

func TestEndsWithN(t *testing.T) {

	var found = [4]string{"ian", "Ian","iuiygaygn", "I d skd a efju N"}
	var missing = [3]string{"ihhhhh", "ina", "xia"}

	for _, val := range found {
		result := EndsWithN( val );
		if result == false {
			t.Errorf("Not found error")
		}
	}
	for _, val := range missing {
		result := EndsWithN( val );
		if result == true {
			t.Errorf("Found error")
		}
	}

}

func TestContainsA(t *testing.T) {

	var found = [4]string{"ian", "Ian","iuiygaygn", "I d skd a efju N"}
	var missing = [3]string{"ihhhhhn", "in", "xin"}

	for _, val := range found {
		result := ContainsA( val );
		if result == false {
			t.Errorf("Not found error")
		}
	}
	for _, val := range missing {
		result := ContainsA( val );
		if result == true {
			t.Errorf("Found error")
		}
	}
}

