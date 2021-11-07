fn max( o: usize) -> usize{
    if o == 2 {
        return 2;
    }
    o +1 
}

fn min( o: usize) -> usize{
    if o == 0 {
        return 0;
    }
    o - 1
}

fn max5( o: usize) -> usize{
    if o == 4 {
        return 4;
    }
    o + 1
}


fn parse( p : [usize;2] , st : &str) -> [usize;2]  {
    let mut r = p;
    for ch in st.chars() {
        match ch {
            'U' => r[1] = min( r[1]),
            'D' => r[1] = max( r[1]),
            'R' => r[0] = max( r[0]),
            'L' => r[0] = min( r[0]),
            _ => println!("ERROR")
        }
    }
    r
}


fn parse5( p : [usize;2] , grid: [[char; 5]; 5], st : &str) -> [usize;2]  {
    let mut r = p;
    let mut s = r.clone();
    for ch in st.chars() {
        match ch {
            'U' => r[0] = min( r[0]),
            'D' => r[0] = max5( r[0]),
            'R' => r[1] = max5( r[1]),
            'L' => r[1] = min( r[1]),
            _ => println!("ERROR")
        }
        r = if grid[r[0]][r[1]] == ' ' {
            s
        } else {
            r
        };
        s = r.clone()
    }
    r
}
fn populate_grid( st: &str) -> [[char;5];5] {
    let mut grid: [[char;5];5] = [[' '; 5]; 5] ;

    let mut i = 0;
    for l in st.lines() {
        let mut j = 0;
        for k in l.chars() {
            grid[i][j] = k;
            j+=1;
        }
        i += 1;
    }
    grid
}

fn main() {

    let grid: [[u8; 3]; 3] = [ [1,2,3],[4,5,6],[7,8,9]];
    let mut pointer : [usize;2] = [1,1];

    let input = "ULL\n\
    RRDDD\n\
    LURDL\n\
    UUUUD";

    for line in input.lines() {
        pointer = parse( pointer, line);
        print!("{}", grid[pointer[1]][pointer[0]]);
    }
    println!();

    let my_str = include_str!("input.txt");
    for line in my_str.lines() {
        pointer = parse( pointer, line);
        print!("{}", grid[pointer[1]][pointer[0]]);
    }
    println!();
    let pad_layout = "  1  \n 234 \n56789\n ABC \n  D  ";
    let grid = populate_grid(pad_layout);
    let mut cursor : [usize; 2] = [2, 0];
    
    for line in input.lines() {
        cursor = parse5( cursor, grid, line);
        print!("{}", grid[cursor[0]][cursor[1]]);
    }
    println!();

    for line in my_str.lines() {
        cursor = parse5( cursor, grid, line);
        print!("{}", grid[cursor[0]][cursor[1]]);
    }
    println!();
}
