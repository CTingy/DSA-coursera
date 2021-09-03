def find_the_lowest_cost_node(costs, processed):
    min_val = float('inf')
    min_idx = 0
    for i in range(len(costs)):
        if costs[i] < min_val and not processed[i]:
            min_val = costs[i]
            min_idx = i
    return min_idx


def Dijkstra(graph, start_node, end_node, n):
    processed = [False for _ in range(n)]
    parents = [0 for _ in range(n)]  # if you want to print he path, you need a parent map
    
    # init costs for start_node
    costs = []
    for i in range(n):
        costs.append(float('inf') if graph[start_node][i] == -1 else graph[start_node][i])
    
    node = find_the_lowest_cost_node(costs, processed)
    while node and not processed[node]:
        cost = costs[node]
        neighbors = graph[node]
        for sub_node in range(1, len(neighbors)):  # skip the first one
            if neighbors[sub_node] == -1:  # no path between node and sub_node
                continue
            new_cost = cost + neighbors[sub_node]
            if new_cost < costs[sub_node]:
                costs[sub_node] = new_cost
                parents[sub_node] = node
        processed[node] = True
        node = find_the_lowest_cost_node(costs, processed)
    return costs[end_node] if costs[end_node] != float('inf') else -1


if __name__ == '__main__':
    # input 1 is the number of vertices, input 2 is the number of edges
    n_vertices , n_graph = map(int, input().split())
    graph = [[-1 for _ in range(n_vertices+1)] for _ in range(n_vertices+1)]
    for _ in range(n_graph):
        start, end, weight = map(int, input().split())            
        graph[start][end] = weight
    
    start_node, end_node = map(int, input().split())
    print(Dijkstra(graph, start_node, end_node, n_vertices+1))
