package test

import (
	"fmt"
	"shop-here-api/src/app/dao"
	"shop-here-api/src/app/modal"
	"testing"
)

var cartTestItem modal.Cart = modal.Cart{
	CartID:    9,
	UserID:    9,
	ProductID: 0,
	Quantity:  0,
}

func TestUpdateCartItem(t *testing.T) {
	dao.StartDb()
	CaseNumber = 1
	fmt.Printf("Running Test Case:  %v \n", CaseNumber)
	CaseNumber++
	// Item name is valid so it will insert the record and give valid PK value
	object := dao.CartRepo{}
	itemReturn, _ := object.UpdateCartItem(cartTestItem)
	if itemReturn.CartID == 0 {
		t.Errorf("TestFailed, got: %d, want: %v .\n", itemReturn.ID, "value more than zero")
	} else {
		fmt.Printf("TestPassed, got: %d, want: %v.\n", itemReturn.ID, "value more than zero")
	}
	fmt.Printf("Running Test Case:  %v \n", CaseNumber)
	CaseNumber++
	// Item name is invalid so it will not insert the record and give zero PK value
	cartTestItem.ProductID = 0
	itemReturn, _ = object.UpdateCartItem(cartTestItem)
	if itemReturn.ProductID > 0 {
		t.Errorf("TestFailed, got: %d, want: %v.\n", itemReturn.ID, 0)
	} else {
		fmt.Printf("TestPassed, got: %d, want: %v.\n", itemReturn.ID, 0)
	}
}

func TestGetCartInfo(t *testing.T) {
	fmt.Printf("Running Test Case:  %v \n", CaseNumber)
	CaseNumber++
	object := dao.CartRepo{}
	itemObj, err := object.GetAllItemInfo(cartTestItem) //first create the item
	// able to get the item data
	if err == nil {
		fmt.Printf("TestPasses, got: %v, want: %v .\n", len(itemObj), "no of items more than 0")
	} else {
		t.Errorf("TestFailed, got: %v, want: %v .\n", len(itemObj), "no of items more than 0")
	}
}

func TestGetCartItems(t *testing.T) {
	fmt.Printf("Running Test Case:  %v \n", CaseNumber)
	CaseNumber++
	object := dao.CartRepo{}
	itemObj, err := object.GetAllItem(cartTestItem) //first create the item
	// able to get the item data
	if err == nil {
		fmt.Printf("TestPasses, got: %v, want: %v .\n", len(itemObj), "no of items more than 0")

		if len(itemObj) > 0 {
			var slice = make([]uint, len(itemObj))
			for i := 0; i < len(itemObj); i++ {
				slice[i] = itemObj[i].ProductID
			}
			itemRepo := dao.ItemRepo{}
			cartItemsReturn, err := itemRepo.GetAllCartItem(slice)
			if err != nil {
				t.Errorf("TestFailed, got: %v, want: %v .\n", len(cartItemsReturn), "no of items more than 0")
			} else {
				fmt.Printf("TestPasses, got: %v, want: %v .\n", len(itemObj), "no of items more than 0")

			}
		}

	} else {
		t.Errorf("TestFailed, got: %v, want: %v .\n", len(itemObj), "no of items more than 0")
	}

}
