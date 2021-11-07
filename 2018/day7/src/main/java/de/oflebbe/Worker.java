package de.oflebbe;

import java.util.Set;

public class Worker {
    final private String what;
    private int counter;
    final RemoveInterface ri;


    Worker(String st, int addToCodepoint,  RemoveInterface ri) {
        this.counter = st.codePointAt( 0) - "A".codePointAt(0) + 1 + addToCodepoint;
        this.ri = ri;
        what = st;
    }
    void tick() {
        counter--;
        if (counter == 0) {
            ri.operation();
        }
    }

    boolean avail() {
        return counter <= 0;
    }

    @Override
    public String toString() {
        if (avail()) {
            return ".";
        }
        return what;
    }
}
