enum Access {
	Full,
}

fn get_tuple() -> (i32, i32, i32) {
	(1, 2, 3)
}

fn main() {
	let numbers = get_tuple();
	let (a, b, c) = get_tuple();
	println!("{:?}", numbers.0);
	println!("{:?}", numbers.1);
	println!("{:?}", numbers.2);
		
	let (employe, access) = ("John", Access::Full);
}
