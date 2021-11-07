package de.oflebbe;

import java.io.*;

public class Counter {
    final BufferedReader in;

    Counter(Reader in) {
        this.in = new BufferedReader( in);
    }

    int sumItAll() {
        String line;
        int sum = 0;
        try {
            line = in.readLine();
            while (line != null) {
                int val = Integer.parseInt(line);
                sum += val;
                line = in.readLine();
            }
        }
        catch (IOException e) {
            e.printStackTrace();
        }
        return sum;
    }
}
