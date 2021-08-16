# Failed

class Node:
    def __init__(self, key=None, left=None, right=None):
        self.key = key
        self.left = left
        self.right = right


# def inorder_traversal(root):
#     if not root:
#         return []
#     res = []

#     res.extend(inorder_traversal(root.left))
#     res.append(root.key)
#     res.extend(inorder_traversal(root.right))
#     return res


def inorder_traversal(root):
    if not root:
        return []
    stack = []

    def inorder(node):
        if node:
            inorder(node.left)
            stack.append(node.key)
            inorder(node.right)
    inorder(root)
    return stack


# def preorder_traversal(root):
#     if not root:
#         return []
#     res = []
#     res.append(root.key)
#     res.extend(preorder_traversal(root.left))
#     res.extend(preorder_traversal(root.right))
#     return res


def preorder_traversal(root):
    if not root:
        return []
    stack = []
    def preorder(node):
        if node:
            stack.append(node.key)
            preorder(node.left)
            preorder(node.right)
    preorder(root)
    return stack


# def postorder_traversal(root):
#     if not root:
#         return []
#     res = []
#     res.extend(postorder_traversal(root.left))
#     res.extend(postorder_traversal(root.right))
#     res.append(root.key)
#     return res


def postorder_traversal(root):
    if not root:
        return []
    stack = []
    def postorder(node):
        if node:
            postorder(node.left)
            postorder(node.right)
            stack.append(node.key)
    postorder(root)
    return stack


if __name__ == '__main__':
    n = int(input())
    nodes = [Node() for _ in range(n)]
    for i in range(n):
        input_ = [int(i) for i in input().split(' ')]
        
        nodes[i].key = input_[0]
        nodes[i].left = None if input_[1] == -1 else nodes[input_[1]]
        nodes[i].right = None if input_[2] == -1 else nodes[input_[2]]
    
    inorder = inorder_traversal(nodes[0])
    preorder = preorder_traversal(nodes[0])
    postorder = postorder_traversal(nodes[0])

    print(' '.join(str(i) for i in inorder))
    print(' '.join(str(i) for i in preorder))
    print(' '.join(str(i) for i in postorder))
