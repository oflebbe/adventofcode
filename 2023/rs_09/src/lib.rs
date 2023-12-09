use std::fs::read_to_string;

pub fn parse(file_name: &str) -> Vec<Vec<i32>> {
    let mut res = Vec::new();

    let contents = read_to_string(file_name).expect("Should have been able to read the file");

    for line in contents.split("\n") {
        let mut vlen = Vec::new();
        for num in line.split(" ") {
            let number: i32 = num.parse().unwrap();
            vlen.push(number);
        }
        res.push(vlen);
    }

    res
}

pub fn differences(numbers: &Vec<i32>) -> (Vec<i32>, Vec<i32>) {
    let mut start = Vec::new();
    let mut work = numbers.clone();
    loop {
        start.push(work[0]);
        let mut tmp = Vec::new();
        let mut is_zero = true;
        for i in 0..work.len() - 1 {
            let diff = work[i + 1] - work[i];
            is_zero = is_zero && (diff == 0);
            tmp.push(diff);
        }
        work = tmp;
        if is_zero {
            break;
        }
    }
    (work, start)
}

pub fn predict_1(numbers: &Vec<i32>) -> i32 {
    let (mut work, start) = differences(numbers);
    work.push(0);

    for c in (0..start.len()).rev() {
        let mut tmp = Vec::new();
        tmp.push(start[c]);
        let mut last = start[c];
        for i in 1..=work.len() {
            last += work[i - 1];
            tmp.push(last);
        }
        work = tmp;
    }

    work[work.len() - 1]
}

pub fn predict_back(numbers: &Vec<i32>) -> i32 {
    let (mut work, start) = differences(numbers);
    work.push(0);

    for c in (0..start.len()).rev() {
        let mut tmp = Vec::new();
        let mut last = start[c] - work[0];
        tmp.push(last);
        for i in 1..=work.len() {
            last += work[i - 1];
            tmp.push(last);
        }
        work = tmp;
    }

    work[0]
}

pub fn star1(input: &Vec<Vec<i32>>) -> i32 {
    let mut sum = 0;
    for row in input {
        sum += predict_1(row);
    }
    sum
}

pub fn star2(input: &Vec<Vec<i32>>) -> i32 {
    let mut sum = 0;
    for row in input {
        sum += predict_back(row);
    }
    sum
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test1() {
        let input = parse("test.txt");
        assert_eq!(predict_1(&input[0]), 18);

        assert_eq!(predict_back(&input[2]), 5);
    }
}
