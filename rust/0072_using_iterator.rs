fn main() {
	let numbers = vec![1, 2, 3, 4];
	let numbers_plus_one: Vec<String> = numbers.iter().map(|n| (n+1).to_string() + " hello").collect();

	for x in numbers_plus_one {
		println!("{:?}", x);
	}

}
