def BFS(graph, start, end):
    distance = [-1 for _ in range(len(graph))]
    queue = [start]
    distance[start] = 0
    while queue:
        node = queue.pop(0)  # take the first one
        for neighbor in graph[node]:
            if distance[neighbor] == -1:
                queue.append(neighbor)
                distance[neighbor] = distance[node] + 1
    return distance[end]


if __name__ == '__main__':
    # input 1 is the number of vertices, input 2 is the number of edges
    n_vertices , n_edges = map(int, input().split())
    graph = [[] for _ in range(n_vertices+1)]
    for _ in range(n_edges):
        start, end = map(int, input().split())            
        graph[start].append(end)
        graph[end].append(start)
    
    start_node, end_node = map(int, input().split())
    res = BFS(graph, start_node, end_node)
    print(res)
