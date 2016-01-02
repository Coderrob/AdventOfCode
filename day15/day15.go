package main

import "fmt"

func main() {
    ingredients := loadIngredientsFromFile("ingredients.txt")
    receipt := bakeTheBestCookie(ingredients)    
    fmt.Println(receipt.getScore())
    
    receipt = bakeTheBestCalorieCookie(ingredients, 500)
    fmt.Println(receipt.getScore())
}