package dsa

import (
	"container/heap"
	"github.com/s-platform/gore/utilities"
	"math"
	"sort"
	"strconv"
	"strings"
)

func LongestCommonPrefix(strArr []string) (prefix string) {
	if len(strArr) == 0 {
		return
	}
	prefix = strArr[0]

	for _, v := range strArr {
		for utilities.Substring(v, 0, utilities.ToUint(len(prefix))) != prefix && len(prefix) > 0 {
			prefix = utilities.Substring(prefix, 0, utilities.ToUint(len(prefix)-1))
		}
	}

	return
}

func LongestCommonPrefixWithSort(strArr []string) (prefix string) {
	if len(strArr) == 0 {
		return
	}
	sort.SliceStable(strArr, func(i, j int) bool {
		return len(strArr[i]) < len(strArr[j])
	})

	first := strArr[0]
	last := strArr[len(strArr)-1]
	for i := 0; i < len(first); i++ {
		if string(first[i]) != string(last[i]) {
			return prefix
		}
		prefix = prefix + string(first[i])
	}
	return
}

func TwoSum(nums []int, target int) (res []int) {

	mNumberIndex := make(map[int]int)
	for index, value := range nums {
		mNumberIndex[value] = index
	}

	for index, v := range nums {
		incomplete := target - v
		if _, exist := mNumberIndex[incomplete]; exist {
			if index == mNumberIndex[incomplete] {
				continue
			}
			return []int{index, mNumberIndex[incomplete]}
		}
	}

	return
}

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	strInt := strconv.Itoa(x)
	if len(strInt)%2 != 0 {
		left := strInt[0 : len(strInt)/2]
		right := strInt[len(strInt)/2+1:]
		return left == utilities.ReverseString(right)
	} else {
		left := strInt[0 : len(strInt)/2]
		right := strInt[len(strInt)/2:]
		return left == utilities.ReverseString(right)
	}
}

func RomanToInt(s string) (value int) {
	rValue := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
	}
	subTractValue := map[string]int{
		"IV": 3,
		"IX": 8,
		"XL": 30,
		"XC": 80,
		"CD": 300,
		"CM": 800,
	}
	if v, exist := rValue[s]; exist {
		return v
	}

	arrCharacter := strings.Split(s, "")
	for index, element := range arrCharacter {
		addValue := rValue[element]
		if index == 0 {
			value += addValue
		} else {
			_, exist := subTractValue[arrCharacter[index-1]+element]
			if exist {
				value += subTractValue[arrCharacter[index-1]+element]
				continue
			}
			_, exist = rValue[element]
			if exist {
				value += rValue[element]
			}
		}
	}

	return
}

func IntToRoman(value int) (roman string) {
	rValue := map[int]string{
		1:    "I",
		2:    "II",
		3:    "III",
		4:    "IV",
		5:    "V",
		6:    "VI",
		7:    "VII",
		8:    "VIII",
		9:    "IX",
		10:   "X",
		20:   "XX",
		30:   "XXX",
		40:   "XL",
		50:   "L",
		60:   "LI",
		70:   "LII",
		80:   "LIII",
		90:   "XC",
		100:  "C",
		200:  "CC",
		300:  "CCC",
		400:  "CD",
		500:  "D",
		600:  "DC",
		700:  "DCC",
		800:  "DCCC",
		900:  "CM",
		1000: "M",
		2000: "MM",
		3000: "MMM",
	}

	l := strconv.Itoa(value)
	for i := len(l) - 1; i >= 0; i-- {
		v := value - value%int(math.Pow(10, float64(i)))
		value -= v
		roman += rValue[v]
	}
	return roman
}

func IsValidParentheses(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]string, 0)
	runes := strings.Split(s, "")
	for _, runeElement := range runes {
		if runeElement == "(" {
			stack = append(stack, ")")
		} else if runeElement == "[" {
			stack = append(stack, "]")
		} else if runeElement == "{" {
			stack = append(stack, "}")
		} else if runeElement == ")" || runeElement == "]" || runeElement == "}" {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if last == runeElement {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	current := new(ListNode)

	tmp := current

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}
	return tmp
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	nums = nums[:j+1]
	return j + 1
}

func MinimumTime(grid [][]int) int {
	// If the initial conditions are not met, return -1
	if grid[1][0] > 1 && grid[0][1] > 1 {
		return -1
	}

	R, C := len(grid), len(grid[0])

	// Helper function to check if a cell is outside the grid
	isOutside := func(i, j int) bool {
		return i < 0 || i >= R || j < 0 || j >= C
	}

	// Helper function to calculate the index of a cell in a 1D representation
	idx := func(i, j int) int {
		return i*C + j
	}

	N := R * C
	time := make([]int, N)
	for i := range time {
		time[i] = math.MaxInt32
	}

	// Priority queue implementation
	pq := &PriorityQueueCell{}
	heap.Init(pq)

	// Start from the top-left corner
	heap.Push(pq, &Item{time: 0, pos: 0})
	time[0] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		t, ij := current.time, current.pos
		i, j := ij/C, ij%C

		// If we reach the bottom-right corner, return the time
		if i == R-1 && j == C-1 {
			return t
		}

		// Explore neighboring cells
		for _, d := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			r, s := i+d[0], j+d[1]
			if isOutside(r, s) {
				continue
			}

			w := 0
			if (grid[r][s]-t)&1 == 0 {
				w = 1
			}
			nextTime := max(t+1, grid[r][s]+w)

			rs := idx(r, s)
			if nextTime < time[rs] {
				time[rs] = nextTime
				heap.Push(pq, &Item{time: nextTime, pos: rs})
			}
		}
	}

	return -1
}

// Helper function for max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Item represents an entry in the priority queue
type Item struct {
	time int // Current time
	pos  int // Position in the grid (1D index)
}

// PriorityQueueCell is a min-heap of Items
type PriorityQueueCell []*Item

func (pq PriorityQueueCell) Len() int { return len(pq) }

func (pq PriorityQueueCell) Less(i, j int) bool {
	return pq[i].time < pq[j].time
}

func (pq PriorityQueueCell) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueueCell) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueueCell) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
