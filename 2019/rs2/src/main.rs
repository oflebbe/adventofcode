


fn main() {
    #[allow(dead_code)]
    let fan = "input";
    
    let i : i32 ;
    i = 2;

    println!("Hello, world! {}", i);

    println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");
    println!("{} of {:b} people know binary, the other half doesn't", 1, 2);


    //#[derive(Debug)]
    struct Structure(i32);

    // However, custom types such as this structure require more complicated
    // handling. This will not work.
    println!("This struct `{:?}` won't print...", Structure(3));
}
