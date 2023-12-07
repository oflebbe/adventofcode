use rs_07::*;

fn main() {
    let list = parse("input.txt");
    println!("{}", eval(&list, false));

    println!("{}", eval(&list, true));
}
