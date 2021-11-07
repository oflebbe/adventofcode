package de.oflebbe;

import org.junit.jupiter.api.Test;

import java.io.StringReader;

import static org.junit.jupiter.api.Assertions.*;


public class DeviceTest {

    @Test
    void testSecond() {
        String nix = "abcde\n" +
                "fghij\n" +
                "klmno\n" +
                "pqrst\n" +
                "fguij\n" +
                "axcye\n" +
                "wvxyz";

        DevicePartTwo zero = new DevicePartTwo(() -> new StringReader(nix));
        assertEquals(  "fgij", zero.parcelCheck());

    }
}
