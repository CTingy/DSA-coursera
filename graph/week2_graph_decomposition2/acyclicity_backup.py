# find the number of cycles
# ref: https://www.quora.com/How-do-I-count-the-number-of-cycles-in-a-directed-graph


# def explore(start_node, edges, visited_edges, dfs_visited_nodes):
#     if start_node in dfs_visited_nodes:
#         return 1
#     dfs_visited_nodes.append(start_node)
#     nodes = edges.get(start_node)
#     if not nodes:
#         return 0
#     cycles = 0
#     for node in nodes:
#         visited_edge = (start_node, node)
#         if visited_edge in visited_edges:
#             continue
#         visited_edges.append(visited_edge)
#         cycles += explore(node, edges, visited_edges, dfs_visited_nodes)
#     return cycles


from collections import defaultdict


class Graph:
    def __init__(self, edges, nodes):
        self.edges = edges
        self.visited = defaultdict(lambda: 0)
        self.count = 0
        self.nodes = nodes
    
    def count_cycle(self):
        for node in self.nodes:
            self.DFS(node)
        return self.count

    def DFS(self, node):
        # if cycle is detected
        if self.visited[node] == -1:
            self.count += 1
            return
        
        # if DFS has been completed on this node
        if self.visited[node] == 1:
            return

        self.visited[node] = -1

        for neighbor_node in self.edges[node]:
            self.DFS(neighbor_node)
        
        # mark DFS on this node as completed
        self.visited[node] = 1


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

    # print(explore(next(iter(edges)), edges, [], []))

    graph = Graph(edges, nodes)
    print(graph.count_cycle())
