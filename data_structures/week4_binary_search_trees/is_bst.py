# Failed

class Node:
    def __init__(self, key=None, left=None, right=None):
        self.key = key
        self.left = left
        self.right = right


def is_bst(root, lower_bound, higher_bound):
    if not root:
        return True
    
    if lower_bound is not None and root.key <= lower_bound:
        return False
    if higher_bound is not None and root.key > higher_bound:
        return False
    
    return is_bst(root.left, lower_bound, root.key) and is_bst(root.right, root.key, higher_bound)


if __name__ == '__main__':
    n = int(input())
    nodes = [Node() for _ in range(n)]
    for i in range(n):
        input_ = [int(i) for i in input().split(' ')]
        
        nodes[i].key = input_[0]
        nodes[i].left = None if input_[1] == -1 else nodes[input_[1]]
        nodes[i].right = None if input_[2] == -1 else nodes[input_[2]]
    
    root = nodes[0]
    res = is_bst(root, None, None)
    if res:
        print('CORRECT')
    else:
        print('INCORRECT')
