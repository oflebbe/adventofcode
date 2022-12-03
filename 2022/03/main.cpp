#include <assert.h>
#include <filesystem>
#include <fstream>
#include <iostream>
#include <list>
#include <set>
#include <string>
#include <string_view>
#include <unordered_set>
#include <vector>

using namespace std;

char eval(const string_view line) {

  assert(line.length() % 2 == 0);
  const int len2 = line.length() / 2;

  auto part1 = string_view(line.begin(), len2);
  auto part2 = string_view(line.begin() + line.length() / 2, len2);
  unordered_set<char> set1(len2);

  for (int i = 0; i < len2; i++) {
    set1.insert(part1[i]);
  }

  for (int i = 0; i < len2; i++) {
    if (set1.contains(part2[i])) {
      return part2[i];
    }
  }
  return ' ';
}

char eval2(const list<string_view> &lines) {
  assert(lines.size() == 3);
  vector<set<char>> sets;
  for (const auto &line : lines) {
    set<char> s;
    for (auto ch : line) {
      s.insert(ch);
    }
    sets.push_back(s);
  }
  set<char> tmp;
  set_intersection(sets[0].cbegin(), sets[0].cend(), sets[1].cbegin(),
                   sets[1].end(), inserter(tmp, tmp.begin()));
  set<char> res;
  set_intersection(tmp.cbegin(), tmp.cend(), sets[2].cbegin(), sets[2].end(),
                   inserter(res, res.begin()));
  assert(res.size() == 1);
  return *res.begin();
}

int prio2(const list<string_view> &lines) {
  char ch = eval2(lines);
  if (ch >= 'a' && ch <= 'z') {
    return ch - 'a' + 1;
  }
  if (ch >= 'A' && ch <= 'Z') {
    return ch - 'A' + 27;
  }
  assert(false);
}

void test2() {
  list<string_view> cases1{"vJrwpWtwJgWrhcsFMMfFFhFp",
                           "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
                           "PmmdzqPrVvPwwTWBwg"};
  list<string_view> cases2{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT",
                           "CrZsJsPPZsGzwwsLwLmpwMDw"};

  assert(eval2(cases1) == 'r');
  assert(eval2(cases2) == 'Z');
}

int prio(const string_view line) {
  char ch = eval(line);
  if (ch >= 'a' && ch <= 'z') {
    return ch - 'a' + 1;
  }
  if (ch >= 'A' && ch <= 'Z') {
    return ch - 'A' + 27;
  }
  assert(false);
}

// for string delimiter
vector<string_view> split(string_view s, string_view delimiter) {
  size_t pos_start = 0, pos_end, delim_len = delimiter.length();
  string_view token;
  vector<string_view> res;

  while ((pos_end = s.find(delimiter, pos_start)) != string::npos) {
    token = s.substr(pos_start, pos_end - pos_start);
    pos_start = pos_end + delim_len;
    res.push_back(token);
  }

  res.push_back(s.substr(pos_start));
  return res;
}

int task1(const filesystem::path &name) {
  ifstream file(name);
  assert(file);

  string content((istreambuf_iterator<char>(file)),
                 (istreambuf_iterator<char>()));

  vector<string_view> lines = split(content, "\n");
  int sum = 0;
  for (const auto &line : lines) {
    sum += prio(line);
  }
  return sum;
}

int task2(const filesystem::path &name) {
  ifstream file(name);
  assert(file);

  string content((istreambuf_iterator<char>(file)),
                 (istreambuf_iterator<char>()));

  vector<string_view> lines = split(content, "\n");
  int sum = 0;
  assert(lines.size() % 3 == 0);
  for (int i = 0; i < lines.size() / 3; i++) {
    list<string_view> l;
    l.push_back(lines[i * 3 + 0]);
    l.push_back(lines[i * 3 + 1]);
    l.push_back(lines[i * 3 + 2]);
    sum += prio2(l);
  }
  return sum;
}

void test() {
  list<string_view> cases{
      "vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
      "PmmdzqPrVvPwwTWBwg",       "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
      "ttgJtRGJQctTZtZT",         "CrZsJsPPZsGzwwsLwLmpwMDw"};
  list<char> expected{'p', 'L', 'P', 'v', 't', 's'};
  auto exp = expected.cbegin();
  for (const auto &line : cases) {
    assert(exp != expected.cend());
    assert(eval(line) == *exp++);
  }

  list<int> exp_prio{16, 38, 42, 22, 20, 19};
  auto exp_p = exp_prio.cbegin();
  for (const auto &line : cases) {
    // cout << *line << " " << *exp << "\n";
    assert(exp_p != exp_prio.cend());
    assert(prio(line) == *exp_p++);
  }
}

int main() {
  test();
  int sum = task1("input.txt");
  cout << sum << "\n";
  test2();
  int sum2 = task2("input.txt");
  cout << sum2 << "\n";

}