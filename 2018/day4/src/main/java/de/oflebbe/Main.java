package de.oflebbe;

import lombok.Cleanup;

import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.List;

public class Main {
    static public void main( String[] argv) {
        long start = System.currentTimeMillis();
        if (argv.length < 1 ) {
            System.err.println("Bitte Eingabedatei angeben");
            System.exit(1);
        }
        try {
            @Cleanup FileReader file =  new FileReader(argv[0]);
            Intersecter one = new Intersecter(file);

            System.out.printf("Overlapping = %s\n", one.processOne());
            List<Integer> res = one.processTwo();

            System.out.printf("Single %d\n", res.size());
            System.out.printf("Single Claim #%d\n", res.get(0));

        }
        catch (IOException e) {
            e.printStackTrace();
        }
        long stop = System.currentTimeMillis();
        System.out.printf("Timing %d\n", (stop-start));
    }
}
