package main

func getSantasFloor(moves string) int {
	var floor int
	
	for i := 0; i < len(moves); i++ {
        if(moves[i] == '(') {
			floor++;
		}	else {
			floor--;
		}		
    }
	
	return floor
}