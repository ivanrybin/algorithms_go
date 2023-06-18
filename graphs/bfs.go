package graphs

import "container/list"

func BFS(graph Graph, start int) (map[int]int, map[int]int) {
	dist := make(map[int]int, len(graph))
	prev := make(map[int]int, len(graph))
	dist[start] = 0
	prev[start] = start
	queue := list.List{}
	queue.PushBack(start)
	for queue.Len() > 0 {
		v := queue.Front()
		queue.Remove(v)
		for _, u := range graph[v.Value.(int)] {
			if _, ok := dist[u]; !ok {
				dist[u] = dist[v.Value.(int)] + 1
				prev[u] = v.Value.(int)
				queue.PushBack(u)
			}
		}
	}
	return dist, prev
}
