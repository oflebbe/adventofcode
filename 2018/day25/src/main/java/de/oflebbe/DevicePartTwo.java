package de.oflebbe;

import java.io.BufferedReader;
import java.io.IOException;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

public class DevicePartTwo {
    private ReaderInterface read;

    DevicePartTwo(ReaderInterface read) {
        this.read = read;
    }

    String parcelCheck() {
       // int maxTrimTry = 20;
        BufferedReader in = null;

        try {
            in = new BufferedReader( read.operation());

            String line = in.readLine();
            List<String> previous  =  new LinkedList<>();

            while (line != null) {
                for (String p : previous) {
                    String same = StringComparer.commonButOne(p , line);
                    if (!same.isEmpty()) {
                        return same;
                    }
                }
                previous.add( line);
                line = in.readLine();
            }
        }
        catch (IOException e) {
            e.printStackTrace();
        }
        return "";
    }
}
