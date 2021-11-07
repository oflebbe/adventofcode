package de.oflebbe;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertTrue;

import java.io.StringReader;


public class CounterTest {

    @Test
    void testZero() {
        String nix = "";

        Counter zero = new Counter( new StringReader( nix));
        assertTrue( zero.sumItAll() == 0);
    }

    @Test
    void testMinusOne() {
        String nix = "-1";

        Counter zero = new Counter( new StringReader( nix));
        assertTrue( zero.sumItAll() == -1);
    }

    @Test
    void testPlusOne() {
        String nix = "+1";

        Counter zero = new Counter( new StringReader( nix));
        assertTrue( zero.sumItAll() == 1);
    }

    @Test
    void testPlusMinusTwo() {
        String nix = "+2\n-2\n";

        Counter zero = new Counter( new StringReader( nix));
        assertTrue( zero.sumItAll() == 0);
    }
}
