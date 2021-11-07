package de.oflebbe;

import lombok.Cleanup;
import lombok.val;
import lombok.var;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.Reader;
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

public class Intersecter {
    final List<Claim> claims;
    final int[][] fabric;
    int w, h;
    Intersecter(Reader r) {
        claims = new ArrayList<>();
        try {
            @Cleanup BufferedReader bufReader = new BufferedReader(r );

            var line = bufReader.readLine();
            while (line != null) {
                claims.add(Claim.parseString(line));
                line = bufReader.readLine();
            }
        }
        catch (IOException e) {
            e.printStackTrace();
        }
        w = 0;
        h = 0;
        for (Claim c : claims) {
            w = Math.max( w, c.W());
            h = Math.max( h, c.H());
        }

        fabric = new int[w][h];
        for (Claim c : claims) {
            for (int x = c.left; x < c.W(); x++) {
                for (int y = c.top; y < c.H(); y++) {
                    fabric[x][y]++;
                }
            }
        }
    }

    int processOne() {
        var count = 0;
        for (int x = 0; x < w; x++) {
            for (int y = 0; y < h; y++) {
                if (fabric[x][y] > 1) {
                    count++;
                }
            }
        }
        return count;
    }

    List<Integer> processTwo() {
        boolean single = true;
        var singleClaims = new LinkedList<Integer>();
        for (Claim c : claims) {
            single = true;
            for (int x = c.left; x < c.W(); x++) {
                if (!single)
                    break;
                for (int y = c.top; y < c.H(); y++) {
                    if (fabric[x][y] > 1) {
                        single = false;
                    }
                }
            }
            if (single) {
                singleClaims.add( c.number);
            }
        }
        return singleClaims;
    }

}
