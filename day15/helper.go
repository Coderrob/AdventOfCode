package main 

import (
    "os"
    "bufio"
    "strconv"
    "strings"
)

type Receipe struct {    
    mixture map[string]int    
    ingredients []Ingredient
    totalCapacity int
    totalDurability int
    totalFlavor int
    totalTexture int
    totalCalories int
}

type Ingredient struct {
    name string
    capacity int
    durability int
    flavor int
    texture int
    calories int
}

func loadIngredientsFromFile(fileName string) []Ingredient {
    file, _ := os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    ingredients := []Ingredient {}
    for scanner.Scan() {
        ingredients = append(ingredients, getIngrediantFromString(scanner.Text()))
    }
    return ingredients
}

func getIngrediantFromString(input string) Ingredient {
    ingredient := Ingredient {}
    
    partialInput := strings.Split(input, ": ")
    ingredient.name = partialInput[0]
    
    partialInput = strings.Split(partialInput[1], ", ")
    
    capacity, _ := strconv.Atoi(strings.Split(partialInput[0], " ")[1])
    ingredient.capacity = capacity
    
    durability, _ := strconv.Atoi(strings.Split(partialInput[1], " ")[1])
    ingredient.durability = durability
    
    flavor, _ := strconv.Atoi(strings.Split(partialInput[2], " ")[1])
    ingredient.flavor = flavor
    
    texture, _ := strconv.Atoi(strings.Split(partialInput[3], " ")[1])
    ingredient.texture = texture
    
    calories, _ := strconv.Atoi(strings.Split(partialInput[4], " ")[1])
    ingredient.calories = calories
    
    return ingredient
}

func (receipe *Receipe) getByName(ingredientName string) Ingredient {
    var ingredient Ingredient
    for index := 0; index < len(receipe.ingredients); index++ {
        if receipe.ingredients[index].name == ingredientName {
            return receipe.ingredients[index]
        }
    }
    return ingredient
}

func (receipe *Receipe) count() int {
    ingredientCount := 0
    for ingredient := range receipe.mixture {
        ingredientCount += receipe.mixture[ingredient]
    }
    return ingredientCount
}

func (receipe *Receipe) add(ingredient Ingredient) {
    receipe.mixture[ingredient.name]++
}

func (receipe *Receipe) remove(ingredient Ingredient) {
    if receipe.mixture[ingredient.name] == 0 {
        return
    }
    receipe.mixture[ingredient.name]--
}

func (receipe *Receipe) removeAll(ingredient Ingredient) {
    if receipe.mixture[ingredient.name] == 0 {
        return
    }
    receipe.mixture[ingredient.name] = 0
}

func (receipe *Receipe) getScore() int {
    receipe.totalDurability = 0
    receipe.totalCalories = 0
    receipe.totalCapacity = 0
    receipe.totalTexture = 0
    receipe.totalFlavor = 0
    
    for ingredientName := range receipe.mixture {
        for index := 0; index < len(receipe.ingredients); index++ {
            ingredientCount := receipe.mixture[ingredientName]            
            if receipe.ingredients[index].name == ingredientName {                            
                receipe.totalDurability += (ingredientCount * receipe.ingredients[index].durability)
                receipe.totalCalories += (ingredientCount * receipe.ingredients[index].calories)               
                receipe.totalCapacity += (ingredientCount * receipe.ingredients[index].capacity)
                receipe.totalTexture += (ingredientCount * receipe.ingredients[index].texture)
                receipe.totalFlavor += (ingredientCount * receipe.ingredients[index].flavor)
            }
        }
    }

    if receipe.totalCapacity <= 0 || receipe.totalDurability <= 0 || receipe.totalFlavor <= 0 || receipe.totalTexture <= 0 {
        return 0
    }

    return receipe.totalCapacity * receipe.totalDurability * receipe.totalFlavor * receipe.totalTexture
}

func (receipe *Receipe) addIngredient() bool {
    ingredientLength := len(receipe.ingredients)
    var index int
    for index = ingredientLength - 1; index >= 0; index-- {
        if receipe.count() < teaspoonsRequired {
            receipe.add(receipe.ingredients[index])
            break
        } else {
            receipe.removeAll(receipe.ingredients[index])
        }
    }
    return index != -1
}

func (receipe *Receipe) copyReceipe() Receipe {
    newReceipe := Receipe {}
    newReceipe.mixture = map[string]int {}
    newReceipe.ingredients = receipe.ingredients[:]
    
    for key, value := range receipe.mixture {
        newReceipe.mixture[key] = value
    }
    return newReceipe
}