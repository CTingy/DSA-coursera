# pass

def max_pairwise_product(numbers):
    if len(numbers) < 2:
        return 0
    
    for i in range(1, len(numbers)):
        if numbers[i] > numbers[0]:
            numbers[i], numbers[0] = numbers[0], numbers[i]
    
    for i in range(2, len(numbers)):
        if numbers[i] > numbers[1]:
            numbers[i], numbers[1] = numbers[1], numbers[i]

    return numbers[0] * numbers[1]


if __name__ == '__main__':
    input_n = int(input())
    input_numbers = [int(x) for x in input().split()]
    print(max_pairwise_product(input_numbers))
