package de.oflebbe;

import lombok.Cleanup;

import java.io.FileReader;
import java.io.IOException;

public class Main {
    static public void main( String[] argv) {
        long start = System.currentTimeMillis();
        if (argv.length < 1 ) {
            System.err.println("Bitte Eingabedatei angeben");
            System.exit(1);
        }
        try {
            @Cleanup FileReader file =  new FileReader(argv[0]);
            Polymer one = new Polymer(file);

            System.out.printf("Length = %d\n", one.react());
            System.out.printf("Min Length = %d\n", one.whichChars());
        }
        catch (IOException e) {
            e.printStackTrace();
        }
        long stop = System.currentTimeMillis();
        System.out.printf("Timing %d\n", (stop-start));
    }
}
