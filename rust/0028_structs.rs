

enum Color {
	White,
	Black,
}

enum Material {
	Aluminium,
	Iron,
	Nylon,
}

struct Bolt {
	length: f32,
	diameter: f32,
	material: Material,
	color: Color	
}

fn get_mass(b: Bolt) -> f32{
	(b.diameter * b.diameter * 3.1415 / 4.0) * b.length
}

fn main() {
	let mut bolt = Bolt{
		length: 200.0,
		diameter: 140.0,
		material: Material::Nylon,
		color: Color::White,
	};
	println!("{:?}", get_mass(bolt));
}
