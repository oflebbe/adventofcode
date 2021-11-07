
fn possible( st :&str) -> bool {
    let mut a : Vec<_> = st.split_whitespace().map(|x| x.parse::<i32>().unwrap()).collect();
    a.sort();
    a[0] + a[1] > a[2] 
}

fn possible3( stack :Vec<&str>) -> i64 {
    let alle : Vec<_> = stack.iter().flat_map( |line| line.split_whitespace()).
        map(|x| x.parse::<i32>().unwrap()).collect();
    let mut eins : Vec<_>  = alle.iter().step_by(3).collect();
    let mut zwei : Vec<_>  = alle.iter().skip(1).step_by(3).collect();
    let mut drei : Vec<_>  = alle.iter().skip(2).step_by(3).collect();
    eins.sort();
    zwei.sort();
    drei.sort();

    let mut count = 0;
    if eins[0] + eins[1] > *eins[2] {
        count +=1;
    }
    if zwei[0] + zwei[1] > *zwei[2] {
        count +=1;
    }
    if drei[0] + drei[1] > *drei[2] {
        count +=1;
    }
    count
}

fn main() {
    println!("{}", possible("25 10 5"));

    let my_str = include_str!("input.txt");
    let mut count = 0;
    let mut poss = 0;
    for line in my_str.lines() {
        count +=1;
        if possible( line) {
            poss+=1;
        }
    }
    println!("{} {}", count, poss);

    let mut count = 0;
    let mut stack : Vec<&str> = vec!();
    let mut poss = 0;
    for line in my_str.lines() {
        stack.push( line.clone());
        count += 1;
        if count % 3 == 0 {
            poss += possible3( stack);
            stack = vec!();
        }
    }
    println!("{} {}", count, poss);


}
