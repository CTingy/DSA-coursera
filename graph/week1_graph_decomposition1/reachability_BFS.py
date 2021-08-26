# WARN: counter result cannot represent the all the possibilities.
# If want to find all possible route, one pass is not enough. 

def BFS(graph, start, end):
    counter = [0 for _ in range(len(graph))]
    queue = [start]
    counter[start] = 0
    while queue:
        node = queue.pop(0)  # take the first one
        for neighbor in graph[node]:
            if counter[neighbor] == 0 and neighbor != start:
                queue.append(neighbor)
            counter[neighbor] += 1
    return counter[end]


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
    if res >= 1:
        print(1)
    else:
        print(0)
