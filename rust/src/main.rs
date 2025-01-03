use main::mods::file_mod::{generate_matrix_to_file, read_matrix_from_file, MatrixSize, LOG_DIR, RESOURCES_DIR};
use std::io::{self, Write};
use std::fs;

fn get_matrix_size_from_user(matrix_name: &str) -> MatrixSize {
    let mut input = String::new();
    loop {
        print!("Enter size for {} matrix (small, middle, big): ", matrix_name);
        io::stdout().flush().expect("Failed to flush stdout");
        io::stdin().read_line(&mut input).expect("Failed to read line");
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

fn get_num_workers_from_user() -> usize {
    let mut input = String::new();
    loop {
        print!("How many logical CPU cores do you have? ");
        io::stdout().flush().expect("Failed to flush stdout");
        io::stdin().read_line(&mut input).expect("Failed to read line");
        match input.trim().parse::<usize>() {
            Ok(num) if num > 0 => return num,
            _ => println!("Invalid input. Please enter a positive integer."),
        }
        input.clear();
    }
}

fn setup_directories() -> io::Result<()> {
    fs::create_dir_all(LOG_DIR)?;
    fs::create_dir_all(RESOURCES_DIR)?;
    Ok(())
}

fn generate_and_read_matrices(matrix_size_a: usize, matrix_size_b: usize) -> io::Result<(Vec<Vec<i32>>, Vec<Vec<i32>>)> {
    let source_matrix_a = format!("{}/matrix_a.txt", RESOURCES_DIR);
    let source_matrix_b = format!("{}/matrix_b.txt", RESOURCES_DIR);

    generate_matrix_to_file(&source_matrix_a, matrix_size_a)?;
    generate_matrix_to_file(&source_matrix_b, matrix_size_b)?;

    let matrix_a = read_matrix_from_file(&source_matrix_a)?;
    let matrix_b = read_matrix_from_file(&source_matrix_b)?;

    Ok((matrix_a, matrix_b))
}

fn main() -> io::Result<()> {
    let matrix_size_a = get_matrix_size_from_user("Matrix A") as usize;
    let matrix_size_b = get_matrix_size_from_user("Matrix B") as usize;
    let num_workers = get_num_workers_from_user();

    setup_directories()?;

    let (matrix_a, matrix_b) = generate_and_read_matrices(matrix_size_a, matrix_size_b)?;

    match main::mods::calculation_mod::calculate_matrix(&matrix_a, &matrix_b, num_workers) {
        Ok(_result) => {
            println!("Matrix multiplication succeeded.");
        }
        Err(e) => {
            println!("Error calculating matrix: {}", e);
        }
    }

    Ok(())
}