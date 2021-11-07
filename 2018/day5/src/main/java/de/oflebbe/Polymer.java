package de.oflebbe;

import lombok.Cleanup;
import lombok.val;
import lombok.var;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.Reader;
import java.util.*;

public class Polymer {
    String polymer = "";

    Polymer(Reader r) {
        try {
            polymer = new BufferedReader(r).readLine();
        }
        catch (IOException e) {
            e.printStackTrace();
        }

    }

    static boolean react(char one, char two) {
        if (one == two) {
            return false;
        }
        return (Character.toUpperCase( one ) == two || Character.toLowerCase( one) == two);

    }

    int react() {
        return reactOne( polymer);
    }


    int reactOne( String polymer) {
        int result = 0;
        char[] polyArray = polymer.toCharArray();
        int len = polyArray.length;
        int ptr = 0;
        while (ptr < len-1) {
            if (react( polyArray[ptr], polyArray[ptr+1])) {
                System.arraycopy(polyArray, ptr + 2, polyArray, ptr, len - ptr - 2);
                len -= 2;
                if (ptr > 0)
                    ptr--;
            } else {
                ptr++;
            }
        }
        return len;
    }


    static String removeSpecificChar(String polymer, char what) {
       return polymer.replace( Character.toString( what), "").replace( Character.toString( Character.toUpperCase( what)), "");
    }

    int whichChars() {
        Set<Character> chars = new HashSet<>();

        for (char ch : polymer.toLowerCase().toCharArray()) {
            chars.add(ch);
        }
        int min = polymer.length();
        char minChar = ' ';
        for (char ch : chars) {
            int size = reactOne( removeSpecificChar( polymer, ch));
            if (size < min) {
                min = size;
                minChar = ch;
            }
        }
        return min;
    }



}
