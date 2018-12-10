package main

import (
	fileutil "code-advent/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Guard watches the room
type Guard struct {
	id              int
	minutesAsleep   int
	sleepiestMinute int
	sameMinuteCount int
	sleepSchedule   map[int]int
}

func main() {

	input, err := fileutil.GetStringSliceFromFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	start := time.Now()

	guards := parseGuards(input)
	sleepiestGuard := getSleepiestGuard(guards)
	consistantlySleepyGuard := getConsistantlySleepyGuard(guards)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

	fmt.Println(sleepiestGuard.id * sleepiestGuard.sleepiestMinute)
	fmt.Println(consistantlySleepyGuard.id * consistantlySleepyGuard.sleepiestMinute)
}

func getSleepiestGuard(guards map[int]Guard) Guard {
	sleepiestGuard := guards[0]

	for _, nextGuard := range guards {
		if nextGuard.minutesAsleep > sleepiestGuard.minutesAsleep {
			sleepiestGuard = nextGuard
		}
	}
	return sleepiestGuard
}

func getConsistantlySleepyGuard(guards map[int]Guard) Guard {
	consistantlySleepyGuard := guards[0]

	for _, nextGuard := range guards {
		if nextGuard.sameMinuteCount > consistantlySleepyGuard.sameMinuteCount {
			consistantlySleepyGuard = nextGuard
		}
	}
	return consistantlySleepyGuard
}

func parseGuards(input []string) map[int]Guard {
	guards := make(map[int]Guard)
	var currentGuardID int
	var startedSleepingAt int

	sort.Strings(input)
	for _, line := range input {
		time, _ := strconv.Atoi(line[15:17])
		action := line[19:]
		guard := guards[currentGuardID]

		if action == "falls asleep" {
			startedSleepingAt = time

		} else if action == "wakes up" {
			if guard.sleepSchedule == nil {
				guard.sleepSchedule = make(map[int]int)
			}

			for i := startedSleepingAt; i < time; i++ {
				guard.sleepSchedule[i]++
				guard.minutesAsleep++

				if guard.sleepSchedule[i] > guard.sleepSchedule[guard.sleepiestMinute] {
					guard.sleepiestMinute = i
				}

				if i == guard.sleepiestMinute {
					guard.sameMinuteCount++
				}
			}

			guards[currentGuardID] = guard

		} else {
			tokens := strings.Split(action, " ")
			currentGuardID, _ = strconv.Atoi(tokens[1][1:])

			_, ok := guards[currentGuardID]
			if !ok {
				guards[currentGuardID] = Guard{currentGuardID, 0, 0, 0, make(map[int]int)}
			}

		}
	}

	return guards
}
