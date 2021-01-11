
fn solve( input: &str) -> u64 {
    let mut lines = input.lines();
    let time : u64 = lines.next().unwrap().parse().unwrap();
    let bus_lines = lines.next().unwrap();
    let mut earliest = 0;
    let mut min_time : u64 = u64::MAX;
    for bus in bus_lines.split(',') {
        if bus == "x" {
            continue;
        }
        let nr : u64 = bus.parse().unwrap();
        let diff = (nr * ((time / nr) + 1 )) - time;
        if diff < min_time {
            earliest = nr * diff;
            min_time = diff;
        }
    }
    earliest
}

fn parse3( input : &str) -> (Vec<u64>, Vec<u64>) {
    let mut offset = 0;
    let mut diffs : Vec<u64>  = Vec::new();
    let mut nums : Vec<u64>  = Vec::new();
    for bus in input.split(',') {
        if bus == "x" {
            offset += 1;
            continue;
        }
        diffs.push( offset);
        offset+=1;
        let nr : u64 = bus.parse().unwrap();
        nums.push( nr);
    }
   
    print!("{} ", nums[0]);
    for i in 1..nums.len() {
        print!(" -{}- {}", diffs[i], nums[i]);
    }
    println!();

    (nums, diffs)
}

fn parse2( input : &str) -> (Vec<u64>, Vec<u64>) {
    let mut lines = input.lines();
    lines.next();
    parse3( lines.next().unwrap())
}

// Implement variant of solution to chinese remainder theorem
fn solve2( nums : &[u64], diffs : &[u64]) -> u64 {
    let mut step = nums[0];
    let mut v = 0;
    let mut i = 1;
    let mut D = nums[1] - diffs[1];
    loop {
        if v % nums[i] == D {
            step *= nums[i];
            i += 1;
            if i >= nums.len() {
                break;
            }
            // inputs are not sorted, we have to skip forward.
            let mut DD =  nums[i] as i64 - diffs[i] as i64;
            while DD < 0 { 
                DD += nums[i] as i64;
            }
            D = DD as u64;
        }
        v += step;
    }
    return v;
}

fn main() {
    let (n, d) = parse3( "17,x,13,19");
    assert_eq!( 3417,  solve2( &n, &d));

    let test_input="939
7,13,x,x,59,x,31,19";
        let (n, d) = parse2(test_input);
    assert_eq!( 1068781, solve2(&n, &d));
    
    let input = include_str!("input.txt");
    let (n, d) = parse2(input);
    println!("{}", solve2(&n, &d));
}
