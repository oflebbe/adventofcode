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
            Grid g = new Grid( file);
            System.out.printf("Result 1 = %d\n", g.init_field());
            System.out.printf("Result 2 = %d\n", g.countWithDistance( 10000));
        }
        catch (IOException e) {
            e.printStackTrace();
        }
        long stop = System.currentTimeMillis();
        System.out.printf("Timing %d\n", (stop-start));
    }
}
