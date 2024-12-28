mod mods{
    pub mod file_mod;
    pub mod calculation_mod;
}

use mods::file_mod::{generate_matrix_to_file, read_matrix_from_file, MatrixSize};
use std::io::{self, Write};
use std::fs;

fn get_matrix_size_from_user(matrix_name: &str) -> MatrixSize {
    let mut input = String::new();
    loop {
        print!("Enter size for {} matrix (small, middle, big): ", matrix_name);
        io::stdout().flush().unwrap();
        io::stdin().read_line(&mut input).unwrap();
        input = input.trim().to_lowercase();

        match input.as_str() {
            "small" => return MatrixSize::Small,
            "middle" => return MatrixSize::Middle,
            "big" => return MatrixSize::Big,
            _ => println!("Invalid input. Please enter 'small', 'middle', or 'big'."),
        }
        input.clear();
    }
}

fn main() -> io::Result<()> {
    let matrix_size_a = get_matrix_size_from_user("Matrix A") as usize;
    let matrix_size_b = get_matrix_size_from_user("Matrix B") as usize;

    fs::create_dir_all("log")?;
    fs::create_dir_all("resources")?;

    let source_matrix_a = "resources/matrix_a.txt";
    let source_matrix_b = "resources/matrix_b.txt";

    generate_matrix_to_file(source_matrix_a, matrix_size_a)?;
    generate_matrix_to_file(source_matrix_b, matrix_size_b)?;

    let matrix_a = read_matrix_from_file(source_matrix_a)?;
    let matrix_b = read_matrix_from_file(source_matrix_b)?;

    match mods::calculation_mod::calculate_matrix(&matrix_a, &matrix_b) {
        Ok(result) => {}
        Err(e) => {
            println!("Error calculating matrix: {}", e);
        }
    }

    Ok(())
}