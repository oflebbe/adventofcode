package de.oflebbe;

import org.junit.jupiter.api.Test;

import java.io.StringReader;

import static org.junit.jupiter.api.Assertions.*;


public class DeviceTest {

    @Test
    void testSecond() {
        String nix = "+3\n+3\n+4\n-2\n-4";

        Device zero = new Device(() -> new StringReader(nix));
        assertEquals(   10, zero.findTrim());

    }

    @Test
    void testThird() {
        String nix = "-6\n+3\n+8\n+5\n-6";

        Device zero = new Device(() -> new StringReader(nix));
        assertEquals( 5, zero.findTrim());
    }

    @Test
    void testFourth() {
        String nix = "+7\n+7\n-2\n-7\n-4";

        Device zero = new Device(() -> new StringReader(nix));
        assertEquals( 14, zero.findTrim());
    }

}
