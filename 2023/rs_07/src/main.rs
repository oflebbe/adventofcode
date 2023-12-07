use rs_07::*;

fn main() {
    let list = parse("input.txt", false);
    println!("{}", eval(&list));

    let list = parse("input.txt", true);
    println!("{}", eval(&list));
}
