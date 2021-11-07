package de.oflebbe;

import org.junit.jupiter.api.Test;

import java.io.StringReader;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class IntersecterTest {

    @Test
    void testIntersector() {
        String input = "#1 @ 1,3: 4x4\n" +
                "#2 @ 3,1: 4x4\n" +
                "#3 @ 5,5: 2x2\n";

        Intersecter intersect = new Intersecter( new StringReader( input));
        assertEquals(intersect.processOne(), 4);
    }

    @Test
    void testIntersectorTwo() {
        String input = "#1 @ 1,3: 4x4\n" +
                "#2 @ 3,1: 4x4\n" +
                "#3 @ 5,5: 2x2\n";

        Intersecter intersect = new Intersecter( new StringReader( input));
        List<Integer> result = intersect.processTwo();
        assertEquals( result.size(), 1);
        assertEquals( (int) result.get(0), 3);
    }
}
