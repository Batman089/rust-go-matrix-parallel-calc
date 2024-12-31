#[cfg(test)]
mod integration_tests {
    use std::fs;
    use std::path::Path;
    use std::process::Command;

    #[test]
    fn test_full_workflow() {
        // Ensure the resources directory exists
        fs::create_dir_all("resources").unwrap();

        // Generate matrices
        let output = Command::new("cargo")
            .args(&["run", "--", "generate", "small", "generated/resources/matrix_a.txt"])
            .output()
            .expect("Failed to generate matrix A");
        assert!(output.status.success());

        let output = Command::new("cargo")
            .args(&["run", "--", "generate", "big", "generated/resources/matrix_b.txt"])
            .output()
            .expect("Failed to generate matrix B");
        assert!(output.status.success());

        // Check if files are created
        assert!(Path::new("generated/resources/matrix_a.txt").exists());
        assert!(Path::new("generated/resources/matrix_b.txt").exists());

        // Perform matrix calculation
        let output = Command::new("cargo")
            .args(&["run", "--", "calculate", "generated/resources/matrix_a.txt", "generated/resources/matrix_b.txt", "4"])
            .output()
            .expect("Failed to calculate matrices");
        assert!(output.status.success());

        // Check if log file is created
        assert!(Path::new("log/calc_time_log").exists());

        // Clean up
        fs::remove_file("generated/resources/matrix_a.txt").unwrap();
        fs::remove_file("generated/resources/matrix_b.txt").unwrap();
        fs::remove_file("generated/log/calc_time_log").unwrap();
    }
}