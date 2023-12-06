use std::fs;

fn f1( field : & mut[i32;7], h : i32, mut pos: i32, c: char) -> (bool, i32) {

    let p = pos as usize;
    
    if h <= field[p] ||h <= field[p+1] || h <= field[p+2]  ||h <= field[p+3]  {
        field[p] = h+1;
        field[p+1] = h+1;
        field[p+2] = h+1;
        field[p+3] = h+1;
        return (true, pos)
    }
    if c == '>' && pos < 3 {
        pos +=1;
    }
    if c == '<' && pos > 0 {
        pos -= 1;
    }
    return  (false, pos)
}

fn f2( field : & mut[i32;7], h : i32, mut pos: i32, c: char) -> (bool, i32) {

    let p = pos as usize;
    
    if h+1 <= field[p] ||h <= field[p+1] || h+1 <= field[p+2]  {
        field[p] = h+2;
        field[p+1] = h+3;
        field[p+2] = h+2;

        return (true, pos)
    }
    if c == '>' && pos < 4 {
        pos +=1;
    }
    if c == '<' && pos > 0 {
        pos -= 1;
    }
    return  (false, pos)
}

fn f3( field : & mut[i32;7], h : i32, mut pos: i32, c: char) -> (bool, i32) {

    let p = pos as usize;
    
    if h <= field[p] ||h <= field[p+1] || h <= field[p+2]  {
        field[p] = h+1;
        field[p+1] = h+1;
        field[p+2] = h+3;

        return (true, pos)
    }
    if c == '>' && pos < 4 {
        pos +=1;
    }
    if c == '<' && pos > 0 {
        pos -= 1;
    }
    return  (false, pos)
}

fn f4( field : & mut[i32;7], h : i32, mut pos: i32, c: char) -> (bool, i32) {

    let p = pos as usize;
    
    if h <= field[p]  {
        field[p] = h+4;
        
        return (true, pos)
    }
    if c == '>' && pos < 5 {
        pos +=1;
    }
    if c == '<' && pos > 0 {
        pos -= 1;
    }
    return  (false, pos)
}


fn f5( field : & mut[i32;7], h : i32, mut pos: i32, c: char) -> (bool, i32) {

    let p = pos as usize;
    
    if h <= field[p] || h <= field[p+1]   {
        field[p] = h+2;
        field[p+1] = h+2;
        
        return (true, pos)
    }
    if c == '>' && pos < 4 {
        pos +=1;
    }
    if c == '<' && pos > 0 {
        pos -= 1;
    }
    return  (false, pos)
}


fn part1(file_name: &str) -> i32{
    let contents = fs::read_to_string(file_name).expect("Should have been able to read the file");
    let wind : Vec<char> = contents.chars().collect();
    let mut funcs: Vec<fn (&mut [i32;7], i32, i32, char) -> (bool, i32)> = Vec::new();
    funcs.push(f1);
    funcs.push(f2);
    funcs.push(f3);
    funcs.push(f4);
    funcs.push(f5);
    let mut count_func = 0;
    let mut ground = [0; 7];
    let mut count_wind = 0;
    for _ in 0..2022 {
        let high = ground.iter().max().unwrap();
        let mut pos = 2;
        for level in (0..=high+4).rev() {

           let  (end, npos) = funcs[count_func % 5]( &mut ground, level, pos, wind[ count_wind % wind.len()] );
            if end {
                break
            }
            pos = npos;
            count_wind+=1;

        }
        count_func+=1;
    }
   *ground.iter().max().unwrap()
}

fn main() {
    println!( "{}", part1("test.txt"));
}