# pass all test cases

X = 263
PRIME_NUMBER = 1000000007


class Query:
    def __init__(self, input_string):
        s = input_string.split(' ')
        self.action = s[0]
        self.content = s[1]


class HashChain:
    def __init__(self, m):
        self.size = m
        self.chains = [[] for _ in range(m)]

    def add(self, value):
        key = self.H(value)
        if self.find(value):
            return
        self.chains[key] = [value] + self.chains[key][:]

    def delete(self, value):
        key = self.H(value)
        for i, ele in enumerate(self.chains[key]):
            if ele == value:
                self.chains[key].pop(i)
                return

    def find(self, value):
        if value in self.chains[self.H(value)]:
            return True
        return False

    def check(self, key):
        if self.chains[key]:
            print(' '.join(self.chains[key]))
        else:
            print('')

    def H(self, s):
        res = 0
        for i, char in enumerate(s):
            res = (res + ord(char) * (X**i)) % PRIME_NUMBER
        return res % self.size


def execute_query(hash_chain, query):
    if query.action == 'add':
        hash_chain.add(query.content)
    elif query.action == 'check':
        hash_chain.check(int(query.content))
    elif query.action == 'find':
        if hash_chain.find(query.content):
            print('yes')
        else:
            print('no')
    elif query.action == 'del':
        hash_chain.delete(query.content)


if __name__ == "__main__":
    m = int(input())
    number = int(input())

    queries = []
    for i in range(number):
        queries.append(Query(input()))

    hash_chain = HashChain(m)
    for query in queries:
        execute_query(hash_chain, query)
