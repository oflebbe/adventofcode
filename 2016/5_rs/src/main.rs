use md5;

fn problem( input: &str) -> String {
    let mut count : i64 = 0;
    let mut passwd = String::from("");
    loop {
        let st = format!( "{}{}", input, count);
        let digest = md5::compute(st);
        let dig = format!("{:x}", digest);
        if dig.starts_with("00000") {
            let six = dig.chars().nth(5).unwrap();
            passwd.push( six);
            if passwd.len() == 8 {
                return passwd
            }
        }
        count+=1;
    }
}

fn main() {
    println!( "{}", problem("abc"));
}
