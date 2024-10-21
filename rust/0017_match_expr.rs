fn main() {
	let val = true;
	match val {
		true => println!("Its true"),
		false => println!("Its false")
	}

	let some_int = 1;
	match some_int {
		1 => println!("1"),
		_ => println!("something else"),
	}
}
