use std::io;

fn main() {
    let mut input_text = String::new();
    io::stdin()
        .read_line(&mut input_text)
        .expect("failed to read from stdin");

    let trimmed = input_text.trim();
    match trimmed.parse::<bool>() {
        Ok(val) => {
		if val {
			println!("Hello");
		} else {
			println!("GoodBye");
		}
	},
        Err(..) => println!("this was not a boolean value: {}", trimmed),
    };
}
