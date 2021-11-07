package de.oflebbe;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.Reader;
import java.util.*;

public class Dependencies {
    List<Pair<String, String>> depList = null;
    Set<String> ready;


    Dependencies(Reader r) {

        depList = new ArrayList<>();
        ready = new HashSet<>();
        Set<String> x = new HashSet<>();
        Set<String> y = new HashSet<>();

        try {
            BufferedReader b = new BufferedReader(r);
            String line = b.readLine();
            while (line != null) {
                String[] tok = line.split(" ");
                x.add(tok[1]);
                y.add(tok[7]);
                Pair<String, String> p = new Pair<>(tok[1], tok[7]);
                depList.add( p);
                line = b.readLine();
            }
            y.removeAll(x);
            for ( String s : y) {
                depList.add( new Pair<> (s, ""));
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }


    String ordering() throws  Exception{
        StringBuffer retBuffer = new StringBuffer();
        Set<String> allResults = new HashSet<>();
        for (Pair<String, String> p : depList) {
            allResults.add(p.y);
        }

        while (depList.size() > 0) {
            String ret = nextOption();

            retBuffer.append( ret);

            allResults.remove( ret);

            removeOption( ret);
        }
        if (allResults.size() != 1) {
            System.out.println("Mehr Übrig\n");
        }
        retBuffer.append( allResults.iterator().next());
        return retBuffer.toString();
    }


    String nextOption() throws Exception {
        Set<String> all = new HashSet<String>();

        for (Pair<String, String> p : depList) {
            all.add(p.x);
        }
        for (Pair<String, String> p : depList) {
            all.remove(p.y);
        }

        String[] arrayString = all.toArray(new String[all.size()]);
        if (arrayString.length == 0)
            return null;
        Arrays.sort(arrayString);

        return arrayString[0];
    }


    List<String> nextOptions( Set<String> alreadyStarted, Set<String> allResults) throws Exception {
        Set<String> all = new HashSet<String>();

        for (Pair<String, String> p : depList) {
            all.add(p.x);
        }
        for (Pair<String, String> p : depList) {
            all.remove(p.y);
        }
        for (String s : alreadyStarted) {
            all.remove( s);
        }

        String[] arrayString = all.toArray(new String[all.size()]);
        if (arrayString.length == 0)
            return new ArrayList<>();
        Arrays.sort(arrayString);

        List<String> ar = Arrays.asList( arrayString);
        List<String> ret = new ArrayList<String>();
        ret.addAll( ar);
        return ret;
    }

    void removeOption(String what) {
        Iterator<Pair<String, String>> ptr = depList.iterator();
        while (ptr.hasNext()) {
            Pair<String, String> e = ptr.next();
            if (e.x.equals(what)) {
                ptr.remove();
            }
        }
    }



    int scheduler(int workers, int offset) throws  Exception{
        StringBuffer retBuffer = new StringBuffer();
        Set<String> allResults = new HashSet<>();
        Set <String> alreadyStarted = new HashSet<>();
        for (Pair<String, String> p : depList) {
            allResults.add(p.y);
        }
        int counter = 0;

        Worker[] work = new Worker[workers];


        while (true) {
            boolean allAvail = true;
            counter++;
            for (Worker w : work) {
                if (w != null) {
                    w.tick();
                    allAvail = (allAvail && w.avail());
                }
            }

            List<String> options = nextOptions(alreadyStarted, allResults);
            if (allAvail && options.size() == 0) {
                break;
            }

            for (int i = 0; i < work.length; i++) {
                if (work[i] == null || work[i].avail()) {

                    if (options.size() == 0) {
                        break;
                    } else {
                        String what = options.get(0);

                        work[i] = new Worker(what, offset, () -> removeOption(what));
                        alreadyStarted.add( what);
                        options.remove(0);
                    }
                }
            }
            System.out.printf("%03d ", counter);
            for (Worker w : work) {
                if (w != null) {
                    System.out.printf("%s ", w.toString());
                   } else {
                    System.out.print(". ");
                }
            }
            System.out.println();

        }
        return counter -1 ;
    }
}
