package de.oflebbe;

import java.io.*;
import java.util.*;

import lombok.Cleanup;

public class Grid {
    ArrayList<Pair<Integer, Integer>> coords;

    Grid(Reader f) {
        coords = new ArrayList<>();

        try {
            @Cleanup BufferedReader b = new BufferedReader(f);
            String line = b.readLine();
            while (line != null) {
                String[] tok = line.split( ", ");
                if (tok.length != 2) {
                    System.err.println("more then 2 tokens");
                }
                int x = Integer.parseInt( tok[0]);
                int y = Integer.parseInt( tok[1]);
                coords.add(new Pair<>(x,y));

                line = b.readLine();
            }
        }
        catch (IOException e) {
            e.printStackTrace();
        }
    }


    int manhatten( Pair<Integer ,Integer> one, Pair<Integer ,Integer> two) {
        return Math.abs( one.x - two.x) + Math.abs( one.y - two.y);
    }


    int[][] field;
    int left;
    int right;
    int bottom;
    int top;

    int init_field() {
        int minx=Integer.MAX_VALUE, maxx=Integer.MIN_VALUE, miny=Integer.MAX_VALUE, maxy=Integer.MIN_VALUE;
        for (Pair<Integer, Integer> p : coords) {
            if (minx > p.x) {

            }

            minx = Math.min( p.x, minx);
            miny = Math.min( p.y, miny);
            maxx = Math.max( p.x, maxx);
            maxy = Math.max( p.y, maxy);
        }
        int w = maxx-minx;
        int h = maxy-miny;

        left = minx - w;
        right = maxx + w;
        bottom = miny - w;
        top = maxx + w;

        field = new int[right-left+1][top-bottom+1];


        for (int x = left; x <= right ; x++) {
            for (int y = bottom; y <= top ; y++) {
                int min = -1;
                int dist = Integer.MAX_VALUE;
                boolean draw = false;
                Pair<Integer, Integer> point = new Pair<Integer, Integer>(x, y);
                for (int p = 0; p < coords.size(); p++) {
                    int s = manhatten(coords.get(p), point);
                    if (s < dist) {
                        min = p;
                        draw = false;
                        dist = s;
                    }  else if (s == dist) {
                        draw = true;
                    }
                }
                if (draw) {
                    min = Integer.MIN_VALUE;
                }
                field[x - left][y - bottom] = min;
            }
        }

        // find val on edges

        Set<Integer> edge = new HashSet<>();
        for (int y = 0; y <= top - bottom; y++) {
            edge.add(field[0][y]);
            edge.add(field[right - left][y]);
        }
        for (int x = 0; x <= right-left; x++) {
            edge.add(field[x][0]);
            edge.add(field[x][top-bottom]);
        }

        Map<Integer, Integer> counter = new HashMap<>();
        for (int x = 0; x <= right-left; x++) {
            for (int y = 0; y <= top - bottom; y++) {
                counter.put(field[x][y], counter.getOrDefault(field[x][y], 0) + 1);
            }
        }

        int maxCounter = 0;
        for (int i = 0; i < coords.size(); i++) {
            if (edge.contains(i)) {
                continue;
            }
            maxCounter = Math.max(counter.getOrDefault(i, 0), maxCounter);
        }
        return maxCounter;
    }


    int countWithDistance( int minDistance) {
        int counter = 0;
        for (int x = left; x <= right; x++) {
            for (int y = bottom; y <= top; y++) {
                int dist = 0;
                Pair<Integer, Integer> p = new Pair<Integer, Integer>(x, y);
                for (Pair<Integer, Integer> q : coords) {
                    dist += manhatten(p, q);
                }
                if (dist < minDistance) {
                    counter++;
                }
            }
        }
        return counter;
    }
}
