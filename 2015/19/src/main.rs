use std::collections::HashSet;

fn parse(lines: &str) -> (Vec<(&str, &str)>, &str) {
    let mut rules = Vec::new();
    let parts: Vec<&str> = lines.split("\n\n").collect();
    for line in parts[0].split("\n") {
        let mut splitter = line.split(" => ");
        let left = splitter.next().unwrap();
        let right = splitter.next().unwrap();
        rules.push((left, right));
    }
    (rules, parts[1])
}

fn replaceAll(line: &str, (l, r): (&str, &str)) -> HashSet<String> {
    let indices: Vec<(usize, &str)> = line.match_indices(l).collect();
    let mut ret = HashSet::new();
    for i in indices {
        let mut new_line = String::from(&line[..i.0]);
        new_line += r;
        new_line += &line[i.0 + l.len()..];

        ret.insert(new_line);
    }
    ret
}

fn replace_one(line: &str, rules: &Vec<(&str, &str)>, count: u32) -> u32 {
    for r in rules {
        if let Some(_) = line.find(r.1) {
            let new_line = line.replacen(r.1, r.0, 1);
            if new_line == "e" {
                return count;
            }
            let c = replace_one(&new_line, rules, count + 1);
            if c > 0 {
                return c;
            }
        }
    }
    0
}

fn main() {
    let (rules, input) = parse(include_str!("test.txt"));
    println!("{:?}", input);
    let mut hash = HashSet::new();
    for r in rules {
        let h = replaceAll(input, r);
        hash.extend(h);
    }
    println!("{:?}", hash);
    println!("{:?}", hash.len());
    let (rules, input) = parse(include_str!("test.txt"));
    let mut hash = HashSet::new();
    for r in rules {
        let h = replaceAll("HOHOHO", r);
        hash.extend(h);
    }
    println!("{:?}", hash.len());

    let (rules, input) = parse(include_str!("input.txt"));
    let mut hash = HashSet::new();
    for r in rules {
        let h = replaceAll(input, r);
        hash.extend(h);
    }
    println!("{:?}", hash.len());
    let (rules, input) = parse(include_str!("test.txt"));
    println!("{}", replace_one(&input, &rules, 1));

    let input = "HOHOHO";
    println!("{}", replace_one(&input, &rules, 1));

    let (rules, input) = parse(include_str!("input.txt"));
    println!("{}", replace_one(&input, &rules, 1));
}
