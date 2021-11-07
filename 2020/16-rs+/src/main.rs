use regex::Regex;

fn parse( input: &str) {
    let v : Vec<&str>  = input.split("\n\n").collect();
    let r = Regex::new(r"(\w+): (\d+)-(\d+) or (\d+)-(\d+)").unwrap();

    for i 
    let caps = r.captures(input);
    if let Some(cap) = caps {
        let word = cap.get(1).unwrap().as_str().to_string();
        let a : i32 = cap.get(2).unwrap().as_str().parse().unwrap();
        let b : i32 = cap.get(3).unwrap().as_str().parse().unwrap();
        let c : i32 = cap.get(4).unwrap().as_str().parse().unwrap();
        let d : i32 = cap.get(5).unwrap().as_str().parse().unwrap();
        println!("{} {} {} {}", a,b, c,d)
    }
    println!("{}", v[0])


}


fn main() {
    let input = include_str!("example.txt");
    parse(input)
}
