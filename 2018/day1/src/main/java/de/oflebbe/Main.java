package de.oflebbe;

import java.io.FileReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.MalformedURLException;
import java.net.URL;

public class Main {
    static public void main( String[] argv) {
        if (argv.length < 1 ) {
            System.err.println("Bitte Eingabedatei angeben");
            System.exit(1);
        }
        try {
            FileReader input = new FileReader( argv[0]);

            Counter one = new Counter( input);
            System.out.printf( "Sum is %d\n", one.sumItAll());

        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
