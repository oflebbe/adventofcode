package de.oflebbe;

import java.io.FileReader;
import java.io.IOException;

public class Main {
    static public void main( String[] argv) {
        if (argv.length < 1 ) {
            System.err.println("Bitte Eingabedatei angeben");
            System.exit(1);
        }
        Device one = new Device( () -> new FileReader( argv[0]));
        final int trim = one.findTrim();
        System.out.printf("Trim = %d", trim);


    }
}
