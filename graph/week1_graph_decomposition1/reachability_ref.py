
# ref: https://github.com/Sonia-96/Coursera-Data_Structures_and_Algorithms/blob/master/3-Algorithms%20on%20Graphs/Week1-Undirected%20Graphs/1-reachability.py

def reach(adj, visited, x, y):
    visited[x][0] = True
    for vertex in adj[x]:
        visited[vertex][1] += 1
        if not visited[vertex][0]:
            reach(adj, visited, vertex, y)


if __name__ == '__main__':
    n_vertices, n_edges = map(int, input().split())
    edges = []
    adjacency_list = [[] for _ in range(n_vertices + 1)]
    for i in range(n_edges):
        edges.append(tuple(map(int, input().split())))
    for (a, b) in edges:
        adjacency_list[a].append(b)
        adjacency_list[b].append(a)
    u, v = map(int, input().split())
    visited = [[False, 0] for _ in range(n_vertices + 1)]  # index 0 mark visited, # index 1 count
    reach(adjacency_list, visited, u, v)
    print(visited[v][1])
