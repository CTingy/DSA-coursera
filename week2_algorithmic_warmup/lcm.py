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
    if num_gcd == 1:
        print(input_list[0] * input_list[1])
    else:
        print(input_list[0] * (input_list[1]/num_gcd))
