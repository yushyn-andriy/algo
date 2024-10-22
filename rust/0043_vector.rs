

fn main() {

	let array = vec![1, 2, 3];
	let mut numbers = Vec::new();

	numbers.push(1);
	numbers.push(2);
	numbers.push(3);
	numbers.pop();
	numbers.len();
	
	let two = numbers[1];
	for num in numbers {
		println!("{:?}", num);
	}	

}
