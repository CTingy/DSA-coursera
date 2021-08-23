def reachability(node1, node2, edges, visited_edges):
    nodes = edges.get(node1)
    if not nodes:
        return 0
    res = 0
    for node in nodes:
        visited_edge1, visited_edge2 = (node1, node), (node2, node)
        if visited_edge1 in visited_edges or visited_edge2 in visited_edges:
            continue
        visited_edges.append(visited_edge1)
        visited_edges.append(visited_edge2)
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
    print(reachability(nodes[0], nodes[1], edges, []))
