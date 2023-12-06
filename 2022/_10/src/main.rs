use std::fs;

fn task1(file_name: &str) -> i32 {
    let mut strength = 0;
    let mut cycle = 1;

    decoder(file_name, |x: i32| {
        match cycle {
            20 | 60 | 100 | 140 | 180 | 220 =>  strength += x * cycle,
            _ => {}
        }

        cycle += 1;
    });

    strength
}

fn task2(file_name: &str) -> [char; 6*40]  {
   
    let mut image: [char;  6*40] = [' '; 6*40];
    let mut cycle = 0;

    decoder(file_name, |x: i32| {
        let pos = cycle % 40;

        if pos >= x - 1 && pos <= x+1 {
            image[(cycle % (6*40)) as usize] = 'X';
        } 

        cycle += 1;
    });
    image
}

fn decoder<F: FnMut(i32)>(file_name: &str, mut cycler: F) {
    let contents = fs::read_to_string(file_name).expect("Should have been able to read the file");

    let lines = contents.split("\n").collect::<Vec<&str>>();

    let mut x = 1;
    for line in lines {
        let toks: Vec<&str> = line.split(' ').collect();
        match toks[0] {
            "noop" => cycler(x),
            "addx" => {
                cycler(x);
                let delta = toks[1].parse::<i32>().unwrap();
                cycler(x);
                x += delta;
            }
            _ => panic!("opcode"),
        }
    }
}

fn main() {
    println!("{}", task1("input.txt"));

    let image = task2("input.txt");
    for i in 0..6 {
        println!( "{}", image[i*40..(i+1)*40].iter().collect::<String>())
    }
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test() {
        let count = task1("test.txt");

        assert_eq!(count, 13140);
    }
}
