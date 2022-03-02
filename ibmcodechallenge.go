package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	salesTax      = 0.10
	importDutyTax = 0.05
)

var (
	productTable *ProductTable
	basket       *Basket
)

func main() {

	//////////////////////////////////////////////////////////////////////////////////////////////
	//
	// fallback function for handling unhandled exceptions
	//

	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()

	//////////////////////////////////////////////////////////////////////////////////////////////
	//
	// usage directions
	//

	fmt.Println("")
	fmt.Println("IBM Code Challenge for dedellis@us.ibm.com.........")
	fmt.Println("")
	fmt.Println("PROBLEM TWO: SALES TAXES")
	fmt.Println("")

	usageStr := "	usage:   numitems [imported] product name at price \n"
	usageStr += "	example: 1 imported box of chocolates at 10.00 \n"
	usageStr += "	( press enter twice after the last item to print receipt ) \n"
	fmt.Println(usageStr)
	fmt.Println("")

	//////////////////////////////////////////////////////////////////////////////////////////////

	//simulate a Product Table query to the database
	productTable = GetProductTable()

	//cache of the basket containing lineitems as users input from the console
	basket = newBasket()

	//////////////////////////////////////////////////////////////////////////////////////////////
	//
	// loop around console input to obtain the line items to parse into LineItem struct
	// then add to the basket struct
	// handle any errors that bubble up from functions
	//

	anotherInput := true

	for anotherInput {

		bio := bufio.NewReader(os.Stdin)
		input, _ := bio.ReadString('\n')

		input = strings.TrimSpace(strings.Trim(input, "\n"))

		switch input {

		case "":
			anotherInput = false
			PrintRecieptFromBasket()

		default:
			if parsedLineItem, err := ParseLineItemInput(input); err == nil {
				AddLineItemToBasket(parsedLineItem)
			} else {
				fmt.Println(err.Error())
				fmt.Println("...the program will now exit")
				os.Exit(1)
			}

		}

	}

}

// ParseLineItemInput is a function that
// 1.  parses the input string and validates inputs then creates a LineItem from the inputs to return to the main function
// 2.  compares to a database product catalog for the purposes of flagging which products should apply sales and import duty tax
func ParseLineItemInput(input string) (parsedLineItem LineItem, returnerr error) {

	parsedLineItem = LineItem{}

	// handle the price after the " at "  seperately from all things before the " at "
	lineItemPriceSplit := strings.Split(input, " at ")

	// if the format of the line does not return two elements , 1 to the left of the " at " and 1 to the right then it is bad input and we throw an error
	if len(lineItemPriceSplit) != 2 {

		errorStr := "...line item input format exception: \n"
		errorStr += "...please input the LINE ITEM in the correct format detailed below \n"
		errorStr += "...	usage:   numitems [imported] product name at price \n"
		errorStr += "...	example: 1 imported box of chocolates at 10.00 \n"

		returnerr = errors.New(errorStr)
		return
	}

	// parse out the price from the right side of the " at " in the input string
	// add this to the parsedLineItem which will be returned
	lineItemPriceStr := lineItemPriceSplit[1]
	lineItemPriceStr = strings.TrimSpace(strings.Trim(lineItemPriceStr, "\n"))

	// if the conversion of the string price is not a float it is bad formating and we throw an exception
	if lineItemPrice, err := strconv.ParseFloat(lineItemPriceStr, 64); err == nil {
		//add the price to the parsedLineItem
		parsedLineItem.price = lineItemPrice
	} else {
		errorStr := "...price parsing exception: \n"
		errorStr += "...please input the PRICE as a floating point number in the correct format detailed below \n"
		errorStr += "...	usage:   numitems [imported] product name at price \n"
		errorStr += "...	example: 1 imported box of chocolates at 10.00 \n"

		returnerr = errors.New(errorStr)
		return
	}

	// the left side of the input string before the " at "
	amountImportedProductStr := lineItemPriceSplit[0]

	if amountImportedProductStr == "" {

		errorStr := "...line item input format exception: \n"
		errorStr += "...please input the NUMBER of ITEMS, PRODUCT NAME and IMPORTED FLAG in the correct format detailed below \n"
		errorStr += "...	usage:   numitems [imported] product name at price \n"
		errorStr += "...	example: 1 imported box of chocolates at 10.00 \n"

		returnerr = errors.New(errorStr)
		return
	}

	// extract the "imported" flag first
	lineItemImported := strings.Contains(amountImportedProductStr, "imported")

	if lineItemImported {
		// flag the lineitem for applying the duty tax later
		// Import duty is an additional sales tax applicable on all imported goods at a rate of 5%, with no exemptions.
		parsedLineItem.applyImportDutyTax = true
	}

	//remove the "imported" flag if it exists so we can pricess the amount and product
	amountImportedProductStr = strings.Replace(amountImportedProductStr, "imported", "", 1)
	amountProductSplit := strings.SplitN(amountImportedProductStr, " ", 2)

	if len(amountProductSplit) != 2 {

		errorStr := "...line item input format exception: \n"
		errorStr += "...please input the NUMBER of ITEMS and PRODUCT NAME in the correct format detailed below \n"
		errorStr += "...	usage:   numitems [imported] product name at price \n"
		errorStr += "...	example: 1 imported box of chocolates at 10.00 \n"

		returnerr = errors.New(errorStr)
		return
	}

	amountStr := amountProductSplit[0]
	amountStr = strings.TrimSpace(amountStr)

	if amount, err := strconv.ParseInt(amountStr, 10, 64); err == nil {
		//add the amount to the lineitem struct
		parsedLineItem.amount = amount
	} else {
		errorStr := "...amount parsing exception: \n"
		errorStr += "...please input the AMOUNT as an integer in the correct format detailed below \n"
		errorStr += "...	usage:   numitems [imported] product name at price \n"
		errorStr += "...	example: 1 imported box of chocolates at 10.00 \n"

		returnerr = errors.New(errorStr)
		return
	}

	productStr := strings.TrimSpace(amountProductSplit[1])
	parsedLineItem.product = productStr

	//compare with the product table query as to satisfy the rule basic sales tax is applicable at a rate of 10% on all goods, except books, food, and medical products that are exempt.
	if productTable.productRows[productStr] == other {
		parsedLineItem.applySalesTax = true
	}

	return

}

// PrintRecieptFromBasket ...
// business rule: When I purchase items I receive a receipt which lists the name of all the
// items and their price (including tax), finishing with the total cost of the
// items, and the total amounts of sales taxes paid.  The rounding rules for
// sales tax are that for a tax rate of n%, a shelf price of p contains
// (np/100 rounded up to the nearest 0.05) amount of sales tax.
func PrintRecieptFromBasket() {

	totalFinalTx := 0.0
	totalFinalPrice := 0.0

	//fmt.Println(basket.lineItems)

	for _, li := range basket.lineItems {
		liImportedStr := ""
		liImportDutyTax := 0.0
		liSalesTax := 0.0
		liFinalTax := 0.0
		liFinalPrice := 0.0

		//calculate sales and duty tax if necessary
		if li.applyImportDutyTax {
			liImportedStr = "imported"
			liImportDutyTax = li.price * importDutyTax
		}

		if li.applySalesTax {
			liSalesTax = li.price * salesTax
		}

		liFinalTax = liImportDutyTax + liSalesTax

		liFinalTaxRounded := RoundToUnit(liFinalTax, 5)

		liFinalPrice = li.price + liFinalTaxRounded

		totalFinalTx += liFinalTaxRounded
		totalFinalPrice += liFinalPrice

		// uncomment for debugging purposes
		// fmt.Println("li price: ", li.price)
		// fmt.Println("liImportDutyTax: ", liImportDutyTax)
		// fmt.Println("liSalesTax: ", liSalesTax)
		// fmt.Println("liFinalTax: ", liFinalTax)
		// fmt.Println("liFinalTaxRounded: ", liFinalTaxRounded)
		// fmt.Println("liFinalPrice: ", liFinalPrice)

		// fmt.Println(RoundToUnit(0.5625, 5))
		// fmt.Println(RoundToUnit(11.8125, 5))
		// fmt.Println(RoundToUnit(7.125, 5))

		// output the lineitem formatted
		if li.applyImportDutyTax {
			fmt.Printf("%.d %s %s: %.2f\n", li.amount, liImportedStr, li.product, liFinalPrice)
		} else {
			fmt.Printf("%.d %s: %.2f\n", li.amount, li.product, liFinalPrice)
		}

	}

	fmt.Printf("%s: %.2f\n", "Sales Taxes", totalFinalTx)
	fmt.Printf("%s: %.2f\n", "Total", totalFinalPrice)

	return
}

// RoundToUnit ...
// The rounding rules for
// sales tax are that for a tax rate of n%, a shelf price of p contains
// (np/100 rounded up to the nearest 0.05) amount of sales tax.
func RoundToUnit(input float64, unit int64) (backToOriginalFloat float64) {

	if input > 0.0 {
		var inputexp = int64(input * 100.0)
		var roundup = inputexp % unit
		if roundup > 0 {
			roundedNum := inputexp + (unit - roundup)
			backToOriginalFloat = float64(roundedNum) / 100.0
		} else {
			backToOriginalFloat = input
		}

	} else {
		backToOriginalFloat = input
	}

	return
}

// ProductType ...
type ProductType string

// ProductType ...
const (
	book    ProductType = "book"
	food    ProductType = "food"
	medical ProductType = "medical"
	other   ProductType = "other"
)

// Product ...
type Product struct {
	productName string
	productType ProductType
}

// LineItem ...
type LineItem struct {
	amount             int64
	product            string
	price              float64
	applySalesTax      bool
	applyImportDutyTax bool
}

// AddLineItemToBasket ...
func AddLineItemToBasket(lineItem LineItem) (currentCount int) {
	basket.lineItems = append(basket.lineItems, lineItem)
	currentCount = len(basket.lineItems)
	return
}

func newBasket() (theBasket *Basket) {
	theBasket = &Basket{}
	theBasket.lineItems = []LineItem{}
	return
}

// Basket ...
type Basket struct {
	lineItems []LineItem
}

// GetProductTable ...
func GetProductTable() (theProductTable *ProductTable) {
	theProductTable = newProductTable()
	theProductTable.productRows["book"] = book
	theProductTable.productRows["music CD"] = other
	theProductTable.productRows["chocolate bar"] = food
	theProductTable.productRows["box of chocolates"] = food
	theProductTable.productRows["bottle of perfume"] = other
	theProductTable.productRows["packet of headache pills"] = medical
	return
}

func newProductTable() (theProductTable *ProductTable) {
	theProductTable = &ProductTable{}
	theProductTable.productRows = map[string]ProductType{}
	return
}

// ProductTable ...
type ProductTable struct {
	productRows map[string]ProductType
}
