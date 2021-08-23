# find the number of cycles
# ref: https://www.quora.com/How-do-I-count-the-number-of-cycles-in-a-directed-graph


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

        if self.edges.get(node):
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

    graph = Graph(edges, nodes)
    count = graph.count_cycle()
    if count == 0:
        print(0)
    else:
        print(1)
