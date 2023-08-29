import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotEquals;
import static org.junit.Assert.assertNull;
import static org.junit.Assert.assertTrue;

import java.util.ArrayList;

import org.junit.Test;

public class SimpleGraph_test {

    @Test
    public void TestCreate() {
        SimpleGraph sut = new SimpleGraph(10);

        assertTrue(10 == sut.vertex.length);
    }

    @Test
    public void TestAddVertexInEmptyGraph() {
        SimpleGraph sut = new SimpleGraph(10);
        sut.AddVertex(0);

        assertEquals(sut.vertex[0].Value, 0);
    }

    @Test
    public void TestAddTwiceVertexInEmptyGraph() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        for (int i = 0; i < size; i++) {
            sut.AddVertex(i);
        }

        for (int i = 0; i < size; i++) {
            assertEquals(sut.vertex[i].Value, i);
        }

    }

    @Test
    public void TestAddInFullfilled() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        for (int i = 0; i < size; i++) {
            sut.AddVertex(i);
        }

        sut.AddVertex(11);
        for (int i = 0; i < size; i++) {
            assertEquals(sut.vertex[i].Value, i);
        }

    }

    @Test
    public void TestDeleteFromEmpty() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        sut.RemoveVertex(0);

        assertNull(sut.vertex[0]);

    }

    @Test
    public void TestDeleteSingleExistingFromEmpty() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        sut.AddVertex(0);
        sut.RemoveVertex(0);

        assertNull(sut.vertex[0]);

    }

    @Test
    public void TestDeleteAllFromFullfilled() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        for (int i = 0; i < size; i++) {
            sut.AddVertex(i);
        }

        for (int i = 0; i < size; i++) {
            sut.RemoveVertex(i);
        }

        for (int i = 0; i < size; i++) {
            assertNull(sut.vertex[i]);
        }
    }

    @Test
    public void TestAddEdgeInEmpty() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        sut.AddVertex(0);
        sut.AddVertex(1);
        sut.AddEdge(0, 1);

        assertEquals(sut.m_adjacency[0][1], 1);
    }

    @Test
    public void TestAddEdgeInLoopInEmpty() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        sut.AddVertex(0);
        sut.AddEdge(0, 0);

        assertEquals(sut.m_adjacency[0][0], 1);
    }

    @Test
    public void TestAddEdgeInNonexistingVertexInEmpty() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        sut.AddVertex(0);
        sut.AddEdge(0, 1);

        assertEquals(sut.m_adjacency[0][1], 0);
    }

    @Test
    public void TestRemoveVertexWithEdge() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        sut.AddVertex(0);
        sut.AddVertex(1);
        sut.AddEdge(0, 1);

        sut.RemoveVertex(1);

        assertEquals(0, sut.m_adjacency[0][1]);
        assertEquals(0, sut.m_adjacency[1][0]);
    }

    @Test
    public void TestRemoveNotExistingVertex() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        sut.AddVertex(0);
        sut.AddVertex(1);
        sut.AddEdge(0, 1);

        sut.RemoveVertex(2);

        assertEquals(1, sut.m_adjacency[0][1]);
        assertEquals(1, sut.m_adjacency[1][0]);
    }

    @Test
    public void TestRemoveVertexWithoutEdge() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        sut.AddVertex(0);
        sut.AddVertex(1);
        sut.AddVertex(2);
        sut.AddEdge(1, 2);

        sut.RemoveVertex(0);

        assertEquals(0, sut.m_adjacency[0][0]);
        assertEquals(1, sut.m_adjacency[1][2]);
        assertEquals(1, sut.m_adjacency[2][1]);
    }

    @Test
    public void TestRemoveEdgeInEmptyGraph() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        sut.RemoveEdge(0, 1);

        assertEquals(0, sut.m_adjacency[0][1]);
    }

    @Test
    public void TestRemoveExistingEdgeInGraph() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        for (int i = 0; i < size; i++) {
            sut.AddVertex(i);
        }

        for (int i = 0; i < size - 1; i++) {
            sut.AddEdge(i, i + 1);
        }

        for (int i = 0; i < size; i++) {
            sut.RemoveEdge(i, size - 1 - i);
            ;
        }

        for (int i = 0; i < size; i++) {
            assertEquals(sut.m_adjacency[i][size - 1 - i], 0);
        }
    }

    @Test
    public void TestDFSExisting() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        for (int i = 0; i < size; i++) {
            sut.AddVertex(i);
        }

        sut.AddEdge(0, 1);
        sut.AddEdge(1, 5);
        sut.AddEdge(1, 3);
        sut.AddEdge(3, 6);
        sut.AddEdge(6, 4);
        sut.AddEdge(4, 5);
        sut.AddEdge(5, 9);

        ArrayList<Vertex> res = sut.DepthFirstSearch(0, 9);
        assertEquals(6, res.size());
    }

    @Test
    public void TestDFSExisting_2() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        for (int i = 0; i < size; i++) {
            sut.AddVertex(i);
        }

        sut.AddEdge(0, 2);
        sut.AddEdge(2, 4);
        sut.AddEdge(2, 5);
        sut.AddEdge(2, 7);

        sut.AddEdge(7, 1);

        ArrayList<Vertex> res = sut.DepthFirstSearch(0, 1);
        assertNotEquals(0, res.size());
    }

     @Test
    public void TestDFSExisting_1() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        for (int i = 0; i < size; i++) {
            sut.AddVertex(i);
        }

        for (int i = 0; i < size-1; i++) {
            sut.AddEdge(i, i+1);
        }

        ArrayList<Vertex> res = sut.DepthFirstSearch(0, 9);
        assertEquals(10, res.size());
        assertEquals(0, res.get(0).Value);
        assertEquals(9, res.get(9).Value);
    }

    @Test
    public void TestDFSNotExisting() {
        int size = 10;
        SimpleGraph sut = new SimpleGraph(size);
        for (int i = 0; i < size; i++) {
            sut.AddVertex(i);
        }

        sut.AddEdge(0, 1);
        sut.AddEdge(1, 5);
        sut.AddEdge(1, 3);
        sut.AddEdge(3, 6);
        sut.AddEdge(6, 4);
        sut.AddEdge(4, 5);
        sut.AddEdge(5, 2);

        ArrayList<Vertex> res = sut.DepthFirstSearch(0, 9);
        assertTrue(res.size() == 0);
    }
}
