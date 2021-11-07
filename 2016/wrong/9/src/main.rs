use std::collections::HashSet;
use std::collections::HashMap;

fn recurse(dist: u32, city: &str, cities: &HashSet<String>, d : &HashMap<String, u32>) -> u32 {
    let mut my_cities : HashSet<String> = cities.clone();
    my_cities.remove( city);
    let min : Option<u32> ;
    if my_cities.len() != 0 {
        
        for c in my_cities.iter() {
            let name = format!("{}_{}", city, c);
            let dd = d.get(&name).unwrap();

            let way = recurse(dist + dd, c, &my_cities, d);
            min = match way {
                Some(x) => Some( if x < min {
                                min = x;
                            } 
                            x),
                _ => None;
            }

            
        }
    }
    min
}

fn main() {

   
    let input = "London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141";

    let mut cities = HashSet::new();
    let mut d : HashMap<String, u32>= HashMap::new();
    for line in input.lines() {
        let tok : Vec<&str>= line.split_ascii_whitespace().collect();
        cities.insert(tok[0]);
        cities.insert(tok[2]);
        let dist = tok[4].parse::<u32>().unwrap();
        let name = format!("{}_{}", tok[0], tok[2]);

        d.insert(name, dist);
        let name = format!("{}_{}", tok[2], tok[0]);
        d.insert(name, dist);
    }

    let start = cities.iter().next().unwrap();

   


    println!("Hello, world!");
}
