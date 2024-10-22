

fn add(a: i32, b:i32) -> i32 {
	a + b 
}

#[derive(Debug)]
enum Light {
	Red,
}

#[derive(Debug)]
struct Person {
	age:i32,
}

fn display_light(l: &Light) {
	println!("{:?}", l);
}

fn incr_age(p: &mut Person) {
	p.age = p.age + 1;
	println!("{:?}", p);
}

fn main() {
	let a:i32 = 1;
	let b:i32 = 2;
	add(a, b);
	add(a, b);
	let l = Light::Red;
	display_light(&l);
	display_light(&l);

	let mut p = Person{age:33};
	incr_age(&p);
	incr_age(&p);
	incr_age(&p);
}
