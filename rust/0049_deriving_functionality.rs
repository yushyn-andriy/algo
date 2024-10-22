#[derive(Debug, Copy, Clone)]
enum Role {
	Guest,
	User,
	Admin,
	Sudo,
}

#[derive(Debug, Copy, Clone)]
struct User {
	role: Role,
}

fn print(u: User) {
	println!("{:?}", u);
}


fn main() {
	let u = User{role: Role::Sudo};
	println!("{:?}", u);
	print(u);
	print(u);
}


