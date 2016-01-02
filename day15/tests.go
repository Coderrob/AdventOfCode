package main

import "fmt"

func runCookieTests() {
    ingredients := loadIngredientsFromFile("examples.txt")
    receipt := bakeTheTestCookie(ingredients)
    fmt.Println(receipt.getScore() == 62842880)
    
    receipt = bakeAnotherTestCookie(ingredients)
    fmt.Println(receipt.getScore() == 62842880)
    
    receipt = bakeTheBestCalorieCookie(ingredients, 500)
    fmt.Println(receipt.getScore() == 57600000)
}

func bakeTheTestCookie(ingredients []Ingredient) Receipe {
    receipe := Receipe {}
    receipe.mixture = map[string]int {}
    receipe.ingredients = ingredients   
    
    butterscotch := receipe.getByName("Butterscotch")
    for index := 0; index < 44; index++ {
        receipe.add(butterscotch)
    }
    
    cinnamon := receipe.getByName("Cinnamon")
    for index := 0; index < 56; index++ {
        receipe.add(cinnamon)
    }
        
    return receipe
}

func bakeAnotherTestCookie(ingredients []Ingredient) Receipe {
    receipe := Receipe {}
    receipe.mixture = map[string]int {}
    receipe.ingredients = ingredients       
    receipe.mixture["Butterscotch"] = 44
    receipe.mixture["Cinnamon"] = 56
    return receipe
}