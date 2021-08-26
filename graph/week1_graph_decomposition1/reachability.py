def reachability(node1, node2, edges, visited_edges):
    nodes = edges.get(node1)
    if not nodes:
        return 0
    res = 0
    for node in nodes:
        visited_edge = (node1, node)
        if visited_edge in visited_edges:
            continue
        visited_edges.append(visited_edge)
        if node == node2:
            res += 1
        else:
            res += reachability(node, node2, edges, visited_edges)
    return res


if __name__ == '__main__':
    # input 1 is the number of vertices, input 2 is the number of edges
    input_ = [int(i) for i in input().strip().split(' ')]
    edges = {}
    for _ in range(input_[1]):
        input_edge = [int(i) for i in input().strip().split(' ')]
        if edges.get(input_edge[0]):
            edges[input_edge[0]].append(input_edge[1])
        else:
            edges[input_edge[0]] = [input_edge[1]]
        
        if edges.get(input_edge[1]):
            edges[input_edge[1]].append(input_edge[0])
        else:
            edges[input_edge[1]] = [input_edge[0]]
    nodes = [int(i) for i in input().strip().split(' ')]
    res = reachability(nodes[0], nodes[1], edges, [])
    if res >= 1:
        print(1)
    else:
        print(0)
