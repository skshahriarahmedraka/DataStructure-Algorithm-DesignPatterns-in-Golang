1. Spanning Tree
   A spanning tree of an undirected graph G is a connected subgraph that covers all the graph nodes with the minimum possible number of edges. In general, a graph may have more than one spanning tree.
   
2. Minimum Spanning Tree
   If the graph is edge-weighted, we can define the weight of a spanning tree as the sum of the weights of all its edges. A minimum spanning tree is a spanning tree whose weight is the smallest among all possible spanning trees.
   
3. Shortest Path Tree
   In the shortest path tree problem, we start with a source node s.
   For any other node v in graph G, the shortest path between s and v is a path such that the total weight of the edges along this path is minimized. Therefore, the objective of the shortest path tree problem is to find a spanning tree such that the path from the source node s to any other node v is the shortest one in G.
   We can solve this problem with Dijkstra’s algorithm:

   |E|C(|V|-1) -Number of cycle 
   
   dijkstra may fail if there is any negetive weight of an edge 
   bellmanford will fail if weight o


