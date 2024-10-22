

fn main() {
	let numbers = vec![10, 20, 30, 40];
	for num in &numbers {
		match num {
			30 => println!("{:?}", "thirty"),
			_ => println!("{:?}", num)
		}
	}
	println!("{:?}", numbers.len());
	

}
