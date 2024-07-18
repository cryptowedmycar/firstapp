package main

import "fmt"

func canAttendMeetings(intervals [][]int) bool {
	var canAttend bool

	earlyStart, lateEnd := 0, -1

	for i := 0; i < len(intervals); i++ {

		if intervals[i][0] <= earlyStart {
			earlyStart = intervals[i][0]
		}

		if intervals[i][1] > lateEnd {
			lateEnd = intervals[i][1]
		}

	}

	fmt.Println(earlyStart)
	fmt.Println(lateEnd)

	//fmt.Println(canAttend)
	return canAttend
}

func main() {
	var intervals = [][]int{{0, 30}, {5, 10}, {15, 20}}

	canAttendMeetings(intervals)

	//fmt.Println(intervals)

}
