
fn passwordCheck1(st : &str, low: i32, high:i32, character: char) -> bool {
    let mut count = 0;
    for ch in st.chars() {
        if character == ch {
            count+=1;
        }
    }
    count >= low && count <= high
}

fn passwordCheck2(st : &str, low: i32, high:i32, character: char) -> bool {
    let mut count = 0;
    for i in [ low, high].iter() {
        if st.chars().nth((*i - 1)as usize).unwrap() == character {
            count += 1;
        }
    }

    count == 1
}

fn problem( input: &str) -> (u32, u32) {
    let mut count1 = 0;
    let mut count2 = 0;
    for l in input.lines() {
        let tok : Vec<&str> = l.split(|c| c == ' ' || c == '-' || c == ':' ).collect();
        let low =  tok[0].parse::<i32>().unwrap();
        let high = tok[1].parse::<i32>().unwrap();
        let ch = tok[2].chars().next().unwrap();
        let passwd = tok[4];
        
        
        if passwordCheck1(passwd, low, high, ch) {
            count1 += 1;
        }   
        if passwordCheck2(passwd, low, high, ch) {
            count2 += 1;
        }   
    }
    (count1, count2)
}

fn main() {
    let input = "1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc";

    let (one, two) = problem( input);
    
    println!("correct: {} {}\n", one, two);

    let input_real = include_str!("input.txt");
    let (one, two) = problem( input_real);
    println!("correct: {} {}\n",one, two);
}
