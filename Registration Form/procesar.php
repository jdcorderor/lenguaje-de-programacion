<?php
    // Configuration to show errors during development
    error_reporting(E_ALL);
    ini_set("display_errors", 1);

    // Function to show error or success messages
    function showMessage($message, $type = 'error') {
        $color = ($type === 'success') ? '#4CAF50' : '#f44336';
        echo "<!DOCTYPE html>
        <html lang=\"en\">
        <head>
            <meta charset=\"UTF-8\">
            <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">
            <link rel=\"stylesheet\" href=\"styles.css\">
            <title>Registration Result</title>
        </head>
        <body>
            <div class=\"message\">
                <h2>$message</h2>
                <a href=\"index.php\" class=\"back-button\">Back to form</a>
            </div>
        </body>
        </html>";
        exit;
    }

    // Verify that a POST request was received
    if ($_SERVER["REQUEST_METHOD"] !== "POST") {
        showMessage("Invalid access method. Please use the registration form");
    }

    // Verify that all required fields are present and not empty
    $name = isset($_POST["nombre"]) ? trim($_POST["nombre"]) : "";
    $email = isset($_POST["email"]) ? trim($_POST["email"]) : "";
    $password = isset($_POST["password"]) ? $_POST["password"] : "";

    // Validate that no field is empty
    if (empty($name)) {
        showMessage("Error: The name field is required");
    }

    if (empty($email)) {
        showMessage("Error: The email field is required");
    }

    if (empty($password)) {
        showMessage("Error: The password field is required");
    }

    // Validate email format
    if (!filter_var($email, FILTER_VALIDATE_EMAIL)) {
        showMessage("Error: The email format is not valid");
    }

    // Validate password strength
    if (strlen($password) < 8) {
        showMessage("Error: Password must be at least 8 characters long");
    }
    
    // Check for at least one special character
    if (!preg_match('/[^a-zA-Z0-9]/', $password)) {
        showMessage("Error: Password must contain at least one special character (e.g., !@#$%^&*)");
    }

    // Use password_hash() to securely hash the password
    $hashedPassword = password_hash($password, PASSWORD_DEFAULT);

    // Create an associative array with the new user"s data
    $newUser = [
        "id" => uniqid(),
        "nombre" => $name,
        "email" => $email,
        "password" => $hashedPassword,
        "registration_datetime" => date("Y-m-d H:i:s")
    ];

    // Define the path to the JSON file
    $usersFile = "usuarios.json";

    // Read existing data from JSON file
    $existingUsers = [];

    if (file_exists($usersFile)) {
        $fileContent = file_get_contents($usersFile);
        if ($fileContent !== false && !empty($fileContent)) {
            $existingUsers = json_decode($fileContent, true);
            
            // If there"s a JSON decoding error, initialize an empty array
            if ($existingUsers === null) {
                $existingUsers = [];
            }
        }
    }

    // Check if email already exists
    foreach ($existingUsers as $user) {
        if ($user["email"] === $email) {
            showMessage("Error: A user with this email is already registered");
        }
    }

    // Add the new user to the main array
    $existingUsers[] = $newUser;

    // Encode to JSON with pretty print
    $jsonData = json_encode($existingUsers, JSON_PRETTY_PRINT | JSON_UNESCAPED_UNICODE);

    // Verify that encoding was successful
    if ($jsonData === false) {
        showMessage("Error: Could not encode data to JSON");
    }

    // Write data to file
    $result = file_put_contents($usersFile, $jsonData, LOCK_EX);

    // Verify that writing was successful
    if ($result === false) {
        showMessage("Error: Could not save data to file");
    }

    // If we got here, registration was successful
    showMessage("Success! User registered successfully", "success");
?>
