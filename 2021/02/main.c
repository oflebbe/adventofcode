#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main
( int argc, char *argv[]) {
    if (argc < 2) {
        fprintf(stderr, "need filename");
        exit(1);
    }
    FILE *fp =fopen( argv[1], "r");
    if (fp == NULL) {
        fprintf(stderr, "open %s failed", argv[1]);
        exit(1);
    }
    char line[100];
    memset( line, 0, sizeof(line));
    int forward = 0;
    int depth = 0;

    int aim = 0;
    int depth2 = 0;
    while (NULL != fgets( line, 100, fp)) {
        int value;
        char word[20];
        if (2 != sscanf(line, "%20s %d", word, &value)) {
            fprintf(stderr, "scan failed\n");
            exit(1);
        }
        if (strcmp( word, "forward") == 0) {
            forward += value;
            depth2 += aim * value;
        } else if (strcmp( word, "up") == 0) {
            depth -= value;
            aim -= value;
        } else if (strcmp( word, "down") == 0) {
            depth += value;
            aim += value;
        } else {
            fprintf(stderr, "word not found %s\n", word);
            exit(1);
        }
    }
    printf("part1 forward: %d depth: %d, check: %d\n", forward, depth, forward*depth);
    printf("part2 forward: %d depth: %d, check: %d\n", forward, depth2, forward*depth2);
}




