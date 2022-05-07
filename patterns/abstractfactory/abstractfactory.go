/* an abstract factory is a creational design pattern 
*  that allows you to produce families of ralted objects 
*  without specifying their concrete class.
*  Defines an interface but then a concretefactory will 
*  take care of creating the actual product
*/

package main

import "fmt"

type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}