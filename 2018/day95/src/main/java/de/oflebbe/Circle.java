package de.oflebbe;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

public class Circle {
    private List<Integer> circle;
    private List<Integer> score ;
    private Marble marble = new Marble();
    private  int current;
    private int elfs;
    private int max;


    Circle(int elfs, int max) {
        this.elfs = elfs;
        this.max = max;
        circle = new LinkedList<>();
        circle.add( 0);
        score = new ArrayList<>( elfs);

        for (int i = 0; i < elfs; i++) {
            score.add( 0);
        }
        current = 0;
    }

    int timestep() {
        for (;;) {
            int m = marble.next();
            if (m % 71790 == 0) {
                System.out.println(m / 71790);
            }
            if (m % 23 == 0) {
                current = (current - 7 + circle.size()) % circle.size();

                int s = score.get(m % elfs);
                s += m + circle.remove(current);
                score.set(m % elfs, s);
            } else {
                current = (current + 1) % circle.size();
                circle.add(current+1, m);
                current++;
            }
            if ( m == max) {
                int m2 = 0;
                for (int i : score) {
                    m2 = Math.max( m2, i);
                }
                return m2;
            }
        }

    }
}
