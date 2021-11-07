package de.oflebbe;

import org.junit.jupiter.api.Test;

import java.io.StringReader;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class SchedulerTest {

    @Test
    void schedulerTest() {
        String input = "Step C must be finished before step A can begin.\n" +
                "Step C must be finished before step F can begin.\n" +
                "Step A must be finished before step B can begin.\n" +
                "Step A must be finished before step D can begin.\n" +
                "Step B must be finished before step E can begin.\n" +
                "Step D must be finished before step E can begin.\n" +
                "Step F must be finished before step E can begin.";

        Dependencies d = new Dependencies( new StringReader( input));
        try {
            assertEquals(15, d.scheduler(2, 0));
        }
        catch (Exception e) {
            e.printStackTrace();
        }
    }
}
