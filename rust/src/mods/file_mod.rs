use rand::Rng;
use std::fs::File;
use std::io::{self, BufRead, BufReader, BufWriter, Write};
use std::time::Instant;

pub enum MatrixSize {
    Small = 1000,
    Middle = 5000,
    Big = 10000,
}

pub fn generate_matrix_to_file(filename: &str, size: usize) -> io::Result<()> {
    let start_time = Instant::now();
    let mut file = create_file(filename)?;
    write_matrix_to_file(&mut file, size)?;
    log_generation_time(start_time)?;
    Ok(())
}

fn create_file(filename: &str) -> io::Result<BufWriter<File>> {
    let file = BufWriter::new(File::create(filename)?);
    Ok(file)
}

fn write_matrix_to_file(file: &mut BufWriter<File>, size: usize) -> io::Result<()> {
    let mut rng = rand::thread_rng();
    for _ in 0..size {
        let row: Vec<String> = (0..size).map(|_| rng.gen_range(0..100).to_string()).collect();
        writeln!(file, "{}", row.join(" "))?;
    }
    file.flush()?;
    Ok(())
}

fn log_generation_time(start_time: Instant) -> io::Result<()> {
    let duration = start_time.elapsed();
    println!("Matrix generation time: {:?}", duration);

    let mut log_file = BufWriter::new(File::create("./generated/log/generate_matrix_files_log.txt")?);
    writeln!(log_file, "Matrix generation Start time: {:?}", start_time)?;
    writeln!(log_file, "Matrix generation End time: {:?}", Instant::now())?;
    writeln!(log_file, "Matrix generation duration time: {:?}", duration)?;
    log_file.flush()?;
    Ok(())
}

pub fn read_matrix_from_file(filename: &str) -> io::Result<Vec<Vec<i32>>> {
    let file = BufReader::new(File::open(filename)?);
    let mut matrix = Vec::new();

    for line in file.lines() {
        let line = line?;
        let row: Vec<i32> = line.split_whitespace().map(|s| s.parse().unwrap()).collect();
        matrix.push(row);
    }

    Ok(matrix)
}