use regex::Regex;
use std::collections::HashMap;
use std::collections::HashSet;
#[macro_use]
extern crate lazy_static;


lazy_static! {
    static ref RE: Regex = Regex::new(r"[[:upper:]]").unwrap();
}

// visit remaining cave , returning finds so far
fn descend(
    cave: &HashMap<&str, Vec<&str>>,
    small_visited: &HashSet<&str>,
    twice: &str,
    pos: &str,
) -> u32 {
    let mut found = 0;
    if pos == "end" {
        return 1;
    }
    let c = cave[pos].clone();
    for option in c {
        if option.to_string() == "start" {
            continue;
        }

        if small_visited.contains(option) {
            if twice == "" {
                found += descend(cave, small_visited, option, option);
            } else {
                continue;
            }
        } else {
            if !RE.is_match(option) {
                let mut visited = small_visited.clone();
                visited.insert(option);
                found += descend(cave, &visited, twice, option);
            } else {
                found += descend(cave, small_visited, twice, option);
            }
        }
    }
    found
}

fn parse(input: &str) -> HashMap<&str, Vec<&str>> {
    let mut a: HashMap<&str, Vec<&str>> = HashMap::new();

    let lines: Vec<_> = input.split("\n").collect();
    for l in lines {
        let pair: Vec<_> = l.split("-").collect();
        if let Some(list) = a.get_mut(pair[0]) {
            list.push(pair[1]);
        } else {
            a.insert(pair[0], vec![pair[1]]);
        }
        if let Some(list) = a.get_mut(pair[1]) {
            list.push(pair[0]);
        } else {
            a.insert(pair[1], vec![pair[0]]);
        }
    }
    a
}

fn main() {
    let input = include_str!("input.txt");
    let a = parse(input);
    let mut visited = HashSet::new();
    visited.insert("start");
    let res = descend(&a, &visited, "start", "start");
    println!("{}", res);
    let mut visited = HashSet::new();
    visited.insert("start");
    let res = descend(&a,&visited, "", "start");
    println!("{}", res);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn main() {
        let input = "start-A\n\
        start-b\n\
        A-c\n\
        A-b\n\
        b-d\n\
        A-end\n\
        b-end";
        let a = parse(input);
        let mut visited = HashSet::new();
        visited.insert("start");
        let res = descend(&a, &visited, "start", "start");
        assert_eq!(10, res);
    }
    #[test]
    fn part2() {
        let input = "start-A\n\
        start-b\n\
        A-c\n\
        A-b\n\
        b-d\n\
        A-end\n\
        b-end";
        let a = parse(input);
        let mut visited = HashSet::new();
        visited.insert("start");
        let res = descend(&a, &visited, "", "start");
        assert_eq!(36, res);
    }
}
