package de.oflebbe;

import lombok.EqualsAndHashCode;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

@EqualsAndHashCode
public class Claim {
    final int number;
    final int left;
    final int top;
    final int width;
    final int height;
    static Pattern pat = Pattern.compile("^#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)");


    Claim( int number, int left, int top, int width, int height) {
        this.number = number;
        this.left = left;
        this.top = top;
        this.width = width;
        this.height = height;
    }

    Claim( String number, String left, String top, String width, String height) {
        this.number = Integer.parseInt(number);
        this.left = Integer.parseInt(left);
        this.top = Integer.parseInt(top);
        this.width = Integer.parseInt(width);
        this.height = Integer.parseInt(height);
    }

    static Claim parseString( String line) {

        Matcher matcher = pat.matcher( line);
        if (!matcher.matches()) {
            return null;
        }
        return new Claim( matcher.group(1), matcher.group(2), matcher.group(3), matcher.group(4), matcher.group(5));
    }

    int W() { return left + width; }
    int H() { return top + height; }
}
