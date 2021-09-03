# ref: http://web.ntnu.edu.tw/~algo/Path2.html

from collections import defaultdict


def detect_negative_cycles(graph, n_vertices):
    dist = {i : [0, 0] for i in range(1 + n_vertices)}
    max_relax_time = n_vertices - 1
    queue = list(dist.keys())
    while queue:
        node = queue.pop()
        for n, weight in graph[node].items():
            total_dist = dist[node][0] + weight
            if total_dist < dist[n][0]:
                dist[n][0] = total_dist
                dist[n][1] += 1
                queue.append(n)
                if dist[n][1] > max_relax_time:
                    return True
    return False


if __name__ == '__main__':
    # input 1 is the number of vertices, input 2 is the number of edges
    n_vertices , n_graph = map(int, input().split())
    
    # In the begining, I use array as a hash map, but when edges number is low, but n is large
    # use array would increase the time usage of array iteration
    graph = defaultdict(dict)
    
    for _ in range(n_graph):
        start, end, weight = map(int, input().split())
        graph[start][end] = weight
    
    is_negative_cycle_exists = detect_negative_cycles(graph, n_vertices)
    if is_negative_cycle_exists:
        print(1)
    else:
        print(0)


# procedure BellmanFord(list vertices, list edges, vertex source)
#    // 讀入邊和節點的列表並對distance和predecessor寫入最短路徑

#    // 初始化圖
#    for each vertex v in vertices:
#        if v is source then distance[v] := 0
#        else distance[v] := infinity
#        predecessor[v] := null

#    // 對每一條邊重複操作
#    for i from 1 to size(vertices)-1:
#        for each edge (u, v) with weight w in edges:
#            if distance[u] + w < distance[v]:
#                distance[v] := distance[u] + w
#                predecessor[v] := u

#    // 檢查是否有負權重的回路
#    for each edge (u, v) with weight w in edges:
#        if distance[u] + w < distance[v]:
#            error "圖包含負權重的回路"
