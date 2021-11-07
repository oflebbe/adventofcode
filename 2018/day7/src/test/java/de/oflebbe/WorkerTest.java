package de.oflebbe;

import org.junit.jupiter.api.Test;

import java.io.StringReader;
import java.util.HashSet;
import java.util.Set;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class WorkerTest {
    Set<String> toRemove = new HashSet<>();

    void func( String what) {
        toRemove.remove( what);
    }

    @Test
    void testWorker() {
        toRemove.add("A");
        toRemove.add("B");


        Worker w = new Worker( "A", 10, ()->func("A"));
        for (int j = 0; j < 10; j++) {
            assertEquals(false, w.avail());
            assertEquals( 2, toRemove.size());
            w.tick();
        }
        assertEquals(false, w.avail());
        assertEquals( 2, toRemove.size());
        w.tick();
        assertEquals(true, w.avail());
        assertEquals( 1, toRemove.size());
        assertEquals( true, toRemove.contains("B"));

    }

}
