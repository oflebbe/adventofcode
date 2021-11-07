package de.oflebbe;

import lombok.Cleanup;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.Reader;
import java.util.*;
import java.util.regex.Matcher;
import java.util.regex.Pattern;


class SleepDetector {
    ArrayList<String> notes = new ArrayList<>();
    final Pattern minRegexFalls =  Pattern.compile("^\\[.*:(\\d+)\\] falls.*");
    final Pattern minRegexWakes = Pattern.compile("^\\[.*:(\\d+)\\] wakes.*");
    final Pattern guardRegex = Pattern.compile(".*Guard #(\\d+) begins shift");
    final Map<Integer,ArrayList<Integer>> timetable = new HashMap<>();

    SleepDetector( Reader r) {
        try {
            @Cleanup BufferedReader in = new BufferedReader(r);
            String line = in.readLine();
            while (line != null) {
                notes.add(line);
                line = in.readLine();
            }

            notes.sort(null);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    void processNotes() {
        int guard = -1;

        ListIterator<String> ptr = notes.listIterator();
        ArrayList<Integer> guardTimetable = null;
        while (ptr.hasNext()) {
            String line = ptr.next();
            while (true) {
                Matcher guardMatch = guardRegex.matcher(line);
                if (guardMatch.matches()) {
                    guard = Integer.parseInt(guardMatch.group(1));
                    guardTimetable = timetable.get(guard);
                    if (guardTimetable == null) {
                        guardTimetable = new ArrayList<Integer>(60);
                        for (int i = 0; i < 60; i++) {
                            guardTimetable.add(0);
                        }
                    }

                    line = ptr.next();
                } else {
                    break;
                }
            }
            Matcher matcher = minRegexFalls.matcher(line);
            if (!matcher.matches())
                System.err.println("minRegexFalls DOes not match");
            int start = Integer.parseInt(matcher.group(1));
            line = ptr.next();
            matcher = minRegexWakes.matcher(line);
            if (!matcher.matches())
                System.err.println("minRegexWakes does not match");

            int stop = Integer.parseInt(matcher.group(1));
            for (int i = start; i < stop; i++) {
                guardTimetable.set(i, guardTimetable.get(i) + 1);
            }
            timetable.put( guard, guardTimetable);
        }
    }

    int findMostSleeping() {
        int max = 0;
        int maxGuard = 0;

        for (int guard : timetable.keySet()) {
            ArrayList<Integer> guardTimetable = timetable.get(guard);
            int sum = 0;
            for (int e : guardTimetable) {
                sum += e;
            }
            if (sum > max) {
                max = sum;
                maxGuard = guard;
            }
        }
        // System.out.printf("max %d maxGuard %d", max, maxGuard);
        return maxGuard;
    }

    int findMaxminute( int guard) {
        ArrayList<Integer> guardTimetable = timetable.get(guard);
        int maxMinute = 0;
        int minute = -1;
        for (int i = 0; i < 60; i++) {
            int e = guardTimetable.get( i);
            if (maxMinute < e) {
                maxMinute = e;
                minute = i;
            }
        }
        return minute;
    }

    // 0 guard
    // 1 minute
    int findMostSleepSameMinute() {
        int max = 0;
        int maxGuard = -1;
        int maxMinute = -1;

        for (int guard : timetable.keySet()) {
            ArrayList<Integer> guardTimetable = timetable.get(guard);

            for (int i = 0; i < 60; i++) {
                int e = guardTimetable.get(i);
                if (max < e) {
                    maxMinute = i;
                    maxGuard = guard;
                    max = e;
                }
            }
        }
        // System.out.printf("max %d maxGuard %d", max, maxGuard);
        return maxMinute * maxGuard;
    }
}