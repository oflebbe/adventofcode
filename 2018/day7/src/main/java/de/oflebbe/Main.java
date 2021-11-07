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
            Dependencies p = new Dependencies( file);
            // p.findFirst();

            //System.out.printf("Start = %s\n",p.ordering( ));
           // p = new Dependencies( file);
            System.out.printf("Seconds = %d\n",p.scheduler( 5, 60));

        }
        catch (Exception e) {
            e.printStackTrace();
        }
        long stop = System.currentTimeMillis();
        System.out.printf("Timing %d\n", (stop-start));
    }
}
