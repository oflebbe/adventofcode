use std::cmp;

fn problem( input : &str, num : usize ) -> u64 {
    let mut v = Vec::new();
    for l in input.lines() {
        let value =  l.parse::<u64>().unwrap();
        v.push( value)
    }
    let mut target : u64 = 0;
    for i in num..v.len() {
        let mut found = false;
        for j in i-num..i {
            for k in j+1..i {
                if v[j] + v[k] == v[i] {
                    found = true;
                }
            }
        }
        if !found { 
            target = v[i] as u64;
            break;
        }
    }
    if target == 0 {
        panic!("Bug");
    } 
    for i in 0 .. v.len() {
        let mut sum = 0;
        for j in 0..v.len()-i {
            sum += v[i+j];
            if sum > target {
                break;
            } else if sum == target {
                // found now search for largest and smalles int
                // can be done with iterators, though.
                let mut mi = v[i];
                let mut ma = v[i];
                for k in i+1..i+j+1 {
                    mi = cmp::min( mi, v[k]);
                    ma = cmp::max( ma, v[k]);
                }
                return mi + ma;
            }
        }
    }
    panic!("Bug2");
}

fn main() {

    let test = "35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576";
    let res = problem(test, 5);
    println!("correct: {}\n", res );

    let input = include_str!("input.txt");
    let res = problem(input, 25 );
    println!("correct: {}\n", res );
}
