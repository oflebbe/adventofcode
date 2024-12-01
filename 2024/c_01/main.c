#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>
#include <assert.h>

int compare(const void *a, const void *b) {
    const int *a1 = a;
    const int *b1 = b;
    return (*a1 > *b1);
}

int main(int argc, char *argv[]) {
    if (argc < 2) {
        fprintf(stderr,"not enough argumnts\n");
        exit(1);
    }
    FILE *fp = fopen( argv[1], "r");
    if (!fp) {
        fprintf(stderr, "open file %s: %s\n", argv[1], strerror(errno));
        exit(1);
    }
    if (-1 == fseek( fp, 0, SEEK_END)) {
        fprintf(stderr, "seek file %s: %s\n", argv[1], strerror(errno));
        exit(1);
    }
    const long count = ftell( fp);
    if (count < 0) {
        fprintf(stderr, "ftell file %s: %s\n", argv[1], strerror(errno));
        exit(1);
    }
    if (-1 == fseek( fp, 0, SEEK_SET)) {
        fprintf(stderr, "seek file %s: %s\n", argv[1], strerror(errno));
        exit(1);
    }

    char *buffer = calloc( count+1, 1);
    if (buffer == NULL) {
        fprintf(stderr, "calloc %ld: %s\n", count, strerror(errno));
        exit(1);
    }
    const int el = fread( buffer, 1, count, fp);
    if (el != count) {
        fprintf(stderr, "fread file %s: %s\n", argv[1], strerror(errno));
    }
    fclose( fp);

    int lines = 0;
    for (char *ptr = buffer; *ptr != 0; ptr++ ) {
        if (*ptr == '\n') {
            lines++;
        }
    }
    assert(lines < count);
    int *col1 = calloc( lines, sizeof(int));
    int *col2 = calloc( lines, sizeof(int));
    if (col1==NULL || col2==NULL) {
        fprintf(stderr, "calloc err");
        exit(1);
    }
    char *line = strtok( buffer, "\n");
    int counter= 0;
    while (line != NULL) {
        if (strlen(line) == 0) {
            break;
        }
        assert( counter < lines);
        int one, two;
        int c = sscanf(line, "%d %d", &one, &two);
        if (c != 2) {
            fprintf(stderr, "sscanf error on %s\n", line);
            exit(1);
        }
        col1[counter] = one;
        col2[counter] = two;
        counter++;
        line = strtok(NULL, "\n");
    }
    assert(counter == lines);

    qsort( col1, counter, sizeof(int), compare);
    qsort( col2, counter, sizeof(int), compare);
    
    int sum = 0;
    for (int i= 0; i < lines; i++) {
        sum += abs(col1[i]-col2[i]);
    }
    printf("Task1 : sum is %d\n", sum);


    int ptr2 = 0;
    int total_score = 0;
    int last_score = 0;
    int last_value = -1;

    // since we are sorted, we can go line by line
    for (int i=0; i < lines; i++) {
        // already scored
        if (col1[i]== last_value) {
            total_score += last_score;
            continue;
        }
        if (ptr2 >= lines) {
            break;
        } 
        while (col2[ptr2] < col1[i]) {
            ptr2++;
            if (ptr2 >= lines) {
                break;
            }
        }
        int score = 0;

        if (ptr2 >= lines) {
            break;
        }
        while (col2[ptr2] == col1[i]) {
            score++;
            ptr2++;
            if (ptr2>=lines) {
                break;
            }
        }
        
        total_score += score * col1[i];
        last_value= col1[i];
       last_score = score * col1[i];
    }

    printf("Task2: total score %d\n", total_score);
    free(buffer);
    free(col1);
    free(col2);
}