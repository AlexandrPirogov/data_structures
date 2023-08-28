import java.util.*;

class Vertex {
    public int Value;

    public Vertex(int val) {
        Value = val;
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

    // здесь и далее, параметры v -- индекс вершины
    // в списке vertex
    public void RemoveVertex(int v) {
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

        if (!found){
            return;
        }

        for (int i = 0; i < vertex.length; i++) {
            m_adjacency[i][index] = 0;
            m_adjacency[index][i] = 0;
        }

    }

    public boolean IsEdge(int v1, int v2) {
        // true если есть ребро между вершинами v1 и v2
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
        // удаление ребра между вершинами v1 и v2
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
}