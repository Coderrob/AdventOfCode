package main

/*
--- Day 13: Knights of the Dinner Table ---

In years past, the holiday feast with your family hasn't gone so well. Not everyone gets along! This year, you resolve, will be different. You're going to find the optimal seating arrangement and avoid all those awkward conversations.

You start by writing up a list of everyone invited and the amount their happiness would increase or decrease if they were to find themselves sitting next to each other person. You have a circular table that will be just big enough to fit everyone comfortably, and so each person will have exactly two neighbors.

For example, suppose you have only four attendees planned, and you calculate their potential happiness as follows:

Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.
Then, if you seat Alice next to David, Alice would lose 2 happiness units (because David talks so much), but David would gain 46 happiness units (because Alice is such a good listener), for a total change of 44.

If you continue around the table, you could then seat Bob next to Alice (Bob gains 83, Alice gains 54). Finally, seat Carol, who sits next to Bob (Carol gains 60, Bob loses 7) and David (Carol gains 55, David gains 41). The arrangement looks like this:

     +41 +46
+55   David    -2
Carol       Alice
+60    Bob    +54
     -7  +83
After trying every other seating arrangement in this hypothetical scenario, you find that this one is the most optimal, with a total change in happiness of 330.

What is the total change in happiness for the optimal seating arrangement of the actual guest list?
*/

func getHappiestArrangement(fileName string) int {
    preferences := loadAttendeeList(fileName)
    totalHappiness := 0
    attendeeNames := getUniqueAttendeeNames(preferences)
    seatPermutations, _ := NewPerm(attendeeNames, nil)
    
    for seatPermutation, err := seatPermutations.Next(); err == nil; seatPermutation, err = seatPermutations.Next() {
        happiness := getSeatArrangementHappinessAmount(seatPermutation.([]string), preferences)
		if totalHappiness == 0 || totalHappiness < happiness {
			totalHappiness = happiness
		}
    }
    return totalHappiness
}

func getSeatArrangementHappinessAmount(seating []string, preferences []SeatPreference) int {
	happiness := 0
    seatingLength := len(seating)
	for index := 1; index < seatingLength; index++ {		
		currentSeat := seating[index - 1]
		nextSeat := seating[index]
        happiness += getSeatingHappinessValue(currentSeat, nextSeat, preferences)
	}
    happiness += getSeatingHappinessValue(seating[seatingLength - 1], seating[0], preferences)
	return happiness
}

func getSeatingHappinessValue(currentSeat string, nextSeat string, preferences []SeatPreference) int {
    happiness := 0
	for _, preference := range preferences {
        if (preference.name == currentSeat && preference.nextToAttendee == nextSeat) || 
            (preference.nextToAttendee == currentSeat && preference.name == nextSeat) {
            happiness += preference.happiness
        }
	}
    return happiness
}