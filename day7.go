package AdventOfCode

import (
	"fmt"
	"sort"
	"strings"
)

type Node struct {
	Name string
	From []string
	To []string
	Steps int
}

type Worker struct {
	StepsLeft int
	Assigned string
	MaxSteps int
	Node Node
}

func Main7() {
	graph := make(map[string]Node)
	data := strings.Split(Day7Data, "\n")
	for _, strEdge := range data {
		from, to := GetEdge(strEdge)
		graph[to] = Node{Name:to, From: append(graph[to].From, from), To: graph[to].To, Steps: 60 + int(rune(to[0])) - 64}
		graph[from] = Node{Name:from, From: graph[from].From, To: append(graph[from].To, to), Steps: 60 + int(rune(from[0])) - 64}
	}
	fmt.Println(MultipleWorkers(graph, 5))
}

func MultipleWorkers(graph map[string]Node, numWorker int) int {
	steps := 0
	workers := make([]Worker, numWorker)
	for len(graph) > 0 || IsSomeoneWorking(workers) {
		for id := range workers {
			if workers[id].StepsLeft == 0 {
				taskNode, temp_G := GetTaskAndRemove(graph)
				if taskNode.Name != "" {
					graph = temp_G
					workers[id].Node = taskNode
					workers[id].MaxSteps = taskNode.Steps
					workers[id].Assigned = taskNode.Name
					workers[id].StepsLeft = taskNode.Steps
				}
			}
			if !(workers[id].StepsLeft == 0) {
				workers[id].StepsLeft -= 1
			}
			if workers[id].StepsLeft == 0 && workers[id].Assigned != "" {
				fmt.Println(workers)
				graph = RemoveFrom(graph, workers[id].Assigned, workers[id].Node.To)
				workers[id] = Worker{}
			}
		}
		steps += 1
	}

	return steps
}

func RemoveFrom(graph map[string]Node, nodeName string, nodes []string) map[string]Node {
	for _, toNode := range nodes {
		node := graph[toNode]
		node.From = RemoveNode(nodeName, node.From)
		graph[toNode] = node
	}
	return graph
}

func GetTaskAndRemove(graph map[string]Node) (Node, map[string]Node) {
	startNodes := make([]string, 0)
	for nodeName, node := range graph {
		if len(node.From) == 0 {
			startNodes = append(startNodes, nodeName)
		}
	}

	if len(startNodes) == 0 {
		return Node{}, graph
	}

	sort.Strings(startNodes)
	toDelete := startNodes[0]
	node := graph[toDelete]
	delete(graph, toDelete)
	return node, graph
}

func IsSomeoneWorking(workers []Worker) bool {
	for _, worker := range workers {
		if !(worker.StepsLeft == 0) {
			return true
		}
	}
	return false
}

func MultipleTraverse(graph map[string]Node) int {
	var workers [5]Worker
	var startNodes []string
	steps := 0
	for _, node := range graph {
		fmt.Println(node)
	}
	for len(graph) > 0 {
		for nodeName, node := range graph {
			if len(node.From) == 0 && !strings.Contains(strings.Join(startNodes, ""), nodeName) {
				startNodes = append(startNodes, nodeName)
			}
		}
		for i, worker := range workers {
			if worker.StepsLeft == 0 && worker.MaxSteps > 60 {
				for _, toNode := range worker.Node.To {
					node := graph[toNode]
					node.From = RemoveNode(worker.Assigned, node.From)
					graph[toNode] = node
				}
			}
			if worker.StepsLeft == 0 && len(startNodes) > 0 {
				sort.Strings(startNodes)
				toDelete := startNodes[0]
				startNodes = startNodes[1:]
				node := graph[toDelete]

				worker.Assigned = toDelete
				worker.MaxSteps = node.Steps
				worker.StepsLeft = worker.MaxSteps
				worker.Node = node

				if node.Name != "" {
					fmt.Println("Before", workers)
				}

				workers[i] = worker

				if node.Name != "" {
					fmt.Println(node)
					fmt.Println(workers)
					fmt.Println(startNodes)
				}
				delete(graph, toDelete)
			}
			if worker.StepsLeft > 0 {
				workers[i].StepsLeft -= 1
			}
		}
		steps += 1
	}
	return steps
}

func Traverse(graph map[string]Node) string {
	var startNodes []string
	var b strings.Builder
	for _, node := range graph {
		fmt.Println(node)
	}
	for len(graph) > 0 {
		startNodes = make([]string, 0)
		for nodeName, node := range graph {
			if len(node.From) == 0 {
				startNodes = append(startNodes, nodeName)
			}
		}
		fmt.Println(startNodes)

		sort.Strings(startNodes)
		toDelete := startNodes[0]

		for _, toNode := range graph[toDelete].To {
			node := graph[toNode]
			node.From = RemoveNode(toDelete, node.From)
			graph[toNode] = node
		}

		delete(graph, toDelete)
		fmt.Println(toDelete)
		b.WriteString(toDelete)
	}

	return b.String()
}

func RemoveNode(s string, nodes []string) []string {
	var ans []string
	for _, node := range nodes {
		if node == s {
			continue
		}
		ans = append(ans, node)
	}
	return ans
}

func PopLowerNode(nodes []string) (string, []string) {
	lower := "Z"
	for _, node := range nodes {
		if rune(node[0]) < rune(lower[0]) {
			lower = node
		}
	}
	return lower, RemoveNode(lower, nodes)
}

func GetEdge(s string) (string, string) {
	data := strings.Split(s, " ")
	return data[1], data[7]
}

const Test7Data = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

const Day7Data = `Step V must be finished before step H can begin.
Step U must be finished before step R can begin.
Step E must be finished before step D can begin.
Step B must be finished before step R can begin.
Step W must be finished before step X can begin.
Step A must be finished before step P can begin.
Step T must be finished before step L can begin.
Step F must be finished before step C can begin.
Step P must be finished before step Y can begin.
Step N must be finished before step G can begin.
Step R must be finished before step S can begin.
Step D must be finished before step C can begin.
Step O must be finished before step K can begin.
Step L must be finished before step J can begin.
Step J must be finished before step H can begin.
Step M must be finished before step I can begin.
Step G must be finished before step K can begin.
Step Z must be finished before step Q can begin.
Step X must be finished before step Q can begin.
Step H must be finished before step I can begin.
Step K must be finished before step Y can begin.
Step Q must be finished before step S can begin.
Step I must be finished before step Y can begin.
Step S must be finished before step Y can begin.
Step C must be finished before step Y can begin.
Step T must be finished before step S can begin.
Step P must be finished before step S can begin.
Step I must be finished before step S can begin.
Step V must be finished before step O can begin.
Step O must be finished before step Q can begin.
Step T must be finished before step R can begin.
Step E must be finished before step J can begin.
Step F must be finished before step S can begin.
Step O must be finished before step H can begin.
Step Z must be finished before step S can begin.
Step D must be finished before step Z can begin.
Step F must be finished before step K can begin.
Step W must be finished before step P can begin.
Step G must be finished before step I can begin.
Step B must be finished before step T can begin.
Step G must be finished before step Y can begin.
Step X must be finished before step S can begin.
Step B must be finished before step K can begin.
Step V must be finished before step A can begin.
Step U must be finished before step N can begin.
Step T must be finished before step P can begin.
Step V must be finished before step D can begin.
Step G must be finished before step X can begin.
Step B must be finished before step D can begin.
Step R must be finished before step J can begin.
Step M must be finished before step Z can begin.
Step U must be finished before step Z can begin.
Step U must be finished before step G can begin.
Step A must be finished before step C can begin.
Step H must be finished before step Q can begin.
Step X must be finished before step K can begin.
Step B must be finished before step S can begin.
Step Q must be finished before step C can begin.
Step Q must be finished before step Y can begin.
Step R must be finished before step I can begin.
Step V must be finished before step Q can begin.
Step A must be finished before step D can begin.
Step D must be finished before step S can begin.
Step K must be finished before step S can begin.
Step G must be finished before step C can begin.
Step D must be finished before step O can begin.
Step R must be finished before step H can begin.
Step K must be finished before step Q can begin.
Step W must be finished before step R can begin.
Step H must be finished before step Y can begin.
Step P must be finished before step J can begin.
Step N must be finished before step Z can begin.
Step J must be finished before step K can begin.
Step W must be finished before step M can begin.
Step A must be finished before step Z can begin.
Step V must be finished before step W can begin.
Step J must be finished before step X can begin.
Step U must be finished before step F can begin.
Step P must be finished before step L can begin.
Step W must be finished before step G can begin.
Step T must be finished before step F can begin.
Step R must be finished before step C can begin.
Step R must be finished before step O can begin.
Step Z must be finished before step C can begin.
Step E must be finished before step S can begin.
Step L must be finished before step I can begin.
Step U must be finished before step O can begin.
Step W must be finished before step K can begin.
Step K must be finished before step I can begin.
Step O must be finished before step M can begin.
Step V must be finished before step M can begin.
Step V must be finished before step Z can begin.
Step A must be finished before step I can begin.
Step F must be finished before step J can begin.
Step F must be finished before step O can begin.
Step M must be finished before step C can begin.
Step Q must be finished before step I can begin.
Step H must be finished before step S can begin.
Step U must be finished before step A can begin.
Step J must be finished before step S can begin.
Step P must be finished before step Z can begin.`
