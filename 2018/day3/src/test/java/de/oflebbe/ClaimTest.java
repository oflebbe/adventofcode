package de.oflebbe;

import org.junit.jupiter.api.Test;

import java.io.StringReader;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class ClaimTest {

    @Test
    void testClaim() {
        String line = "#31 @ 875,804: 20x16";
        Claim parsed = Claim.parseString( line);
        Claim set = new Claim(31, 875, 804, 20, 16);
        assertEquals( set, parsed);
    }
}
