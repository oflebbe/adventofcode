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
            SleepDetector one = new SleepDetector(file);

            one.processNotes();
            int most = one.findMostSleeping();
            int when = one.findMaxminute( most);
            System.out.printf("%d\n", most);
            System.out.printf("%d\n", when);
            System.out.printf("Result %d\n", most*when);

            System.out.printf("Result2 %d\n", one.findMostSleepSameMinute());
        }
        catch (IOException e) {
            e.printStackTrace();
        }
        long stop = System.currentTimeMillis();
        System.out.printf("Timing %d\n", (stop-start));
    }
}
