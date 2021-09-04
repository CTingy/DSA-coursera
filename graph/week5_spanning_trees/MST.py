import math


class UnionSet:
    parents = {}
    rank = {}
    
    @classmethod
    def make_union_set(cls, node):
        cls.parents[node] = node
        cls.ranks[node] = 0

    @classmethod
    def find(cls, node):
        while node != cls.parents[node]:
            node = cls.find(cls.parents[node])
        return node
    
    @classmethod
    def merge(cls, node1, node2):
        if cls.ranks[node1] > cls.ranks[node2]:
            cls.parents[node2] = node1
        else:
            cls.parents[node1] = node2
            if cls.ranks[node1] == cls.ranks[node2]:
                cls.ranks[node2] += 1


class MySet:
    sets = []
    
    @classmethod
    def make_union_set(cls, node):
        cls.sets.append(set([node]))

    @classmethod
    def find(cls, node):
        # return set index
        for i in range(len(cls.sets)):
            if node in cls.sets[i]:
                return i

    @classmethod
    def merge(cls, node1, node2):
        node1_set_idx = cls.find(node1)
        node2_set_idx = cls.find(node2)
        
        cls.sets[node1_set_idx].update(cls.sets[node2_set_idx])
        cls.sets.pop(node2_set_idx)


def get_distance(x1, y1, x2, y2):
    return math.sqrt(math.pow(x1 - x2, 2) + math.pow(y1 - y2, 2))


def Kruskal(edges):
    sorted_edges = sorted(edges, key=lambda x: x[0])
    total_distance = 0
    for values in sorted_edges:
        distance, start_node, end_node = values[0], values[1], values[2]
        if MySet.find(start_node) == MySet.find(end_node):
            continue
        # locate in different tree
        MySet.merge(start_node, end_node)
        total_distance += distance
        # else: in same tree, do nothing
    return total_distance


if __name__ == '__main__':
    n = int(input())
    edges = []  # [[distance, node1(x1, y1), node2(x2, y2)], [], []]
    nodes = set()

    for i in range(n):
        x, y = map(int, input().split())
        MySet.make_union_set((x, y))
        if i != 0:
            for node in nodes:
                distance = get_distance(x, y, node[0], node[1])
                edges.append([distance, (x, y), node])
        nodes.add((x, y))
    
    parents_set = Kruskal(edges)
    print('{0:.9f}'.format(parents_set))
