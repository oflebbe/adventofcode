#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int score(char opp, char own) {
  assert(opp >= 'A' && opp <= 'C');
  assert(own >= 'X' && own <= 'Z');

  int ownScore = own - 'X' + 1;

  if (opp - 'A' == own - 'X') {
    return 3 + ownScore;
  }
  if (opp == 'A' && own == 'Y') {
    return 6 + ownScore;
  }
  if (opp == 'B' && own == 'Z') {
    return 6 + ownScore;
  }
  if (opp == 'C' && own == 'X') {
    return 6 + ownScore;
  }
  return 0 + ownScore;
}

int realScore(char opp, char goal) {
  assert(opp >= 'A' && opp <= 'C');
  assert(goal >= 'X' && goal <= 'Z');

  switch (goal) {
  case 'Y':
    return 3 + opp - 'A' + 1;

  case 'X':
    switch (opp) {
    case 'A':
      return 3; // C
    case 'B':
      return 1; // A
    case 'C':
      return 2; // B
    }
  case 'Z':
    switch (opp) {
    case 'A':
      return 6+2; // B
    case 'B':
      return 6+3; // C
    case 'C':
      return 6+1; // A
    }
  }
  fprintf(stderr,"internal error");
  exit(2);
}

int test() {
  FILE *fp = fopen("test.txt", "r");
  if (fp == NULL) {
    fprintf(stderr, "could not open");
    exit(1);
  }
  char *buf = NULL;
  size_t len = 0;

  int result[] = {8, 1, 6};
  int count = 0;
  int globalScore = 0;
  while (getline(&buf, &len, fp) == 4) {
    int s = score(buf[0], buf[2]);
    if (s != result[count]) {
      fprintf(stderr, "Error line %d", count + 1);
    }
    globalScore += s;

    count++;
  }
  free(buf);
  fclose(fp);
  if (globalScore != 15) {
    fprintf(stderr, "score wrong\n");
  }
  return globalScore;
}

int test2() {
  FILE *fp = fopen("test.txt", "r");
  if (fp == NULL) {
    fprintf(stderr, "could not open");
    exit(1);
  }
  char *buf = NULL;
  size_t len = 0;

  int result[] = {4, 1, 7};
  int count = 0;
  int globalScore = 0;
  while (getline(&buf, &len, fp) == 4) {
    int s = realScore(buf[0], buf[2]);
    if (s != result[count]) {
      fprintf(stderr, "Error line %d", count + 1);
    }
    globalScore += s;

    count++;
  }
  free(buf);
  fclose(fp);
  if (globalScore != 12) {
    fprintf(stderr, "score wrong\n");
  }
  return globalScore;
}

int main() {
  test();

  FILE *fp = fopen("input.txt", "r");
  if (fp == NULL) {
    fprintf(stderr, "could not open");
    exit(1);
  }
  char *buf = NULL;
  size_t len = 0;

  int count = 0;
  int globalScore = 0;
  while (getline(&buf, &len, fp) == 4) {
    int s = score(buf[0], buf[2]);
    globalScore += s;

    count++;
  }
  
  printf("Score: %d\n", globalScore);

  test2();
  rewind(fp);

  int count2 = 0;
  int globalScore2 = 0;
  while (getline(&buf, &len, fp) == 4) {
    int s2 = realScore(buf[0], buf[2]);
    globalScore2 += s2;

    count2++;
  }

   printf("Real Score: %d\n", globalScore2);
  free(buf);
}