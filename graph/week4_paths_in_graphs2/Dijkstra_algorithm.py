# best practice is to use priority queue

from collections import defaultdict


def find_the_lowest_cost_node(dist, processed):
    min_val = float('inf')
    min_node = None
    for k, v in dist.items():
        if v < min_val and k not in processed:
            min_val = v
            min_node = k
    return min_node


def Dijkstra(graph, start_node, end_node):
    # init
    dist = {}
    processed = set([start_node])
    for k, v in graph[start_node].items():
        dist[k] = v
    
    node = find_the_lowest_cost_node(dist, processed)
    while node:
        for n, weight in graph[node].items():
            total_dist = dist[node] + weight
            if not dist.get(n) or total_dist < dist[n]:
                dist[n] = total_dist
        processed.add(node)
        node = find_the_lowest_cost_node(dist, processed)
    
    if not dist or not dist.get(end_node):
        return -1
    return dist[end_node]


if __name__ == '__main__':
    # input 1 is the number of vertices, input 2 is the number of edges
    n_vertices , n_graph = map(int, input().split())
    
    # In the begining, I use array as a hash map, but when edges number is low, but n is large
    # use array would increase the time usage of array iteration
    graph = defaultdict(dict)
    for _ in range(n_graph):
        start, end, weight = map(int, input().split())
        graph[start][end] = weight
    
    start_node, end_node = map(int, input().split())
    print(Dijkstra(graph, start_node, end_node))
