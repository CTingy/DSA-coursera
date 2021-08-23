clock = 0

def topological_sort(n, edges, visited):
    nodes_order = {node: 0 for node in range(1, n+1)}
    for node in range(1, n+1):
        explore(node, edges, nodes_order, visited)

    # print(nodes_order)
    return sorted(nodes_order, key=lambda x: nodes_order[x], reverse=True)


def explore(node, edges, nodes_order, visited):
    global clock
    # print(visited, node, neighbor_nodes, nodes_order, clock)
    
    if visited[node]:
        return

    neighbor_nodes = edges[node]
    if not neighbor_nodes:
        visited[node] = True
        return

    for neighbor_node in neighbor_nodes:
        explore(neighbor_node, edges, nodes_order, visited)
    
    # post order
    clock += 1
    nodes_order[node] = clock
    visited[node] = True


if __name__ == '__main__':
    # input 1 is the number of vertices, input 2 is the number of edges
    n, n_edges = map(int, input().split(' '))
    edges = [[] for _ in range(n+1)]
    visited = [False for _ in range(n+1)]  # use hash map to speed up checking visited
    for _ in range(n_edges):
        start, end = map(int, input().split(' '))
        edges[start].append(end)

    res = topological_sort(n, edges, visited)
    print(' '.join([str(_) for _ in res]))


# ref: https://github.com/Sonia-96/Coursera-Data_Structures_and_Algorithms/blob/master/3-Algorithms%20on%20Graphs/Week2-Directed%20Graphs/2-topological%20sort.py
