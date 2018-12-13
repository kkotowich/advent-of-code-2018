package main

import (
	fileutil "code-advent/utils"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

// State of each step
type State int

// State of each step
const (
	NotStarted State = 0
	InProgress State = 1
	Done       State = 2
)

// Step determines previous and next step
type Step struct {
	id             rune
	nextSteps      []rune
	prevSteps      []rune
	state          State
	runningTime    int
	timeToComplete int
}

// Worker is a resource to work on a step
type Worker struct {
	id     int
	isIdle bool
	stepID rune
}

func main() {
	defer log.Printf("main took %s", time.Since(time.Now()))

	input, err := fileutil.GetStringSliceFromFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	stepTree := buildStepTree(input)
	orderedSteps := orderSteps(stepTree)
	result1 := getStepOrder(orderedSteps)

	stepTree = buildStepTree(input)
	workers := createWorkers(2)

	workingSteps := []rune{orderedSteps[0].id}

	result2 := doWork(workingSteps, stepTree, workers)

	fmt.Println(result1)
	fmt.Println(result2)
}

func doWork(workingStepIDs []rune, steps map[rune]Step, workers []Worker) int {
	elapsedSeconds := 0
	isDone := false

	for !isDone {
		for _, id := range workingStepIDs {
			workingStep := steps[id]

			switch workingStep.state {

			case NotStarted:
				allowWork := true

				for _, prevStep := range workingStep.prevSteps {
					if steps[prevStep].state != Done {
						allowWork = false
						break
					}
				}

				if allowWork {
					for workerIndex := range workers {
						if workers[workerIndex].isIdle {
							workingStep.state = InProgress
							workingStep.runningTime = 0
							workers[workerIndex].isIdle = false
							workers[workerIndex].stepID = workingStep.id
							break
						}
					}
				}

			case InProgress:
				workingStep.runningTime++
				if workingStep.runningTime == workingStep.timeToComplete {
					workingStep.state = Done

					for _, nextStep := range workingStep.nextSteps {
						allowAdd := true
						for _, workingStepID := range workingStepIDs {
							if workingStepID == nextStep {
								allowAdd = false
							}
						}
						if allowAdd {
							workingStepIDs = append(workingStepIDs, steps[nextStep].id)

							allowWork := true

							for _, prevStep := range steps[nextStep].prevSteps {
								if steps[prevStep].state != Done {
									allowWork = false
									break
								}
							}

							if allowWork {
								for workerIndex := range workers {
									if workers[workerIndex].isIdle {
										newStep := steps[nextStep]
										newStep.state = InProgress
										newStep.runningTime = 0
										workers[workerIndex].isIdle = false
										workers[workerIndex].stepID = newStep.id
										steps[nextStep] = newStep

										fmt.Println("hiiii")
										break
									}
								}
							}

						}

					}

					s := strings.Split(string(workingStepIDs), "")
					sort.Strings(s)
					workingStepIDs = []rune(strings.Join(s, ""))

					for workerIndex := range workers {
						if !workers[workerIndex].isIdle {
							workers[workerIndex].isIdle = true
							workers[workerIndex].stepID = 0
							break
						}
					}
				}
			case Done:
			}

			steps[workingStep.id] = workingStep
		}
		// fmt.Println(steps)
		// fmt.Println(strings.Split(string(workingStepIDs), ""))

		// for _, step := range steps {
		// 	fmt.Println(step)
		// }
		// fmt.Println(workers)
		// fmt.Println(len(workingStepIDs))
		// fmt.Println(len(steps))

		// fmt.Println("time")
		// fmt.Println(elapsedSeconds)
		// fmt.Println("workers")
		for _, worker := range workers {
			if worker.stepID == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(worker.stepID))
			}
		}
		fmt.Println("")

		// fmt.Println("---------")

		elapsedSeconds++
		// isDone = elapsedSeconds == 15
		isDone = isAllWorkersIdle(workers) && len(workingStepIDs) == len(steps)
	}
	return elapsedSeconds
}

func isAllWorkersIdle(workers []Worker) bool {
	for _, worker := range workers {
		if !worker.isIdle {
			return false
		}
	}
	return true
}

func createWorkers(workerCount int) []Worker {
	var workers []Worker
	for i := 0; i < workerCount; i++ {
		workers = append(workers, Worker{i, true, 0})
	}
	return workers
}

func buildStepTree(input []string) map[rune]Step {
	steps := make(map[rune]Step)
	for _, line := range input {
		parentStep, childStep := parseLine(line)

		_, ok := steps[parentStep.id]
		if !ok {
			steps[parentStep.id] = parentStep
		} else {
			existingStep := steps[parentStep.id]

			existingNextSteps := existingStep.nextSteps
			for _, stepID := range parentStep.nextSteps {
				existingNextSteps = append(existingNextSteps, stepID)
			}
			existingStep.nextSteps = existingNextSteps

			s := strings.Split(string(existingStep.nextSteps), "")
			sort.Strings(s)
			existingStep.nextSteps = []rune(strings.Join(s, ""))

			steps[parentStep.id] = existingStep
		}

		_, ok = steps[childStep.id]
		if !ok {
			steps[childStep.id] = childStep
		} else {
			existingStep := steps[childStep.id]

			existingPrevSteps := existingStep.prevSteps
			for _, stepID := range childStep.prevSteps {
				existingPrevSteps = append(existingPrevSteps, stepID)
			}
			existingStep.prevSteps = existingPrevSteps

			s := strings.Split(string(existingStep.prevSteps), "")
			sort.Strings(s)
			existingStep.prevSteps = []rune(strings.Join(s, ""))

			steps[childStep.id] = existingStep
		}
	}
	return steps
}

func parseLine(input string) (Step, Step) {
	tokens := strings.Split(input, " ")

	parentTTC := int([]rune(tokens[1])[0]) - 64
	childTTC := int([]rune(tokens[7])[0]) - 64

	parentStep := Step{[]rune(tokens[1])[0], []rune(tokens[7]), nil, NotStarted, 0, parentTTC}
	childStep := Step{[]rune(tokens[7])[0], nil, []rune(tokens[1]), NotStarted, 0, childTTC}

	return parentStep, childStep
}

func getStepOrder(input []Step) string {
	var stepString string

	for _, step := range input {
		stepString = stepString + string(step.id)
	}

	return string(stepString)
}

func orderSteps(input map[rune]Step) []Step {
	remainingSteps := input
	var sortedSteps []Step

	for len(remainingSteps) > 0 {
		firstStep := findNextStep(remainingSteps)
		sortedSteps = append(sortedSteps, firstStep)
		delete(remainingSteps, firstStep.id)
	}

	return sortedSteps
}

func findNextStep(input map[rune]Step) Step {
	var topStepIDs []rune

	for _, currentStep := range input {
		isFirst := true
		for _, nextStep := range input {
			for _, nextStep := range nextStep.nextSteps {
				if currentStep.id == nextStep {
					isFirst = false
				}
			}
		}
		if isFirst {
			topStepIDs = append(topStepIDs, currentStep.id)
		}
	}

	s := strings.Split(string(topStepIDs), "")
	sort.Strings(s)
	topStepIDs = []rune(strings.Join(s, ""))

	return input[topStepIDs[0]]
}

func printStep(step Step) {
	fmt.Println(string(step.id))
	fmt.Println(string(step.nextSteps))
}
