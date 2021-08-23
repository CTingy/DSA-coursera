# find the number of cycles

def explore(start_node, edges, visited_edges, dfs_visited_nodes):
    nodes = edges.get(start_node)
    if not nodes:
        return 0
    if start_node in dfs_visited_nodes:
        print(start_node, dfs_visited_nodes, visited_edges)
        return 1
    dfs_visited_nodes.append(start_node)
    cycles = 0
    for node in nodes:
        visited_edge = (start_node, node)
        if visited_edge in visited_edges:
            continue
        visited_edges.append(visited_edge)
        cycles += explore(node, edges, visited_edges, dfs_visited_nodes)
    return cycles


if __name__ == '__main__':
    # input 1 is the number of vertices, input 2 is the number of edges
    input_ = [int(i) for i in input().strip().split(' ')]
    edges = {}
    nodes = set()
    for _ in range(input_[1]):
        input_edge = [int(i) for i in input().strip().split(' ')]
        if edges.get(input_edge[0]):
            edges[input_edge[0]].append(input_edge[1])
        else:
            edges[input_edge[0]] = [input_edge[1]]
        nodes.add(input_edge[0])

    print(explore(next(iter(edges)), edges, [], []))


"""
test case: output should be: 5, but got 6

5 10             
4 1
5 4
4 5
3 1
5 2
1 4
2 3
4 2
3 4
3 5
"""
