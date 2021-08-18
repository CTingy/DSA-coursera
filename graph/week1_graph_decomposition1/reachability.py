def reachability(node1, node2, edges, res=0, visited=None):
    if not visited:
        visited = []
    visited.append(node1)
    nodes = edges[node1]
    for node in nodes:
        if node in visited:
            continue
        if node == node2:
            res += 1
        else:
            res += reachability(node, node2, edges, res, visited)
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
    print(reachability(nodes[0], nodes[1], edges))
