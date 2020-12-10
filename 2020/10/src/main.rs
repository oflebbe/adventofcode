fn problem( input : &str) -> (i32, i32) {
    let mut adapters = Vec::new();
    adapters.push(0);
    for line in input.lines() {
        let jolts = line.parse::<u32>().unwrap();
        adapters.push( jolts);
    }
    adapters.sort();
    adapters.push(adapters[adapters.len()-1] + 3);
    let mut count = 0;
    let mut start = 1;
    while 3 >= adapters[start] {
        count += problem2( &adapters[start+1..adapters.len()], adapters[start],  &vec![0]);
        start+=1;
    }
    let mut distribution = vec![0,0,0,0];
    for i in 0..adapters.len()-1 {
        distribution[ (adapters[i+1] - adapters[i]) as usize] += 1;
    }
    return (distribution[3] * distribution[1], count)
}

fn problem2( ad : &[u32], last : u32, prev: &Vec<u32>) -> i32 {
    if ad.len() <= 2 {
      //  print!("{:?},{},{:?}\n", prev, last, &ad);
       return 1
    }
    let mut count = 0;
    let mut start = 0;
    while start < ad.len() && last + 3 >= ad[start] {
        let mut newprev : Vec<u32> = prev.to_vec();
        newprev.push( last);
        count += problem2( &ad[start+1..ad.len()], ad[start], &newprev);
        start+=1;
    }
    count
}

fn main() {
  let i1 = "28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3";
    let (a,b) = problem( i1);
    print!("{} {}\n", a,b,); 

let i2 = "16
10
15
5
1
11
7
19
6
12
4";
let (a,b) = problem( i2);
print!("{} {}\n",a,b );

 let input = include_str!("input.txt");
    let (a,b) = problem( input);
    print!("{} {}\n", a,b) 
}
