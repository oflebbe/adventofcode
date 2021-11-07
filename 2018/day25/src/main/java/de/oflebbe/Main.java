package de.oflebbe;

import java.io.FileReader;
import java.io.IOException;

public class Main {
    static public void main( String[] argv) {
        if (argv.length < 1 ) {
            System.err.println("Bitte Eingabedatei angeben");
            System.exit(1);
        }
        DevicePartTwo one = new DevicePartTwo( () -> new FileReader( argv[0]));
        final String right = one.parcelCheck();
        System.out.printf("Checksum = %s", right);


    }
}
