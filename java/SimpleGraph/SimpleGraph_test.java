import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNull;
import static org.junit.Assert.assertTrue;

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
        sut.AddEdge(0,1);

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
        sut.AddEdge(1,2);

        sut.RemoveVertex(0);

        assertEquals(0, sut.m_adjacency[0][0]);
        assertEquals(1, sut.m_adjacency[1][2]);
        assertEquals(1, sut.m_adjacency[2][1]);
    }
}