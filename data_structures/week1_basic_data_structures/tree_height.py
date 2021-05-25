from collections import defaultdict


def get_tree_height(node_list, nodes_dict):
    height = 1

    # get_next_level_nodes
    while node_list:
        sub_node_list = []
        for node in node_list:
            sub_node_list.extend(nodes_dict[node])
        node_list = sub_node_list
        height += 1

    return height


if __name__ == "__main__":
    n = input()
    elements = [int(_) for _ in input().split(' ')]
    nodes_dict = defaultdict(list)
    for i in range(len(elements)):
        if elements[i] == -1:
            root = i
        else:
            nodes_dict[elements[i]].append(i)

    print(get_tree_height(nodes_dict[root], nodes_dict))
