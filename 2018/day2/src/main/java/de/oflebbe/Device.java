package de.oflebbe;

import java.io.*;
import java.util.HashSet;
import java.util.Set;

public class Device {
    private ReaderInterface read;

    Device( ReaderInterface read) {
        this.read = read;
    }

    int findTrim() {
       // int maxTrimTry = 20;
        BufferedReader in = null;

        String line;
        int sum = 0;
        Set<Integer> alreadySeen = new HashSet<>();
        alreadySeen.add( 0);

        try {
            in = new BufferedReader( read.operation());

            for (;;) {
                line = in.readLine();
                while (line != null) {
                    int val = Integer.parseInt(line);
                    sum += val;
                    /*System.out.println(sum);
                    if (maxTrimTry-- < 0) {
                        return 999;
                    }*/
                    if (alreadySeen.contains(sum)) {
                        in.close();
                        return sum;
                    }
                    alreadySeen.add( sum);
                    line = in.readLine();
                }
                in.close();
                in = new BufferedReader( read.operation());
            }
        }
        catch (IOException e) {
            e.printStackTrace();
        }
        return 0;
    }
}
