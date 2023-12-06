
#include <charconv>
#include <filesystem>
#include <fstream>
#include <iostream>
#include <list>
#include <optional>
#include <set>
#include <string>
#include <string_view>
#include <unordered_set>
#include <vector>
#include <assert.h>

#include <stdlib.h>

int numOfHeapAllocations = 0;


using namespace std;

void* operator new(size_t size)
{
    numOfHeapAllocations++;
    return malloc(size);
}


// for string delimiter
template <typename T> vector<T> split(T s, string_view delimiter) {
  size_t pos_start = 0, pos_end, delim_len = delimiter.length();
  T token;
  vector<T> res;

  while ((pos_end = s.find(delimiter, pos_start)) != string::npos) {
    token = s.substr(pos_start, pos_end - pos_start);
    pos_start = pos_end + delim_len;
    res.push_back(token);
  }

  res.push_back(s.substr(pos_start));
  return res;
}

std::optional<int> to_int(const std::string_view &input) {
  int out;
  const std::from_chars_result result =
      std::from_chars(input.data(), input.data() + input.size(), out);
  if (result.ec == std::errc::invalid_argument ||
      result.ec == std::errc::result_out_of_range) {
    return std::nullopt;
  }
  return out;
}

vector<string> generate_input(const filesystem::path &name) {
  ifstream file(name);
  assert(file);

  string content((istreambuf_iterator<char>(file)),
                 (istreambuf_iterator<char>()));

  vector<string> lines = split<string>(content, "\n");
  cout << numOfHeapAllocations << endl;
  return lines;
}

pair<int, int> parse_elf(string_view s) {
  auto range = split(s, "-");
  auto rl = to_int(range[0]);
  auto rr = to_int(range[1]);
  assert(rl && rr);
  return pair { *rl, *rr };
}

bool fully( const pair<int,int>& a, pair<int,int>& b) {
  assert( a.first <= a.second);
  assert( b.first <= b.second);
  
  if (a.first == b.first || a.second == b.second) {
    return true;
  }
  if ( a.first < b.first ) {
    return a.second > b.second;
  } 
  return a.second < b.second;
}

bool partial( const pair<int,int>& a, const pair<int,int>& b) {
  assert( a.first <= a.second);
  assert( b.first <= b.second);
  
  if ( a.second < b.first || a.first > b.second ) {
    return false;
  } 
  return true;
}
pair<int,int> test(const filesystem::path &name) {
  vector<string> lines = generate_input(name);
  int count1 = 0;
  int count2 = 0;
  for (const auto &line : lines) {
    auto toks = split<string_view>(line, ",");
    assert(toks.size() == 2);

    pair<int,int> one = parse_elf( toks[0]);
    pair<int,int> two = parse_elf( toks[1]);
    if (fully( one, two )) {
      count1++;
    }
    if (partial(one,two)) {
      count2++;
    }
  }
  return pair{count1,count2};
}

int main() {
   auto count = test("test.txt");
  assert(count.first == 2&& count.second == 4);
  /*count = test("input.txt");
  auto heaps =numOfHeapAllocations;
  cout << count.first << ":" <<count.second << endl;
  cout << heaps<<endl; */
  cout << numOfHeapAllocations<<endl;
  
}