static test_input: &str = ".#.#.#
...##.
#....#
..#...
#.#..#
####..";

fn read(input: &str) -> Vec<Vec<bool>> {
    let mut array = Vec::new();
    for line in input.split('\n') {
        let mut line_vec = Vec::new();
        for j in line.chars() {
            line_vec.push(j == '#')
        }
        array.push(line_vec)
    }

    array
}

fn display(a: &Vec<Vec<bool>>) {
    for line in a {
        for v in line {
            if *v {
                print!("#");
            } else {
                print!(".");
            }
        }
        println!();
    }
    println!();
}

fn get(a: &Vec<Vec<bool>>, i: i8, j: i8) -> i8 {
    if i >= 0 && i < a.len() as i8 && j >= 0 && j < a.len() as i8 {
        return if a[j as usize][i as usize] { 1 } else { 0 };
    }
    0
}

fn generation(a: &Vec<Vec<bool>>) -> Vec<Vec<bool>> {
    let mut ret = Vec::<Vec<bool>>::new();

    let L = a.len();
    for _ in 0..L {
        let mut line = Vec::new();
        for _ in 0..L {
            line.push(false);
        }
        ret.push(line);
    }
    for j in 0..L {
        for i in 0..L {
            let _i = i as i8;
            let _j = j as i8;
            let mut sum = get(&a, _i + 1, _j + 1);
            sum += get(&a, _i + 0, _j + 1);
            sum += get(&a, _i - 1, _j + 1);
            sum += get(&a, _i - 1, _j);
            sum += get(&a, _i + 1, _j);
            sum += get(&a, _i + 1, _j - 1);
            sum += get(&a, _i + 0, _j - 1);
            sum += get(&a, _i - 1, _j - 1);
            if a[j][i] {
                ret[j][i] = sum == 2 || sum == 3
            } else {
                ret[j][i] = sum == 3
            }
        }
    }
    ret
}

fn set_corner(a: &mut Vec<Vec<bool>>) {
    a[0][0] = true;
    let L = a.len();
    a[L-1][0] = true;
    a[0][L-1] = true;
    a[L-1][L-1] = true;
}

fn count(a: &Vec<Vec<bool>>) -> i32 {
    let mut sum = 0;
    let L = a.len();
    for j in 0..L {
        for i in 0..L {
            if a[j][i] {
                sum+=1;
            }
        }
    }
    sum
}

fn test_main() {
    let mut a = read(test_input);
    display(&a);
    for _ in 0..4 {
        a = generation(&a);
        display(&a);
    }
    println!("{}", count(&a));
}

fn test2_main() {
    let mut a = read(test_input);
    set_corner(&mut a);
    display(&a);
    for _ in 0..5 {
        a = generation(&a);
        set_corner(&mut a);
        display(&a);
    }
    println!("{}", count(&a));
}


fn main() {
    test_main();
    let input = include_str!("input.txt");
    let mut a = read(input);
    for _ in 0..100 {
        a = generation(&a);
    }
    println!("{}", count(&a));

    test2_main();

    let mut a = read(input);
    set_corner(&mut a);
    for _ in 0..100 {
        a = generation(&a);
        set_corner(&mut a);
    }
    println!("{}", count(&a));
}