use std::fs::File;
use std::io::{BufWriter, Write};
use std::sync::{Arc, Mutex};
use std::thread;
use std::time::Instant;

pub fn calculate_matrix(matrix_a: &[Vec<i32>], matrix_b: &[Vec<i32>], num_workers: usize) -> Result<Vec<Vec<i32>>, String> {
    if let Some(err) = pre_check(matrix_a, matrix_b, num_workers) {
        return Err(err);
    }

    let (start_time, mut log_file) = create_log_file()?;
    let result = initialize_result_matrix(matrix_a, matrix_b);

    perform_parallel_multiplication(matrix_a, matrix_b, num_workers, &result);

    let duration = start_time.elapsed();
    log_calculation_time(&mut log_file, start_time, duration)?;

    Ok(Arc::try_unwrap(result).unwrap().into_inner().unwrap())
}

fn pre_check(matrix_a: &[Vec<i32>], matrix_b: &[Vec<i32>], num_workers: usize) -> Option<String> {
    if num_workers <= 0 {
        return Some("Number of workers must be greater than zero".to_string());
    }

    if matrix_a.is_empty() || matrix_b.is_empty() {
        return Some("Matrix is empty".to_string());
    }

    if matrix_a[0].len() != matrix_b.len() {
        return Some("Matrix multiplication is not possible due to dimension mismatch".to_string());
    }

    None
}

fn create_log_file() -> Result<(Instant, BufWriter<File>), String> {
    std::fs::create_dir_all("log").map_err(|e| e.to_string())?;
    let log_file = File::create("log/calc_time_log").map_err(|e| e.to_string())?;
    Ok((Instant::now(), BufWriter::new(log_file)))
}

fn initialize_result_matrix(matrix_a: &[Vec<i32>], matrix_b: &[Vec<i32>]) -> Arc<Mutex<Vec<Vec<i32>>>> {
    Arc::new(Mutex::new(vec![vec![0; matrix_b[0].len()]; matrix_a.len()]))
}

fn perform_parallel_multiplication(matrix_a: &[Vec<i32>], matrix_b: &[Vec<i32>], num_workers: usize, result: &Arc<Mutex<Vec<Vec<i32>>>>) {
    let chunk_size = matrix_a.len() / num_workers;
    let mut handles = vec![];

    for i in 0..num_workers {
        let result = Arc::clone(result);
        let matrix_a = matrix_a.to_vec();
        let matrix_b = matrix_b.to_vec();
        let handle = thread::spawn(move || {
            let start_row = i * chunk_size;
            let end_row = if i == num_workers - 1 {
                matrix_a.len()
            } else {
                start_row + chunk_size
            };

            for row in start_row..end_row {
                for col in 0..matrix_b[0].len() {
                    for k in 0..matrix_b.len() {
                        result.lock().unwrap()[row][col] += matrix_a[row][k] * matrix_b[k][col];
                    }
                }
            }
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }
}

fn log_calculation_time(log_file: &mut BufWriter<File>, start_time: Instant, duration: std::time::Duration) -> Result<(), String> {
    writeln!(log_file, "Matrix multiplication Start time: {:?}", start_time).map_err(|e| e.to_string())?;
    writeln!(log_file, "Matrix multiplication End time: {:?}", Instant::now()).map_err(|e| e.to_string())?;
    writeln!(log_file, "Matrix multiplication duration time: {:?}", duration).map_err(|e| e.to_string())?;
    Ok(())
}