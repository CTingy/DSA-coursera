def find_brackets(text):

    brackets = {"(": ")", "[": "]", "{": "}"}
    back_brackets = brackets.values()
    stack = []
    positions = []
    for i in range(len(text)):
        character = text[i]
        if character in brackets.keys():
            stack.append(brackets[character])
            positions.append(i + 1)
        elif character in back_brackets:
            if not stack or stack.pop() != character:
                return i + 1
            positions.pop()
    
    if len(stack) == 0:
        return 0
    return positions.pop()


if __name__ == '__main__':
    text = input()
    index = find_brackets(text)
    if index == 0:
        print("Success")
    else:
        print(index)
