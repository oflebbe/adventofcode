//use std::fmt;
use std::fs;

struct Pairtuple(String, String);

fn load(file_name: &str) -> Vec<Pairtuple> {
    let contents = fs::read_to_string(file_name).expect("Should have been able to read the file");

    let lines = contents.split("\n\n").collect::<Vec<&str>>();

    let mut input = Vec::<Pairtuple>::new();
    for pairs in lines {
        let p = pairs
            .split("\n")
            .filter(|l| l.len() > 1)
            .collect::<Vec<&str>>();
        input.push(Pairtuple(p[0].to_string(), p[1].to_string()));
    }
    input
}

type List = Vec<Data>;

enum Data {
    N(i32),
    L(Box<List>),
}
/* 
impl fmt::Display for Data {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match &self {
            Data::N(a) => write!(f, "{}", a),
            Data::L(a) => {
                _ = write!(f, "[");
                let mut first = true;
                for v in &**a {
                    if first {
                        first = false;
                    } else {
                        _ = write!(f, ",");
                    }
                    _ = write!(f, "{}", v);

                }
                write!(f, "]")
            }
        }
    }
}*/

fn parse(line: &[u8], index: usize) -> (Data, usize) {
    let mut ptr = index;
    if line[ptr] == b'[' {
        let mut l = List::new();
        ptr += 1;
        if line[ptr] == b']' {
            let var_name = (Data::L(Box::new(l)), ptr + 1);
            return var_name;
        }
        loop {
            let (s, p) = parse(line, ptr);
            ptr = p;
            l.push(s);

            if line[ptr] == b']' {
                ptr += 1;
                break;
            }
            if line[ptr] != b',' {
                panic!();
            }
            ptr += 1;
        }

        let ret = (Data::L(Box::new(l)), ptr);
        return ret;
    } else if line[ptr].is_ascii_digit() {
        let mut next = ptr + 1;
        while line[next].is_ascii_digit() {
            next += 1;
        }
        let v = std::str::from_utf8(&line[ptr..next])
            .unwrap()
            .parse::<i32>()
            .unwrap();

        return (Data::N(v), next);
    } else {
        panic!();
    }
}

fn compare_lists(a: &List, b: &List) -> i32 {
    for i in 0..usize::min(a.len(), b.len()) {
        let diff = compare(&a[i], &b[i]);
        if diff < 0 {
            return -1;
        } else if diff > 0 {
            return 1;
        }
    }
    a.len() as i32 - b.len() as i32
}

fn compare(left: &Data, right: &Data) -> i32 {
   // println!("Compare: {} {}", left, right);
    return match (left, right) {
        (Data::N(a), Data::N(b)) => a - b,
        (Data::L(a), Data::L(b)) => compare_lists(&*a, &*b),
        (Data::N(a), Data::L(_)) => compare(&Data::L(Box::new(vec![Data::N(*a)])), right),
        (Data::L(_), Data::N(b)) => compare(left, &Data::L(Box::new(vec![Data::N(*b)]))),
    };
}

fn func1(filename: &str) {
    let input = load(filename);
    let mut sum = 0;
    let mut index = 1;
    for tuple in input {
      
        let (left, _) = parse(&tuple.0.as_bytes(), 0);

        let (right, _) = parse(&tuple.1.as_bytes(), 0);

        if compare(&left, &right) < 0 {
            sum += index;
        }
        index += 1;
    }

    println!("{} ", sum)
}

fn func2(filename: &str) {
    let input = load(filename);


    let mut all = Vec::new();

    for tuple in input {
        let (left, _) = parse(&tuple.0.as_bytes(), 0);
        let (right, _) = parse(&tuple.1.as_bytes(), 0);
        all.push( left);
        all.push(right);
    }
    
    let (two, _) = parse( "[[2]]".as_bytes(), 0);
    let (six, _) = parse( "[[6]]".as_bytes(), 0);
    
    let mut two_count = 1;
    let mut six_count = 2;
    
    for v in &all {
        if compare( v, &two) < 0 {
            two_count +=1;
        }
        if compare( v, &six) < 0 {
            six_count +=1;
        }
    }

    println!("{}", two_count*six_count)
}

fn main() {
    func1("test.txt");
    func1("input.txt");
    func2("test.txt");
    func2("input.txt");
}
