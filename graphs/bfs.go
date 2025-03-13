package graphs

import "container/list"

// BFS O(V+E) / O(V) (time / mem).
func BFS(graph Graph, start int) (map[int]int, map[int]int) {
	dist := make(map[int]int, len(graph))
	prev := make(map[int]int, len(graph))
	dist[start] = 0
	prev[start] = start
	queue := list.List{}
	queue.PushBack(start)
	for queue.Len() > 0 {
		v := queue.Remove(queue.Front()).(int)
		for _, u := range graph[v] {
			if _, ok := dist[u]; !ok {
				dist[u] = dist[v] + 1
				prev[u] = v
				queue.PushBack(u)
			}
		}
	}
	return dist, prev
}
