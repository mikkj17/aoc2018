package day04

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/mikkj17/aoc2018/utils"
	"github.com/mikkj17/aoc2018/utils/sets"
)

const dateFormat = "2006-01-02 15:04"
const beginPattern = `^\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2})\] Guard #(\d+) begins shift$`
const sleepPattern = `^\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2})\] falls asleep$`
const wakeUpPattern = `^\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2})\] wakes up$`

func parse(inp string) []Record {
	lines := strings.Split(inp, "\n")
	slices.Sort(lines)

	beginRegex := regexp.MustCompile(beginPattern)
	sleepRegex := regexp.MustCompile(sleepPattern)
	wakeUpRegex := regexp.MustCompile(wakeUpPattern)
	var currentGuard int
	var currentTimestamp time.Time
	var currentAction Action
	records := []Record{}

	for _, line := range lines {
		match := beginRegex.FindStringSubmatch(line)
		if len(match) > 0 {
			currentGuard = utils.ToInt(match[2])
			currentAction = begins
			currentTimestamp, _ = time.Parse(dateFormat, match[1])
		}

		match = sleepRegex.FindStringSubmatch(line)
		if len(match) > 0 {
			currentAction = sleeps
			currentTimestamp, _ = time.Parse(dateFormat, match[1])
		}

		match = wakeUpRegex.FindStringSubmatch(line)
		if len(match) > 0 {
			currentAction = wakesUp
			currentTimestamp, _ = time.Parse(dateFormat, match[1])
		}

		records = append(records, Record{
			guardId:   currentGuard,
			action:    currentAction,
			timeStamp: currentTimestamp,
		})
	}

	return records
}

func sleepyGuards(records []Record) (int, map[int][]Nap) {
	guardTimers := map[int]int{}
	naps := map[int][]Nap{}
	var sleepStart time.Time

	for _, record := range records {
		if record.action == sleeps {
			sleepStart = record.timeStamp
		} else if record.action == wakesUp {
			guardTimers[record.guardId] += int(record.timeStamp.Sub(sleepStart).Minutes())
			naps[record.guardId] = append(naps[record.guardId], Nap{
				guardId:     record.guardId,
				startMinute: sleepStart.Minute(),
				endMinute:   record.timeStamp.Minute(),
			})
		}
	}

	sleeper, _ := utils.MaxValue(guardTimers)

	return sleeper, naps
}

func minuteCounter(naps []Nap) map[int]int {
	minutesAsleep := map[int]int{}
	for _, nap := range naps {
		for minute := nap.startMinute; minute < nap.endMinute; minute++ {
			minutesAsleep[minute]++
		}
	}

	return minutesAsleep
}

func partOne(inp string) int {
	records := parse(inp)
	sleeper, naps := sleepyGuards(records)
	minutesAsleep := minuteCounter(naps[sleeper])
	maxMinute, _ := utils.MaxValue(minutesAsleep)

	return sleeper * maxMinute
}

func partTwo(inp string) int {
	records := parse(inp)
	distinctGuards := sets.FromSlice(utils.Map(records, func(record Record) int {
		return record.guardId
	}))

	_, naps := sleepyGuards(records)
	maxMinutes := map[int]int{}
	for _, guardId := range sets.ToSlice(distinctGuards) {
		minutesAsleep := minuteCounter(naps[guardId])
		if len(minutesAsleep) == 0 {
			continue
		}
		_, count := utils.MaxValue(minutesAsleep)
		maxMinutes[guardId] = count
	}

	maxGuard, _ := utils.MaxValue(maxMinutes)
	minutesAsleep := minuteCounter(naps[maxGuard])
	maxMinute, _ := utils.MaxValue(minutesAsleep)

	return maxGuard * maxMinute
}

func Solve() {
	input := utils.ReadInput(4)
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
