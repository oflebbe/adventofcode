package de.oflebbe;

import org.junit.jupiter.api.Test;

import java.io.StringReader;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class PolymerTest {

    @Test
    void testIntersector() {
        String input = "dabAcCaCBAcCcaDA";

        Polymer intersect = new Polymer(new StringReader(input));
        assertEquals(10, intersect.react());
        assertEquals( 4, intersect.whichChars());
    }

}
