package de.oflebbe;

public class StringComparer {

    static String commonButOne(String a, String b) {
        if (a.length() != b.length()) {
            return "";
        }
        StringBuffer result = new StringBuffer(a.length()-1);
        int counter = 0;
        for (int index = 0; index < a.length(); index++) {
            if (a.charAt(index) != b.charAt(index)) {
                counter++;
                if (counter > 1)
                    return "";
            } else {
                result.append( a.charAt(index));
            }

        }
        if (counter != 1)
            return "";
        return result.toString();
    }
}
