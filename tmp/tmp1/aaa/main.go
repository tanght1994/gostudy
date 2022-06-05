package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

func aaa() *TreeNode {
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n3 := &TreeNode{3, nil, nil}
	n5 := &TreeNode{5, nil, nil}
	n6 := &TreeNode{6, nil, nil}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n5
	n2.Right = n6
	return n1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}

func makegraph(beginWord string, endWord string, wordList []string) (graph [][]int, begin, end int) {
	// wordid = {"word1": 0, "word2": 1, ...}
	wordid := map[string]int{}
	id := 0
	for _, v := range wordList {
		if _, ok := wordid[v]; ok {
			wordid[v] = id
			id++
		}
	}
	if _, ok := wordid[beginWord]; !ok {
		wordid[beginWord] = id
	}

	if _, ok := wordid[endWord]; !ok {
		return nil, 0, 0
	}

	// 去重之后的字符串list
	words := make([]string, len(wordid))
	for word := range wordid {
		words = append(words, word)
	}

	/*
		graph格式如下
		[
			[3, 4],  		// 0 可以转化为 3, 4
			[],  			// 1 无法转到任何地方
			[0, 1, 3, 4],  	// 2 可以转化为 0, 1, 3, 4
			[0, 4],  		// 3 可以转化为 0, 4
			[0, 3],  		// 4 可以转化为 0, 3
		]
	*/
	graph = make([][]int, len(words))

	for i := range words {
		word1 := words[i]
		transform := make([]int, 0)
		for j := i + 1; j < len(words); j++ {
			word2 := words[j]
			if cantrans(word1, word2) {
				transform = append(transform, wordid[word2])
			}
		}
		graph[i] = transform
	}

	return graph, wordid[beginWord], wordid[endWord]
}

func ladderLength1(beginWord string, endWord string, wordList []string) int {
	return 0
	// if beginWord == endWord {
	// 	return 1
	// }

	// graph, begin, end := makegraph(beginWord, endWord, wordList)
	// if graph == nil {
	// 	return 0
	// }

	// min := math.MaxInt
	// visited := []int{}
	// level := 1

	// paths1 := graph[begin]
	// paths2 := graph[begin]

	// return min
}

func cantrans(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}
		diff++
		if diff == 2 {
			break
		}
	}
	return diff <= 1
}

func indexof(array []int, target int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			return i
		}
	}
	return -1
}
