


fn add(a:i32, b:i32) -> i32 {
	a + b
}

fn display_result(a:i32, b:i32, sum:i32) {
	println!("{a:?} + {b:?} is {sum:?}", a=a, b=b, sum=sum);
}

fn main() {
	let a = 2;
	let b = 3;
	let x = add(a, b);
	display_result(a, b, x);
}
