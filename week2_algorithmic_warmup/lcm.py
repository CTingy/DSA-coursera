def gcd(smaller, bigger):
    while True:
        remainder = bigger % smaller
        if remainder == 0:
            return smaller
        bigger, smaller = smaller, remainder


if __name__ == '__main__':
    input_str = input()
    input_list = [int(i) for i in input_str.split(' ')]
    input_list.sort()
    num_gcd = gcd(input_list[0], input_list[1])
    print(int(input_list[0] * (input_list[1]/num_gcd)))
