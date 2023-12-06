use regex::Regex;
use std::cmp;
use std::collections::HashMap;
use std::collections::HashSet;
use std::fs;

#[derive(Debug)]
struct Valve {
    Rate: i32,
    Valves: Vec<String>,
}

fn read_val(file_name: &str) -> HashMap<String, Valve> {
    let contents = fs::read_to_string(file_name).expect("Should have been able to read the file");
    let lines: Vec<&str> = contents.split("\n").collect();
    let re = Regex::new(r"^Valve ([A-Z]*) has flow rate=(\d*); tunnels* leads* to valves* (.*)$")
        .unwrap();

    let mut Valves = HashMap::new();
    for line in lines {
        let cap = re.captures(line).unwrap();
        let name = String::from(cap.get(1).unwrap().as_str());
        let rate = cap.get(2).unwrap().as_str().parse::<i32>().unwrap();
        let valves = cap
            .get(3)
            .unwrap()
            .as_str()
            .split(", ")
            .map(|x| String::from(x))
            .collect();
        let valve = Valve {
            Rate: rate,
            Valves: valves,
        };
        Valves.insert(name, valve);
    }
    Valves
}

fn count_closest(v: &HashMap<String, Valve>, a: &str, b: &str, s: &HashSet<&str>) -> Option<i32> {
    let mut best = 0;

    let mut first = true;
    let mut s2 = s.clone();
    s2.insert(a);

    for n in v[a].Valves.iter() {
        if n == b {
            return Some(1);
        }
        let nstr = n.as_str();
        if s2.contains(nstr) {
            continue;
        }
        if let Some(c) = count_closest(v, n.as_str(), b, &s2) {
            if first {
                best = c;
                first = false;
            } else {
                best = cmp::min(c, best);
            }
        }
    }
    if !first {
        return Some(best + 1);
    }
    None
}

fn find_max(
    myMap: &HashMap<(&str, &str), i32>,
    ventils: &HashMap<&str, i32>,
    current: &str,
    remainingTime: i32,
) -> i32 {
    let mut v2 = ventils.clone();

    v2.remove(current);
    if remainingTime <= 0 {
        return 0;
    }
    let remaining_time2 = remainingTime - 1;
    let flow = ventils[current] * remainingTime;
    let mut maxFlow = 0;
    for iter in v2.keys() {
        if ventils.contains_key(*iter) {
            maxFlow = cmp::max(
                maxFlow,
                find_max(
                    myMap,
                    &v2,
                    *iter,
                    remaining_time2 - myMap[&(current, *iter)],
                ),
            )
        }
    }
    maxFlow + flow
}

fn part1(filename: &str) -> i32 {
    let v = read_val(filename);

    let mut ventils = HashMap::new();
    for k in v.keys() {
        if v[k].Rate > 0 {
            ventils.insert(k.as_str(), v[k].Rate);
        }
    }
    ventils.insert("AA", 0);
    let mut myMap = HashMap::new();
    for i in ventils.keys() {
        for j in ventils.keys() {
            if *i == *j {
                continue;
            }
            if let Some(c) = count_closest(&v, *i, *j, &HashSet::new()) {
                myMap.insert((*i, *j), c);
            } else {
                panic!();
            }
        }
    }
    find_max(&myMap, &ventils, "AA", 30)
}

fn find_max_loop(myMap: &HashMap<(&str, &str), i32>, ventils: &HashMap<&str, i32>) -> i32 {
    let base = 1 << (ventils.len() - 1);
    let mut max = 0;
    let v_row: Vec<&str> = ventils.keys().map(|x| *x).collect();
    for i in 0..base {
        let mut v1 = HashMap::new();
        let mut v2 = HashMap::new();

        for k in 0..v_row.len() {
            let key = v_row[k];
            if (1 << k) & i == 0 {
                v1.insert(key, ventils[key]);
            } else {
                v2.insert(key, ventils[key]);
            }
        }
        v1.insert("AA", 0);
        v2.insert("AA", 0);

        let a1 = find_max(myMap, &v1, "AA", 26);
        let a2 = find_max(myMap, &v2, "AA", 26);
        max = cmp::max(max, a1 + a2);
        let r = 1;
    }
    max
}

fn part2(filename: &str) -> i32 {
    let v = read_val(filename);

    let mut ventils = HashMap::new();
    for k in v.keys() {
        if v[k].Rate > 0 {
            ventils.insert(k.as_str(), v[k].Rate);
        }
    }
    ventils.insert("AA", 0);

    let mut myMap = HashMap::new();
    for i in ventils.keys() {
        for j in ventils.keys() {
            if *i == *j {
                continue;
            }
            if let Some(c) = count_closest(&v, *i, *j, &HashSet::new()) {
                myMap.insert((*i, *j), c);
            } else {
                panic!();
            }
        }
    }
    ventils.remove("AA");

    find_max_loop(&myMap, &ventils)
}

fn main() {
    println!("{}", part1("test.txt"));
    println!("{}", part1("input.txt"));

    println!("{}", part2("test.txt"));
    println!("{}", part2("input.txt"));
}
