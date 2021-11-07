package de.oflebbe;

import org.junit.jupiter.api.Test;

import java.io.StringReader;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class GridTest {

    @Test
    void testGrid() {
        String input = "1, 1\n" +
                "1, 6\n" +
                "8, 3\n" +
                "3, 4\n" +
                "5, 5\n" +
                "8, 9";

        Grid g = new Grid( new StringReader( input));
        assertEquals( g.init_field(),17);
        assertEquals( g.countWithDistance( 32), 16);

    }

}
