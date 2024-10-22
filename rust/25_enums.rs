

enum Color {
	Red,
	Blue,
	White
}

fn detect_color(c: Color) {
	match c {
		Color::Red => println!("red"),
		Color::Blue=> println!("blue"),
		Color::White => println!("white"),
	}
}

fn main() {
	let c = Color::Red;
	detect_color(c);
}
