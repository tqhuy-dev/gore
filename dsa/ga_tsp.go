package dsa

import (
	"fmt"
	"math/rand"
	"time"
)

// Kích thước quần thể & số thế hệ
const (
	popSize      = 100  // Kích thước quần thể
	genCount     = 1000 // Số thế hệ tối đa
	mutationRate = 0.1  // Xác suất đột biến
)

// Thành phố với tọa độ x, y
type City struct {
	x, y float64
}

// Một cá thể (giải pháp) là một hoán vị các thành phố
type Individual struct {
	route   []int   // Hoán vị các thành phố
	fitness float64 // Tổng khoảng cách của tuyến đường
}

// Danh sách thành phố (giả lập)
var cities = []City{
	{2, 3}, {5, 8}, {6, 3}, {1, 9}, {3, 7}, {7, 2}, {8, 6}, {4, 4},
}

// Tính khoảng cách giữa hai thành phố
func distance(a, b City) float64 {
	dx, dy := a.x-b.x, a.y-b.y
	return (dx*dx + dy*dy) // Tránh dùng sqrt để tăng tốc
}

// Tính độ thích nghi (fitness) của một cá thể (tổng khoảng cách)
func calcFitness(ind *Individual) {
	totalDist := 0.0
	for i := 0; i < len(ind.route)-1; i++ {
		totalDist += distance(cities[ind.route[i]], cities[ind.route[i+1]])
	}
	totalDist += distance(cities[ind.route[len(ind.route)-1]], cities[ind.route[0]]) // Quay lại điểm xuất phát
	ind.fitness = totalDist
}

// Tạo một cá thể ngẫu nhiên
func createRandomIndividual() Individual {
	route := rand.Perm(len(cities)) // Hoán vị ngẫu nhiên
	ind := Individual{route: route}
	calcFitness(&ind)
	return ind
}

// Chọn lọc theo phương pháp Tournament Selection
func tournamentSelection(population []Individual) Individual {
	best := population[rand.Intn(len(population))]
	for i := 0; i < 5; i++ { // Chọn 5 cá thể ngẫu nhiên, lấy cá thể tốt nhất
		competitor := population[rand.Intn(len(population))]
		if competitor.fitness < best.fitness {
			best = competitor
		}
	}
	return best
}

// Lai ghép (Order Crossover - OX1)
func crossover(parent1, parent2 Individual) Individual {
	size := len(parent1.route)
	start, end := rand.Intn(size), rand.Intn(size)

	if start > end {
		start, end = end, start
	}

	child := make([]int, size)
	copy(child[start:end+1], parent1.route[start:end+1])

	used := make(map[int]bool)
	for _, v := range child[start : end+1] {
		used[v] = true
	}

	idx := (end + 1) % size
	for _, gene := range parent2.route {
		if !used[gene] {
			child[idx] = gene
			used[gene] = true
			idx = (idx + 1) % size
		}
	}

	newInd := Individual{route: child}
	calcFitness(&newInd)
	return newInd
}

// Đột biến (Swap Mutation)
func mutate(ind *Individual) {
	if rand.Float64() < mutationRate {
		i, j := rand.Intn(len(ind.route)), rand.Intn(len(ind.route))
		ind.route[i], ind.route[j] = ind.route[j], ind.route[i]
		calcFitness(ind)
	}
}

// Tiến hóa quần thể
func evolve(population []Individual) []Individual {
	newPop := make([]Individual, 0, len(population))

	// Giữ cá thể tốt nhất (Elitism)
	best := population[0]
	for _, ind := range population {
		if ind.fitness < best.fitness {
			best = ind
		}
	}
	newPop = append(newPop, best)

	// Lai ghép và tạo thế hệ mới
	for len(newPop) < len(population) {
		p1, p2 := tournamentSelection(population), tournamentSelection(population)
		child := crossover(p1, p2)
		mutate(&child)
		newPop = append(newPop, child)
	}

	return newPop
}

// Chạy thuật toán GA
func geneticAlgorithm() Individual {
	rand.Seed(time.Now().UnixNano())

	// Khởi tạo quần thể
	population := make([]Individual, popSize)
	for i := 0; i < popSize; i++ {
		population[i] = createRandomIndividual()
	}

	// Tiến hóa
	for gen := 0; gen < genCount; gen++ {
		population = evolve(population)
	}

	// Trả về cá thể tốt nhất
	best := population[0]
	for _, ind := range population {
		if ind.fitness < best.fitness {
			best = ind
		}
	}
	return best
}

func RunGA() {
	bestSolution := geneticAlgorithm()
	fmt.Println("Best Route:", bestSolution.route)
	fmt.Println("Total Distance:", bestSolution.fitness)
}
