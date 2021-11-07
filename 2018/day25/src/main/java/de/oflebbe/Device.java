package de.oflebbe;

import java.io.*;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

public class Device {
    private ReaderInterface read;

    Device( ReaderInterface read) {
        this.read = read;
    }

    int parcelCheck() {
       // int maxTrimTry = 20;
        BufferedReader in = null;

        String line;
        int count2 = 0;
        int count3 = 0;

        try {
            in = new BufferedReader( read.operation());

            line = in.readLine();
            while (line != null) {
                Map<Character, Integer> counter = new HashMap<>();

                for (char ch : line.toCharArray()) {
                    counter.put(ch, counter.getOrDefault(ch, 0) + 1);
                }
                if (counter.containsValue(2)) {
                    count2++;
                }
                if (counter.containsValue(3)) {
                    count3++;
                }
                line = in.readLine();
            }
        }
        catch (IOException e) {
            e.printStackTrace();
        }
        return count2 * count3;
    }
}
