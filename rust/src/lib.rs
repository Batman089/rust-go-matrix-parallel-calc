pub mod mods {
    pub mod file_mod;
    pub mod calculation_mod;
}

pub use mods::file_mod::{generate_matrix_to_file, read_matrix_from_file, MatrixSize};