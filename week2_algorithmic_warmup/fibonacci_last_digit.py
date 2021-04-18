def fibonacci(n):
    if n <= 1:
        return n
    fibonacci_list = [0, 1]
    for i in range(2, n + 1):
        fibonacci_list.append((fibonacci_list[i-1] + fibonacci_list[i-2])%10)
    return fibonacci_list[-1]


if __name__ == '__main__':
    input_n = int(input())
    print(fibonacci(input_n))
