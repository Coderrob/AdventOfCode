package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
)

type Attendee struct {
    attendee string
    leftSeatHappiness int
    rightSeatHappiness int
}

type SeatPreference struct {
    name string
    happiness int
    nextToAttendee string
}

func loadAttendeeList(fileName string) []SeatPreference {
    attendeeList := []SeatPreference {}
    file, _ := os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        seatPreference := getSeatPreferenceFromString(scanner.Text())
        attendeeList = append(attendeeList, seatPreference)
    }
    return attendeeList
}

func getSeatPreferenceFromString(input string) SeatPreference {
    preference := SeatPreference {}    
    partialInput := strings.Split(input, " happiness units by sitting next to ")

    preference.nextToAttendee = strings.Replace(partialInput[1], ".", "", -1)
    
    if strings.Contains(partialInput[0], "lose") {
        partialInput = strings.Split(partialInput[0], " would lose ")
        preference.name = partialInput[0]
        value, _ := strconv.Atoi(partialInput[1])
        preference.happiness = value * -1
    } else {
        partialInput = strings.Split(partialInput[0], " would gain ")
        preference.name = partialInput[0]
        value, _ := strconv.Atoi(partialInput[1])
        preference.happiness = value
    }
    
    return preference
}

func getUniqueAttendeeNames(preferences []SeatPreference) []string {
	var attendeeNames = map[string]bool {}
	for _, preference := range preferences {
		if attendeeNames[preference.name] == false {
			attendeeNames[preference.name] = true
		}
	}
	return getKeysFromStringMap(attendeeNames)
}

func getKeysFromStringMap(input map[string]bool) []string {
	keys := make([]string, 0, len(input))
    for k := range input {
        keys = append(keys, k)
    }
	return keys
}