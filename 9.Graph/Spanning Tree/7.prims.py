class Node :

    def __init__(self, arg_id) :
        self._id = arg_id

class Graph :

    def __init__(self, arg_source, arg_adj_list) :
        self.source = arg_source
        self.adjlist = arg_adj_list

    def PrimsMST(self):

        # Priority queue is implemented as a dictionary with
        # key as an object of 'Node' class and value as the cost of 
        # reaching the node from the source.
        # Since the priority queue can have multiple entries for the
        # same adjacent node but a different cost, we have to use objects as
        # keys so that they can be stored in a dictionary. 
        # [As dictionary can't have duplicate keys so objectify the key]

        # The distance of source node from itself is 0. Add source node as the first node
        # in the priority queue
        priority_queue = { Node(self.source) : 0 }
        added = [False] * len(self.adjlist)
        min_span_tree_cost = 0

        while priority_queue :
            # Choose the adjacent node with the least edge cost
            node = min (priority_queue, key=priority_queue.get)
            cost = priority_queue[node]

            # Remove the node from the priority queue
            del priority_queue[node]

            if added[node._id] == False :
                min_span_tree_cost += cost
                added[node._id] = True
                print("Added Node : " + str(node._id) + ", cost now : "+str(min_span_tree_cost))


                for item in self.adjlist[node._id] :
                    adjnode = item[0]
                    adjcost = item[1]
                    if added[adjnode] == False :
                        priority_queue[Node(adjnode)] = adjcost

        return min_span_tree_cost

def main() :

    g1_edges_from_node = {}

    # Outgoing edges from the node: (adjacent_node, cost) in graph 1.
    g1_edges_from_node[0] = [ (1,1), (2,2), (3,1), (4,1), (5,2), (6,1) ]
    g1_edges_from_node[1] = [ (0,1), (2,2), (6,2) ]
    g1_edges_from_node[2] = [ (0,2), (1,2), (3,1) ]
    g1_edges_from_node[3] = [ (0,1), (2,1), (4,2) ]
    g1_edges_from_node[4] = [ (0,1), (3,2), (5,2) ]
    g1_edges_from_node[5] = [ (0,2), (4,2), (6,1) ]
    g1_edges_from_node[6] = [ (0,1), (2,2), (5,1) ]

    g1 = Graph(0, g1_edges_from_node)
    cost = g1.PrimsMST()
    print("Cost of the minimum spanning tree in graph 1 : " + str(cost) +"\n")

    # Outgoing edges from the node: (adjacent_node, cost) in graph 2.
    g2_edges_from_node = {}
    g2_edges_from_node[0] = [ (1,1), (2,1) ]
    g2_edges_from_node[1] = [ (0,1), (2,2) ]
    g2_edges_from_node[2] = [ (0,1), (1,2), (3,3) ]
    g2_edges_from_node[3] = [ (4,2), (2,3), (5,1) ]
    g2_edges_from_node[4] = [ (3,2), (5,1) ]
    g2_edges_from_node[5] = [ (4,1), (3,1) ]

    g2 = Graph(0, g2_edges_from_node)
    cost = g2.PrimsMST()
    print("Cost of the minimum spanning tree in graph 2 : " + str(cost))

if __name__ == "__main__" :
    main()
