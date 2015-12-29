package main 

import (
    "os"
    "bufio"
    "strconv"
    "strings"
)

type Reindeer struct {
    name string
    speed int
    duration int
    timeout int
}

type ReindeerRaceRecord struct {
    reindeer *Reindeer
    distanceTravelled int
    elapsedSeconds int
    duration int
    timeout int
    points int
}

func (raceRecord *ReindeerRaceRecord) tick() {
    reindeer := raceRecord.reindeer    
    raceRecord.elapsedSeconds++
    
    if raceRecord.duration == 0 && raceRecord.timeout == 0 {
        raceRecord.duration = reindeer.duration
        raceRecord.timeout = reindeer.timeout    
    }
    
    if raceRecord.duration > 0 {
        raceRecord.distanceTravelled += reindeer.speed
        raceRecord.duration--
        return
    }
    
    if raceRecord.timeout > 0 {
        raceRecord.timeout--
        return
    }
}

func loadReindeerList(fileName string) []Reindeer {
    file, _ := os.Open(fileName)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    reindeer := []Reindeer {}
    for scanner.Scan() {
        reindeer = append(reindeer, getReindeerFromString(scanner.Text()))        
    }
    return reindeer
}

func getReindeerFromString(input string) Reindeer {
    reindeer := Reindeer {}
    
    partialInput := strings.Split(input, " can fly ")
    reindeer.name = partialInput[0]
    
    partialInput = strings.Split(partialInput[1], " km/s for ")
    speed, _ := strconv.Atoi(partialInput[0])
    reindeer.speed = speed
    
    partialInput = strings.Split(partialInput[1], " seconds, but then must rest for ")
    duration, _ := strconv.Atoi(partialInput[0])
    reindeer.duration = duration
    
    restingTime := strings.Replace(partialInput[1], " seconds.", "", -1)
    timeout, _ := strconv.Atoi(restingTime)
    reindeer.timeout = timeout
    
    return reindeer
}

func startNewReindeerRace(reindeer []Reindeer) []ReindeerRaceRecord {
    reindeerRaceRecords := []ReindeerRaceRecord {}
    for index := 0; index < len(reindeer); index++ {
        reindeerRaceRecord := ReindeerRaceRecord {}
        reindeerRaceRecord.reindeer = &reindeer[index]
        reindeerRaceRecords = append(reindeerRaceRecords, reindeerRaceRecord)
    }
    return reindeerRaceRecords
}

func getFurthestDistanceTravelled(reindeerRaceRecords []ReindeerRaceRecord) int {
    furthestDistance := 0
    for index := 0; index < len(reindeerRaceRecords); index++ {
        if reindeerRaceRecords[index].distanceTravelled > furthestDistance {
            furthestDistance = reindeerRaceRecords[index].distanceTravelled
        }
    }
    return furthestDistance
}

func getHighestPointsEarned(reindeerRaceRecords []ReindeerRaceRecord) int {
    hightestPoints := 0
    for index := 0; index < len(reindeerRaceRecords); index++ {
        if reindeerRaceRecords[index].points > hightestPoints {
            hightestPoints = reindeerRaceRecords[index].points
        }
    }
    return hightestPoints
}