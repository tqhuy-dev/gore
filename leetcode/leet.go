package leetcode

import "sort"

func Merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	// Merge in reverse order
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// If there are remaining elements in nums2
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}

func GetServerIndex(n int32, arrival []int32, burstTime []int32) []int32 {

	type requestTmp struct {
		index   int32
		arrival int32
		burst   int32
	}

	m := int32(len(arrival))
	requests := make([]requestTmp, m)
	for i := int32(0); i < m; i++ {
		requests[i] = requestTmp{i, arrival[i], burstTime[i]}
	}

	sort.SliceStable(requests, func(i, j int) bool {
		if requests[i].arrival == requests[j].arrival {
			return requests[i].index < requests[j].index
		}
		return requests[i].arrival < requests[j].arrival
	})

	arrCountTimeServer := make([]int32, n)
	result := make([]int32, m)

	for _, req := range requests {
		pendingFlg := false
		for i := 0; i < int(n); i++ {
			if arrCountTimeServer[i] <= req.arrival {
				arrCountTimeServer[i] = req.arrival + req.burst
				result[req.index] = int32(i + 1)
				pendingFlg = true
				break
			}
		}
		if !pendingFlg {
			result[req.index] = -1
		}
	}

	return result
}

func RoadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	// Write your code here
	if c_road >= c_lib {
		return int64(n * c_lib)
	}

	// Build adjacency list
	graph := make(map[int32][]int32)
	for _, edge := range cities {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n+1)

	var bfs func(start int32) int64
	bfs = func(start int32) int64 {
		queue := []int32{start}
		visited[start] = true
		count := int64(1)

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
					count++
				}
			}
		}

		return count
	}

	var totalCost int64
	for city := int32(1); city <= n; city++ {
		if !visited[city] {
			componentSize := bfs(city)
			totalCost += int64(c_lib) + (componentSize-1)*int64(c_road)
		}
	}

	return totalCost
}
