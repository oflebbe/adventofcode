use regex::Regex;
use std::collections::HashMap;
use std::fs;

fn new_pos(headx: i32, heady: i32, mut tailx: i32, mut taily: i32) -> (i32, i32) {
    let drow = (headx - tailx).signum();
    let dcol = (heady - taily).signum();

    if (headx - tailx).abs() + (heady - taily).abs() > 2 {
        if (headx - tailx).abs() > 0 {
            tailx += drow;
        }
        if (heady - taily).abs() > 0 {
            taily += dcol;
        }
    } else {
        if (headx - tailx).abs() > 1 {
            tailx += drow;
        }
        if (heady - taily).abs() > 1 {
            taily += dcol;
        }
    }

    (tailx, taily)
}

fn task1(file_name: &str, len: i32) -> i32 {
    let contents = fs::read_to_string(file_name).expect("Should have been able to read the file");

    let lines = contents.split("\n").collect::<Vec<&str>>();

    let mut grid = HashMap::new();

    let re = Regex::new(r"^(.) (\d*)").unwrap();
    let mut tail = Vec::new();

    for _i in 0..len {
        tail.push((0, 0));
    }
    for line in lines {
        let cap = re.captures(line).unwrap();

        let (dx, dy) = match cap.get(1).unwrap().as_str() {
            "L" => (-1, 0),
            "R" => (1, 0),
            "U" => (0, -1),
            "D" => (0, 1),
            _ => panic!("instruction"),
        };
        let steps = cap.get(2).unwrap().as_str().parse::<i32>().unwrap();
        for _i in 0..steps {
            tail[0].0 += dx;
            tail[0].1 += dy;

            for i in 1..len as usize {
                tail[i] = new_pos(tail[i - 1].0, tail[i - 1].1, tail[i].0, tail[i].1);
            }

            let key = format!(
                "{}_{}",
                tail[(len - 1) as usize].0,
                tail[(len - 1) as usize].1 + 1
            );
            // println!("{} {}", &key, &line);
            grid.insert(key, 1);
        }
    }
    grid.len() as i32
}

fn main() {
    let count = task1("input.txt", 2);

    println!("{}", count);

    let count = task1("input.txt", 10);

    println!("{}", count);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test() {
        let count = task1("test.txt", 2);

        assert_eq!(count, 13);

        let count = task1("test.txt", 10);

        assert_eq!(count, 1);

        let count = task1("larger.txt", 10);

        assert_eq!(count, 36);
    }

    #[test]
    fn test_pos() {
        assert_eq!(new_pos(0, 0, 0, 0), (0, 0));

        assert_eq!(new_pos(1, 0, 0, 0), (0, 0));

        assert_eq!(new_pos(2, 0, 0, 0), (1, 0));

        assert_eq!(new_pos(-1, 0, 0, 0), (0, 0));

        assert_eq!(new_pos(-2, 0, 0, 0), (-1, 0));

        assert_eq!(new_pos(0, -2, 0, 0), (0, -1));
        assert_eq!(new_pos(1, 1, 0, 0), (0, 0));
        assert_eq!(new_pos(2, 1, 0, 0), (1, 1));
        assert_eq!(new_pos(1, 2, 0, 0), (1, 1));
    }
}
