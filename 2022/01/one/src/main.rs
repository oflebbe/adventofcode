use std::fs;

fn calories( file_name: &str) -> Vec<i32> {
    let contents = fs::read_to_string(file_name)
        .expect("Should have been able to read the file");

    contents.split("\n\n").map(|block| {
        block.split('\n').map(|line| {
            line.parse::<i32>().expect("expecting number")
        }).sum()
    }).collect::<Vec<i32>>()
}

fn  max_calories(  cal : &mut Vec<i32>) -> i32 {
    cal.sort();
    cal.reverse();
    cal[0]
}

fn sum_top3( cal : &Vec<i32>) -> i32 {
    cal[0]+cal[1]+cal[2]
}

fn main() {
    let mut cals = calories( "input.txt");
    println!( "{}", max_calories( &mut cals));
    println!( "{}", sum_top3( &cals));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test() {
        let mut c = calories("example.txt");
        let expected =  vec![6000, 4000, 11000, 24000, 10000];
        assert_eq!( c, expected);


        assert_eq!( max_calories(&mut c),24000);

        assert_eq!(sum_top3(&c), 45000);
    }
}
    
