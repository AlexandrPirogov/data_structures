import java.util.*;

class Vertex {
    public int Value;
    public boolean hit;

    public Vertex(int val) {
        Value = val;
        hit = false;
    }
}

class SimpleGraph {
    Vertex[] vertex;
    int[][] m_adjacency;
    int max_vertex;
    int lastDeleted;

    public SimpleGraph(int size) {
        max_vertex = size;
        m_adjacency = new int[size][size];
        vertex = new Vertex[size];
        lastDeleted = 0;
    }

    public void AddVertex(int value) {
        Vertex add = new Vertex(value);
        if (lastDeleted == -1) {
            return;
        }

        vertex[lastDeleted] = add;
        nextFree();
    }

    public void RemoveVertex(int v) {
        if (v < 0 || v > max_vertex - 1) {
            return;
        }
        boolean found = false;
        int index = -1;
        for (int i = 0; i < vertex.length && !found; i++) {
            if (vertex[i] != null && vertex[i].Value == v) {
                found = true;
                vertex[i] = null;
                index = i;
                nextFree();
            }
        }

        if (!found) {
            return;
        }

        for (int i = 0; i < vertex.length; i++) {
            m_adjacency[i][index] = 0;
            m_adjacency[index][i] = 0;
        }

    }

    public boolean IsEdge(int v1, int v2) {
        return m_adjacency[v1][v2] == 1;
    }

    public void AddEdge(int v1, int v2) {
        Vertex first = vertex[v1];
        Vertex second = vertex[v2];
        if (first == null || second == null) {
            return;
        }

        m_adjacency[v1][v2] = 1;
        m_adjacency[v2][v1] = 1;
    }

    public void RemoveEdge(int v1, int v2) {
        Vertex first = vertex[v1];
        Vertex second = vertex[v2];
        if (first == null || second == null) {
            return;
        }

        m_adjacency[v1][v2] = 0;
        m_adjacency[v2][v1] = 0;
    }

    private void nextFree() {
        if (lastDeleted >= vertex.length) {
            lastDeleted = -1;
            return;
        }

        if (lastDeleted < vertex.length - 2 && vertex[lastDeleted + 1] == null) {
            lastDeleted++;
            return;
        }

        boolean found = false;
        lastDeleted = -1;
        for (int i = 0; i < vertex.length && !found; i++) {
            if (vertex[i] == null) {
                lastDeleted = i;
                found = true;
            }
        }
    }

    public ArrayList<Vertex> DepthFirstSearch(int VFrom, int VTo) {
        for (Vertex v : vertex) {
            v.hit = false;
        }

        ArrayDeque<Integer> stack = new ArrayDeque<Integer>();
        boolean found = false;
        stack.add(VFrom);
        vertex[VFrom].hit = true;
        while (!stack.isEmpty() && !found) {
            Integer curr = stack.peekLast();
            if (vertex[curr] == vertex[VTo]) {
                found = true;
                continue;
            }

            int adj = 0;
            for (int i = 0; i < vertex.length && adj == 0; i++) {
                if (!vertex[i].hit && m_adjacency[curr][i] == 1) {
                    vertex[i].hit = true;
                    stack.addLast(i);
                    adj++;
                    break;
                }
            }
            if (adj == 0) {
                stack.removeLast();

            }
        }

        ArrayList<Vertex> res = new ArrayList<Vertex>();
        System.out.printf("From node %d to node %d\n", vertex[VFrom].Value, vertex[VTo].Value);
        for (Integer index : stack) {
            System.out.printf("Node %d\n", vertex[index].Value);
            res.add(vertex[index]);
        }
        return res;
    }

    public ArrayList<Vertex> BreadthFirstSearch(int VFrom, int VTo) {
        for (Vertex v : vertex) {
            v.hit = false;
        }

        int[] prev = new int[vertex.length*2];
        int cursor = 0;

        ArrayList<Integer> queue = new ArrayList<Integer>();
        boolean found = false;
        queue.add(VFrom);
        vertex[VFrom].hit = true;
        while (!queue.isEmpty() && !found) {
            Integer curr = queue.remove(0);

            if (vertex[curr] == vertex[VTo]) {
                found = true;
                break;
            }

            int adj = 0;
            for (int i = 0; i < vertex.length; i++) {
                if (!vertex[i].hit && m_adjacency[curr][i] == 1) {
                    vertex[i].hit = true;
                    queue.add(i);
                    adj++;
                     prev[cursor] = curr;
                    if (vertex[i] == vertex[VTo]) {
                        found = true;
                        prev[cursor + 1] = VTo;
                        break;
                    }
                }
            }
            if (adj > 0) {
                cursor++;

            }
        }

        ArrayList<Vertex> res = new ArrayList<Vertex>();
        if (queue.isEmpty()){
            return res;
        }

        System.out.printf("From node %d to node %d\n", vertex[VFrom].Value, vertex[VTo].Value);
        for (int i = 0 ; i <= cursor; i++) {
            System.out.printf("Node %d\n", vertex[prev[i]].Value);
             res.add(vertex[prev[i]]);
        }

        return res;
    }
}