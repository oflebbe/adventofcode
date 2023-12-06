HashMap
use std::cmp;
use std::collections::HashSet;
use std::fs;

fn merger(ranges: &Vec<(i64, i64)>) -> Vec<(i64, i64)> {
    let mut merge = Vec::new();
    let mut merged = HashSet::new();
    for i in 0..ranges.len() {
        if merged.contains(&i) {
            continue;
        }

        let mut rmin = ranges[i].0;
        let mut rmax = ranges[i].1;
        for j in i + 1..ranges.len() {
            if merged.contains(&j) {
                continue;
            }
            if rmin <= ranges[j].1 && rmax >= ranges[j].0 {
                rmin = cmp::min(rmin, ranges[j].0);
                rmax = cmp::max(rmax, ranges[j].1);
                merged.insert(j);
            }
            if (rmin == ranges[j].1+1) {
                rmin = ranges[j].0;
                merged.insert(j);
            }
            if (rmax == ranges[j].0-1) { 
                rmax = ranges[j].1;
                merged.insert(j);
            }
        }
        merge.push((rmin, rmax));
    }
    merge
}

fn read_val(file_name: &str) -> Vec<[i64; 4]> {
    let contents = fs::read_to_string(file_name).expect("Should have been able to read the file");
    let lines: Vec<&str> = contents.split("\n").collect();
    let re = Regex::new(r"^.*x=(-*\d*), y=(-*\d*):.*x=(-*\d*), y=(-*\d*)$").unwrap();
    let mut values = Vec::new();
    for line in lines {
        let cap = re.captures(line).unwrap();
        let mut val: [i64; 4] = [0; 4];
        for i in 1..=4 {
            val[i - 1] = cap.get(i).unwrap().as_str().parse().unwrap();
        }
        values.push(val);
    }
    values
}

fn task1(file_name: &str, row: i64) -> i64 {
    let values = read_val(file_name);
    let mut ranges: Vec<(i64, i64)> = Vec::new();
    for val in values {
        let dist = (val[0] - val[2]).abs() + (val[1] - val[3]).abs();

        let ydist = (val[1] - row).abs();
        if ydist > dist {
            continue;
        }
        let xmin = val[0] - (dist - ydist);
        let xmax = val[0] + (dist - ydist);

        ranges.push((xmin, xmax));
    }
    // merge ranges
    let mut merged = merger(&ranges);
    while merged.len() < ranges.len() {
        ranges = merged;
        merged = merger(&ranges);
    }
    let mut count = 0;
    for r in ranges {
        count += r.1 - r.0;
    }
    count as i64
}

fn task2(file_name: &str, atmost: i64) -> i64 {
    let values = read_val(file_name);

    for row in 0..=atmost {
        let mut ranges: Vec<(i64, i64)> = Vec::new();
        for val in values.iter() {
            let dist = (val[0] - val[2]).abs() + (val[1] - val[3]).abs();

            let ydist = (val[1] - row).abs();
            if ydist > dist {
                continue;
            }
            
            let xmin = val[0] - (dist - ydist);
            let xmax = val[0] + (dist - ydist);

            ranges.push((xmin, xmax));
        }
        // merge ranges
        let mut merged = merger(&ranges);
        while merged.len() < ranges.len() {
            ranges = merged;
            merged = merger(&ranges);
        }
        if merged.len() > 1 {
            println!("{:?} {}", merged, row);

            return (merged[0].1+1) * 4000000 + row
        }
    }
    0
}

fn main() {
    println!("{}", task1("test.txt", 10));

    println!("{}", task1("input.txt", 2_000_000));

    println!("{}", task2("test.txt", 20));

    println!("{}", task2("input.txt", 4000000));
}
