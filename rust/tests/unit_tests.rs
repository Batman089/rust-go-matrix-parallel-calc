#[cfg(test)]
mod tests {
    use main::mods::{calculation_mod, file_mod};
    use std::fs;

    #[test]
    fn test_small_matrix_multiplied_by_big_matrix_with_valid_worker_number() {
        let matrix_a = vec![vec![1; 1000]; 1000]; // Small matrix
        let matrix_b = vec![vec![1; 10000]; 10000]; // Big matrix
        let num_workers = 4; // Valid worker number

        let result = calculation_mod::calculate_matrix(&matrix_a, &matrix_b, num_workers);
        assert!(result.is_err());
    }

    #[test]
    fn test_valid_small_matrix_with_valid_worker_number() {
        let matrix_a = vec![vec![1; 1000]; 1000]; // Small matrix
        let matrix_b = vec![vec![1; 1000]; 1000]; // Small matrix
        let num_workers = 20; // Valid worker number

        let result = calculation_mod::calculate_matrix(&matrix_a, &matrix_b, num_workers);
        assert!(result.is_ok());
    }

    #[test]
    fn test_valid_matrices_with_invalid_worker_numbers() {
        let matrix_a = vec![vec![1; 100]; 100]; // Small matrix
        let matrix_b = vec![vec![1; 100]; 100]; // Small matrix

        let num_workers = 0; // Invalid worker number
        let result = calculation_mod::calculate_matrix(&matrix_a, &matrix_b, num_workers);
        assert!(result.is_err());
    }

    #[test]
    fn test_unavailable_matrix_names() {
        let invalid_matrix_name = "generated/resources/non_existent_matrix.txt";

        let result = file_mod::read_matrix_from_file(invalid_matrix_name);
        assert!(result.is_err());
    }

    #[test]
    fn test_dimension_mismatch() {
        let matrix_a = vec![vec![1; 1000]; 1000]; // Small matrix
        let matrix_b = vec![vec![1; 500]; 500]; // Incompatible matrix

        let num_workers = 4; // Valid worker number
        let result = calculation_mod::calculate_matrix(&matrix_a, &matrix_b, num_workers);
        assert!(result.is_err());
    }

    #[test]
    fn test_generate_and_read_matrix() {
        let filename = "generated/resources/test_matrix.txt";
        let size = 1000; // Small matrix

        let generate_result = file_mod::generate_matrix_to_file(filename, size);
        assert!(generate_result.is_ok());

        let read_result = file_mod::read_matrix_from_file(filename);
        assert!(read_result.is_ok());

        // Clean up
        fs::remove_file(filename).unwrap();
    }
}