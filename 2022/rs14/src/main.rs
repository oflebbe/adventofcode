use std::{collections::HashMap, fs, sync::Arc};

enum FieldType {
    Rock,
    Sand
}
fn load(file_name: &str) -> HashMap<u16, HashMap<u16, FieldType>> {
    let contents = fs::read_to_string(file_name).expect(
        "Should have been able t
o read the file",
    );
    let mut field = HashMap::<u16, HashMap<u16, FieldType>>::new();
    let lines = contents.split("\n").collect::<Vec<&str>>();

    for &line in &lines {
        let pairs: Vec<&str> = line.split(" -> ").collect();
        let mut first = true;
        let mut oldx = 0;
        let mut oldy = 0;
        for p in &pairs {
            let tok : Vec<&str> = p.split(",").collect();
            let x = tok[0].parse::<u32>().unwrap();
            let y = tok[1].parse::<u32>().unwrap();
            if first {
                oldx = x;
                oldy = y;
                first = false;
                continue;
            }
            if x == oldx {
                for py in std::cmp::min( y, oldy)..=std::cmp::max(y, oldy) {
                    if let Some(v) = field.get_mut( &(x as u16)) {
                        v.insert( py as u16, FieldType::Rock);
                    } else {
                        let mut h = HashMap::new();
                        h.insert( py as u16, FieldType::Rock);
                        field.insert(x as u16, h);
                    }
                }
            } else if y == oldy {
                for px in std::cmp::min( x, oldx)..=std::cmp::max(x, oldx) {
                    if let Some(v) = field.get_mut( &(px as u16)) {
                        v.insert( y as u16, FieldType::Rock);
                    } else {
                        let mut h = HashMap::new();
                        h.insert( y as u16, FieldType::Rock);
                        field.insert(px as u16, h);
                    }
                }
            }
        }
    }

    field
}

enum SandResult {
    YES,
    NO,
    END
}

// return true if Sand was able to fill in
fn insert_sand_if_empty( field: &mut HashMap<u16, HashMap<u16, FieldType>>, p: (u16,u16) ) -> SandResult {
    if let Some( col) = field.get(&p.0) {
        if let Some(_) = col.get(&p.1) {
            return SandResult::NO;
        }
        col.insert( p.1, FieldType::Sand);
        return SandResult::YES;
    }
    SandResult::END
    
}


fn sand( field: & mut HashMap<u16, HashMap<u16, FieldType>> ) -> int {
  let mut count = 1;
  let mut coord = ( 500, 0);

  loop {
    match insert_sand_if_empty(field, (coord.0, coord+1)) {
        SandResult::NO => match insert_sand_if_empty(field, (coord.0-1, coord+1)) {
            SandResult::NO => match insert_sand_if_empty(field, (coord.0+1, coord+1)) {
                SandResult::NO => match insert_sand_if_empty(field, (coord.0, coord)) {
                    SandResult::YES => (),
                    _ => panic!()
                }
                SandResult::END => return count,
                SandResult::YES => (),
            },
            SandResult::YES => (),
            SandResult::END => return count,
        }
        SandResult::YES => (),
        SandResult::END => return count,
    }
    count += 1;
    

                    
        }


    }
        if ! {
            if !insert_sand_if_empty(field, (coord.0+1, coord+1)) {

    }
        if let Some(_) = col0.get( &(coord.1 + 1)) {
            // Oh there uis already smthg
            // try left
            if let Some(col_l) = field.get( &(coord.0 - 1)) {
                if let Some(_) = col_l.get( &(coord.1 + 1)) {



                    col_l.insert( coord.1+1, FieldType::Sand);
                    continue;
                } else {


                    if let Some(col_m) = field.get( &(coord.0 + 1)) {
                        if let Some(_) = col_m.get( &(coord.1 + 1)) {
                            col_m.insert( coord.1+1, FieldType::Sand);
                            continue;
                        } else {
                            col.insert( coord.1, FieldType::Sand);
                        }
                    } else {

                        field.insert( coord.0+1 )
                    }
                }
        }

    } else {
        return count;
    }

  }
}
fn main() {
    println!("Hello, world!");
}
