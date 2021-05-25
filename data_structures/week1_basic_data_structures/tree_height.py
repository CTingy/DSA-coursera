from collections import defaultdict


def get_tree_height(sub_nodes, nodes):
    if not sub_nodes:
        return 1
    max_height = 0
    for node in sub_nodes:
        height = get_tree_height(nodes[node], nodes)
        if height > max_height:
            max_height = height
    return max_height + 1


if __name__ == "__main__":
    n = input()
    elements = [int(_) for _ in input().split(' ')]
    nodes = defaultdict(list)
    for i in range(len(elements)):
        if elements[i] == -1:
            root = i
        else:
            nodes[elements[i]].append(i)

    print(get_tree_height(nodes[root], nodes))
