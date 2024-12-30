use std::fs::File;
use std::io::{BufWriter, Write};
use std::sync::{Arc, Mutex};
use std::thread;
use std::time::Instant;

pub fn calculate_matrix(matrix_a: &[Vec<i32>], matrix_b: &[Vec<i32>], num_workers: usize) -> Result<Vec<Vec<i32>>, String> {
    // Check if matrix multiplication is possible
    if matrix_a[0].len() != matrix_b.len() {
        println!("Matrix multiplication is not possible due to dimension mismatch");
        return Err("Matrix multiplication is not possible due to dimension mismatch".to_string());
    }

    let start_time = Instant::now();

    // Use Arc and Mutex to share the result between threads
    let result = Arc::new(Mutex::new(vec![vec![0; matrix_b[0].len()]; matrix_a.len()]));
    let chunk_size = matrix_a.len() / num_workers;

    let mut handles = vec![];

    // Create threads to calculate the result matrix
    for i in 0..num_workers {
        let result = Arc::clone(&result);
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

    let duration = start_time.elapsed();
    println!("Matrix multiplication completed.");
    println!("Matrix multiplication time: {:?}", duration);

    let mut log_file = BufWriter::new(File::create("log/calc_time_log").map_err(|e| e.to_string())?);
    writeln!(log_file, "Matrix multiplication Start time: {:?}", start_time).map_err(|e| e.to_string())?;
    writeln!(log_file, "Matrix multiplication End time: {:?}", Instant::now()).map_err(|e| e.to_string())?;
    writeln!(log_file, "Matrix multiplication duration time: {:?}", duration).map_err(|e| e.to_string())?;

    Ok(Arc::try_unwrap(result).unwrap().into_inner().unwrap())
}