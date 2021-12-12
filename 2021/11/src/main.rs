#![no_std]

use collections;

fn readin(st: &str) -> Vec<Vec<u8>> {
    let lines: Vec<&str> = st.split('\n').filter(|line| line.len() > 0).collect();

    let mut res: Vec<Vec<_>> = lines
        .iter()
        .map(|line| {
            line.chars()
                .map(|ch| ch.to_digit(10).unwrap() as u8)
                .collect()
        })
        .collect();
    let w = res[0].len();
    let h = res.len();
    res.insert(0, vec![0; w]);
    res.push(vec![0; w]);
    for i in 0..h + 2 {
        res[i].insert(0, 0);
        res[i].push(0);
    }
    res
}

fn step(matrix: &mut Vec<Vec<u8>>) -> usize {
    let mut a: Vec<Vec<_>> = matrix
        .iter()
        .map(|x| x.iter().map(|x| x + 1).collect())
        .collect();
    loop {
        let mut flash = false;
        for h in 1..a.len()-1 {
            for w in 1..a[0].len()-1 {
                if a[h][w] > 9 && a[h][w] < 100 {
                    flash = true;
                    a[h - 1][w - 1] += 1;
                    a[h][w - 1] += 1;
                    a[h + 1][w - 1] += 1;
                    a[h - 1][w] += 1;
                    a[h + 1][w] += 1;
                    a[h - 1][w + 1] += 1;
                    a[h][w + 1] += 1;
                    a[h + 1][w + 1] += 1;
                    a[h][w] = 100 // flashed
                }
            }
        }
        if !flash {
            break;
        }
    }
    for h in 1..a.len()-1 {
        for w in 1..a[0].len()-1 {
            if a[h][w] > 9 {
                matrix[h][w] = 0
            } else {
                matrix[h][w] = a[h][w];
            }
        }
    }
    a.into_iter().flatten().filter(|&x| x >= 100).count()
}

fn main() {
    let mut matrix = readin(include_str!("input.txt"));
    let mut sum = 0;
    for i in 0..100 {
        sum+=step( &mut matrix)
    }
    println!("{}", sum);
let mut count = 100;
    let mut number = 0;
    while number != (matrix.len() -2 )* (matrix.len()-2) {
        count+=1;
        number = step( &mut matrix)
    }
    println!("{}", count)
    
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_part1() {
        let var = "11111
19991
19191
19991
11111";
        let mut matrix = readin(var);
        assert_eq!(9, step(&mut matrix));
    }

    #[test]
    fn test_part1_2() {
        let var = "5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526";
        let mut sum = 0;
        let mut matrix = readin(var);
        for i in 0..10 {
            sum+= step( &mut matrix);
        }
        assert_eq!(204, sum);

    }
}
