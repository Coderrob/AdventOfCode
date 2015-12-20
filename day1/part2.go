package main

func getPositionWhenSantasGoesIntoTheBasement(moves string) int {
	var floor int
	var position int
	var basementEntered bool
	
	for index := 0; index < len(moves) && basementEntered != true; index++ {
        if(moves[index] == '(') {
			floor++;
		}	else {
			floor--;
		}
		
		if (floor == -1) {
			basementEntered = true
			position = index + 1
		}
    }
	
	return position
}