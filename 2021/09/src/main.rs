fn readin(st: &str) -> Option<Vec<Vec<i16>>> {
    let lines: Vec<&str> = st.split('\n').filter(|line| line.len() > 0).collect();
    Some(
        lines
            .iter()
            .map(|line| {
                line.chars()
                    .map(|ch| ch.to_digit(10).unwrap() as i16)
                    .collect()
            })
            .collect(),
    )
}

fn lowpoints(matrix: &Vec<Vec<i16>>) -> u32 {
    let mut sum = 0;
    for h in 0..matrix.len() {
        for w in 0..matrix[0].len() {
            let el = matrix[h][w];
            if h > 0 {
                if matrix[h - 1][w] < el {
                    continue;
                }
            }
            if h < matrix.len() - 1 {
                if matrix[h + 1][w] <= el {
                    continue;
                }
            }
            if w > 0 {
                if matrix[h][w - 1] <= el {
                    continue;
                }
            }
            if w < matrix[0].len() - 1 {
                if matrix[h][w + 1] <= el {
                    continue;
                }
            }
            sum += (el + 1) as u32
        }
    }

    sum
}

fn bassin(matrix: &mut Vec<Vec<i16>>) -> u32 {
    let mut next_bassin = -1;
    let mut found_something = true;
    while found_something {
        found_something = false;
        for h in 0..matrix.len() {
            for w in 0..matrix[0].len() {
                let el = matrix[h][w];
                if el == 9 {
                    continue
                }
                // already in bassin
                if el < 0 {
                    continue
                }
                if h > 0 {
                    if matrix[h - 1][w] < 0 {
                        found_something = true;
                        matrix[h][w] = matrix[h - 1][w];
                        continue;
                    }
                }
                if h < matrix.len() - 1 {
                    if matrix[h + 1][w] < 0 {
                        found_something = true;
                        matrix[h][w] = matrix[h + 1][w];
                        continue;
                    }
                }
                if w > 0 {
                    if matrix[h][w - 1] <  0 {
                        found_something = true;
                        matrix[h][w] = matrix[h][w-1];
                        continue;
                    }
                }
                if w < matrix[0].len() - 1 {
                    if matrix[h][w + 1] < 0 {
                        found_something = true;
                        matrix[h][w] = matrix[h][w+1];
                        continue;
                    }
                }
                
            }
        }
        if found_something {
            // bassin was expanded
            continue;
        }
        // search for new bassin
        found_something = false;
        for h in 0..matrix.len() {
            for w in 0..matrix[0].len() {
                let el = matrix[h][w];
                if el >= 0 && el < 9 {
                    matrix[h][w] = next_bassin;
                    next_bassin -= 1;
                    found_something = true;
                    break;
                }
            }
            if found_something {
                break;
            }
        }
    }
    let mut result = vec![0;-next_bassin as usize];
    for h in 0..matrix.len() {
        for w in 0..matrix[0].len() {
            let el = matrix[h][w];
            if el < 0 {
                result[(-el-1) as usize]+=1;
            } 
        }
    }
    result.sort();
    result.reverse();
    
    
    result.into_iter().take(3).reduce(|x, y| x*y).unwrap()
}

fn main() {
    let mut var = readin(include_str!("input.txt")).unwrap();
    println!("{}", lowpoints(&var));
    println!("{}", bassin(&mut var));
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_part1() {
        let var = include_str!("example.txt");
        if let Some(matrix) = readin(var) {
            assert_eq!(15, lowpoints(&matrix))
        } else {
            assert_eq!(0, 1)
        }
    }

    #[test]
    fn test_part2() {
        let var = include_str!("example.txt");
        if let Some(mut matrix) = readin(var) {
            assert_eq!(1134, bassin(&mut matrix))
        } else {
            assert_eq!(0, 1)
        }
    }
}
