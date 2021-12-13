#include <string>
#include <unordered_map>
#include <set>
#include <list>
#include <sstream>
#include <istream>
#include <iostream>
#include <ctype.h>

using namespace std;

unordered_map<string, list<string> >
parse(  istream &input ) {
    std::string line;

    unordered_map<string, list<string> > cave;
    while (getline( input, line)) {
        int pos = line.find( '-');
        string a( line, 0, pos);
        string b(line, pos+1);
        if (cave.find( a) == cave.end()) {
            list<string> li;
            li.push_back( b);

            cave[ a] = li;
        } else {
            cave[a].push_back(b);
        }
        if (cave.find( b) == cave.end()) {
            list<string> li;
            li.push_back( a);

            cave[ b] = li;
        } else {
            cave[b].push_back(a);
        }
    }
    return cave;
}

int descend(const unordered_map<string, list<string> > &cave, set<string> &smallVisited, const string &twice , const string &pos) {
	int found = 0;
	if (pos == "end") {
		return 1;
	}
    auto a = (*cave.find(pos)).second;
	for (auto option : a) {
		if (option == "start") {
			continue;
		}
        if (smallVisited.find(option) != smallVisited.end() ) {
            if (twice == "") {
                found += descend(cave, smallVisited, option, option);
            } else {
                continue;
            }
        } else {
            if (islower(option[0])) {
				smallVisited.insert(option);
				found += descend(cave, smallVisited, twice, option);
				smallVisited.erase( option);
			} else {
				found += descend(cave, smallVisited, twice, option);
			}
		}
	} 
	return found;
}



int main() {
    
   string input("KF-sr\n"
"OO-vy\n"
"start-FP\n"
"FP-end\n"
"vy-mi\n"
"vy-KF\n"
"vy-na\n"
"start-sr\n"
"FP-lh\n"
"sr-FP\n"
"na-FP\n"
"end-KF\n"
"na-mi\n"
"lh-KF\n"
"end-lh\n"
"na-start\n"
"wp-KF\n"
"mi-KF\n"
"vy-sr\n"
"vy-lh\n"
"sr-mi");
    auto is = istringstream(input);
    auto cave = parse(is);
    
    
    set<string> visited;
    visited.insert("start");
    cout << descend(cave,visited, "start", "start") << "\n";
    cout << descend(cave,visited, "", "start") << "\n";
}